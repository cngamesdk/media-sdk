package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAuaxProjectCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AuaxProjectCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.Name = "测试智投项目"
	req.KolUserId = 2022036857
	req.KolUserType = 1
	req.AutoManageType = 2
	req.SubjectId = 1234567890
	req.OcpxActionType = 191
	req.RoiRatio = 1.5
	req.Description = "测试广告语"
	req.PhotoPackageInfo = []int64{1234567890}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AuaxProjectCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
