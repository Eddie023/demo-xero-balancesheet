import { render, screen, cleanup } from '@testing-library/react'
import BalanceSheetReport from './index'
import { Report, Row} from "./interface"

afterEach(() => {
    cleanup()
})

describe("BalanceSheetReport", () => {
    const mockReport :Report = {
        id: "BalanceSheet",
        name: "Balance Sheet",
        type: "BalanceSheet",
        titles: [
            "Balance Sheet",
            "Demo Org",
            "As at 21 August 2024"
        ],
        date: "21 August 2024",
        rows: [{
                type: "Header",
                title: "",
                cells: [{
                        value: ""
                    },
                    {
                        value: "21 August 2024"
                    },
                    {
                        value: "22 August 2023"
                    }
                ]
            },
            {
                type: "Section",
                title: "Assets",
                rows: []
            },
            {
                type: "Section",
                title: "Bank",
                rows: [{
                        type: "Row",
                        cells: [{
                                value: "My Bank Account",
                            },
                            {
                                value: "126.70",
                            },
                            {
                                value: "99.60",
                            }
                        ]
                    },
                    {
                        type: "Row",
                        cells: [{
                                value: "Sample Business",
                            },
                            {
                                value: "92911.00",
                            },
                            {
                                value: "92911.00",
                            }
                        ]
                    },
                    {
                        type: "SummaryRow",
                        cells: [{
                                value: "Total Bank"
                            },
                            {
                                value: "104076.70"
                            },
                            {
                                value: "104049.60"
                            }
                        ]
                    }
                ]
            }]
    }

    it('should render the report titles correctly', () => {
        render(<BalanceSheetReport report={mockReport}/>)
        
        const balanceSheetReportElm = screen.getAllByTestId("report-title")[1]
        expect(balanceSheetReportElm).toBeInTheDocument();
        expect(balanceSheetReportElm).toHaveTextContent("Demo Org")
    })

    it('should render the tables header correctly', () => {
        render(<BalanceSheetReport report={mockReport}/>)
        
        const headerRowElm = screen.getByTestId("header-row")
        expect(headerRowElm).toBeInTheDocument();
        expect(headerRowElm).toHaveTextContent("21 August 2024")
    })

    it('should render rows correctly', () => {
        render(<BalanceSheetReport report={mockReport}/>)
        
        const rowSectionElms = screen.getAllByTestId("row-section")
        expect(rowSectionElms.length).toBe(4)
    })

    it('should render an empty rows correctly', () => {
        const emptyReport :Report =  {
            id: "",
            name: "",
            type: "",
            date: "",
            titles: ['Empty Report'],
          };
          
          render(<BalanceSheetReport report={emptyReport} />);
          
          expect(screen.getByText('Empty Report')).toBeInTheDocument();
          // No table should be rendered
          expect(screen.queryByRole('table')).not.toBeInTheDocument();
    })
})

