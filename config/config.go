package config

// Config config stcut
type config struct {
	server   ServerConfigInterface
	database DatabaseConfigInterface
}

// Interface config interface
type Config interface {
	Server() ServerConfigInterface
	Database() DatabaseConfigInterface

}

// Initialize initialize config
func Initialize() Config {
	return &config{
		server:   NewServerConfig(),
		database: NewDatabaseConfig(),
	}
}

// Server get server config
func (config *config) Server() ServerConfigInterface {
	return config.server
}

// Database get database config
func (config *config) Database() DatabaseConfigInterface {
	return config.database
}