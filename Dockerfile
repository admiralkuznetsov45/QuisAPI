# stage 1
FROM golang:1.16-alpine AS builder
RUN mkdir /
ADD . /
WORKDIR /
RUN go clean --modcache
RUN go build /main.go
EXPOSE 8080
CMD ["/main"]

# stage 2
FROM alpine:3.14
WORKDIR /root/
COPY --from=builder /main .
EXPOSE 8080
CMD ["./main"]