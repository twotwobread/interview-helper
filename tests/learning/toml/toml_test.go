package learning_test

import (
	"log"
	"testing"

	"github.com/BurntSushi/toml"

	"github.com/stretchr/testify/assert"
)

type Config struct {
	Test struct {
		Value string
	}
}

func TestToml(t *testing.T) {
	var testConfig Config
	if _, err := toml.DecodeFile("config/test.toml", &testConfig); err != nil {
		log.Fatal(err)
		return
	}
	assert.Equal(t, "test", testConfig.Test.Value, "Toml decode TestConfig structure")
}
