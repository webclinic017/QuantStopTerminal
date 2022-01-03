package base

import (
	"context"
)

type HttpAPI interface {
	Get(ctx context.Context, relativePath string, result interface{}) error
	Post(ctx context.Context, relativePath string, content interface{}, result interface{}) error
	Do(ctx context.Context, method string, relativePath string, content interface{}, result interface{}) (capture error)
	Close() error
}
