FROM golang:1.24.5-alpine3.21

WORKDIR /app

COPY . .

RUN go build -o api

EXPOSE 3000

CMD ["./api"]