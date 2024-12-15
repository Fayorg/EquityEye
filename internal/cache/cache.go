package cache

import (
	"EquityEye/types"
	"time"
)

type Cache interface {
	RegisterProvider(provider types.ProviderConfiguration) error // This is used to make sure that the provider is in the db with initial values
	IncreaseUsage(provider types.ProviderConfiguration) error    // This will increase the usage of a provider
	IncreaseUsageBy(provider types.ProviderConfiguration, value int) error
	GetUsage(provider types.ProviderConfiguration) (int, error) // This will get the usage of a provider
	GetProvider(providerName string) (string, error)            // This will get the best available provider taking in account the type
	TemporaryDisableProvider(provider types.ProviderConfiguration, duration time.Duration) error
	IsProviderTemporarilyDisabled(provider types.ProviderConfiguration) (bool, error)
}
