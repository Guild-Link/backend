package hypixel

import (
	"github.com/guild-link/backend/pkg/cache"
	"github.com/guild-link/backend/pkg/hypixel"
	pb "github.com/guild-link/backend/proto/hypixel"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHypixelServer
	client *hypixel.Client
}

func Register(s grpc.ServiceRegistrar, apiKey string, cache *cache.Cache) {
	pb.RegisterHypixelServer(s, &Server{client: hypixel.NewClient(apiKey, cache)})
}
