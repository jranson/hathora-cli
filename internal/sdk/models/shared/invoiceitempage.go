// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type InvoiceItemPage struct {
	InvoiceItems []InvoiceItem `json:"invoiceItems"`
}

func (o *InvoiceItemPage) GetInvoiceItems() []InvoiceItem {
	if o == nil {
		return []InvoiceItem{}
	}
	return o.InvoiceItems
}
