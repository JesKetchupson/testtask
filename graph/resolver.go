package graph

import (
	"github.com/nwk123/testtask/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	sub map[string]chan *model.Post

	posts []*model.Post
}
