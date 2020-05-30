package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	"strconv"


	lib "github.com/hauxe/gom/library"

	"github.com/gocolly/colly"
	"github.com/jeroldleslie/golang-web-scraper/ingestor/ingestor"
	"github.com/jeroldleslie/golang-web-scraper/ingestor/pool"
	"github.com/jeroldleslie/golang-web-scraper/ingestor/service"
	"github.com/jeroldleslie/golang-web-scraper/ingestor/db"
	library "github.com/jeroldleslie/golang-web-scraper/ingestor/lib"

)

type job struct {
	name string
	f    func()
}

// Name get job name
func (j *job) Name() string {
	return j.name
}

// GetContext get http context
func (j *job) GetContext() context.Context {
	return context.Background()
}

// Execute handle job
func (j *job) Execute() error {
	j.f()
	return nil
}

func main() {

	fmt.Println("Wait for a while to bring database up...")
	time.Sleep(time.Millisecond * 20000)


	//numJob := 200
	numWorker,_ := strconv.Atoi(library.GetEnv("NUM_WORKERS","20"))
	worker, err := pool.CreateWorker()
	pageno := 0


	s := &service.Service{DB: db.GetDB()}
	
	
	err = worker.StartServer(worker.SetMaxWorkersOption(numWorker))
	var wg sync.WaitGroup
	if (err == nil) {
		// TODO handle err
	}

	c := colly.NewCollector(
		colly.AllowedDomains("datacvr.virk.dk"),
	)

	c.OnHTML("div.cvr", func(e *colly.HTMLElement) {
		str := e.ChildText("p")
		cvrid := str[3:len(str)]
		f := func() {
			ingestor.IngestorFunction(cvrid, s)
			wg.Done()
		}

		j := job{
			name: "job: " + lib.ToString(pageno),
			f:    f,
		}
		wg.Add(1)
		err := worker.QueueJob(&j, 5000*time.Millisecond)

		if err != nil {
			// TODO something
		}

	})
	quit := true
	c.OnError(func(_ *colly.Response, err error) {
		quit = false
	})

	
	for {
		if quit {
			url := "https://datacvr.virk.dk/data/visninger?page=%d&soeg=&oprettet=null&ophoert=null&branche=&language=en-gb&type=virksomhed&sortering=default"
			url = fmt.Sprintf(url, pageno)
			fmt.Println(url)
			c.Visit(url)
		} else {
			break
		}
		pageno++
	}

	wg.Wait()
	worker.StopServer()
	fmt.Println("Ingestor Finished!")
}
