FROM golang:1.25-trixie

ARG PACKAGES="vim curl less git bash screen"

RUN apt-get update && apt-get install -y $PACKAGES

WORKDIR /app
COPY . /app
