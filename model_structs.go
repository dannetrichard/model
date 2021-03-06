package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Openid             string `json:"openid" gorm:"unique"`
	SessionKey         string `json:"session_key"`
	Unionid            string `json:"unionid"`
	BuyerNick          string `json:"buyer_nick" gorm:"unique"`
	WeChat             string `json:"we_chat" gorm:"unique"`
	Level              int    `json:"level" gorm:"default:0"`
	Coins              int    `json:"coins" gorm:"default:0"`
	WeChatChangedTimes int    `json:"we_chat_changed_times" gorm:"default:0"`
}

type Refund struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	BuyerNick          string `json:"buyer_nick"`
	Created            string `json:"created"`
	Desc               string `json:"desc"`
	GoodStatus         string `json:"good_status"`
	HasGoodReturn      bool   `json:"has_good_return"`
	Modified           string `json:"modified"`
	Num                int    `json:"num"`
	NumIid             int64  `json:"num_iid"`
	Oid                string `json:"oid"`
	OperationContraint string `json:"operation_contraint"`
	OrderStatus        string `json:"order_status"`
	Payment            string `json:"payment"`
	Price              string `json:"price"`
	Reason             string `json:"reason"`
	RefundFee          string `json:"refund_fee"`
	RefundID           string `json:"refund_id" gorm:"unique"`
	RefundPhase        string `json:"refund_phase"`
	Sku                string `json:"sku"`
	Status             string `json:"status"`
	Tid                string `json:"tid"`
	Title              string `json:"title"`
	TotalFee           string `json:"total_fee"`
	CompanyName        string `json:"company_name"`
	Sid                string `json:"sid"`
	GoodReturnTime     string `json:"good_return_time"`
}

type Ship struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	CompanyName string `json:"company_name"`
	Sid         string `json:"sid" gorm:"unique"`
}

type Cat struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Cid       int    `json:"cid"`
	Name      string `json:"name"`
	ParentCid int    `json:"parent_cid"`
	PicURL    string `json:"pic_url"`
	SortOrder int    `json:"sort_order"`
	Type      string `json:"type"`
}

type Item struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	ApproveStatus          string `json:"approve_status"`
	AuctionPoint           int    `json:"auction_point"`
	Cid                    int    `json:"cid"`
	Created                string `json:"created"`
	DelistTime             string `json:"delist_time"`
	Desc                   string `json:"desc" gorm:"type:varchar(5000)"`
	DetailURL              string `json:"detail_url"`
	EmsFee                 string `json:"ems_fee"`
	ExpressFee             string `json:"express_fee"`
	Features               string `json:"features"`
	FreightPayer           string `json:"freight_payer"`
	HasDiscount            bool   `json:"has_discount"`
	HasInvoice             bool   `json:"has_invoice"`
	HasShowcase            bool   `json:"has_showcase"`
	HasWarranty            bool   `json:"has_warranty"`
	Increment              string `json:"increment"`
	InputPids              string `json:"input_pids"`
	InputStr               string `json:"input_str"`
	IsEx                   bool   `json:"is_ex"`
	IsFenxiao              int    `json:"is_fenxiao"`
	IsLightningConsignment bool   `json:"is_lightning_consignment"`
	IsTaobao               bool   `json:"is_taobao"`
	IsTiming               bool   `json:"is_timing"`
	IsVirtual              bool   `json:"is_virtual"`
	IsXinpin               bool   `json:"is_xinpin"`
	ItemWeight             string `json:"item_weight"`
	ListTime               string `json:"list_time"`
	Modified               string `json:"modified"`
	Newprepay              string `json:"newprepay"`
	Nick                   string `json:"nick"`
	Num                    int    `json:"num"`
	NumIid                 int64  `json:"num_iid" gorm:"unique"`
	OneStation             bool   `json:"one_station"`
	PeriodSoldQuantity     int    `json:"period_sold_quantity"`
	PicURL                 string `json:"pic_url"`
	PostFee                string `json:"post_fee"`
	PostageID              int64  `json:"postage_id"`
	Price                  string `json:"price"`
	ProductID              int    `json:"product_id"`
	PropertyAlias          string `json:"property_alias"`
	Props                  string `json:"props"`
	PropsName              string `json:"props_name" gorm:"type:varchar(5000)"`
	SellPromise            bool   `json:"sell_promise"`
	SellerCids             string `json:"seller_cids"`
	SoldQuantity           int    `json:"sold_quantity"`
	StuffStatus            string `json:"stuff_status"`
	SubStock               int    `json:"sub_stock"`
	Title                  string `json:"title"`
	Type                   string `json:"type"`
	ValidThru              int    `json:"valid_thru"`
	Violation              bool   `json:"violation"`
	WapDesc                string `json:"wap_desc"`
	WapDetailURL           string `json:"wap_detail_url"`
	MyDiscount             int    `json:"my_discount"`
}

