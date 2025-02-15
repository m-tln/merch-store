package config

type Config struct {

}

func NewConfig() (*Config, error) {
	return &Config{}, nil
}

func (cfg *Config) GetUserDSN() (string, error) {
	return "", nil
}

func (cfg *Config) GetPurchaseDSN() (string, error) {
	return "", nil
}

func (cfg *Config) GetTransactionDSN() (string, error) {
	return "", nil
}

func (cfg *Config) GetGoodsDSN() (string, error) {
	return "", nil
}

func (cfg *Config) GetPort() (string) {
	return ":8080"
}