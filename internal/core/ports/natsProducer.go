package ports

import (
	"time"
)

type NatsProducer interface {
	UserCreatedRequest(data []byte, timeout time.Duration) (string, error)
	Drain() error
}
