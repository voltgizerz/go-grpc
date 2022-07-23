package utils

import "github.com/api-gateway/config"

var log = config.SetupLog()

// RecoverPanic - recover panic.
func RecoverPanic() {
	if err := recover(); err != nil {
		log.Errorf("panic occurred:", err)
	}
}
