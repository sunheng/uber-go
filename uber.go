package uber

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// v1/path?parameters
	UBERAPI string = "https://api.uber.com/v1/%s?%s"
	price   string = "estimates/price"
	time    string = "estimates/time"
)

type Options struct {
	ServerToken string
}

type Client struct {
	Options *Options
}

func NewClient(options *Options) *Client {
	return &Client{options}
}

func (c *Client) GetPriceEstimates(query *PriceQuery) (PricesResponse, error) {
	param := query.JoinWithServerToken(c.Options.ServerToken)
	requestUrl := fmt.Sprintf(UBERAPI, price, param)

	resp, err := http.Get(requestUrl)
	if err != nil {
		return PricesResponse{}, err
	}
	defer resp.Body.Close()

	var pricesResponse PricesResponse
	if err := json.NewDecoder(resp.Body).Decode(&pricesResponse); err != nil {
		return PricesResponse{}, err
	}

	return pricesResponse, nil
}

func (c *Client) GetTimeEstimates(query *TimeQuery) (TimesResponse, error) {
	param := query.JoinWithServerToken(c.Options.ServerToken)
	requestUrl := fmt.Sprintf(UBERAPI, time, param)

	resp, err := http.Get(requestUrl)
	if err != nil {
		return TimesResponse{}, err
	}
	defer resp.Body.Close()

	var timesResponse TimesResponse
	if err := json.NewDecoder(resp.Body).Decode(&timesResponse); err != nil {
		return TimesResponse{}, err
	}

	return timesResponse, nil
}
