package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	PhotoApi = "https://api.pexels.com/v1"
)

type Client struct {
	Token          string
	hc             http.Client
	RemainingTimes int32
}

func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

type SearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Photos       []Photo `json:"photos"`
}

type Photo struct {
	Id              int32       `json:"id"`
	Width           int32       `json:"width"`
	Height          int32       `json:"height"`
	Url             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerUrl string      `json:"photographer_url"`
	Src             PhotoSource `json:"src"`
}
type PhotoSource struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Large2x   string `json:"large2x"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Potrait   string `json:"potrait"`
	Square    string `json:"square"`
	Landscape string `json:"landscpae"`
	Tiny      string `json:"tiny"`
}
type CuratedResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_Results"`
	NextPage     string  `json:"next_page"`
	Photos       []Photo `json:"photos"`
}

func main() {
	os.Setenv("PexelsToken", "fHXvqxEVlC2jCXW3jOkBTKV7AdTxVk4XZmO51iygnxQVK6awwbyP2UtF")
	var TOKEN = os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)

	// result, err := c.SearchPhotos("wawes", 15, 1)
	result, err := c.getRandomPhoto()

	if err != nil {
		fmt.Errorf("Search error: %v", err)
	}

	// if result.Page == 0 {
	// 	fmt.Errorf("Search result wrong")
	// }
	fmt.Println(result)
}
func (c *Client) SearchPhotos(query string, perPage, page int) (*SearchResult, error) {
	url := fmt.Sprintf(PhotoApi+"/search?query=%s&per_page%d&page=%d", query, perPage, page)
	res, err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result SearchResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) CuratedPhotos(perPage int, page int) (*CuratedResult, error) {
	url := fmt.Sprintf(PhotoApi+"/curated?per_page%d&page=%d", perPage, page)
	res, err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result CuratedResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (c *Client) requestDoWithAuth(method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.Token)
	res, err := c.hc.Do(req)
	if err != nil {
		return res, err
	}
	times, err := strconv.Atoi(res.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return res, err
	} else {
		c.RemainingTimes = int32(times)
	}
	return res, nil
}
func (c *Client) GetPhoto(id int32) (*Photo, error) {
	url := fmt.Sprintf(PhotoApi+"/photos/%d", id)
	res, err := c.requestDoWithAuth("GET", url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result Photo
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (c *Client) getRandomPhoto() (*Photo, error) {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := randomGenerator.Intn(1001)
	result, err := c.CuratedPhotos(1, randomNumber)
	if err == nil && len(result.Photos) == 1 {
		return &result.Photos[0], nil
	}
	fmt.Println(result.Photos)
	return nil, err
}
