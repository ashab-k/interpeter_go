FROM golang:1.23.4-alpine

WORKDIR /app

COPY . .

RUN go build -o interpreter .

ENTRYPOINT ["/app/interpreter"]
