package media_sdk

import (
	"github.com/cngamesdk/media-sdk/media/ocean_engine"
	"go.uber.org/zap"
	"testing"
)

func TestSdk(t *testing.T) {
	var logger *zap.Logger
	oceanEngineCode := ocean_engine.Code
	media, err := GetAdapter(oceanEngineCode, logger)
	if err != nil {
		t.Error(err)
		return
	}
	println(media.Name())

	noExistsCode := "test_code"
	_, errNoExists := GetAdapter(noExistsCode, logger)

	if errNoExists != nil {
		t.Error(errNoExists)
		return
	}
}
