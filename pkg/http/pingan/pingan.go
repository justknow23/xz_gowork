package pingan

import (
	"encoding/json"
	"fmt"
	"gitlab.idc.xiaozhu.com/xz-go/common/client"
	"gitlab.idc.xiaozhu.com/xz-go/common/log"
	"gitlab.idc.xiaozhu.com/xz/lib/component/xzapi"
	"insurance/pkg/constant"
	"insurance/pkg/global"
	"net/url"
	"strings"
)

const ResponseCode999999 = "999999" //正常成功
const ResponseCode888888 = "888888" //重复投保

// GetTokenRequest 获取token参数
type GetTokenRequest struct {
	ClientId     string `json:"client_id"`
	GrantType    string `json:"grant_type"`
	ClientSecret string `json:"client_secret"`
}

type GetTokenResponse struct {
	Ret  string `json:"ret"`
	Data struct {
		ExpiresIn   string `json:"expires_in"`
		Openid      string `json:"openid"`
		AccessToken string `json:"access_token"`
	} `json:"data"`
	Msg string `json:"msg"`
}

// ApplicationInfoListInner 上保险参数
type ApplicationInfoListInner struct {
	Name             string `json:"name"`
	CertificateNo    string `json:"certificateNo"`
	CertificateType  string `json:"certificateType"`
	PersonnelType    string `json:"personnelType"`
	Address          string `json:"address"`
	LinkManName      string `json:"linkManName"`
	InvoicePrintType string `json:"invoicePrintType"`
}

type BaseInfo struct {
	AmountCurrencyCode string `json:"amountCurrencyCode"`
	InsuranceBeginDate string `json:"insuranceBeginDate"`
	InsuranceEndDate   string `json:"insuranceEndDate"`
	TotalActualPremium string `json:"totalActualPremium"`
	ProductCode        string `json:"productCode"`
}
type ExtendInfo struct {
	IsPolicyBeforePayfee int `json:"isPolicyBeforePayfee"`
}

type RiskPackageInfoListInner struct {
	ApplyNum           int    `json:"applyNum"`
	ProductPackageCode string `json:"productPackageCode"`
}
type RiskPersonInfoListInner struct {
	Name               string `json:"name"`
	CertificateNo      string `json:"certificateNo"`
	CertificateType    string `json:"certificateType"`
	PersonnelAttribute string `json:"personnelAttribute"`
	MobileTelephone    string `json:"mobileTelephone"`
	Birthday           string `json:"birthday"`
	Age                string `json:"age"`
	HouseAddress       string `json:"houseAddress"`
}

type RiskGroupInfoListInner struct {
	RiskPackageInfoList []RiskPackageInfoListInner `json:"riskPackageInfoList"`
	RiskPersonInfoList  []RiskPersonInfoListInner  `json:"riskPersonInfoList"`
}

type PingAnRequest struct {
	PartnerCode        string `json:"partnerCode"`
	DepartmentCode     string `json:"departmentCode"`
	TransSerialNo      string `json:"transSerialNo"`
	PrepaidAccountID   string `json:"prepaidAccountId"`
	PrepaidAccountType string `json:"prepaidAccountType"`
	Contract           struct {
		ApplicantInfoList []ApplicationInfoListInner `json:"applicantInfoList"`
		BaseInfo          BaseInfo                   `json:"baseInfo"`
		ExtendInfo        ExtendInfo                 `json:"extendInfo"`
		RiskGroupInfoList []RiskGroupInfoListInner   `json:"riskGroupInfoList"`
	} `json:"contract"`
}

type PingAnResponseResult struct {
	ApplyPolicyNo       string `json:"applyPolicyNo"`
	PolicyNo            string `json:"policyNo"`
	NoticeNo            string `json:"noticeNo"`
	ProductCode         string `json:"productCode"`
	ProductName         string `json:"productName"`
	TotalInsuredAmount  string `json:"totalInsuredAmount"`
	AmountCurrencyCode  string `json:"amountCurrencyCode"`
	TotalActualPremium  string `json:"totalActualPremium"`
	PremiumCurrencyCode string `json:"premiumCurrencyCode"`
	InsuranceBeginDate  string `json:"insuranceBeginDate"`
	InsuranceEndDate    string `json:"insuranceEndDate"`
}

