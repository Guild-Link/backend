package hypixel

import (
	"github.com/guild-link/backend/pkg/cache"
	"github.com/guild-link/backend/pkg/hypixel"
	pb "github.com/guild-link/backend/proto/hypixel"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHypixelServer
	hypixel *hypixel.Client
}

func Register(s grpc.ServiceRegistrar, c *cache.Cache, apiKey, compatURL string) {
	pb.RegisterHypixelServer(s, &Server{
		hypixel: hypixel.NewClient(c, apiKey, compatURL),
	})
}

func profileResponse(profile *hypixel.SkyBlockProfile) *pb.SkyBlockProfile {
	return &pb.SkyBlockProfile{
		Username: profile.Player.Name,
		Uuid:     profile.Player.ID,
		Profile:  profile.Name,
	}
}
