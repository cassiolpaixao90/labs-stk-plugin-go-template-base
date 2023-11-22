FROM golang:{{version_golang}}-alpine as build

ENV CGO_ENABLED=0
ENV GO111MODULE="on"

WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o server

FROM alpine:latest
WORKDIR /app
RUN addgroup --gid 101 -S gouser && adduser --uid 1001 -S gouser -G gouser
RUN chown -R gouser:gouser /app
USER gouser:gouser
COPY --from=build /app/server .
EXPOSE 8080
CMD ["./server"]

