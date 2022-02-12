package types

import (
	"encoding/json"
	"fmt"
)

// WebClientOptions Options:   * `publickey-change-disabled` - changing SSH public keys is not allowed   * `write-disabled` - upload, rename, delete are not allowed even if the user has permissions for these actions   * `mfa-disabled` - enabling multi-factor authentication is not allowed. This option cannot be set if the user has MFA already enabled   * `password-change-disabled` - changing password is not allowed   * `api-key-auth-change-disabled` - enabling/disabling API key authentication is not allowed   * `info-change-disabled` - changing info such as email and description is not allowed   * `shares-disabled` - sharing files and directories with external users is disabled   * `password-reset-disabled` - resetting the password is disabled
type WebClientOptions string

// List of WebClientOptions
const (
	WEBCLIENTOPTIONS_PUBLICKEY_CHANGE_DISABLED    WebClientOptions = "publickey-change-disabled"
	WEBCLIENTOPTIONS_WRITE_DISABLED               WebClientOptions = "write-disabled"
	WEBCLIENTOPTIONS_MFA_DISABLED                 WebClientOptions = "mfa-disabled"
	WEBCLIENTOPTIONS_PASSWORD_CHANGE_DISABLED     WebClientOptions = "password-change-disabled"
	WEBCLIENTOPTIONS_API_KEY_AUTH_CHANGE_DISABLED WebClientOptions = "api-key-auth-change-disabled"
	WEBCLIENTOPTIONS_INFO_CHANGE_DISABLED         WebClientOptions = "info-change-disabled"
	WEBCLIENTOPTIONS_SHARES_DISABLED              WebClientOptions = "shares-disabled"
	WEBCLIENTOPTIONS_PASSWORD_RESET_DISABLED      WebClientOptions = "password-reset-disabled"
)

// AllowedWebClientOptionsEnumValues All allowed values of WebClientOptions enum
var AllowedWebClientOptionsEnumValues = []WebClientOptions{
	"publickey-change-disabled",
	"write-disabled",
	"mfa-disabled",
	"password-change-disabled",
	"api-key-auth-change-disabled",
	"info-change-disabled",
	"shares-disabled",
	"password-reset-disabled",
}

func (v *WebClientOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebClientOptions(value)
	for _, existing := range AllowedWebClientOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebClientOptions", value)
}

// NewWebClientOptionsFromValue returns a pointer to a valid WebClientOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebClientOptionsFromValue(v string) (*WebClientOptions, error) {
	ev := WebClientOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebClientOptions: valid values are %v", v, AllowedWebClientOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebClientOptions) IsValid() bool {
	for _, existing := range AllowedWebClientOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebClientOptions value
func (v WebClientOptions) Ptr() *WebClientOptions {
	return &v
}

type NullableWebClientOptions struct {
	value *WebClientOptions
	isSet bool
}

func (v NullableWebClientOptions) Get() *WebClientOptions {
	return v.value
}

func (v *NullableWebClientOptions) Set(val *WebClientOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableWebClientOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableWebClientOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebClientOptions(val *WebClientOptions) *NullableWebClientOptions {
	return &NullableWebClientOptions{value: val, isSet: true}
}

func (v NullableWebClientOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebClientOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
