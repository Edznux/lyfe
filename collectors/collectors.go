package collectors

import "context"

type Collector interface {
	Collect(ctx context.Context, userid int64) error
	Init() error
}
