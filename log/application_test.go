package log

import (
	"os"
	"testing"
	"time"

	"github.com/goravel/framework/config"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/testing/file"
	testingfile "github.com/goravel/framework/testing/file"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	err := testingfile.CreateEnv()
	assert.Nil(t, err)

	addDefaultConfig()

	app := Application{}
	instance := app.Init()

	instance.Debug("debug")
	instance.Error("error")

	dailyFile := "storage/logs/goravel-" + time.Now().Format("2006-01-02") + ".log"
	singleFile := "storage/logs/goravel.log"
	singleErrorFile := "storage/logs/goravel-error.log"

	assert.FileExists(t, dailyFile)
	assert.FileExists(t, singleFile)
	assert.FileExists(t, singleErrorFile)

	assert.Equal(t, 3, file.GetLineNum(dailyFile))
	assert.Equal(t, 3, file.GetLineNum(singleFile))
	assert.Equal(t, 2, file.GetLineNum(singleErrorFile))

	err = os.Remove(".env")
	assert.Nil(t, err)

	err = os.RemoveAll("storage")
	assert.Nil(t, err)
}

//addDefaultConfig Add default config for test.
func addDefaultConfig() {
	configApp := config.ServiceProvider{}
	configApp.Register()

	facadesConfig := facades.Config
	facadesConfig.Add("logging", map[string]interface{}{
		"default": facadesConfig.Env("LOG_CHANNEL", "stack"),
		"channels": map[string]interface{}{
			"stack": map[string]interface{}{
				"driver":   "stack",
				"channels": []string{"daily", "single", "single-error"},
			},
			"single": map[string]interface{}{
				"driver": "single",
				"path":   "storage/logs/goravel.log",
				"level":  "debug",
			},
			"single-error": map[string]interface{}{
				"driver": "single",
				"path":   "storage/logs/goravel-error.log",
				"level":  "error",
			},
			"daily": map[string]interface{}{
				"driver": "daily",
				"path":   "storage/logs/goravel.log",
				"level":  facadesConfig.Env("LOG_LEVEL", "debug"),
				"days":   7,
			},
		},
	})
}
