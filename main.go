package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"github.com/yen0x/events_perf/model"
	"log"
	"time"
)

func main() {
	count := 1000
	conn, err := amqp.Dial("amqp://user:user@server.lxc:5672/password")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	start := time.Now()
	for i := 0; i < count; i++ {

		body, err := generateEvent()
		failOnError(err, "Failed to serialize event")

		err = ch.Publish(
			"events", // exchange
			"",       // routing key
			false,    // mandatory
			false,    // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		failOnError(err, "Failed to publish a message")
	}
	elapsed := time.Since(start)

	log.Printf(" [x] Sent %d message in %s", count, elapsed)

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func generateEvent() ([]byte, error) {
	uuid1, _ := uuid.NewRandom()
	uuid2, _ := uuid.Parse("1234abcd-022d-4db9-a2d1-a964fa1353e5")
	actor := model.Actor{
		ActorType: "user",
		Data:      json.RawMessage(`{"id": "1234abcd-022d-4db9-a2d1-a964fa1353e5"}`),
	}
	event := model.Event{
		EventId:     uuid1,
		Type:        "user.created",
		Actor:       actor,
		ClientId:    uuid2,
		Application: "rh2",
		CreatedAt:   time.Now().Format(time.RFC3339),
		Data:        json.RawMessage(`{}`),
	}

	return json.Marshal(event)
}
