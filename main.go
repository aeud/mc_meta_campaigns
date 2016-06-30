package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	mcEndpoint   string = "https://us11.api.mailchimp.com"
	mcVersion    string = "3.0"
	count        int32  = 100
	bucketName   string = "lx-ga"
	objectPrefix string = "mc/campaigns"
	projectId    string = "luxola.com:luxola-analytics"
	datasetId    string = "mailchimp"
	tableId      string = "meta_campaigns"
)

var (
	test           bool        = true
	date           string      = "2016-06-20"
	stream         chan *MCJob = make(chan *MCJob)
	f              time.Time
	t              time.Time
	googleCredPath string
	mcToken        string
)

func processDay(s *MCService, d0 time.Time) {
	campaigns := s.GetCampaigns(d0)
	wg := new(sync.WaitGroup)
	for i := 0; i < len(campaigns); i++ {
		wg.Add(1)
		campaign := campaigns[i]
		s.GetMetaCampaign(campaign, func(campaign *MCCampaign) {
			responseToGS(campaign, d0)
			wg.Done()
		})
	}
	wg.Wait()
}

func init() {
	flagFrom := flag.String("f", time.Now().Format("2006-01-02"), "From date")
	flagTo := flag.String("t", time.Now().Format("2006-01-02"), "To date")
	flagGoogleCredPath := flag.String("g", "", "Google JWT")
	flagMCToken := flag.String("m", "", "Mailchimp token")
	flagDelta := flag.Int("d", 0, "Delta")

	var err error

	flag.Parse()
	f, err = time.Parse("2006-01-02", *flagFrom)
	catchError(err)
	f = f.AddDate(0, 0, -1**flagDelta)
	t, err = time.Parse("2006-01-02", *flagTo)
	catchError(err)

	googleCredPath = *flagGoogleCredPath
	mcToken = *flagMCToken
}

func main() {
	service := NewMCService(new(http.Client))
	for i := 0; i < 10; i++ {
		go mcWorker(service, &stream)
	}

	for t.Sub(f).Seconds() >= 0 {
		d0 := f
		fmt.Println(d0.Format("2006-01-02"))
		processDay(service, d0)
		f = f.AddDate(0, 0, 1)
	}
	writeBQTable()
}
