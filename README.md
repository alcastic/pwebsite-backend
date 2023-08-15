# Personal Website Backend

Backend service which allows to website readers send messages to the website owner.

# Development tools

1. Go 1.20.6
1. Docker Server and Client v20.10.24
1. [DbdiagramIO](https://dbdiagram.io/d): online database driagram tool, load the file _docs/db/pwebsite-db.dbdiagram.io.dbml_ on (https://dbdiagram.io/d) in order to generate the entity relationship diagram.
1. [Golang-Migrate](https://github.com/golang-migrate/migrate): Handles db migrations, see the target _migrateup_ and _migratedown_ on the _Makefile_ for more reference.
