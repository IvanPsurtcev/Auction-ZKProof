FROM golang:1.20.4-bullseye

WORKDIR app

COPY . .

RUN go mod download
