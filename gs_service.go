package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	storage "google.golang.org/api/storage/v1"
	"log"
	"net/http"
	"strings"
	"time"
)

func getGSService(client *http.Client) *storage.Service {
	service, err := storage.New(client)
	catchError(err)
	return service
}

func responseToGS(c *MCCampaign, d0 time.Time) {

	fileName := fmt.Sprintf("%v/%v/%v.json.gz", objectPrefix, d0.Format("2006/01/02"), c.Id)

	object := &storage.Object{Name: fileName}

	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	bs, err := json.Marshal(c)
	catchError(err)
	w.Write(bs)
	w.Close()

	file := strings.NewReader(b.String())

	if res, err := getGSService(getGoogleHttpClient()).Objects.Insert(bucketName, object).Media(file).Do(); err == nil {
		fmt.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
	} else {
		log.Printf("Objects.Insert failed: %v\n", err)
		time.Sleep(10 * time.Second)
		responseToGS(c, d0)
	}
}
