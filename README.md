## Tiny url shorter

Tiny api for shorting urls.

## Running
Clone from github
```shell
git clone https://io112/tiny-url-shorter
cd tiny-url-shorter
```
Run project
```shell
go build cmd\url-shorter\main.go
go run cmd\url-shorter\main.go
```

## Environment variables
Change this environment variables to modify configuration
| Variable      | Description                     | Default          |
| ------------- |---------------------------------| -----------------|
| PORT          | webserver port                  | 8080             |
| HOST          | webserver hostname without port | http://localhost |
| DB_NAME       | name of sqlite database         |    urldb.db      |
