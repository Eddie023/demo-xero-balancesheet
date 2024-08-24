package xeroreports

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type XeroClient struct {
	baseURL string
}

func NewXeroClient(baseURL string) *XeroClient {
	return &XeroClient{baseURL: baseURL}
}

func (x *XeroClient) GetBalanceSheet(ctx context.Context, query map[string]string) (Report, error) {
	u, err := url.Parse(fmt.Sprintf("%s/api.xro/2.0/Reports/BalanceSheet", x.baseURL))
	if err != nil {
		return Report{}, err
	}

	q := url.Values{}
	for k, v := range query {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return Report{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return Report{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Report{}, err
	}

	type GetBalanceSheetResponse struct {
		Reports []Report `json:"Reports"`
	}

	var response GetBalanceSheetResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Report{}, err
	}

	if len(response.Reports) == 0 {
		return Report{}, errors.New("empty reports")
	}

	return response.Reports[0], nil
}
