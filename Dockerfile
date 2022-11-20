FROM golang:1.19.3-alpine3.16

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV FROM="echez.kojo@gmail.com"
ENV PASS="xxxxxxx"
ENV TO="echez.kojo@gmail.com"
ENV HOST="smtp.gmail.com"
ENV EMAILPORT="587"
ENV IDENTITY="From: Echez Kojo <echez.kojo@gmail.com>\r\n"
ENV DEVDB="host=localhost port=5432 user=postgres password=password dbname=contact-server sslmode=require"

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]