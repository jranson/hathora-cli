// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/hathora/ci/internal/sdk/internal/utils"
)

type UpdateUserInviteScopesType string

const (
	UpdateUserInviteScopesTypeUserRole     UpdateUserInviteScopesType = "UserRole"
	UpdateUserInviteScopesTypeArrayOfScope UpdateUserInviteScopesType = "arrayOfScope"
)

// UpdateUserInviteScopes - Scopes can only be removed or added if a user has those scopes.
type UpdateUserInviteScopes struct {
	UserRole     *UserRole `queryParam:"inline"`
	ArrayOfScope []Scope   `queryParam:"inline"`

	Type UpdateUserInviteScopesType
}

func CreateUpdateUserInviteScopesUserRole(userRole UserRole) UpdateUserInviteScopes {
	typ := UpdateUserInviteScopesTypeUserRole

	return UpdateUserInviteScopes{
		UserRole: &userRole,
		Type:     typ,
	}
}

func CreateUpdateUserInviteScopesArrayOfScope(arrayOfScope []Scope) UpdateUserInviteScopes {
	typ := UpdateUserInviteScopesTypeArrayOfScope

	return UpdateUserInviteScopes{
		ArrayOfScope: arrayOfScope,
		Type:         typ,
	}
}

func (u *UpdateUserInviteScopes) UnmarshalJSON(data []byte) error {

	var userRole UserRole = UserRole("")
	if err := utils.UnmarshalJSON(data, &userRole, "", true, true); err == nil {
		u.UserRole = &userRole
		u.Type = UpdateUserInviteScopesTypeUserRole
		return nil
	}

	var arrayOfScope []Scope = []Scope{}
	if err := utils.UnmarshalJSON(data, &arrayOfScope, "", true, true); err == nil {
		u.ArrayOfScope = arrayOfScope
		u.Type = UpdateUserInviteScopesTypeArrayOfScope
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpdateUserInviteScopes", string(data))
}

func (u UpdateUserInviteScopes) MarshalJSON() ([]byte, error) {
	if u.UserRole != nil {
		return utils.MarshalJSON(u.UserRole, "", true)
	}

	if u.ArrayOfScope != nil {
		return utils.MarshalJSON(u.ArrayOfScope, "", true)
	}

	return nil, errors.New("could not marshal union type UpdateUserInviteScopes: all fields are null")
}

type UpdateUserInvite struct {
	// Scopes can only be removed or added if a user has those scopes.
	Scopes UpdateUserInviteScopes `json:"scopes"`
	// A user's email.
	UserEmail string `json:"userEmail"`
}

func (o *UpdateUserInvite) GetScopes() UpdateUserInviteScopes {
	if o == nil {
		return UpdateUserInviteScopes{}
	}
	return o.Scopes
}

func (o *UpdateUserInvite) GetUserEmail() string {
	if o == nil {
		return ""
	}
	return o.UserEmail
}
