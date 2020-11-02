#!/bin/bash


docker run --rm \
	--name pbwatch-db \
	-e POSTGRES_USER=polpettone \
	-e POSTGRES_PASSWORD=polpettone \
	-e POSTGRES_DB=pbwatch \
	-p 5432:5432 \
 	-v $HOME/.pbwatch/db:/var/lib/postgresql/data \
	-d \
	postgres:10
