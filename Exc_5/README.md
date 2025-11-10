# Software Architecture for Big Data - Exercise 5
Today we are going to add two new microservices: A reverse proxy and a web server.

Typically the frontend (i.e. HTML, CSS, JS, static content) is served
by [nginx](https://nginx.org/), [Apache HTTP Server](https://httpd.apache.org/), or a similar. In this exercise we are going
to serve the frontend for the order service with [sws](https://github.com/static-web-server/static-web-server)
and add [traefik](https://doc.traefik.io/traefik/) as application proxy / load balancer.

## Todo 
- [ ] Add [sws](https://github.com/static-web-server/static-web-server) to docker compose and serve the `./frontend` folder
- [ ] Add [traefik](https://doc.traefik.io/traefik/) reverse proxy
  - sws should be reachable at http://localhost
  - The orderservice should be reachable at http://orders.localhost

## Tips and Tricks
The documentation for traefik provides a [quickstart example](https://doc.traefik.io/traefik/expose/docker/) for Docker compose.
This can be adapted to serve the orderservice and sws. Traefik must not use any other network than w

The Docker socket mapping on Windows might look different than on Linux and OSX. 
If your other Docker containers cannot be found, have a look [here](https://stackoverflow.com/questions/57466568/how-do-you-mount-the-docker-socket-on-windows/62176649#62176649),
[here](https://github.com/docker/for-win/issues/4642#issuecomment-567811455) or [here](https://community.traefik.io/t/how-to-run-on-windows-host-with-docker-provider/4834/4).

The orderservice uses the port 3000, keep in mind when setting up its traefik labels.

The static web server can be [configured via environment variables](https://static-web-server.net/configuration/environment-variables/),
to expose port 80 and serve the frontend folder.

