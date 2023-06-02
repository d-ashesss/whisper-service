package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	config *Config
	server *grpc.Server
}

func NewApp(config *Config, server *grpc.Server) *App {
	return &App{
		config: config,
		server: server,
	}
}

func (a *App) Run() {
	log.Print("[app] Starting up")
	signalCtx, signalStop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	defer signalStop()

	wg, gCtx := errgroup.WithContext(signalCtx)

	wg.Go(func() error {
		addr := ":" + a.config.Port
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Printf("[app] Unable to start listener at %q: %s", addr, err)
			return err
		}
		log.Printf("[app] Starting gRPC server at %q", addr)
		if err := a.server.Serve(lis); err != nil {
			log.Print("[app] gRPC server has stopped unexpectedly: ", err)
			return err
		}
		return nil
	})
	wg.Go(func() error {
		<-gCtx.Done()
		log.Print("[app] Shutting down gRPC server")
		a.server.GracefulStop()
		return nil
	})
	wg.Go(func() error {
		<-gCtx.Done()
		signalStop()
		return nil
	})

	if err := wg.Wait(); err != nil {
		log.Print("[app] Unexpected exit reason:", err)
	}
}
