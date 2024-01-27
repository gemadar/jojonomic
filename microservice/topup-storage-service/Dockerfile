FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o ./microservice/topup-storage-service

EXPOSE 8082

CMD ["./microservice/topup-storage-service"]