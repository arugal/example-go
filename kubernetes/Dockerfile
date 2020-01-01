FROM golang:1.13

ENV GOPROXY=https://goproxy.io
ENV CGO_ENABLED=0

RUN GOOS=linux go build -o ./app .
RUN mv ./app /app

FROM debian
COPY --from=0 /app /app
ENTRYPOINT /app