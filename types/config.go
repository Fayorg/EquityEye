package types

type Provider struct {
	Name           string `yaml:"name" json:"name"`
	ProviderName   string `yaml:"providerName" json:"providerName"`
	Key            string `yaml:"key" json:"key"`
	Limit          int    `yaml:"limit" json:"limit"`
	LimitTimeframe int    `yaml:"limitTimeframe" json:"limitTimeframe"`
}
