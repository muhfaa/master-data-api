FROM golang:alpine

WORKDIR /master-app

ADD . .

RUN go mod download

RUN go mod verify

RUN go build -o master-app ./app/main.go

ENTRYPOINT ["./master-app"]
