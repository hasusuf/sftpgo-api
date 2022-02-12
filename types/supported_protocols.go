package types

import (
	"encoding/json"
	"fmt"
)

// SupportedProtocols Protocols:   * `SSH` - includes both SFTP and SSH commands   * `FTP` - plain FTP and FTPES/FTPS   * `DAV` - WebDAV over HTTP/HTTPS   * `HTTP` - WebClient/REST API
type SupportedProtocols string

// List of SupportedProtocols
const (
	SUPPORTEDPROTOCOLS_SSH  SupportedProtocols = "SSH"
	SUPPORTEDPROTOCOLS_FTP  SupportedProtocols = "FTP"
	SUPPORTEDPROTOCOLS_DAV  SupportedProtocols = "DAV"
	SUPPORTEDPROTOCOLS_HTTP SupportedProtocols = "HTTP"
)

// AllowedSupportedProtocolsEnumValues All allowed values of SupportedProtocols enum
var AllowedSupportedProtocolsEnumValues = []SupportedProtocols{
	"SSH",
	"FTP",
	"DAV",
	"HTTP",
}

func (v *SupportedProtocols) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupportedProtocols(value)
	for _, existing := range AllowedSupportedProtocolsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupportedProtocols", value)
}

// NewSupportedProtocolsFromValue returns a pointer to a valid SupportedProtocols
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupportedProtocolsFromValue(v string) (*SupportedProtocols, error) {
	ev := SupportedProtocols(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupportedProtocols: valid values are %v", v, AllowedSupportedProtocolsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupportedProtocols) IsValid() bool {
	for _, existing := range AllowedSupportedProtocolsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SupportedProtocols value
func (v SupportedProtocols) Ptr() *SupportedProtocols {
	return &v
}

type NullableSupportedProtocols struct {
	value *SupportedProtocols
	isSet bool
}

func (v NullableSupportedProtocols) Get() *SupportedProtocols {
	return v.value
}

func (v *NullableSupportedProtocols) Set(val *SupportedProtocols) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedProtocols) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedProtocols) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedProtocols(val *SupportedProtocols) *NullableSupportedProtocols {
	return &NullableSupportedProtocols{value: val, isSet: true}
}

func (v NullableSupportedProtocols) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedProtocols) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
