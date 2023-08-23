# Personal Website Backend

Backend service which allows to website readers send messages to the website owner.

# Development tools

1. Go 1.20.6
1. Docker Server and Client v20.10.24
1. [DbdiagramIO](https://dbdiagram.io/d): online database driagram tool, load the file _docs/db/pwebsite-db.dbdiagram.io.dbml_ on (https://dbdiagram.io/d) in order to generate the entity relationship diagram.
1. [Golang-Migrate](https://github.com/golang-migrate/migrate): Handles db migrations, see the target _migrateup_ and _migratedown_ on the _Makefile_ for more reference.

# Configurations

Next environment variables must be defined to run the service:

- DB_DRIVER
- DB_USER
- DB_PASS
- DB_HOST
- DB_PORT
- DB_NAME
- DB_SSLMODE

Sample:

```
export DB_DRIVER=postgres
export DB_USER=postgres
export DB_PASS=password
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=db
export DB_SSLMODE=disable
```
