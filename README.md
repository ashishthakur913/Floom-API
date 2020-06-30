# Floom API

GOLANG | GIN | POSTGRES

## Requirements

You need to have **Go** and **PostgreSQL** installed and configured on your system.

## Setup

Clone the repository and follow the steps below:

## Configuration

The API uses PostgreSQL database and the configuration could be found in **config/database.go**.

Default Credentials:
```
port := "5432"
host := "localhost"
name := "floom"
user := "postgres"
password := "1245"
```

## Running the application

Execute the following commands after Go is installed and database is configured:

1. Build
```
$ go build
```
You should now have a binary installed in the directory with the name 'Floom'.

2. Start the application
```
$ ./Floom
```
The API server should now be running on port :5000. If you wish to change the port you can do so in config/server.go


![API server](/server.png "API server")

