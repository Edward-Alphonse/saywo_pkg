package storage

import "context"

type DBStorage interface {
	//Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key uint64) (any, error)
}
