package hypixel

import (
	"net/http"

	sc "github.com/DuckySoLucky/SkyCrypt-Types"
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

type SkyBlockProfile struct {
	Mojang *mojang.Profile

	Selected bool
	ID       string
	Name     string
	GameMode string

	Data              *sc.Member
	Banking           *sc.Banking
	CommunityUpgrades *sc.CommunityUpgrades
}
