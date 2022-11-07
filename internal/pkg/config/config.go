package config

type HarborConfig interface {
	GetActiveContext() Context
}
