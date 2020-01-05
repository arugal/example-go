FROM golang:1.13

ENV GOPROXY=https://goproxy.io
ENV CGO_ENABLED=0

RUN mkdir /example-go
WORKDIR /example-go
ADD . /example-go
RUN go build -o bin/projectA ./istio-demo/projectA/main.go

FROM alpine:3.10

RUN mkdir /app
WORKDIR /app
COPY --from=0 /example-go/bin/projectA .
ENTRYPOINT ["./projectA"]
EXPOSE 8081