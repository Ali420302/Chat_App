package message

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateMessage(ctx context.Context, message *Message) (*Message, error) {
	var lastInsertId int
	query := "INSERT INTO messages(sender_id, room_id, text, time) VALUES ($1, $2,$3, $4) returning id"
	err := r.db.QueryRowContext(ctx, query, message.SenderID , message.RoomID , message.Text, message.Time).Scan(&lastInsertId)
	if err != nil {
		return &Message{}, err
	}

	message.ID = int64(lastInsertId)
	return message, nil
}

func (r *repository) GetMessage(ctx context.Context, email string) (*Message, error) {
	u := Message{}
	query := "SELECT id, sender_id, room_id, text, timestamp FROM messages WHERE room_id = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.SenderID, &u.RoomID,&u.Text,&u.Time)
	if err != nil {
		return &Message{}, nil
	}

	return &u, nil
}