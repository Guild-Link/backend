package hypixel

import (
	"net/http"

	"github.com/guild-link/backend/pkg/cache"
	"github.com/guild-link/backend/pkg/compatlink"
	"github.com/guild-link/backend/pkg/mojang"
)

type Client struct {
	apiKey string
	http   http.Client

	cache  *cache.Cache
	mojang *mojang.Client
	compat *compatlink.Client
}
