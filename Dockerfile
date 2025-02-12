FROM postgres:latest
FROM golang:latest

RUN apt-get update && apt-get install -y pgcli
ENV PATH=$PATH:/usr/bin/pgcli

COPY . .
WORKDIR /cmd

RUN go mod download
RUN go build -o app

CMD [ "./app" ]

