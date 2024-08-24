# Xero-Accounting 

Simple application to demonstrate fetching reports from Xero API and displaying in the table format in the frontend. 

## Running the application 
1. Both Backend and Frontend contain their own `docker-compose.yml` files which can be run separately or there is a single `make up` command in the root directory which runs both FE and BE. 

## Running the tests 

### Backend 

1. Run `make test` command inside backend directory

### Frontend 

1. Run `npm run test` command inside frontend directory

# Design Decisions 

## Architecture 
The codebase for Backend is maintained such that it follows the [Clean Code Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)  which emphasizes the separation of concerns,  making codebase more maintainable, testable and scalable. 

## Testing Philosophy 
1. Keeping things simple by not relying heavily on third party tools for generating mocks, stubs. 
2. Creating tests that focus on the public methods of a package by separting with `_test` pkg. This is done to emphasize testing the external behaviour and API of the package rather than internal implementation details. Go projects typically test at package level. 

## Possible enhancements
1. Creating a generic reports component to display multiple type of report in the FE. 
2. Handling optional query params that can be passed to XERO API endpoint. 
3. Separation of JSON serialization from internal platform entity. 