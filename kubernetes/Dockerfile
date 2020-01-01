FROM golang:1.13

ENV GOPROXY=https://goproxy.io
ENV CGO_ENABLED=0

RUN mkdir /app
COPY main.go /app/main.go
WORKDIR /app
RUN GOOS=linux go build -o ./client .

FROM debian
COPY --from=0 /app/client /client
ENTRYPOINT /client