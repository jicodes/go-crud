# A CRUD API project in Go/Gin with Postgres and Gorm

## Installation packages

```sh
# Install Gin and Gorm
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

# Install godotenv for environment variables
go get github.com/joho/godotenv
```

## Run the project
```sh
# since the main and routes are in different files, you need to run both files like this
go run main.go routes.go

# or you can run the project using the following command, to run all the Go files in the current directory:
go run .

# run the migration to create the table in the database
go run migrate/migrate.go
```