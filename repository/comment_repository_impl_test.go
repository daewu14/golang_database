package repository

import (
	belajar_database "belajar-database"
	"belajar-database/entity"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestCommentRepositoryImpl_Insert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email: "daewu@mail.com",
		Comments: "Haloo oni adalah komentar",
	}
	ress, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(ress)
}

func TestCommentRepositoryImpl_FindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_database.GetConnection())
	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestCommentRepositoryImpl_FindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_database.GetConnection())
	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}