FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o ./microservice/topup-service

EXPOSE 8082

CMD ["./microservice/topup-service"]