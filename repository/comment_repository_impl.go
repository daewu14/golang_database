package repository

import (
	"belajar-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "insert into comments(email, comments) values(?, ?)"
	ress, err := repo.DB.ExecContext(ctx, query, comment.Email, comment.Comments)
	if err != nil {
		return comment, err
	}
	id, err := ress.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	query := "select id, email, comments from comments where id = ? limit 1"
	ress, err := repo.DB.QueryContext(ctx, query, id)
	defer ress.Close()
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}

	if ress.Next() {
		// Data Ada
		ress.Scan(&comment.Id, &comment.Email, &comment.Comments)
		return comment, nil
	} else {
		// Data Tidak Ada
		return comment, errors.New("Id "+strconv.Itoa(int(id))+" not found!")
	}
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "select id, email, comments from comments"
	ress, err := repo.DB.QueryContext(ctx, query)
	defer ress.Close()
	if err != nil {
		return nil, err
	}
	var comments = []entity.Comment{}
	for ress.Next() {
		comment := entity.Comment{}
		ress.Scan(&comment.Id, &comment.Email, &comment.Comments)
		comments = append(comments, comment)
	}
	return comments, nil
}
