package rss

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
	"testing"
)

type MockClient struct {
	GetFunc func(url string) (*http.Response, error)
}

func (m *MockClient) Get(url string) (*http.Response, error) {
	return DoGetFunc(url)
}

var (
	DoGetFunc func(url string) (*http.Response, error)
)

func TestParse(t *testing.T) {
	Client = &MockClient{}
	testRss := rssFeed{
		Rss: "test RSS",
		Channel: Channel{
			Items: []Item{
				{Title: "test title", Content: "test conternt", Link: "http://test.url", PubTime: "Sat, 29 Jul 2023 10:54:27 GMT"},
			}},
	}
	b, err := xml.Marshal(testRss)
	if err != nil {
		t.Fatal(err)
	}
	r := io.NopCloser(bytes.NewReader([]byte(b)))
	DoGetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       r,
			Status:     "200 OK",
			Proto:      "HTTP/1.1",
		}, nil
	}
	news, err := Parse("test.url")
	if err != nil {
		t.Fatal(err)
	}
	if len(news) == 0 {
		t.Fatal("data not unmarshal")
	}
	t.Logf("received %v news: %v", len(news), news)
}
