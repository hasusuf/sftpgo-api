package types

import (
	"encoding/json"
	"fmt"
)

// MFAProtocols Protocols:   * `SSH` - includes both SFTP and SSH commands   * `FTP` - plain FTP and FTPES/FTPS   * `HTTP` - WebClient/REST API
type MFAProtocols string

// List of MFAProtocols
const (
	MFAPROTOCOLS_SSH  MFAProtocols = "SSH"
	MFAPROTOCOLS_FTP  MFAProtocols = "FTP"
	MFAPROTOCOLS_HTTP MFAProtocols = "HTTP"
)

// AllowedMFAProtocolsEnumValues All allowed values of MFAProtocols enum
var AllowedMFAProtocolsEnumValues = []MFAProtocols{
	"SSH",
	"FTP",
	"HTTP",
}

func (v *MFAProtocols) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MFAProtocols(value)
	for _, existing := range AllowedMFAProtocolsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MFAProtocols", value)
}

// NewMFAProtocolsFromValue returns a pointer to a valid MFAProtocols
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMFAProtocolsFromValue(v string) (*MFAProtocols, error) {
	ev := MFAProtocols(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MFAProtocols: valid values are %v", v, AllowedMFAProtocolsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MFAProtocols) IsValid() bool {
	for _, existing := range AllowedMFAProtocolsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MFAProtocols value
func (v MFAProtocols) Ptr() *MFAProtocols {
	return &v
}

type NullableMFAProtocols struct {
	value *MFAProtocols
	isSet bool
}

func (v NullableMFAProtocols) Get() *MFAProtocols {
	return v.value
}

func (v *NullableMFAProtocols) Set(val *MFAProtocols) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAProtocols) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAProtocols) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAProtocols(val *MFAProtocols) *NullableMFAProtocols {
	return &NullableMFAProtocols{value: val, isSet: true}
}

func (v NullableMFAProtocols) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAProtocols) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
