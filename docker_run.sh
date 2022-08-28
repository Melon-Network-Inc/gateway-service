#!/bin/sh

docker run --name database -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=melon_service -d postgres:13.8
docker run --name cahce -p 6379:6379 -d redis:7.0.4
make docker-push
make docker-run