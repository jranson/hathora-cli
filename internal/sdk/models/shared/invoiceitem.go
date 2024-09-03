// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type InvoiceItem struct {
	Amount      float64 `json:"amount"`
	UnitPrice   float64 `json:"unitPrice"`
	Quantity    float64 `json:"quantity"`
	Unit        string  `json:"unit"`
	ProductName string  `json:"productName"`
}

func (o *InvoiceItem) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *InvoiceItem) GetUnitPrice() float64 {
	if o == nil {
		return 0.0
	}
	return o.UnitPrice
}

func (o *InvoiceItem) GetQuantity() float64 {
	if o == nil {
		return 0.0
	}
	return o.Quantity
}

func (o *InvoiceItem) GetUnit() string {
	if o == nil {
		return ""
	}
	return o.Unit
}

func (o *InvoiceItem) GetProductName() string {
	if o == nil {
		return ""
	}
	return o.ProductName
}
