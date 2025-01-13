// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetUpcomingInvoiceTotalGlobals struct {
	OrgID *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *GetUpcomingInvoiceTotalGlobals) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

type GetUpcomingInvoiceTotalRequest struct {
	OrgID *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *GetUpcomingInvoiceTotalRequest) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

// GetUpcomingInvoiceTotalResponseBody - Ok
type GetUpcomingInvoiceTotalResponseBody struct {
	Value float64 `json:"value"`
}

func (o *GetUpcomingInvoiceTotalResponseBody) GetValue() float64 {
	if o == nil {
		return 0.0
	}
	return o.Value
}
