package insurance

import (
	"context"
	"gitlab.idc.xiaozhu.com/xz/lib/component/xzapi"
	"insurance/pkg/errors"
	"insurance/pkg/json"
	"net/http"
)

type TenantData struct {
	ID         string `json:"id"`
	Sex        string `json:"sex"`
	RealName   string `json:"realName"`
	Name       string `json:"name"`
	NationID   string `json:"nationId"`
	IDCardNo   string `json:"IDCardNo"`
	PassportNo string `json:"passportNo"`
	Mobile     string `json:"mobile"`
	Oversea    string `json:"oversea"`
}

type BookOrderClientContent struct {
	LuID             string       `json:"luId"`
	SkuID            string       `json:"skuId"`
	SubmitterID      string       `json:"submitterId"`
	SubmitterName    string       `json:"submitterName"`
	SubmitterMobile  string       `json:"submitterMobile"`
	LandlordID       string       `json:"landlordId"`
	CheckInDay       string       `json:"checkInDay"`
	CheckOutDay      string       `json:"checkOutDay"`
	RoomNum          int          `json:"roomNum"`
	BookOrderID      string       `json:"bookOrderId"`
	CurrentState     string       `json:"currentState"`
	LastState        string       `json:"lastState"`
	ActualTotalPrice int          `json:"actualTotalPrice"`
	CancelType       string       `json:"cancelType"`
	CancelReason     string       `json:"cancelReason"`
	CreateTime       string       `json:"createTime"`
	PayTime          string       `json:"payTime"`
	PayMethod        string       `json:"payMethod"`
	BookFromEnv      string       `json:"bookFromEnv"`
	Remark           string       `json:"remark"`
	BookFlow         string       `json:"bookFlow"`
	CheckInTime      string       `json:"checkInTime"`
	BizType          int          `json:"bizType"`
	TenantData       []TenantData `json:"tenantData"`
	ShowTotalPrice   int          `json:"showTotalPrice"`
	CleanFee         int          `json:"cleanFee"`
	RoomFee          struct {
		ShowPrice                int     `json:"showPrice"`
		UsePromotionFromLandlord int     `json:"usePromotionFromLandlord"`
		UseCoupon                float64 `json:"useCoupon"`
		UseCash                  float64 `json:"useCash"`
	} `json:"roomFee"`
	LodgeUnitInfo struct {
		LuID              string `json:"luId"`
		StoreID           string `json:"storeId"`
		SkuID             string `json:"skuId"`
		Name              string `json:"name"`
		StoreName         string `json:"storeName"`
		LodgeUnitFromType string `json:"lodgeUnitFromType"`
	} `json:"lodgeUnitInfo"`
	PartnerID string `json:"partnerId"`
	NationID  int    `json:"nationId"`
	CityID    int    `json:"cityId"`
}

type BookOrderClientResponse struct {
	Status    int                    `json:"status"`
	Content   BookOrderClientContent `json:"content"`
	ErrorMsg  string                 `json:"errorMsg"`
	TimeStamp int                    `json:"timeStamp"`
}

// OrderClient 需调用 NewOrderClient 获得初始化
type OrderClient struct {
	xzapi.Client
}

// NewOrderClient -
func NewOrderClient(ctx context.Context) *OrderClient {
	bc := &OrderClient{}
	bc.Name = "service_bookorder"
	bc.Ctx = ctx
	return bc
}

// GetBookOrderInfo returns a single luMetrics data
func (bo *OrderClient) GetBookOrderInfo(orderId string) (*BookOrderClientResponse, error) {
	var content BookOrderClientResponse
	client, err := bo.Client.Client().Get("/getBookOrderInfo/" + orderId)
	if err != nil {
		return &content, err
	}
	body, _ := client.GetBody()
	err = json.FromJSONString(string(body), &content)
	if err != nil {
		return &content, err
	}
	if content.Status == http.StatusOK {
		return &content, nil
	}
	return &content, errors.New(content.ErrorMsg)
}
