// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CustomerPortalURL struct {
	ReturnURL string `json:"returnUrl"`
}

func (o *CustomerPortalURL) GetReturnURL() string {
	if o == nil {
		return ""
	}
	return o.ReturnURL
}
