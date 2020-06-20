## Download go modules
```shell script
go mod download
```

## Export OS Environment
This project use `ElephantSQL` to be a database. It needs a URL to connect the database. We pass the URL value via the OS environment with below command
```shell script
export DATABASE_URL=postgres://txapurbo:EaW4RkUG9Oxw1kwhmIGyJTp1SALtleaP@john.db.elephantsql.com:5432/txapurbo
```

## Start a server
```shell script
go run main.go
```
