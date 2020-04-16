FROM golang:latest as builder
ARG SERVICE
ARG VERSION
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags "-w -X main.version=${VERSION}" -i -o ./build/bin ./cmd/${SERVICE}

FROM alpine:latest
LABEL maintainer="Mohsen Samiei <eng.msamiee@gmail.com>"
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /src/build ./
CMD ["./bin"]