package adder

// это локальный пакет, но он что то не работает
import "grpcadder/pkg/api"

// GRPCServer ...
type GRPCServer struct {}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error)
	return &api.AddResponse{Result: req.Getx() + req.GetY()}, nil
}

