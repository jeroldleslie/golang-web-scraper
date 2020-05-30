package service

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/jeroldleslie/golang-web-scraper/ingestor/model"
)


type Service struct {
	DB *pg.DB
}

func (s Service) InsertCompanyInfo(companyInfo *model.CompanyInfo) {
	err := s.DB.Insert(companyInfo)
    if err != nil {
        fmt.Println(err.Error())
    }
}
