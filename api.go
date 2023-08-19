package main

import (
	"encoding/json"
	"fmt"
	"io"
	"lanc/dto"
	"net/http"
)

const apiBaseUrl = "https://lobste.rs"

func getHottest(pageNum int) ([]dto.ShortPost, error) {
	endpoint := "/hottest.json"
	if pageNum > 1 {
		endpoint = fmt.Sprintf("/page/%d.json", pageNum)
	}
	return getUnified[[]dto.ShortPost](endpoint)
}

func getNewest(pageNum int) ([]dto.ShortPost, error) {
	if pageNum < 1 {
		pageNum = 1
	}
	endpoint := fmt.Sprintf("/newest/page/%d.json", pageNum)
	return getUnified[[]dto.ShortPost](endpoint)
}

func getPost(shortId string) (dto.Post, error) {
	endpoint := "/s/" + shortId + ".json"
	return getUnified[dto.Post](endpoint)
}

func getUnified[T any](endpoint string) (T, error) {
	var post T // -> post := zero
	res, err := get(endpoint)
	if err != nil {
		return post, err
	}
	dec := json.NewDecoder(res)
	if err = dec.Decode(&post); err != nil {
		return post, err
	}
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
