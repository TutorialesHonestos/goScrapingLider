package web

type ResponseLider struct {
	Products []Products `json:"products"`
}
type Price struct {
	BasePriceReference int `json:"BasePriceReference"`
	BasePriceSales     int `json:"BasePriceSales"`
	BasePriceTLMC      int `json:"BasePriceTLMC"`
}
type Filters struct {
	DescripciN                     string `json:"Descripción,omitempty"`
	Producto                       string `json:"Producto,omitempty"`
	Marca                          string `json:"Marca,omitempty"`
	Modelo                         string `json:"Modelo,omitempty"`
	TipoDePantalla                 string `json:"Tipo de Pantalla,omitempty"`
	TamaODePantalla                string `json:"Tamaño de Pantalla,omitempty"`
	DefiniciN                      string `json:"Definición,omitempty"`
	ResoluciN                      string `json:"Resolución,omitempty"`
	TecnologA                      string `json:"Tecnología,omitempty"`
	TasaDeRefresco                 string `json:"Tasa de Refresco,omitempty"`
	Hdmi                           string `json:"HDMI,omitempty"`
	Usb                            string `json:"USB,omitempty"`
	AltoXAnchoXProfundidadSinBase  string `json:"Alto x Ancho x Profundidad (Sin Base),omitempty"`
	AltoXAnchoXProfundidadConBase  string `json:"Alto x Ancho x Profundidad (Con Base),omitempty"`
	Peso                           string `json:"Peso,omitempty"`
	PaSDeOrigen                    string `json:"País de origen,omitempty"`
	RazNSocialDelProveedor         string `json:"Razón Social del Proveedor,omitempty"`
	DomicilioDelProveedor          string `json:"Domicilio del Proveedor,omitempty"`
	GarantAOtorgadaPorElFabricante string `json:"Garantía otorgada por el Fabricante,omitempty"`
	CategorA                       string `json:"Categoría,omitempty"`
}
type Products struct {
	ID                  string        `json:"ID"`
	Brand               string        `json:"brand"`
	Destacado           bool          `json:"destacado"`
	DisplayName         string        `json:"displayName"`
	Gtin13              string        `json:"gtin13"`
	ItemNumber          string        `json:"itemNumber"`
	Keyword             []interface{} `json:"keyword"`
	LeadTime            int           `json:"leadTime"`
	Max                 int           `json:"max"`
	Posicion            int           `json:"posicion"`
	Sku                 int           `json:"sku"`
	VendorID            interface{}   `json:"vendorId"`
	Size                int           `json:"size"`
	Categorias          []string      `json:"categorias"`
	Price               Price         `json:"price"`
	Discount            int           `json:"discount"`
	Filters             []Filters     `json:"filters"`
	MakePublic          bool          `json:"makePublic"`
	Etiqueta            string        `json:"etiqueta"`
	Tags                []interface{} `json:"tags"`
	CampaignID          string        `json:"campaignId"`
	AppID               string        `json:"appId"`
	NotAllowedStoresHD  []int         `json:"notAllowedStoresHD"`
	NotAllowedStoresS2S []interface{} `json:"notAllowedStoresS2S"`
	AllowedStoresHD     []int         `json:"allowedStoresHD"`
	AllowedStoresS2S    []int         `json:"allowedStoresS2S"`
	StockInicial        interface{}   `json:"stockInicial"`
	Available           bool          `json:"available"`
	ObjectID            string        `json:"objectID"`
	Checked             bool          `json:"checked,omitempty"`
	InStore             bool          `json:"inStore,omitempty"`
}
