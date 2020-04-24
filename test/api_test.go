package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/mohsensamiei/link-shortener/api"
)

func Benchmark(b *testing.B) {
	var slug string
	client := http.Client{}
	create := int(b.N * 5 / 100)
	redirect := int(b.N*85/100) + create
	for i := 0; i < b.N; i++ {
		if i <= create {
			body, _ := json.Marshal(api.LinkCreateRequest{
				Url: "http://localhost:8080",
			})
			req, _ := http.NewRequest(http.MethodPost, "http://localhost:3080/links", bytes.NewReader(body))
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", fmt.Sprint("Bearer ", ctx.Value(tokenKey{})))
			res, err := client.Do(req)
			if err != nil {
				b.Error(err)
			} else if res.StatusCode != http.StatusCreated {
				b.Errorf("return status code %v", res.StatusCode)
			}
			bytes, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()

			link := new(api.LinkCreateResponse)
			_ = json.Unmarshal(bytes, link)
			slug = link.Slug
		} else if i <= redirect {
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprint("http://localhost:3080/r/", slug), nil)
			res, err := client.Do(req)
			if err != nil {
				b.Error(err)
			} else if res.StatusCode != http.StatusOK {
				b.Errorf("return status code %v", res.StatusCode)
			}
		} else {
			req, _ := http.NewRequest(http.MethodGet, "http://localhost:4080/visits/url", nil)
			res, err := client.Do(req)
			if err != nil {
				b.Error(err)
			} else if res.StatusCode != http.StatusOK {
				b.Errorf("return status code %v", res.StatusCode)
			}
		}
	}
}
