FROM golang:1.16-alpine as builder

ENV CGO_ENABLED=1
ENV GO111MODULE=on

ENV environment=local

# Maintainer info
LABEL maintainer="Bronsun Mateusz Broncel"

RUN apk add --no-cache git  git gcc g++

# install msgfmt
RUN apk -U add gettext

# Change workdir to expeceted
WORKDIR /app

# Downloading all go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Chnage to workdir where is a server
WORKDIR /app/

# Building go project 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/


# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/config/*.env ./config/  

EXPOSE 8080

CMD ["sh","-c","./main -e ${environment}"]