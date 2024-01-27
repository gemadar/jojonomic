FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o ./microservice/buyback-storage-service

EXPOSE 8082

CMD ["./microservice/buyback-storage-service"]