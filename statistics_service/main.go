package main

import (
	"log"
	"net/http"
	"os"

	"github.com/segmentio/kafka-go"

	st "statistics_service/pkg"
)

const (
	topic   = "my-topic"
	groupID = "consumer-group-id"
)

func main() {
	kafkaAddr, ok := os.LookupEnv("KAFKA_SERVER")
	if !ok {
		log.Fatalf("KAFKA_SERVER not set")
	}

	l := log.New(os.Stdout, "kafka reader: ", 0)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaAddr},
		Topic:   topic,
		Logger:  l,
		GroupID: groupID,
	})

	clickhouse, err := st.NewClickhouse()
	if err != nil {
		log.Fatalf("failed to create clickhouse: %v", err)
	}

	statsListener := st.NewStatisticsListener(reader, clickhouse)

	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

		http.ListenAndServe(":80", nil)
	}()

	grpcServer, err := st.NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	go grpcServer.Listen()

	statsListener.Listen()
}
