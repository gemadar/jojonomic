FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o ./microservice/buyback-service

EXPOSE 8082

CMD ["./microservice/buyback-service"]