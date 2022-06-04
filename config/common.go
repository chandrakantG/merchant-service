package config

type ServiceConfig struct {
	ServiceHost string
	ServicePort string
	JwtSecret   string
}

var serviceConfig ServiceConfig

func initServiceConfig() {
	serviceConfig = ServiceConfig{
		ServiceHost: getString("SERVICE_HOST"),
		ServicePort: getString("SERVICE_PORT"),
		JwtSecret:   getString("JWT_SECRET"),
	}
}

func GetServiceConfig() ServiceConfig {
	return serviceConfig
}
