package server_test

import (
	"context"

	"github.com/eddie023/accounting/pkg/server"
)

type fakeAuthenticator struct{}

func (f *fakeAuthenticator) Authenticate(token string) error {
	return nil
}

type fakeBalanceSheetProvider struct {
	Response server.Report
	Err      error
}

func (f *fakeBalanceSheetProvider) BuildReport(ctx context.Context, query map[string]string) (server.Report, error) {
	return f.Response, f.Err
}
