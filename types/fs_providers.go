package types

import (
	"encoding/json"
	"fmt"
)

// FsProviders Filesystem providers:   * `0` - Local filesystem   * `1` - S3 Compatible Object Storage   * `2` - Google Cloud Storage   * `3` - Azure Blob Storage   * `4` - Local filesystem encrypted   * `5` - SFTP
type FsProviders int32

// List of FsProviders
const (
	FSPROVIDERS__0 FsProviders = 0
	FSPROVIDERS__1 FsProviders = 1
	FSPROVIDERS__2 FsProviders = 2
	FSPROVIDERS__3 FsProviders = 3
	FSPROVIDERS__4 FsProviders = 4
	FSPROVIDERS__5 FsProviders = 5
)

// AllowedFsProvidersEnumValues All allowed values of FsProviders enum
var AllowedFsProvidersEnumValues = []FsProviders{
	0,
	1,
	2,
	3,
	4,
	5,
}

func (v *FsProviders) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FsProviders(value)
	for _, existing := range AllowedFsProvidersEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FsProviders", value)
}

// NewFsProvidersFromValue returns a pointer to a valid FsProviders
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFsProvidersFromValue(v int32) (*FsProviders, error) {
	ev := FsProviders(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FsProviders: valid values are %v", v, AllowedFsProvidersEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FsProviders) IsValid() bool {
	for _, existing := range AllowedFsProvidersEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FsProviders value
func (v FsProviders) Ptr() *FsProviders {
	return &v
}

type NullableFsProviders struct {
	value *FsProviders
	isSet bool
}

func (v NullableFsProviders) Get() *FsProviders {
	return v.value
}

func (v *NullableFsProviders) Set(val *FsProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableFsProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableFsProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFsProviders(val *FsProviders) *NullableFsProviders {
	return &NullableFsProviders{value: val, isSet: true}
}

func (v NullableFsProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFsProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
