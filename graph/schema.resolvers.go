package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"example.com/gqlgen-demo/graph/generated"
	"example.com/gqlgen-demo/graph/model"
)

func (r *queryResolver) Offer(ctx context.Context) (*model.Offer, error) {
	r.Lock()
	defer r.Unlock()
	return r.offer, nil
}

func (r *queryResolver) Event(ctx context.Context, id int) (*model.Event, error) {
	r.Lock()
	defer r.Unlock()
	for _, e := range r.offer.Events {
		if e.ID == id {
			return e, nil
		}
	}
	return nil, fmt.Errorf("event not found")
}

func (r *subscriptionResolver) Offer(ctx context.Context) (<-chan *model.Offer, error) {
	s := r.broker.addSubscriber(ctx)
	return s.ch, nil
}

func (r *subscriptionResolver) Event(ctx context.Context, id int) (<-chan *model.Event, error) {
	s := r.broker.addSubscriber(ctx)
	ch := make(chan *model.Event)
	go func() {
		for o := range s.ch {
			for _, e := range o.Events {
				if id == e.ID {
					ch <- e
				}
			}
		}
	}()
	return ch, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
