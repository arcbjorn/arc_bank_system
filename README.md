## Arc bank system

### Overview

There are 2 containers to run: `api` and `postgres`.

### Development

```shell
# run api and postgres containers
make up

# run all migrations
make migrate_up

# regenerate ORM code (just in case)
make generate_orm

# shut down & delete api and postgres containers
make down
```

Refer to `Makefile` for other commands.
