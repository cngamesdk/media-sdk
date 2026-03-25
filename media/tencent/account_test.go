package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

func TestOrganizationAccountRelationGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.OrganizationAccountRelationGetReq{}
	req.AccessToken = "123"
	req.PaginationMode = model.PaginationModeNormal
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OrganizationAccountRelationGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}
