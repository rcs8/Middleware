package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

type ClientRabbitMQ struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	requestQueue *amqp.Queue
	replyQueue   *amqp.Queue
}

func (c *ClientRabbitMQ) MakeRequest() ([]float64, error) {
	msgsFromServer, err := c.channel.Consume(c.replyQueue.Name, "", true, false, false, false, nil)

	if err != nil {
		return nil, err
	}

	message := c.prepareArgs()

	msgRequestBytes, err := json.Marshal(message)

	if err != nil {
		return nil, err
	}

	err = c.channel.Publish("", c.requestQueue.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes})

	if err != nil {
		return nil, err
	}

	var response Reply

	err = json.Unmarshal((<-msgsFromServer).Body, &response)

	if err != nil {
		return nil, err
	}

	return response.Result, err
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func (c *ClientRabbitMQ) MakeRequestBenchmark() ([]float64, int64, error) {
	msgsFromServer, err := c.channel.Consume(c.replyQueue.Name, "", true, false, false, false, nil)

	if err != nil {
		panic(err) // debug
		return nil, 0, err
	}

	startTime := time.Now()

	message := c.prepareArgs()

	msgRequestBytes, err := json.Marshal(message)

	if err != nil {
		return nil, 0, err
	}

	corrID := randomString(32)
	err = c.channel.Publish("", c.requestQueue.Name, false, false, amqp.Publishing{
		ContentType:   "text/plain",
		CorrelationId: corrID,
		ReplyTo:       "amq.rabbitmq.reply-to",
		Body:          msgRequestBytes,
	})

	if err != nil {
		return nil, 0, err
	}

	var response Reply

	for msg := range msgsFromServer {
		if corrID == msg.CorrelationId {
			err = json.Unmarshal(msg.Body, &response)
			break
		}
	}

	totalTime := time.Now().Sub(startTime).Microseconds()

	if err != nil {
		return nil, 0, err
	}

	return response.Result, totalTime, err
}

func (c *ClientRabbitMQ) Close() {
	(*c.conn).Close()
	(*c.channel).Close()
}

func NewClientRabbitMQ(address string) (*ClientRabbitMQ, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	requestQueue, err := ch.QueueDeclare("request", false, false, false, false, nil)

	if err != nil {
		return nil, err
	}

	replyQueue, err := ch.QueueDeclare("amq.rabbitmq.reply-to", false, false, false, false, nil)

	if err != nil {
		return nil, err
	}

	return &ClientRabbitMQ{
		conn:         conn,
		channel:      ch,
		requestQueue: &requestQueue,
		replyQueue:   &replyQueue,
	}, err
}

func (c *ClientRabbitMQ) prepareArgs() Request {
	return Request{
		Op: "sqrt",
		A:  4,
		B:  3,
		C:  -5,
	}
}
