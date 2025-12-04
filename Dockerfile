FROM golang:1.25-trixie

ARG PACKAGES="vim curl less git bash screen"

RUN apt-get update && apt-get install -y $PACKAGES

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app
