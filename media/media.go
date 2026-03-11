package media

import (
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/utils"
)

type Media struct {
	Config *config.Config
	Client *utils.HTTPClient
}
