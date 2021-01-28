package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/JesKetchupson/testtask/graph/generated"
	"github.com/JesKetchupson/testtask/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	p := &model.Post{
		ID:              uuid.New().String(),
		Description:     input.Description,
		Title:           input.Title,
		PublicationDate: time.Now(),
	}
	r.posts = append(r.posts, p)
	go func() {
		if r.sub["sub"] != nil {
			r.sub["sub"] <- p
		}
	}()
	return p, nil
}

func (r *queryResolver) Posts(ctx context.Context, first *int, offset *int) ([]*model.Post, error) {

	*offset += *first
	if cap(r.posts) < *offset {
		*offset = cap(r.posts)
	}
	if 0 > *offset {
		*offset = 0
	}

	if len(r.posts) < *first {
		*first = 0
	}
	if 0 > *first {
		*first = 0
	}

	return r.posts[*first:*offset], nil
}

func (r *subscriptionResolver) PostAdded(ctx context.Context, channName string) (<-chan *model.Post, error) {
	events := make(chan *model.Post, 1)
	r.sub["sub"] = events
	return events, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
