// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type FleetsPage struct {
	Fleets []Fleet `json:"fleets"`
}

func (o *FleetsPage) GetFleets() []Fleet {
	if o == nil {
		return []Fleet{}
	}
	return o.Fleets
}
