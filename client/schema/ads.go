package schema

import "fmt"

type GetAdsResponse struct {
	Advertisement   bool          `json:"advertisement"`
	ConfigEtag      string        `json:"config_etag"`
	CounterMap      CounterMap    `json:"counter_map"`
	Gallery         []interface{} `json:"gallery"`
	ListAds         []ListAd      `json:"list_ads"`
	NextPage        string        `json:"next_page"`
	ProximitySlices []interface{} `json:"proximity_slices"`
	Sorting         string        `json:"sorting"`
}

type CounterMap struct {
	All int64 `json:"all"`
}

type ListAd struct {
	Ad       *Ad               `json:"ad"`
	LabelMap map[string]string `json:"labelmap"`
}

type Ad struct {
	Address         *ParamValue            `json:"address,omitempty"`
	AddressNumber   *ParamValue            `json:"address_number,omitempty"`
	Description     string                 `json:"body"`
	CanEdit         bool                   `json:"can_edit"`
	IsCompanyAd     bool                   `json:"company_ad"`
	Category        CategoryClass          `json:"category"`
	Images          []Image                `json:"images,omitempty"`
	ListID          string                 `json:"list_id"`
	ListPrice       *ListPrice             `json:"list_price,omitempty"`
	ListTime        ListTime               `json:"list_time"`
	Locations       []AdLocation           `json:"locations"`
	Prices          []ListPrice            `json:"prices,omitempty"`
	ShareLink       string                 `json:"share_link"`
	Subject         string                 `json:"subject"`
	User            UserClass              `json:"user"`
	Thumbnail       *Image                 `json:"thumbnail,omitempty"`
	Type            *ParamValue            `json:"type,omitempty"`
	AdDetails       *map[string]MixedParam `json:"ad_details,omitempty"`
	PaymentDelivery *bool                  `json:"payment_delivery"`
	PhoneHidden     *bool                  `json:"phone_hidden,omitempty"`
	HasPlates       *bool                  `json:"has_plates,omitempty"`
	HighlightPrice  *bool                  `json:"highlight_price,omitempty"`
	PriceAutofact   *ListPrice             `json:"price_autofact,omitempty"`
	ContractType    *ParamValue            `json:"contract_type,omitempty"`
	JobCategory     *MultipleParam         `json:"job_category,omitempty"`
	WorkingDay      *ParamValue            `json:"working_day,omitempty"`
	Condition       *ParamValue            `json:"condition,omitempty"`
	Gender          *ParamValue            `json:"gender,omitempty"`
	InternalMemory  *ParamValue            `json:"internal_memory,omitempty"`
	Cartype         *ParamValue            `json:"cartype,omitempty"`
	VehicleBrand    *ParamValue            `json:"vehicle_brand,omitempty"`
	VehicleModel    *ParamValue            `json:"vehicle_model,omitempty"`
	VehicleVersion  *ParamValue            `json:"vehicle_version,omitempty"`
	Bathrooms       *ParamValue            `json:"bathrooms,omitempty"`
	EstateType      *ParamValue            `json:"estate_type,omitempty"`
	Rooms           *ParamValue            `json:"rooms,omitempty"`
	Size            *ParamValue            `json:"size,omitempty"`
	Cubiccms        *ParamValue            `json:"cubiccms,omitempty"`
	Mileage         *ParamValue            `json:"mileage,omitempty"`
	Fuel            *ParamValue            `json:"fuel,omitempty"`
	Gearbox         *ParamValue            `json:"gearbox,omitempty"`
	FootwearGender  *ParamValue            `json:"footwear_gender,omitempty"`
	FootwearSize    *ParamValue            `json:"footwear_size,omitempty"`
	FootwearType    *MultipleParam         `json:"footwear_type,omitempty"`
	ServiceType     *MultipleParam         `json:"service_type,omitempty"`
	UtilSize        *ParamValue            `json:"util_size,omitempty"`
}

type MixedParam struct {
	Single   ParamValue   `json:"single"`
	Multiple []ParamValue `json:"multiple"`
}

type MultipleParam struct {
	Multiple []ParamValue `json:"multiple"`
}

type ParamValue struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

type CategoryClass struct {
	Code  string  `json:"code"`
	Color string  `json:"color"`
	Label string  `json:"label"`
	Icon  *string `json:"icon,omitempty"`
}

type Image struct {
	BaseURL string `json:"base_url"` // i.e: https://img.yapo.cl
	MediaID string `json:"media_id"` // i.e: public/media/ad/2153839858,
	Path    string `json:"path"`     // i.e: 21/2153839858.jpg
	Width   int64  `json:"width"`    // i.e: 640
	Height  int64  `json:"height"`   // i.e: 480
}

// GetUrl example: https://img.yapo.cl/images/25/2540986382.jpg
func (i *Image) GetUrl() string {
	return fmt.Sprintf("%s/images/%s", i.BaseURL, i.Path)
}

type ListPrice struct {
	Currency   Currency   `json:"currency"`
	Label      string     `json:"label"`
	PriceValue float64    `json:"price_value"`
	OldPrice   *ListPrice `json:"old_price,omitempty"`
}

type ListTime struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}

type AdLocation struct {
	Code      string           `json:"code"`
	Key       string           `json:"key"`
	Label     GeoRegion        `json:"label"`
	Locations []LocationDetail `json:"locations"`
}

type LocationDetail struct {
	Code        string           `json:"code"`
	Key         string           `json:"key"`
	Label       string           `json:"label"`
	Locations   []LocationDetail `json:"locations,omitempty"`
	Coordinates *Coordinates     `json:"coordinates,omitempty"`
}

type Coordinates struct {
	Lat    float64 `json:"lat"`
	Long   float64 `json:"long"`
	Radius *int64  `json:"radius,omitempty"`
}

type UserClass struct {
	Account Account `json:"account"`
}

type Account struct {
	Name         string  `json:"name"`
	ProfileToken string  `json:"profile_token"`
	Since        *string `json:"since,omitempty"`
	UserID       string  `json:"user_id"`
}

type Currency string

const (
	CurrencyCLP = Currency("CLP")
	CurrencyUF  = Currency("UF")
)

type GeoRegion string

const (
	RegionMetropolitana    = GeoRegion("Región Metropolitana")
	ITarapaca              = GeoRegion("I Tarapacá")
	IIAntofagasta          = GeoRegion("II Antofagasta")
	IIIAtacama             = GeoRegion("III Atacama")
	IVCoquimbo             = GeoRegion("IV Coquimbo")
	VValparaiso            = GeoRegion("V Valparaíso")
	VIOHiggins             = GeoRegion("VI O'Higgins")
	VIIMaule               = GeoRegion("VII Maule")
	VIIIBiobio             = GeoRegion("VIII Biobío")
	IXAraucania            = GeoRegion("IX Araucanía")
	XLosLagos              = GeoRegion("X Los Lagos")
	XIAisen                = GeoRegion("XI Aisén")
	XIIMagallanesAntartica = GeoRegion("XII Magallanes & Antártica")
	XIVLosRios             = GeoRegion("XIV Los Ríos")
	XVAricaParinacota      = GeoRegion("XV Arica & Parinacota")
	XVINuble               = GeoRegion("XVI Ñuble")
)
