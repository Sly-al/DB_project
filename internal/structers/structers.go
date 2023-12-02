package structers

type Car struct {
	Id           int
	Name         string
	IsNew        string
	Brand        string
	Engine       string
	Color        string
	Transmission string
	Body         string
	Price        float64
}

type Merchant struct {
	Id      int
	Name    string
	Country string
	City    string
	Address string
}

type Supplier struct {
	Id      int
	Brand   string
	Country string
	City    string
	Address string
}

type Equipment struct {
	Id           int
	Engine       string
	Color        string
	Transmission string
	Body         string
}

type Machine struct {
	Id          int
	Name        string
	IsNew       bool
	BrandId     int
	EquipmentId int
}

type Catalog struct {
	Id         int
	MerchantId int
	Price      int
	Sale       int
	ProductId  int
}

type Client struct {
	Id       int
	Login    string
	Password string
	Surname  string
	Name     string
	Status   string
}
