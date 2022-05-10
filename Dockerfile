FROM golang:1.18-bullseye

ENV GIN_MODE=debug

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /followsmartmoney-gin cmd/app/main.go

EXPOSE 80

CMD ["/followsmartmoney-gin"]