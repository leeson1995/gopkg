package env

import "os"

const (
	APP_ENV     = "app_env"
	PRODUCT_ENV = "product"
	TEST_ENV    = "test"
)

var (
	env = TEST_ENV
)

func init() {
	nenv := os.Getenv(APP_ENV)
	if nenv == PRODUCT_ENV {
		env = PRODUCT_ENV
	}
}

func IsProduct() bool {
	return env == PRODUCT_ENV
}

func IsTest() bool {
	return env == TEST_ENV
}

func GetEnv() string {
	return env
}
