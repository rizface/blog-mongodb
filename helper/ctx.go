package helper

import (
	"context"
	"time"
)

func CreateCtx(dur time.Duration) (context.Context,context.CancelFunc) {
	ctx,cancel := context.WithTimeout(context.Background(),dur * time.Second)
	return ctx,cancel
}
