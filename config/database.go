package config

// DatabaseConfigInterface database config interface
type DatabaseConfigInterface interface {
	Host() string
	Port() string
	Name() string
	User() string
	Password() string
}

// Database database config struct
type Database struct {
	port     string
	host     string
	name     string
	user     string
	password string
}

// NewDatabaseConfig create database instance
func NewDatabaseConfig() *Database {
	port := "5432"
	host := "localhost" //Test DB
	name := "floom"
	user := "postgres"
	password := "1245"

	database := &Database{
		port:     port,
		host:     host,
		name:     name,
		user:     user,
		password: password,
	}
	return database
}

// Host get database host
func (database *Database) Host() string {
	return database.host
}

// Port get database port number
func (database *Database) Port() string {
	return database.port
}

// Name get database name
func (database *Database) Name() string {
	return database.name
}

// User get databsae user name
func (database *Database) User() string {
	return database.user
}

// Password get database user password
func (database *Database) Password() string {
	return database.password
}
