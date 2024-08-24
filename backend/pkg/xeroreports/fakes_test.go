package xeroreports_test

import (
	"context"

	"github.com/eddie023/accounting/pkg/xeroreports"
)

type fakeXeroClient struct {
	Response xeroreports.Report
	Err      error
}

func (f *fakeXeroClient) GetBalanceSheet(ctx context.Context, query map[string]string) (xeroreports.Report, error) {
	return f.Response, f.Err
}
