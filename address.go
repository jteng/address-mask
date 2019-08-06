package addressmask

type Address struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	AddressLine3 string `json:"addressLine3"`
	AddressLine4 string `json:"addressLine4"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zipCode"`
	Country      string `json:"country"`
}
