FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o ./microservice/cek-mutasi-service

EXPOSE 8082

CMD ["./microservice/cek-mutasi-service"]