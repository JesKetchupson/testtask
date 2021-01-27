package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/nwk123/testtask/graph/generated"
	"github.com/nwk123/testtask/graph/model"
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
			<-r.sub["sub"]
			r.sub["sub"] <- p
		}
	}()
	return p, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return r.posts, nil
}

func (r *subscriptionResolver) PostAdded(ctx context.Context, channName string) (<-chan *model.Post, error) {
	return r.sub["sub"], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver {
	postCh := make(chan *model.Post, 1)
	r.sub["sub"] = postCh
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
