package statistics_listener

import (
	"context"
	"log"

	"statistics_service/pkg/pb"

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

		event := &pb.Event{}
		if err := proto.Unmarshal(msg.Value, event); err != nil {
			log.Printf("failed to unmarshal event: %v", err)
			continue
		}

		switch e := event.EventType.(type) {
		case *pb.Event_ViewEvent:
			viewEvent := e.ViewEvent
			err := sl.clickhouse.SaveView(viewEvent.UserId, viewEvent.PostId)
			if err != nil {
				log.Printf("failed to save view: %v", err)
			}
		case *pb.Event_LikeEvent:
			likeEvent := e.LikeEvent
			err := sl.clickhouse.SaveLike(likeEvent.UserId, likeEvent.PostId)
			if err != nil {
				log.Printf("failed to save like: %v", err)
			}
		default:
			log.Printf("unknown event type: %T", e)
		}
	}
}
