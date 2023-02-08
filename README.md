# Sales project

This project is a particular implementation of reading a data file in csv format and inserting its records.

## Requirements
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Golang](https://go.dev/doc/install) 
- [Migrate](https://github.com/golang-migrate/migrate)
- [Make](https://community.chocolatey.org/packages/make) only windows users this guide

## How to use it

**Installing and provisioning the db**

In the first instance it is recommended to use docker to raise the service and the necessary dependencies

The project requires a postgresql database to insert the sales records to be processed, so follow the recommendations and install make and run the command
```
make postgres
```

Intall migrate for created new database and run the migrations with the comand

for created new database this comand
```
make createdb
```

for run migrations
```
make migrateup
```

**NOTE**
if you already have installed a local postgresql database higher than 12.x ignore these steps,
and use the sql found in migrations/db to create the table in a database named store

**Configuring file**

in the repository there is a directory called files, inside it you must place the csv file you want to insert

## Export environment variables.
```
export $(cat .env.example | grep -v ^# | xargs)
```
there are several variables that can be configured such as the connection to the db or even the csv file, for this you must set the environment variables of the .env.example file

**NOTE**
only run local, for the docker using -e NAME_ENVIROMENT=$VALUE

## Run with golang instaled

stand at the root of the project
and run

```
go run ./cmd/main.go
```

## Using Docker
to build image
```
docker build -t sales-project .
```

to run container
```
docker run --name sales sales-project
```

to run container if the db postgresql is containirized
```
docker run -e DB_HOST=host.docker.internal  --name sales sales-project
```