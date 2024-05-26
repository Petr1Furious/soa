package statistics_listener

import (
	"context"
	"log"

	kpb "statistics_service/pkg/kafka_pb"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type StatisticsListener struct {
	kafkaReader *kafka.Reader
	clickhouse  *Clickhouse
}

func NewStatisticsListener(kafkaReader *kafka.Reader, clickhouse *Clickhouse) *StatisticsListener {
	return &StatisticsListener{
		kafkaReader: kafkaReader,
		clickhouse:  clickhouse,
	}
}

func (sl *StatisticsListener) Listen() {
	for {
		msg, err := sl.kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}

		event := &kpb.Event{}
		if err := proto.Unmarshal(msg.Value, event); err != nil {
			log.Printf("failed to unmarshal event: %v", err)
			continue
		}

		switch e := event.EventType.(type) {
		case *kpb.Event_ViewEvent:
			viewEvent := e.ViewEvent
			err := sl.clickhouse.SaveView(viewEvent.UserId, viewEvent.PostId, viewEvent.AuthorId)
			if err != nil {
				log.Printf("failed to save view: %v", err)
			}
		case *kpb.Event_LikeEvent:
			likeEvent := e.LikeEvent
			err := sl.clickhouse.SaveLike(likeEvent.UserId, likeEvent.PostId, likeEvent.AuthorId)
			if err != nil {
				log.Printf("failed to save like: %v", err)
			}
		default:
			log.Printf("unknown event type: %T", e)
		}
	}
}
