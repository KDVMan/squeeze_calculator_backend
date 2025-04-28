package models_config

type BinanceModel struct {
	FuturesLimit      int     `yaml:"futures_limit" env-required:"true"`
	FuturesCommission float64 `yaml:"futures_commission" env-required:"true"`
}
