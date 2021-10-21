package insurance

import (
	"context"
	"gitlab.idc.xiaozhu.com/xz-go/common/client"
	"gitlab.idc.xiaozhu.com/xz/lib/component/xzapi"
	"insurance/pkg/errors"
)

type LodgeAddrInfoResponse struct {
	Status    int         `json:"status"`
	Content   struct{}    `json:"content"`
	ErrorMsg  interface{} `json:"errorMsg"`
	Timestamp string      `json:"timestamp"`
}

type LodgeAddrInfoBody struct {
	AddressID      string `json:"addressId"`
	NationID       string `json:"nationId"`
	NationName     string `json:"nationName"`
	DistrictID     string `json:"districtId"`
	DistrictName   string `json:"districtName"`
	ProvinceID     string `json:"provinceId"`
	ProvinceName   string `json:"provinceName"`
	CityID         string `json:"cityId"`
	CityName       string `json:"cityName"`
	Longitude      string `json:"longitude"`
	Latitude       string `json:"latitude"`
	DetailAddress  string `json:"detailAddress"`
	RoomNumber     string `json:"roomNumber"`
	DisplayAddress string `json:"displayAddress"`
	LockID         string `json:"lockId"`
}

// LodgeClient 需调用 NewOrderClient 获得初始化
type LodgeClient struct {
	xzapi.Client
}

// NewLodgeClient -
func NewLodgeClient(ctx context.Context) *LodgeClient {
	bc := &LodgeClient{}
	bc.Name = "service_lodge"
	bc.Ctx = ctx
	return bc
}

// GetLodgeUnitInfoByLuId returns a single luMetrics data
func (bo *LodgeClient) GetLodgeUnitInfoByLuId(luId string) (*LodgeAddrInfoBody, error) {
	var content map[string]LodgeAddrInfoBody
	_, err := bo.ClientWithParseContent(&content).Get("/lucenter/getAddrInfoByLuIds", client.Options{
		Query: map[string]interface{}{
			"luIds": luId,
		},
	})
	if err != nil {
		return nil, err
	}
	if rs, ok := content[luId]; ok {
		return &rs, nil
	}
	return nil, errors.ErrorData
}
