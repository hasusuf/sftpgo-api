package types

import (
	"encoding/json"
	"fmt"
)

// UserType This is an hint for authentication plugins. It is ignored when using SFTPGo internal authentication
type UserType string

// List of UserType
const (
	USERTYPE_EMPTY     UserType = ""
	USERTYPE_LDAP_USER UserType = "LDAPUser"
	USERTYPE_OS_USER   UserType = "OSUser"
)

// AllowedUserTypeEnumValues All allowed values of UserType enum
var AllowedUserTypeEnumValues = []UserType{
	"",
	"LDAPUser",
	"OSUser",
}

func (v *UserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserType(value)
	for _, existing := range AllowedUserTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserType", value)
}

// NewUserTypeFromValue returns a pointer to a valid UserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserTypeFromValue(v string) (*UserType, error) {
	ev := UserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserType: valid values are %v", v, AllowedUserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserType) IsValid() bool {
	for _, existing := range AllowedUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserType value
func (v UserType) Ptr() *UserType {
	return &v
}

type NullableUserType struct {
	value *UserType
	isSet bool
}

func (v NullableUserType) Get() *UserType {
	return v.value
}

func (v *NullableUserType) Set(val *UserType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserType(val *UserType) *NullableUserType {
	return &NullableUserType{value: val, isSet: true}
}

func (v NullableUserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
