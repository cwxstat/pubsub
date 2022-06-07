package pub

import (
	"sync"
	"time"
)

type MSG struct {
	Message   string
	timeStamp time.Time
	Duration  time.Duration
	Action    string
	Act       int
}

type Queue struct {
	Message []MSG
}

type Pub struct {
	topic map[string]Queue
	sync.Mutex
}

func NewPub() *Pub {
	p := &Pub{}
	p.topic = map[string]Queue{}
	return p
}

func (p *Pub) Add(topic string, msg MSG) {
	msg.timeStamp = time.Now()

	if queue, ok := p.topic[topic]; ok {
		queue.Message = append(queue.Message, msg)
		p.topic[topic] = queue
	} else {
		queue := Queue{}
		queue.Message = []MSG{}
		queue.Message = append(queue.Message, msg)
		p.topic[topic] = queue

	}

}
