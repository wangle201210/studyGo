package service

import (
	"github.com/wangle201210/studyGo/tools/bilibili/dal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	dal *dal.Query
}

func New() (s *Service) {
	s = new(Service)
	db, err := gorm.Open(mysql.Open("root:bilibili@(127.0.0.1:32444)/bilibili?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	query := dal.Use(db)
	s.dal = query
	return
}
