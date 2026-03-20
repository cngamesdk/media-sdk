package toutiao

import (
	"github.com/cngamesdk/media-sdk/media/toutiao/model"
	"testing"
)

func TestClickMaros(t *testing.T) {
	macros := model.ClickMacros
	println(macros.BuildQueryString())

	macros.Reset(model.ProjectID, "first_level_id")
	println(macros.BuildQueryString())

	macros.Add("ext1", "EXT1")
	println(macros.BuildQueryString())

	macros.Add("ext2", "EXT2")
	println(macros.BuildUrl("https://www.xxx.com/?game_id=123&agent_id=123&site_id=123"))
}
