package models_config

type ApiModel struct {
	Url string `yaml:"url" env-required:"true"`
}
