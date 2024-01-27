package entity

type ApplicationConfig struct {
	Port string
}

type GithubConfig struct {
	BaseUrl      string
	ClientID     string
	ClientSecret string
}

type DatabaseConfig struct {
	Path string
	Name string
}

type Config struct {
	GithubConfig
	ApplicationConfig
	DatabaseConfig
}
