# Floom API

GOLANG | GIN | POSTGRES

## Installation

You need to have Go installed on your system. For details visit [this link](https://golang.org/doc/install) and install the distribution for your system. You can also follow [this](https://www.linode.com/docs/development/go/install-go-on-ubuntu/) short article for a quick setup.

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

2. Start the application
```
$ ./Floom
```
The API server should now be running on port :5000. If you wish to change the port you can do so in config/server.go


![API server](/server.png "API server")

