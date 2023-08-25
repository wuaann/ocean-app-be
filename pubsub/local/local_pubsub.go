package pubsublocal

import (
	"context"
	"ocean-app-be/common"
	"ocean-app-be/pubsub"

	"log"
	"sync"
)

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewPubSub() *localPubSub {
	ps := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	ps.run()

	return ps
}

func (ps *localPubSub) Publish(
	ctx context.Context,
	topic pubsub.Topic,
	data *pubsub.Message,
) error {
	data.SetChannel(topic)

	go func() {

		defer common.AppRecover()
		ps.messageQueue <- data // enqueue message
		log.Println("New event published", data.String(), "with data", data.Data())
	}()

	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock()

	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {

					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}
}

func (ps *localPubSub) run() error {
	log.Println("Pubsub started")

	go func() {
		for { // for forever
			mess := <-ps.messageQueue
			log.Println("Message dequeue:", mess)

			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						c <- mess
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
