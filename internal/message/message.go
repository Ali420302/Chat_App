package message

import "context"

type Message struct {
	SenderID       int  `json:"sender_id"`
	RoomID       int  `json:"room_id"`
	ID       int64  `json:"id"`
	Text string `json:"text"`
	Time    string `json:"time"`
}

type CreateMessageReq struct {
	SenderID       int  `json:"sender_id"`
	RoomID       int  `json:"room_id"`
	Text string `json:"text"`
	Time    string `json:"time"`
}

type CreateMessageRes struct {
	SenderID       int  `json:"sender_id"`
	RoomID       int  `json:"room_id"`
	ID       int64  `json:"id"`
	Text string `json:"text"`
	Time    string `json:"time"`
}



type Repository interface {
	CreateMessage(ctx context.Context, message *Message) (*Message, error)
	GetMessage(ctx context.Context, room_id string) (*Message, error)
	
}

type Service interface {
	CreateMessage(c context.Context, req *CreateMessageReq) (*CreateMessageRes, error)
	
}
