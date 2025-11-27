FROM joseluisq/static-web-server:latest
WORKDIR /public
COPY ../frontend .
ENV SERVER_PORT=80
ENV SERVER_ROOT="/public"