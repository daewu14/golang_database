package repository

import (
	"belajar-database/entity"
	"context"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, Id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}