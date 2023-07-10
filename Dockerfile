FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o=./bin/api ./cmd/api/main.go

EXPOSE 3030
CMD [ "./app/bin/api" ]

FROM golang:alpine

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build ./cmd/api

EXPOSE 3030

# Run the executable
CMD [ "/build" ]
