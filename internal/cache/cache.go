package cache

import "EquityEye/types"

type Cache interface {
	RegisterProvider(provider types.Provider) error     // This is used to make sure that the provider is in the db with initial values
	IncreaseUsage(provider types.Provider) (int, error) // This will increase the usage of a provider
	GetUsage(provider types.Provider) (int, error)      // This will get the usage of a provider
	GetProvider(providerName []string) types.Provider   // This will get the best available provider taking in account the type
}
