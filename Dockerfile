FROM golang AS builder
WORKDIR /go/src/github.com/htw-swa-jk-nk-ns/service-calculate
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/htw-swa-jk-nk-ns/service-calculate/app .
CMD ["./app"]