# Xero-Accounting 

Simple application to demonstrate fetching reports from Xero API and displaying in the table format in the frontend. 

You can watch the loom demo here: https://www.loom.com/share/592cf3f5d76b4b78a4f188b96e47edf8?sid=6bddf62a-16b0-4eba-af8f-b82f201b3770

## Running the application 
1. Run `make run` from backend directory to start the backend services. 
2. Run `npm run start` from frontend directory to start react app. 

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