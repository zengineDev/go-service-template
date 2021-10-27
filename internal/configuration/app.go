package configuration

type AppConfig struct {
	Name        string      `json:"name" yaml:"name"`
	Version     string      `json:"version" yaml:"version"`
	Port        int         `json:"port" yaml:"port"`
	Environment Environment `json:"environment" yaml:"environment"`
	Domain      string      `json:"domain" yaml:"domain"`
	Debug       bool        `json:"debug" yaml:"debug"`
}
