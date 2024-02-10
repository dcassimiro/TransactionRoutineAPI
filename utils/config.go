package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config configuration interface
type Config interface {
	GetInt(key string) int
	GetBool(key string) bool
	GetString(key string) string
	GetFloat64(key string) float64
	GetDuration(key string) time.Duration
	Close()
}

type configImpl struct {
	vmain     *viper.Viper
	vreplacer *viper.Viper

	kill chan bool
}

func (c *configImpl) GetBool(key string) bool {
	return c.vmain.GetBool(key)
}

func (c *configImpl) GetString(key string) string {
	value := c.vmain.GetString(key)
	for k, v := range c.vreplacer.AllSettings() {
		old := "$" + k + "$"
		new, ok := v.(string)
		if !ok {
			continue
		}
		value = strings.ReplaceAll(value, old, new)
	}
	return value
}

func (c *configImpl) GetDuration(key string) time.Duration {
	return c.vmain.GetDuration(key)
}

func (c *configImpl) GetInt(key string) int {
	return c.vmain.GetInt(key)
}

func (c *configImpl) GetFloat64(key string) float64 {
	return c.vmain.GetFloat64(key)
}

// Close closes the watch function
func (c *configImpl) Close() {
	c.kill <- true
}

// Watch modified
func Watch(fn func(c Config, quit chan bool)) {
	quit, kill := make(chan bool), make(chan bool)
	vmain, vreplacer := viper.New(), viper.New()

	c := &configImpl{
		vmain:     vmain,
		vreplacer: vreplacer,
		kill:      kill,
	}

	//Replace the _ with .
	vmain.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// binds environment variables
	vmain.AutomaticEnv()

	// start the server
	go fn(c, quit)

	<-kill
}
