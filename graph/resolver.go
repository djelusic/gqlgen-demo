package graph

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"example.com/gqlgen-demo/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	offer           *model.Offer
	updateChan      chan *model.Offer
	broker *broker
	sync.Mutex
}

func NewResolver() *Resolver {
	o := &model.Offer{
		Events: model.EventsMock,
	}
	r := &Resolver{
		offer:           o,
		updateChan:      make(chan *model.Offer),
	}
	r.broker = newBroker(r)
	go r.updateLoop()
	return r
}

func (r *Resolver) updateLoop() {
	for range time.NewTicker(1 * time.Second).C {
		r.Lock()
		// pick a random odd to update
		var odds []*model.Odd
		for _, e := range r.offer.Events {
			for _, m := range e.Markets {
				for _, o := range m.Odds {
					odds = append(odds, o)
				}
			}
		}
		od := odds[rand.Intn(len(odds))]
		newOdd := rand.Float64() * 10
		od.Value = &newOdd
		var ev *model.Event
		var ma *model.Market
		for _, e := range r.offer.Events {
			for _, m := range e.Markets {
				for _, o := range m.Odds {
					if o.ID == od.ID {
						ev = e
						ma = m
						break
					}
				}
			}
		}
		if ev == nil || ma == nil {
			return
		}
		evDiff := &model.Event{
			ID: ev.ID,
			Markets: []*model.Market{
				{
					ID: ma.ID,
					Odds: []*model.Odd{od},
				},
			},
		}
		r.Unlock()
		r.updateChan <- &model.Offer{
			Events: []*model.Event{evDiff},
		}
	}
}

type subscriber struct {
	ch chan *model.Offer
}

type broker struct {
	subId int
	subscribers map[int]*subscriber
	resolver *Resolver
	sync.Mutex
}

func (b *broker) loop() {
	for u := range b.resolver.updateChan {
		b.Lock()
		for _, s := range b.subscribers {
			s.ch <- u
		}
		b.Unlock()
	}
}

func newBroker(r *Resolver) *broker {
	b := &broker{
		subId: 0,
		subscribers: make(map[int]*subscriber),
		resolver: r,
	}
	go b.loop()
	return b
}

func (b *broker) addSubscriber(ctx context.Context) *subscriber {
	b.Lock()
	defer b.Unlock()
	s := &subscriber{
		ch: make(chan *model.Offer),
	}
	sid := b.subId
	b.subscribers[sid] = s
	b.subId++
	go func() {
		<- ctx.Done()
		b.Lock()
		defer b.Unlock()
		delete(b.subscribers, sid)
	}()
	return s
}
