FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o ./microservice/input-harga-service

EXPOSE 8082

CMD ["./microservice/input-harga-service"]