package grpchandler

import (
	"context"
	"strconv"

	"github.com/TomasCruz/users/internal/core/entities"
	"github.com/TomasCruz/users/internal/handlers/grpchandler/users"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (g *GRPCHandler) GetUserByID(ctx context.Context, req *users.UserIDReqMsg) (*users.UserRespMsg, error) {
	g.logger.Info(nil, "GRPCHandler.GetUserByID")

	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	user, err := g.svc.GetUserByID(uid)
	if err != nil {
		g.logger.Error(err, "GRPCHandler.GetUserByID")
		return nil, err
	}

	return &users.UserRespMsg{
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

func (g *GRPCHandler) ListUser(ctx context.Context, req *users.ListUserReqMsg) (*users.ListUserRespMsg, error) {
	g.logger.Info(nil, "GRPCHandler.ListUser")

	countries := req.GetCountry()
	pageSize := req.GetPageSize()
	pageNumber := req.GetPageNumber()

	filter := map[string]map[string]struct{}{}
	filter["country"] = extractParticularFilter(countries)

	if pageSize != 0 {
		filter["page-size"] = map[string]struct{}{strconv.FormatInt(pageSize, 10): {}}
	}

	if pageNumber != 0 {
		filter["page-number"] = map[string]struct{}{strconv.FormatInt(pageNumber, 10): {}}
	}

	userFilter := entities.ExtractUserFilter(filter)
	ps, pn := entities.ExtractPagination(filter, nil, nil)

	userList, totalCount, err := g.svc.ListUser(userFilter, ps, pn)
	if err != nil {
		g.logger.Error(err, "HTTPHandler.ListUserHandler")
		return nil, err
	}

	list := make([]*users.UserRespMsg, 0, len(userList))
	for _, currUser := range userList {
		list = append(list, &users.UserRespMsg{
			Id:        string(currUser.UserID.String()),
			FirstName: currUser.FirstName,
			LastName:  currUser.LastName,
			PswdHash:  currUser.PswdHash,
			Email:     currUser.Email,
			Country:   currUser.Country,
			CreatedAt: timestamppb.New(currUser.CreatedAt),
			UpdatedAt: timestamppb.New(currUser.UpdatedAt),
		})
	}

	return &users.ListUserRespMsg{Users: list, TotalCount: totalCount, ResultCount: int64(len(userList))}, nil
}
