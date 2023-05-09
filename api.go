package main

import (
	"encoding/json"
	"fmt"
	"io"
	"lanc/models"
	"net/http"
)

const apiBaseUrl = "https://lobste.rs"

func getHottest(n int) ([]models.ShortPost, error) {
	endpoint := "/hottest.json"
	if n > 1 {
		endpoint = fmt.Sprintf("/page/%d.json", n)
	}
	return getUnified[[]models.ShortPost](endpoint)
}

func getNewest(n int) ([]models.ShortPost, error) {
	if n < 1 {
		n = 1
	}
	endpoint := fmt.Sprintf("/newest/page/%d.json", n)
	return getUnified[[]models.ShortPost](endpoint)
}

func getPost(shortId string) (models.Post, error) {
	endpoint := "/s/" + shortId + ".json"
	return getUnified[models.Post](endpoint)
}

var k int = 0

func getUnified[T any](endpoint string) (T, error) {
	var post T
	res, err := get(endpoint)
	if err != nil {
		return post, err
	}
	dec := json.NewDecoder(res)
	dec.Decode(&post)
	return post, nil
}

func get(endpoint string) (io.Reader, error) {
	client := http.DefaultClient
	res, err := client.Get(apiBaseUrl + endpoint)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
