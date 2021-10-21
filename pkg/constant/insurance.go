package constant

const (
	// -上平安保险常量
	PinganClientId     = "" //详见配置文件
	PinganGrantType    = ""    //详见配置文件
	PinganClientSecret = ""              //详见配置文件

	PartnerCode        = "P_XZ_GA"                     //合作伙伴代码  P_XZ_GA|P_XZ_GA21030
	DepartmentCode     = ""                            //机构代码 21030 详见配置文件
	AmountCurrencyCode = "01"                          //币种
	OurToTheseName     = "北京快跑信息科技有限公司"                //名称
	CertificateNo      = "911101075996456643"          //证件号
	CertificateType    = "04"                          //证件类型
	PersonnelType      = "0"                           //个团标志[1个人,0团体]
	OurToTheseAddress  = "北京市石景山区石景山路乙18号院5号楼13层1517室" //投保人地址
	LinkManName        = "王连涛"                         //联系人名称
	InvoicePrintType   = "01"                          //发票打印类型
	PersonnelAttribute = "100"                         // [非空]人员属性 100（主真实被保人）、200（主虚拟被保人）

	PrepaidAccountId   = "" //详见配置文件
	PrepaidAccountType = "" //详见配置文件

	TenantAccidentProductCode          = "" //房客product code 详见配置文件
	TenantAccidentProductPackageCode   = "" // [非空]房客套餐编码 详见配置文件
	LandlordAccidentProductCode        = "" //房东product code 详见配置文件
	LandlordAccidentProductPackageCode = "" //  [非空]房东套餐编码 详见配置文件

	IsPolicyBeforePayfee = 1 //是否是见费出单(1-是;0-否)
	ApplyNum             = 1 //[非空]投保份数

	TotalActualPremium = "0.25" //房客每天的保费
	InsuranceBasePrice = "0.57" //房东每天的保费

	DefaultPingAnErrorCode = "500000"
	DefaultPingAnErrorMsg  = "第三方接口返回异常，请及时确认 "

	// IdCard 证件类型
	IdCard       = "IDcard"
	MilitaryCard = "militarycard"
	Passport     = "passport"

	DefaultZero        = 0
	DefaultEmptyString = ""
	DefaultOperate     = 1
	SexBaseNum         = 1
	SexMan             = "man"
	SexWoman           = "woman"

	DefaultCheckInTime  = "14:00:00"
	DefaultCheckOutTime = "12:00:00"

	DefaultPage         = 1
	DefaultPageSize     = 100 //底层最大100限制
	DefaultGoRoutineMax = 8
)
