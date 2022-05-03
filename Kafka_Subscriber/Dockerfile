FROM golang:1.16
WORKDIR /Kafka_Subscriber
COPY . .
RUN  go mod download
CMD ["go","run","main.go"]