package mojang

import (
	"context"

	"github.com/guild-link/backend/pkg/cache"
	"github.com/guild-link/backend/pkg/mojang"
	pb "github.com/guild-link/backend/proto/mojang"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedMojangServer
	client *mojang.Client
}

func Register(s grpc.ServiceRegistrar, cache *cache.Cache) {
	pb.RegisterMojangServer(s, &Server{client: mojang.NewClient(cache)})
}

func (s *Server) GetProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileReponse, error) {
	p, err := s.client.GetProfile(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	return &pb.ProfileReponse{
		Name: p.Name,
		Id:   p.ID,
	}, nil
}
