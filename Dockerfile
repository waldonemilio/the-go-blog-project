FROM golang:alpine AS builder

# Use the 'app' directory as the working directory
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . . 

RUN go mod tidy
