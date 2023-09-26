package api

import (
	"apigateway/pkg/storage"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

// results from microservices
type res struct {
	err     error
	post    storage.Post
	comment []storage.Comment
}

// отправляет запрос в микросервис новостей для получения новости
func postReq(wg *sync.WaitGroup, ctx context.Context, id int, result chan<- res) {
	defer wg.Done()
	url := fmt.Sprintf("http://localhost:2000/news/%d", id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		result <- res{err: err}
		return
	}
	requestId := ctx.Value(requestIdHeader)
	req.Header.Add(requestIdHeader, fmt.Sprint(requestId))
	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		result <- res{err: err}
		return
	}

	dec := json.NewDecoder(resp.Body)
	var post storage.Post
	err = dec.Decode(&post)
	if err != nil {
		result <- res{err: err}
		return
	}

	// log.Printf("postReq: %v\n", post)
	result <- res{post: post}
}

// отправляет запрос в микросервис комментариев для получения комментариев
func commentsReq(wg *sync.WaitGroup, ctx context.Context, id int, result chan<- res) {
	defer wg.Done()
	url := fmt.Sprintf("http://localhost:2010/comments/%d", id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		result <- res{err: err}
		return
	}
	requestId := ctx.Value(requestIdHeader)
	req.Header.Add(requestIdHeader, fmt.Sprint(requestId))
	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		result <- res{err: err}
		return
	}
	dec := json.NewDecoder(resp.Body)
	var comments []storage.Comment
	err = dec.Decode(&comments)
	if err != nil {
		result <- res{err: err}
		return
	}

	// log.Printf("commentsRequest: %v \n", comments)
	result <- res{comment: comments}
}

// возвращает настроеного клиента для запроса
func newClient() *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
	return client
}
