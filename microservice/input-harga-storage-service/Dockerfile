FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o ./microservice/input-harga-storage-service

EXPOSE 8082

CMD ["./microservice/input-harga-storage-service"]