type PingAnResponse struct {
	Ret       string `json:"ret"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
	Data      struct {
		ResponseCode string               `json:"responseCode"`
		ResponseMsg  string               `json:"responseMsg"`
		Result       PingAnResponseResult `json:"result"`
	} `json:"data"`
}

// PingAnClient 需调用 NewOrderClient 获得初始化
type PingAnClient struct {
	xzapi.Client
}

// GetToken  获取授权token 可增加缓存减少http调用
func (bo *PingAnClient) GetToken() (GetTokenResponse, error) {
	var content GetTokenResponse
	var uri url.URL
	q := uri.Query()
	q.Add("client_id", global.Settings.PinganClientId)
	q.Add("grant_type", global.Settings.PinganGrantType)
	q.Add("client_secret", global.Settings.PinganClientSecret)
	queryStr := q.Encode()
	log.Info("pingan.GetToken query: ", queryStr)
	clientGet, err := bo.Client.Client().Get("/oauth/oauth2/access_token?" + queryStr)
	if err != nil {
		return content, err
	}
	body, err := clientGet.GetBody()
	log.Info("pingan.GetToken body: ", body)
	if err != nil {
		return content, err
	}
	err = json.Unmarshal(body, &content)
	if err != nil {
		return content, err
	}
	return content, nil
}

// DoInsurance  上保险
func (bo *PingAnClient) DoInsurance(req PingAnRequest, accessToken, requestId, productCode, productPackageCode string, riskPersonInfo []RiskPersonInfoListInner) (PingAnResponse, error) {
	var (
		content                     PingAnResponse
		ApplicationInfoListInnerArr []ApplicationInfoListInner
		uri                         url.URL
		clientPost                  *client.Response
		body                        client.ResponseBody
	)
	// need insuranceBeginDate insuranceEndDate totalActualPremium
	// need riskPersonInfoList
	//赋默认值
	req.PartnerCode = constant.PartnerCode
	req.DepartmentCode = global.Settings.PingAnDepartmentCode
	req.PrepaidAccountType = global.Settings.PingAnPrepaidAccountType
	req.PrepaidAccountID = global.Settings.PingAnPrepaidAccountId

	ailn := ApplicationInfoListInner{
		Name:             constant.OurToTheseName,
		CertificateNo:    constant.CertificateNo,
		CertificateType:  constant.CertificateType,
		PersonnelType:    constant.PersonnelType,
		Address:          constant.OurToTheseAddress,
		LinkManName:      constant.LinkManName,
		InvoicePrintType: constant.InvoicePrintType,
	}
	req.Contract.ApplicantInfoList = append(ApplicationInfoListInnerArr, ailn)

	req.Contract.BaseInfo.AmountCurrencyCode = constant.AmountCurrencyCode
	req.Contract.BaseInfo.ProductCode = productCode
	req.Contract.ExtendInfo.IsPolicyBeforePayfee = constant.IsPolicyBeforePayfee

	riskPackageInfo := []RiskPackageInfoListInner{
		{
			ApplyNum:           constant.ApplyNum,
			ProductPackageCode: productPackageCode,
		},
	}
	riskGroupInfo := []RiskGroupInfoListInner{
		{
			RiskPackageInfoList: riskPackageInfo,
			RiskPersonInfoList:  riskPersonInfo,
		},
	}
	req.Contract.RiskGroupInfoList = riskGroupInfo

	reqJson, err := json.Marshal(req)
	if err != nil {
		return content, err
	}
	log.Infof("app.DoInsurance request: %+v\n", string(reqJson))
	//拼接query
	q := uri.Query()
	q.Add("access_token", accessToken)
	q.Add("request_id", requestId)
	queryStr := q.Encode()
	//请求
	clientPost, err = bo.Client.Client().Post("/open/appsvr/property/openapi/P_XZ_GA/applyGA?"+queryStr, client.Options{
		JSON: req,
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
	})
	log.Info("pingan.DoInsurance clientPost======: ", clientPost, err)
	if err != nil {
		return content, err
	}
	body, err = clientPost.GetBody()
	log.Info("pingan.DoInsurance body======: ", body, err)
	if err != nil {
		return content, err
	}
	err = json.Unmarshal(body, &content)
	log.Info("pingan.DoInsurance content======: ", content, err)
	if err != nil {
		return content, err
	}
	log.Info("pingan.DoInsurance return======: ", content)
	return content, nil
}

func (bo *PingAnClient) CheckFailCode(code string) bool {
	codes := global.Settings.FailCodeList
	log.Info(fmt.Sprintf("pingan.CheckFailCode %v", codes))
	if codes != "" {
		codesList := strings.Split(codes, ",")
		for _, v := range codesList {
			if v == code {
				return true
			}
		}
	}
	return false
}
