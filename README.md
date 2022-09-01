## Tour API
The API provides ability to create/read/delete tour entries

## How to run

1. Make sure you have "Go" installed
2. Create an .env file (on the main.go level)
    - specify DATABASE_URL
2. Run command to build:  `go build main.go`
3. Run: `./main`

### Goal
There is only one reason - is education. There will be a good way to learn Go :D

### Tasks

- [x] Design database tables
- [x] Create the API
- [x] Add db migration tool
- [x] Add a created time as a column for the all tables.
- [ ] Add an API validation before trying to create DB records
- [ ] DON'T have "Alter table" feature (as a column changed I have to drop the table) 
