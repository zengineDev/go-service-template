package configuration

type Environment string

const (
	Production  Environment = "production"
	Development Environment = "development"
	Testing     Environment = "testing"
	Staging     Environment = "staging"
)

const defaultShutdownTimeoutSeconds = 15
const defaultPostgresPort = 5432
