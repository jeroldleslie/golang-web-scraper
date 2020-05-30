package ingestor

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jeroldleslie/golang-web-scraper/ingestor/model"
	"github.com/jeroldleslie/golang-web-scraper/ingestor/service"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getValue(str string,key string) string {
	index := strings.Index(str, key)
	if index == 0 {
		rawBytes := []byte(str)
	
		lines := strings.Split(string(rawBytes), "\n")
		
		str = strings.Join(lines[:], " ")

		value := strings.Split(str,key)[1]
		return strings.Trim(value, " ")
	}
	return ""
}


func IngestorFunction(cvrid string, s *service.Service) {
	c := colly.NewCollector(
		colly.AllowedDomains("datacvr.virk.dk"),
	)

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println(err.Error())
	})

	companyInfo := &model.CompanyInfo{}
	c.OnHTML("body", func(e *colly.HTMLElement) {

		e.ForEach("div.dataraekker", func(_ int, elem *colly.HTMLElement) {
			 
			str := elem.ChildText("div")
			

		cvrid := getValue(str,"CVR number")
		if cvrid != "" {
			companyInfo.CvrId = cvrid
		}

		address := getValue(str,"Address")
		if address != "" {
			companyInfo.Address = address
		}

		postal_code_and_city := getValue(str,"Postal code and city")
		if postal_code_and_city != "" {
			companyInfo.PostalCodeAndCity = postal_code_and_city
		}

		startdate := getValue(str,"Start date")
		if startdate != "" {
			companyInfo.StartDate = startdate
		}

		businesstype := getValue(str,"Business type")
		if businesstype != "" {
			companyInfo.BusinessType = businesstype
		}

		advertisingprotection := getValue(str,"Advertising protection")
		if advertisingprotection != "" {
			companyInfo.AdvertisingProtection = advertisingprotection
		}

		status := getValue(str,"Status")
		if status != "" {
			companyInfo.Status = status
		}

		telephone := getValue(str,"Telephone")
		if telephone != "" {
			companyInfo.Telephone = telephone
		}

		fax := getValue(str,"Fax")
		if fax != "" {
			companyInfo.Fax = fax
		}

		email := getValue(str,"Email")
		if email != "" {
			companyInfo.Email = email
		}

		municipality := getValue(str,"Municipality")
		if municipality != "" {
			companyInfo.Municipality = municipality
		}

		activityCode := getValue(str,"Activity code")
		if activityCode != "" {
			companyInfo.ActivityCode = activityCode
		}

		secondaryNames := getValue(str,"Secondary names")
		if secondaryNames != "" {
			companyInfo.SecondaryNames = secondaryNames
		}

		financialYear := getValue(str,"Financial year")
		if financialYear != "" {
			companyInfo.FinancialYear = financialYear
		}

		latestArticlesOfAssociation := getValue(str,"Latest articles of association")
		if latestArticlesOfAssociation != "" {
			companyInfo.LatestArticlesOfAssociation = latestArticlesOfAssociation
		}

		classesOfShares := getValue(str,"Classes of shares")
		if classesOfShares != "" {
			companyInfo.ClassesOfShares = classesOfShares
		}

		registeredCapital := getValue(str,"Registered capital")
		if registeredCapital != "" {
			companyInfo.RegisteredCapital = registeredCapital
		}

		firstAccountingPeriod := getValue(str,"First accounting period")
		if firstAccountingPeriod != "" {
			companyInfo.FirstAccountingPeriod = firstAccountingPeriod
		}
		
		})
		
		s.InsertCompanyInfo(companyInfo)
		
	})
	
	url := "https://datacvr.virk.dk/data/index.php?enhedstype=virksomhed&id=%s&type=virksomhed&sortering=default&language=en-gb&q=visenhed"
	url = fmt.Sprintf(url, cvrid)
	fmt.Println(url)
	c.Visit(url)
}
