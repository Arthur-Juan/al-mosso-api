FROM golang:1.20.10-alpine

WORKDIR app/

COPY . .

RUN go get -d -v ./...

RUN go build -o build/main cmd/api/main.go

EXPOSE 8080

CMD ["./build/main"]