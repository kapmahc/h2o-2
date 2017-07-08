#!/bin/sh
dsn="postgres://postgres@localhost:5432/h2o_test?sslmode=disable"
src="file://db/postgres/migrations"

migrate -database $dsn -source $src $1
#migrate -database $dsn -source $src down
