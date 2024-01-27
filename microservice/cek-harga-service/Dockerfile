FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o ./microservice/cek-harga-service

EXPOSE 8082

CMD ["./microservice/cek-harga-service"]