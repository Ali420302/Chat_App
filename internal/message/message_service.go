package message

import (
	"context"
	
	"time"

)

const (
	secretKey = "secret"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateMessage(c context.Context, req *CreateMessageReq) (*CreateMessageRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	

	u := &Message{
		SenderID: int(req.SenderID),
		RoomID:    int(req.RoomID),
		Text:    req.Text,
		Time:    req.Time,

	}

	r, err := s.Repository.CreateMessage(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateMessageRes{
		ID:       int64(r.ID),
		SenderID: r.SenderID,
		RoomID:    r.RoomID,
		Text:    r.Text,
		Time:    r.Time,

	}

	return res, nil
}

	

	
