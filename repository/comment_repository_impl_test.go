package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	golangdatabase "github.com/nanandt/golang-database"
	"github.com/nanandt/golang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnecttion())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "test@repository.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnecttion())
	comment, err := commentRepository.FindById(context.Background(), 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnecttion())
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
