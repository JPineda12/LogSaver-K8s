package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	pb "github.com/diemorales96/Sopes1_ProyectoF2/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedJuegosServer
}
type GameResult struct {
	Game_id       int32  `json:"game_id"`
	Players       int32  `json:"players"`
	Game_name     string `json:"game_name"`
	Winner_number int32  `json:"winner_number"`
	Queue         string `json:"queue"`
}

var topic = "logs"
var broker1Address = "my-cluster-kafka-bootstrap.kafka"

func (s *server) AddLog(ctx context.Context, in *pb.JuegoRequest) (*pb.RequestReply, error) {
	result := GameResult{}
	if in.GetGameId() == 1 {
		result = gameNumberOne(in)
	} else if in.GetGameId() == 2 {
		result = gameNumberTwo(in)
	} else if in.GetGameId() == 3 {
		result = gameNumberThree(in)
	} else if in.GetGameId() == 4 {
		result = gameNumberFour(in)
	} else if in.GetGameId() == 5 {
		result = gameNumberFive(in)
	}
	// intialize the writer with the broker addresses, and the topic
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker1Address,
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	msg, err2 := json.Marshal(result)
	if err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	delivery_chan := make(chan kafka.Event, 10000)
	_ = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg},
		delivery_chan,
	)
	return &pb.RequestReply{Message: "Log Enviado a Kafka"}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterJuegosServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func gameNumberOne(prueba *pb.JuegoRequest) GameResult {
	//var ganador int
	ganador := rand.Intn(int(prueba.Players))
	resultado := GameResult{}
	resultado.Game_id = prueba.GameId
	resultado.Players = prueba.Players
	resultado.Game_name = "Random"
	resultado.Winner_number = int32(ganador)
	resultado.Queue = "Kafka"
	return resultado
}

func gameNumberTwo(prueba *pb.JuegoRequest) GameResult {
	ganador := int32(prueba.Players / 2)
	resultado := GameResult{}
	resultado.Game_id = prueba.GameId
	resultado.Players = prueba.Players
	resultado.Game_name = "Half Game"
	resultado.Winner_number = ganador
	resultado.Queue = "Kafka"
	return resultado
}

func gameNumberThree(prueba *pb.JuegoRequest) GameResult {
	ganador := prueba.Players
	resultado := GameResult{}
	resultado.Game_id = prueba.GameId
	resultado.Players = prueba.Players
	resultado.Game_name = "The Last One"
	resultado.Winner_number = ganador
	resultado.Queue = "Kafka"
	return resultado
}

func gameNumberFour(prueba *pb.JuegoRequest) GameResult {
	var ganador int
	for {
		ganador = rand.Intn(int(prueba.Players))
		if (ganador)%2 == 0 {
			break
		}
	}

	resultado := GameResult{}
	resultado.Game_id = prueba.GameId
	resultado.Players = prueba.Players
	resultado.Game_name = "Even Game"
	resultado.Winner_number = int32(ganador)
	resultado.Queue = "Kafka"
	return resultado
}

func gameNumberFive(prueba *pb.JuegoRequest) GameResult {
	var ganador int
	for {
		ganador = rand.Intn(int(prueba.Players))
		if (ganador)%2 != 0 {
			break
		}
	}
	resultado := GameResult{}
	resultado.Game_id = prueba.GameId
	resultado.Players = prueba.Players
	resultado.Game_name = "Odd Game"
	resultado.Winner_number = int32(ganador)
	resultado.Queue = "Kafka"
	return resultado
}
