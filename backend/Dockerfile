FROM golang:1.20
# RUN apk update &&  apk add git
# RUN go get github.com/cosmtrek/air@v1.29.0
WORKDIR /app
RUN apt-get update -y&& apt-get install -y git curl vim libmagickwand-dev
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app
# EXPOSE 8000
# CMD ["app"]