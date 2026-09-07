package cache

import (
	"time"

	"github.com/valkey-io/valkey-go/valkeyaside"
)

type Cache struct {
	client  valkeyaside.CacheAsideClient
	ttl     time.Duration
	timeout time.Duration
}
