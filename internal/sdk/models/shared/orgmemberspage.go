// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type OrgMembersPage struct {
	Members []OrgMember `json:"members"`
}

func (o *OrgMembersPage) GetMembers() []OrgMember {
	if o == nil {
		return []OrgMember{}
	}
	return o.Members
}
