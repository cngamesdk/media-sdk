package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestWordInfoNegativeWordUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.WordInfoNegativeWordUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 90000344
	req.UnitIds = []int64{761270}
	req.NegativeWord = kuaishouModel.NegativeWord{
		ExactWords:  []string{"增加精确否词"},
		PhraseWords: []string{"增加短语否词"},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.WordInfoNegativeWordUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
