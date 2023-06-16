# Whisper Service

[![test status](https://github.com/d-ashesss/whisper-service/workflows/test/badge.svg?branch=main)](https://github.com/d-ashesss/whisper-service/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/d-ashesss/whisper-service)](https://goreportcard.com/report/github.com/d-ashesss/whisper-service)
[![MIT license](https://img.shields.io/github/license/d-ashesss/whisper-service?color=blue)](https://opensource.org/licenses/MIT)
[![Go version](https://img.shields.io/github/go-mod/go-version/d-ashesss/whisper-service)](https://github.com/d-ashesss/whisper-service/blob/main/go.mod)
[![latest tag](https://img.shields.io/github/v/tag/d-ashesss/whisper-service?include_prereleases&sort=semver)](https://github.com/d-ashesss/whisper-service/tags)
[![feline reference](https://img.shields.io/badge/may%20contain%20cat%20fur-%F0%9F%90%88-blueviolet)](https://github.com/d-ashesss/whisper-service)

gRPC service for a self-hosted [openai/whisper](https://github.com/openai/whisper).

## Running with Docker

To get and run the service from Docker Hub:

```shell
docker pull ashesss/openai-whisper-service:latest
docker run -p 8080:8080 ashesss/openai-whisper-service:latest
```

### GPU

To utilize the GPU for the service, first make sure that nvidia plugin for docker is installed then start ther container with GPU support:

```shell
docker run -p 8080:8080 --gpus all ashesss/openai-whisper-service:latest
```

### Model Cache

ASR models will be downloaded for each new container, if you want to cache them on persistent volume, mount it as follows:

```shell
docker run -p 8080:8080 -v /path/to/cached/models:/root/.cache/whisper ashesss/openai-whisper-service:latest
```

## Basic Usage

With Go you can simply import client from this repo:

```go
import whisperpb "github.com/d-ashesss/whisper-service/proto"

func main() {
	conn, _ := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	c := whisperpb.NewWhisperServiceClient(conn)
	stream, err := c.Transcribe(context.Background())
	...
}
```

For other languages use [whisper.proto](https://github.com/d-ashesss/whisper-service/blob/main/proto/whisper.proto) file to generate gRPC client.
