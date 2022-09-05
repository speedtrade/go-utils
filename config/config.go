package config

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type providers struct {
	providers map[string]*viper.Viper
	mu        sync.Mutex
}

var baseConfigPath string
var p *providers

// Init is used to initialize the configurations
func Init(path string) {
	baseConfigPath = path
	p = &providers{
		providers: make(map[string]*viper.Viper),
	}
}

// Get is used to get the instance to the config provider for the configuration name
func Get(configFileName string) (*viper.Viper, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	// see for an existing provider
	if provider, ok := p.providers[configFileName]; ok {
		// provider already exists
		// use it
		return provider, nil
	}

	// try to get the provider
	provider := viper.New()
	provider.SetConfigName(configFileName)
	provider.AddConfigPath(baseConfigPath)
	err := provider.ReadInConfig()
	if err != nil {
		// config not found
		return nil, fmt.Errorf("config %s not found", configFileName)
	}

	// add a watcher for this provider
	provider.WatchConfig()

	//look for changes in the config file
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// successfully found config, store it for future use
	p.providers[configFileName] = provider

	return provider, nil
}