type Trade struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	AdjustFee       string `json:"adjust_fee"`
	BuyerNick       string `json:"buyer_nick"`
	BuyerRate       bool   `json:"buyer_rate"`
	ConsignTime     string `json:"consign_time"`
	Created         string `json:"created"`
	DiscountFee     string `json:"discount_fee"`
	EndTime         string `json:"end_time"`
	Modified        string `json:"modified"`
	PayTime         string `json:"pay_time"`
	Payment         string `json:"payment"`
	PostFee         string `json:"post_fee"`
	ReceivedPayment string `json:"received_payment"`
	SellerRate      bool   `json:"seller_rate"`
	ShippingType    string `json:"shipping_type"`
	Status          string `json:"status"`
	Tid             string `json:"tid" gorm:"unique"`
	Title           string `json:"title"`
	TotalFee        string `json:"total_fee"`
	Type            string `json:"type"`
}

type Order struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Tid               string `json:"tid"`
	BuyerNick         string `json:"buyer_nick"`
	RefundTime        string `json:"refund_time"`
	RefundWeChat      string `json:"refund_we_chat"`
	RefundFee         string `json:"refund_fee"`
	AdjustFee         string `json:"adjust_fee"`
	BuyerRate         bool   `json:"buyer_rate"`
	Cid               int    `json:"cid"`
	ConsignTime       string `json:"consign_time"`
	DiscountFee       string `json:"discount_fee"`
	DivideOrderFee    string `json:"divide_order_fee"`
	EndTime           string `json:"end_time"`
	InvoiceNo         string `json:"invoice_no"`
	IsDaixiao         bool   `json:"is_daixiao"`
	IsOversold        bool   `json:"is_oversold"`
	LogisticsCompany  string `json:"logistics_company"`
	Num               int    `json:"num"`
	NumIid            int64  `json:"num_iid"`
	Oid               string `json:"oid" gorm:"unique"`
	OidStr            string `json:"oid_str"`
	OrderFrom         string `json:"order_from"`
	Payment           string `json:"payment"`
	PicPath           string `json:"pic_path"`
	Price             string `json:"price"`
	RefundID          string `json:"refund_id"`
	RefundStatus      string `json:"refund_status"`
	SellerRate        bool   `json:"seller_rate"`
	SellerType        string `json:"seller_type"`
	SkuID             string `json:"sku_id"`
	SkuPropertiesName string `json:"sku_properties_name"`
	Status            string `json:"status"`
	Title             string `json:"title"`
	TotalFee          string `json:"total_fee"`
}

type AdminUser struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email" gorm:"unique"`
	AccessToken   string `json:"accessToken"`
	RefreshToken  string `json:"refreshToken"`
	Roles         string `json:"roles"`
	Pic           string `json:"pic"`
	Fullname      string `json:"fullname"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Occupation    string `json:"occupation"`
	CompanyName   string `json:"companyName"`
	Phone         string `json:"phone"`
	Language      string `json:"language"`
	TimeZone      string `json:"timeZone"`
	Website       string `json:"website"`
	EmailSettings struct {
		EmailNotification       bool `json:"emailNotification"`
		SendCopyToPersonalEmail bool `json:"sendCopyToPersonalEmail"`
		ActivityRelatesEmail    struct {
			YouHaveNewNotifications       bool `json:"youHaveNewNotifications"`
			YouAreSentADirectMessage      bool `json:"youAreSentADirectMessage"`
			SomeoneAddsYouAsAsAConnection bool `json:"someoneAddsYouAsAsAConnection"`
			UponNewOrder                  bool `json:"uponNewOrder"`
			NewMembershipApproval         bool `json:"newMembershipApproval"`
			MemberRegistration            bool `json:"memberRegistration"`
		} `json:"activityRelatesEmail"`
		UpdatesFromKeenthemes struct {
			NewsAboutKeenthemesProductsAndFeatureUpdates       bool `json:"newsAboutKeenthemesProductsAndFeatureUpdates"`
			TipsOnGettingMoreOutOfKeen                         bool `json:"tipsOnGettingMoreOutOfKeen"`
			ThingsYouMissedSindeYouLastLoggedIntoKeen          bool `json:"thingsYouMissedSindeYouLastLoggedIntoKeen"`
			NewsAboutMetronicOnPartnerProductsAndOtherServices bool `json:"newsAboutMetronicOnPartnerProductsAndOtherServices"`
			TipsOnMetronicBusinessProducts                     bool `json:"tipsOnMetronicBusinessProducts"`
		} `json:"updatesFromKeenthemes"`
	} `json:"emailSettings" gorm:"embedded`
	Communication struct {
		Email bool `json:"email"`
		Sms   bool `json:"sms"`
		Phone bool `json:"phone"`
	} `json:"communication" gorm:"embedded`
	Address struct {
		AddressLine string `json:"addressLine"`
		City        string `json:"city"`
		State       string `json:"state"`
		PostCode    string `json:"postCode"`
	} `json:"address" gorm:"embedded`
	SocialNetworks struct {
		LinkedIn  string `json:"linkedIn"`
		Facebook  string `json:"facebook"`
		Twitter   string `json:"twitter"`
		Instagram string `json:"instagram"`
	} `json:"socialNetworks" gorm:"embedded`
}
