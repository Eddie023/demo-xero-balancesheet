import { render, screen } from '@testing-library/react';
import App from './App';

test('renders report-name', () => {
  render(<App />);
  
  const reportName = screen.getByTestId('report-name');
  expect(reportName).toBeInTheDocument();
});
