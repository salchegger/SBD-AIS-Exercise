FROM golang:1.25 AS builder
WORKDIR /app
COPY .. .
RUN sh /app/scripts/build-application.sh

FROM alpine AS run
WORKDIR /
COPY --from=builder /app/ordersystem /app/ordersystem
# EXPOSE doesn't actually do anything!
EXPOSE 3000
CMD ["/app/ordersystem"]