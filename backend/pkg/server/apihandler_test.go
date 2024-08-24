package server_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/eddie023/accounting/pkg/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type handlerFixture struct {
	path   string
	method string

	query []string

	gotStatusCode int
	gotBody       string

	fakeGetBalanceSheetResponse server.Report
	fakeGetBalanceSheetErr      error

	fakeAuthenticator        server.TokenAuthorizer
	fakeBalanceSheetProvider server.BalanceSheetReportProvider
}

func newHandlerFixture(method string, query []string) *handlerFixture {
	return &handlerFixture{
		method: method,
		query:  query,
	}
}

func (f *handlerFixture) execute(t *testing.T) {
	f.fakeAuthenticator = &fakeAuthenticator{}

	f.fakeBalanceSheetProvider = &fakeBalanceSheetProvider{
		Response: f.fakeGetBalanceSheetResponse,
		Err:      f.fakeGetBalanceSheetErr,
	}

	handler := server.NewHandler(f.fakeAuthenticator, f.fakeBalanceSheetProvider)

	u, err := url.Parse(f.path)
	require.NoError(t, err)
	// make sure query param has key,value pair
	require.Equal(t, 0, len(f.query)%2)

	q := url.Values{}
	for i := 0; i < len(f.query); i += 2 {
		q.Add(f.query[i], f.query[i+1])
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(context.Background(), f.method, u.String(), nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	f.gotStatusCode = res.StatusCode
	bodyBuf, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	f.gotBody = string(bodyBuf)
}

func TestGetEquityPositionHandler(t *testing.T) {
	tests := []struct {
		name                         string
		httpMethod                   string
		queryParam                   []string
		balanceSheetProviderErr      error
		balanceSheetProviderResponse server.Report
		wantStatusCode               int
		wantBody                     string
	}{
		{
			name:       "valid request returning rows",
			httpMethod: http.MethodGet,
			balanceSheetProviderResponse: server.Report{
				ID:     "BalanceSheet",
				Name:   "BalanceSheet",
				Type:   "BalanceSheet",
				Titles: []string{"Demo Org"},
				Date:   "21 August 2024",
			},
			wantStatusCode: http.StatusOK,
			wantBody:       `{"id":"BalanceSheet","name":"BalanceSheet","type":"BalanceSheet","titles":["Demo Org"],"date":"21 August 2024"}`,
		},
		{
			name:       "valid request returning without rows",
			httpMethod: http.MethodGet,
			balanceSheetProviderResponse: server.Report{
				ID:     "BalanceSheet",
				Name:   "BalanceSheet",
				Type:   "BalanceSheet",
				Titles: []string{"Demo Org"},
				Date:   "21 August 2024",
				Rows: &[]server.Row{
					{
						Type:  "Section",
						Title: "Bank",
						Rows: &[]server.Row{
							{
								Type: "Row",
								Cells: &[]server.Cell{
									{
										Value: "My Bank Account",
									},
									{
										Value: "100.00",
									},
								},
							},
						},
					},
					{
						Type: "Header",
						Cells: &[]server.Cell{
							{
								Value: "21 Aug 2024",
							},
							{
								Value: "22 August 2023",
							},
						},
					},
				},
			},
			wantStatusCode: http.StatusOK,
			wantBody: `{
				"id": "BalanceSheet",
				"name": "BalanceSheet",
				"type": "BalanceSheet",
				"titles": ["Demo Org"],
				"date": "21 August 2024",
				"rows": [
					{
						"type": "Section",
						"title": "Bank",
						"rows": [
							{
								"type": "Row",
								"title": "",
								"cells": [
									{
										"value": "My Bank Account"
									},
									{
										"value": "100.00"
									}
								]
							}
						]
					},
					{
						"type": "Header",
						"title": "",
						"cells": [
							{
								"value": "21 Aug 2024"
							},
							{
								"value": "22 August 2023"
							}
						]
					}
				]
				}
			`,
		},
		{
			name:                    "should not expose any internal information for internal error",
			httpMethod:              http.MethodGet,
			wantStatusCode:          http.StatusInternalServerError,
			balanceSheetProviderErr: fmt.Errorf("failed to get report"),
			wantBody:                `{"error":"Internal Server Error"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := newHandlerFixture(tc.httpMethod, tc.queryParam)
			f.fakeGetBalanceSheetResponse = tc.balanceSheetProviderResponse
			f.fakeGetBalanceSheetErr = tc.balanceSheetProviderErr

			f.path = "/api/reports/balance-sheet"

			f.execute(t)

			assert.Equal(t, tc.wantStatusCode, f.gotStatusCode)
			assert.JSONEq(t, tc.wantBody, f.gotBody)
		})
	}
}
