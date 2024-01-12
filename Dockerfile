FROM golang:1.21.5-alpine as builder
WORKDIR /build
COPY go.mod . 
RUN go mod download
COPY . .
ARG service
RUN go build -o $service ./cmd/$service/main.go

FROM alpine:3
ARG service
COPY --from=builder /build/$service /bin/main
ENTRYPOINT ["/bin/main"]