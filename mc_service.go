package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type MCService struct {
	Client *http.Client
}

func NewMCService(c *http.Client) *MCService {
	s := new(MCService)
	s.Client = c
	return s
}

func (s *MCService) GetCampaigns(d0 time.Time) []*MCCampaign {
	d1 := d0.AddDate(0, 0, 1)
	v := new(struct {
		Campaigns []*MCCampaign `json:"campaigns"`
	})
	f := "campaigns/?fields=campaigns.id&since_send_time=%v&before_send_time=%v&count=%v"
	path := fmt.Sprintf(f, d0.Format("2006-01-02"), d1.Format("2006-01-02"), count)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	stream <- NewMCJob(path, func(j *MCJob) {
		ResponseToInterface(j.Response, v)
		wg.Done()
	})
	wg.Wait()
	return v.Campaigns
}
func (s *MCService) GetMetaCampaign(campaign *MCCampaign, done func(campaign *MCCampaign)) {
	path := fmt.Sprintf("campaigns/%v", campaign.Id)
	stream <- NewMCJob(path, func(j *MCJob) {
		ResponseToInterface(j.Response, campaign)
		done(campaign)
	})
}

type MCJob struct {
	Path     string
	Response *http.Response
	Done     func(j *MCJob)
}

func NewMCJob(path string, done func(j *MCJob)) *MCJob {
	j := new(MCJob)
	j.Path = path
	j.Done = done
	return j
}

func mcGet(s *MCService, path string) *http.Response {
	// Build the URL
	url := fmt.Sprintf("%v/%v/%v", mcEndpoint, mcVersion, path)
	req, err := http.NewRequest("GET", url, nil)
	catchError(err)

	req.Header.Add("Authorization", fmt.Sprintf("Basic %v", mcToken))

	resp, err := s.Client.Do(req)
	catchError(err)

	// If there are more than 10 simultaneous requests to the API
	if resp.StatusCode == 429 {
		fmt.Println("Wait 10 sec")
		time.Sleep(10 * time.Second)
		return mcGet(s, path)
	}
	return resp
}

func ResponseToInterface(resp *http.Response, i interface{}) {
	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	catchError(decoder.Decode(i))
}
