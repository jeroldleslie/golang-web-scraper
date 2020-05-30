package model

import (
	"fmt"
	"time"
)

type CompanyInfo struct {
	tableName struct{} `sql:"company_info"`
	CvrId string `sql:",pk"`
	CreatedAt time.Time `sql:"default:now()"`
	Address string
	PostalCodeAndCity string
	StartDate string
	BusinessType string
	AdvertisingProtection string
	Status string
	Telephone string
	Fax string
	Email string
	Municipality string
	ActivityCode string
	SecondaryNames string
	FinancialYear string
	LatestArticlesOfAssociation string
	ClassesOfShares string
	RegisteredCapital string
	FirstAccountingPeriod string
}

func (ci CompanyInfo) String() string {
	return fmt.Sprintf("CompanyInfo<%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s>", 
	ci.CvrId, 
	ci.Address, 
	ci.PostalCodeAndCity,
	ci.StartDate,
	ci.BusinessType,
	ci.AdvertisingProtection,
	ci.Status,
	ci.Telephone,
	ci.Email,
	ci.Fax,
	ci.Municipality,
	ci.ActivityCode,
	ci.SecondaryNames,
	ci.FinancialYear,
	ci.LatestArticlesOfAssociation,
	ci.ClassesOfShares,
	ci.RegisteredCapital,
	ci.FirstAccountingPeriod,)
}