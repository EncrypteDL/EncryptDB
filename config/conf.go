package config

import (
	"context"
)

//Package config provides utilities for loading configuration from multiple sources that can be used to configure the SDK's API clients, and utilities.

// A Config represents a generic configuration value or set of values.
type Config interface{}

// A loader is used to load external configuration data and returns it as
// a generic Config type.
type loader func(context.Context, configs) (Config, error)

type configs []Config

// AppendFromLoaders iterates over the slice of loaders passed in calling each
// loader function in order.
func (c configs) AppendFromLoaders(ctx context.Context, loader []loader) (Config, error) {
	for _, fn := range loader {
		cfg, err := fn(ctx, c)
		if err != nil {
			return nil, err
		}
		c = append(c, cfg)
	}
	return c, nil
}

// ResolveConfig calls the provide function passing slice of configuration sources.
func (c configs) ResolveConfig(f func(configs []interface{}) error) error {
	var cfgs []interface{}
	for i := range c {
		cfgs = append(cfgs, i)
	}
	return f(cfgs)
}
