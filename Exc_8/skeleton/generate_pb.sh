#!/bin/sh
# Return on any error
set -e

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/orders.proto
