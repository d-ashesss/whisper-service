FROM golang:1.20 AS builder
COPY . /tmp/src
WORKDIR /tmp/src
RUN go build -o whisper-service .

FROM python:3.11-slim
RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -y ffmpeg git
RUN pip install --no-cache-dir git+https://github.com/openai/whisper.git
COPY --from=builder /tmp/src/whisper-service /bin/whisper-service
ENTRYPOINT ["/bin/whisper-service"]
