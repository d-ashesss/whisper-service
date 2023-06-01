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
	server *grpc.Server
}

func NewApp(server *grpc.Server) *App {
	return &App{
		server: server,
	}
}

func (a *App) Run() {
	log.Print("[app] Starting up")
	signalCtx, signalStop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	defer signalStop()

	wg, gCtx := errgroup.WithContext(signalCtx)

	wg.Go(func() error {
		lis, err := net.Listen("tcp", ":10000")
		if err != nil {
			log.Print("[app] Unable to open socket:", err)
			return err
		}
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
