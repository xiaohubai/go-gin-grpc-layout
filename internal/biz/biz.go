package biz

import "github.com/xiaohubai/go-gin-grpc-layout/internal/data"

type Usecase struct {
	db *data.Data
}

func NewUsecase() *Usecase {
	return &Usecase{
		db: data.NewData(),
	}
}
