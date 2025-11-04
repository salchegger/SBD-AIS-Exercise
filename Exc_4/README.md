# Software Architecture for Big Data - Exercise 4

We are going to add the PostgreSQL and orderservice containers
from the previous exercise into a docker compose file.

## Todo
- [ ] Create a docker compose file
    - Share a network between DB and our application
    - Expose the correct ports (:3000 and :5432 respectively)
    - Add environment variables
        - Set the DB_HOST environment variable to your Postgres instance name
    - Add Dockerfile to your `ordersystem` service entry
    - Specify `command` to start your binary on container run

Environment variables:

```env
POSTGRES_DB=order
POSTGRES_USER=docker
POSTGRES_PASSWORD=docker
POSTGRES_TCP_PORT=5432
DB_HOST=<your-postgres-container-name>
```

## Tips and Tricks

To find out more about docker-compose files, have a look over at
the [Docker Compose documentation](https://docs.docker.com/compose/)
and [Docker Compose Getting started](https://docs.docker.com/compose/gettingstarted/).

To manage local Docker containers in a convenient way,
you can use [lazydocker](https://github.com/jesseduffield/lazydocker).

Docker commands are documented here: [Docker](https://docs.docker.com/reference/cli/docker/) ([Cheatsheet](https://docs.docker.com/get-started/docker_cheatsheet.pdf)), [Docker Compose](https://docs.docker.com/reference/cli/docker/compose/)

