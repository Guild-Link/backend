package hypixel

import (
	"github.com/guild-link/hypixel-go/pkg/cache"
	"github.com/guild-link/hypixel-go/pkg/hypixel"
	pb "github.com/guild-link/hypixel-go/proto/hypixel"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHypixelServer
	client *hypixel.Client
}

func Register(s grpc.ServiceRegistrar, apiKey string, cache *cache.Cache) {
	pb.RegisterHypixelServer(s, &Server{client: hypixel.NewClient(apiKey, cache)})
}
