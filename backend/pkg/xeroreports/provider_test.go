package xeroreports_test

import (
	"context"
	"errors"
	"testing"

	"github.com/eddie023/accounting/pkg/server"
	"github.com/eddie023/accounting/pkg/xeroreports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildReport(t *testing.T) {
	tests := []struct {
		name               string
		wantErr            bool
		errMsg             string
		getBalanceSheetErr error
		report             xeroreports.Report
		want               server.Report
	}{
		{
			name: "xero report with rows",
			report: xeroreports.Report{
				ReportID:     "BalanceSheet",
				ReportName:   "BalanceSheet",
				ReportType:   "BalanceSheet",
				ReportTitles: &[]string{"Demo Org"},
				ReportDate:   "21 August 2024",
				Rows: &[]xeroreports.Row{
					{
						RowType: "Section",
						Title:   "Bank",
						Rows: &[]xeroreports.Row{
							{
								RowType: "Row",
								Cells: &[]xeroreports.Cell{
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
						RowType: "Header",
						Cells: &[]xeroreports.Cell{
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
			want: server.Report{
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
		},
		{
			name: "xero report without rows",
			report: xeroreports.Report{
				ReportID:     "BalanceSheet",
				ReportName:   "BalanceSheet",
				ReportType:   "BalanceSheet",
				ReportTitles: &[]string{"Demo Org"},
				ReportDate:   "21 August 2024",
			},
			want: server.Report{
				ID:     "BalanceSheet",
				Name:   "BalanceSheet",
				Type:   "BalanceSheet",
				Titles: []string{"Demo Org"},
				Date:   "21 August 2024",
			},
		},
		{
			name:               "xero client returns error",
			getBalanceSheetErr: errors.New("some error"),
			wantErr:            true,
			errMsg:             "getting balance sheet report from client: some error",
		},
		{
			name: "xero client report failed validation",
			report: xeroreports.Report{
				ReportType: "Profit And Loss",
			},
			wantErr: true,
			errMsg:  "failed validation: report must be of type 'BalanceSheet'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeXeroClient := &fakeXeroClient{
				Response: tt.report,
				Err:      tt.getBalanceSheetErr,
			}

			xeroBalanceSheetReporter := xeroreports.NewXeroReportsProvider(fakeXeroClient)
			got, err := xeroBalanceSheetReporter.BuildReport(context.Background(), nil)
			if tt.wantErr {
				assert.EqualError(t, err, tt.errMsg)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
