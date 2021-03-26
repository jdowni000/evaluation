
FROM golang:alpine

ADD . /app

WORKDIR /app/cmd/app

RUN go build -v

ENTRYPOINT [ "/app/cmd/app/app" ]

