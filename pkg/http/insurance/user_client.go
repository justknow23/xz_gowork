package insurance

import (
	"context"
	"gitlab.idc.xiaozhu.com/xz-go/common/client"
	"gitlab.idc.xiaozhu.com/xz/lib/component/xzapi"
	"insurance/pkg/errors"
	"insurance/pkg/json"
	"net/http"
)

type UserClientResponse struct {
	Status  int `json:"status"`
	Content struct {
		UserID                       string      `json:"userId"`
		Business                     string      `json:"business"`
		RealName                     string      `json:"realName"`
		NickName                     string      `json:"nickName"`
		Sex                          string      `json:"sex"`
		UserRole                     string      `json:"userRole"`
		MobileNation                 int         `json:"mobileNation"`
		Mobile                       string      `json:"mobile"`
		Mobilebak                    string      `json:"mobilebak"`
		HeadImgID                    string      `json:"headImgId"`
		MobileCode                   string      `json:"mobileCode"`
		ProvinceID                   int         `json:"provinceId"`
		CityID                       int         `json:"cityId"`
		Email                        string      `json:"email"`
		UserFrom                     string      `json:"userFrom"`
		NationID                     int         `json:"nationId"`
		SelfIntro                    interface{} `json:"selfIntro"`
		Birthday                     string      `json:"birthday"`
		Education                    string      `json:"education"`
		Profession                   string      `json:"profession"`
		AgeGroup                     interface{} `json:"ageGroup"`
		Age                          int         `json:"age"`
		IDCardNo                     string      `json:"idCardNo"`
		PassportNo                   string      `json:"passportNo"`
		CreateTime                   string      `json:"createTime"`
		IsPersonalLandlord           bool        `json:"isPersonalLandlord"`
		IsOrgLandlord                bool        `json:"isOrgLandlord"`
		IsLandlord                   bool        `json:"isLandlord"`
		WeixinUnionID                string      `json:"weixinUnionId"`
		WeixinAuth                   string      `json:"weixinAuth"`
		FuwuchuangOpenID             string      `json:"fuwuchuangOpenId"`
		FuwuchuangAuth               string      `json:"fuwuchuangAuth"`
		SinaWeiboOpenID              string      `json:"sinaWeiboOpenId"`
		SinaWeiboAuth                string      `json:"sinaWeiboAuth"`
		ZhiMaOpenID                  string      `json:"zhiMaOpenId"`
		XiaoBaiOpenID                string      `json:"xiaoBaiOpenId"`
		HeadImgURL                   string      `json:"headImgUrl"`
		HeadImgURLMedium             string      `json:"headImgUrl_Medium"`
		HeadImgURLSmall              string      `json:"headImgUrl_Small"`
		RealNameStatus               bool        `json:"realNameStatus"`
		OrderNumsAsSubmitter         int         `json:"orderNumsAsSubmitter"`
		TenantCommentNums            int         `json:"tenantCommentNums"`
		ReplyRate                    int         `json:"replyRate"`
		ConfirmRate                  int         `json:"confirmRate"`
		ConfirmMinutes               int         `json:"confirmMinutes"`
		IsVipNow                     int         `json:"isVipNow"`
		StudentIdentifyStatus        bool        `json:"studentIdentifyStatus"`
		StudentIdentifyLastAuditTime string      `json:"studentIdentifyLastAuditTime"`
		FkRealNameVerify             bool        `json:"fkRealNameVerify"`
		FdRealNameVerify             bool        `json:"fdRealNameVerify"`
	} `json:"content"`
	TimeStamp int `json:"timeStamp"`
}

// UserClient 需调用 NewUserClient 获得初始化
type UserClient struct {
	xzapi.Client
}

// NewUserClient -
func NewUserClient(ctx context.Context) *UserClient {
	bc := &UserClient{}
	bc.Name = "service_user"
	bc.Ctx = ctx
	return bc
}

// GetUserInfo returns a single luMetrics data
func (bo *UserClient) GetUserInfo(userId, level string) (*UserClientResponse, error) {
	var content UserClientResponse
	client, err := bo.Client.Client().Get("/user/getById", client.Options{
		Query: map[string]interface{}{
			"id":        userId,
			"infoLevel": level,
		},
	})
	if err != nil {
		return nil, err
	}
	body, _ := client.GetBody()
	err = json.FromJSONString(string(body), &content)
	if err != nil {
		return &content, err
	}
	if content.Status == http.StatusOK {
		return &content, nil
	}
	return &content, errors.New("/user/getById接口调用Err")
}
