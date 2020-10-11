package main

import (
	"encoding/json"
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

	message := Request{Op: "sqrt", args: c.prepareArgs()}

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

func (c *ClientRabbitMQ) MakeRequestBenchmark() ([]float64, int64, error) {
	msgsFromServer, err := c.channel.Consume(c.replyQueue.Name, "", true, false, false, false, nil)

	if err != nil {
		return nil, 0, err
	}

	startTime := time.Now()

	message := Request{Op: "sqrt", args: c.prepareArgs()}

	msgRequestBytes, err := json.Marshal(message)

	if err != nil {
		return nil, 0, err
	}

	err = c.channel.Publish("", c.requestQueue.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes})

	if err != nil {
		return nil, 0, err
	}

	var response Reply

	err = json.Unmarshal((<-msgsFromServer).Body, &response)

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
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/vhost")

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

	replyQueue, err := ch.QueueDeclare("response", false, false, false, false, nil)

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

func (c *ClientRabbitMQ) prepareArgs() Args {
	return Args{
		A: 4,
		B: 3,
		C: -5,
	}
}
