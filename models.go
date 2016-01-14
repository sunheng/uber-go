package uber

import "fmt"

type PriceQuery struct {
	StartLatitude  float64
	StartLongitude float64
	EndLatitude    float64
	EndLongitude   float64
}

func (p *PriceQuery) JoinWithServerToken(serverToken string) string {
	q := "start_latitude=%f&start_longitude=%f&end_latitude=%f&end_longitude=%f"
	return fmt.Sprintf("server_token=%s&%s", serverToken, fmt.Sprintf(q, p.StartLatitude, p.StartLongitude, p.EndLatitude, p.EndLongitude))
}

type PricesResponse struct {
	Prices []Price `json:"prices,omitempty"`
}

type Price struct {
	ProductID       string  `json:"product_id,omitempty"`
	CurrencyCode    string  `json:"currency_code,omitempty"`
	DisplayName     string  `json:"display_name,omitempty"`
	Estimate        string  `json:"estimate,omitempty"`
	LowEstimate     int     `json:"low_estimate,omitempty"`
	HighEstimate    int     `json:"high_estimate,omitempty"`
	SurgeMultiplier float64 `json:"surge_multiplier,omitempty"`
	Duration        int     `json:"duration,omitempty"`
	Distance        float64 `json:"distance,omitempty"`
}

type TimeQuery struct {
	StartLatitude  float64
	StartLongitude float64
}

func (t *TimeQuery) JoinWithServerToken(serverToken string) string {
	q := "start_latitude=%f&start_longitude=%f"
	return fmt.Sprintf("server_token=%s&%s", serverToken, fmt.Sprintf(q, t.StartLatitude, t.StartLongitude))
}

type TimesResponse struct {
	Times []Time `json:"times,omitempty"`
}

type Time struct {
	ProductID   string `json:"product_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Estimate    int    `json:"estimate,omitempty"`
}
