package types

import (
	"encoding/json"
	"fmt"
)

// LoginMethods Available login methods. To enable multi-step authentication you have to allow only multi-step login methods   * `publickey`   * `password`   * `keyboard-interactive`   * `publickey+password` - multi-step auth: public key and password   * `publickey+keyboard-interactive` - multi-step auth: public key and keyboard interactive   * `TLSCertificate`   * `TLSCertificate+password` - multi-step auth: TLS client certificate and password
type LoginMethods string

// List of LoginMethods
const (
	LOGINMETHODS_PUBLICKEY                     LoginMethods = "publickey"
	LOGINMETHODS_PASSWORD                      LoginMethods = "password"
	LOGINMETHODS_KEYBOARD_INTERACTIVE          LoginMethods = "keyboard-interactive"
	LOGINMETHODS_PUBLICKEYPASSWORD             LoginMethods = "publickey+password"
	LOGINMETHODS_PUBLICKEYKEYBOARD_INTERACTIVE LoginMethods = "publickey+keyboard-interactive"
	LOGINMETHODS_TLS_CERTIFICATE               LoginMethods = "TLSCertificate"
	LOGINMETHODS_TLS_CERTIFICATEPASSWORD       LoginMethods = "TLSCertificate+password"
)

// AllowedLoginMethodsEnumValues All allowed values of LoginMethods enum
var AllowedLoginMethodsEnumValues = []LoginMethods{
	"publickey",
	"password",
	"keyboard-interactive",
	"publickey+password",
	"publickey+keyboard-interactive",
	"TLSCertificate",
	"TLSCertificate+password",
}

func (v *LoginMethods) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoginMethods(value)
	for _, existing := range AllowedLoginMethodsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoginMethods", value)
}

// NewLoginMethodsFromValue returns a pointer to a valid LoginMethods
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoginMethodsFromValue(v string) (*LoginMethods, error) {
	ev := LoginMethods(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LoginMethods: valid values are %v", v, AllowedLoginMethodsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoginMethods) IsValid() bool {
	for _, existing := range AllowedLoginMethodsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LoginMethods value
func (v LoginMethods) Ptr() *LoginMethods {
	return &v
}

type NullableLoginMethods struct {
	value *LoginMethods
	isSet bool
}

func (v NullableLoginMethods) Get() *LoginMethods {
	return v.value
}

func (v *NullableLoginMethods) Set(val *LoginMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginMethods(val *LoginMethods) *NullableLoginMethods {
	return &NullableLoginMethods{value: val, isSet: true}
}

func (v NullableLoginMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
