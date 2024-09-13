package grpchandler

import (
	"context"

	"github.com/TomasCruz/users/internal/handlers/grpchandler/users"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (g *GRPCHandler) GetUserByID(ctx context.Context, req *users.UserIDMsg) (*users.UserMsg, error) {
	g.logger.Info(nil, "GRPCHandler.GetUserHandler")

	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	user, err := g.cr.GetUserByID(uid)
	if err != nil {
		g.logger.Error(err, "GRPCHandler.GetUserHandler")
		return nil, err
	}

	return &users.UserMsg{
		Id:        user.UserID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		PswdHash:  user.PswdHash,
		Email:     user.Email,
		Country:   user.Country,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}
