package repository

import (
	"context"

	audit "github.com/biletskyi-mykyta-go-projects/crud-logger/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Audit struct {
	db *mongo.Database
}

func NewAudit(db *mongo.Database) *Audit {
	return &Audit{
		db: db,
	}
}

func (r *Audit) Insert(ctx context.Context, item audit.LogItem) error {
	_, err := r.db.Collection("logs").InsertOne(ctx, item)

	return err
}
