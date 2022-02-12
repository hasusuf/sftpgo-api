package types

import (
	"encoding/json"
)

// CryptFsConfig Crypt filesystem configuration details
type CryptFsConfig struct {
	Passphrase *Secret `json:"passphrase,omitempty"`
}

// NewCryptFsConfig instantiates a new CryptFsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptFsConfig() *CryptFsConfig {
	this := CryptFsConfig{}
	return &this
}

// NewCryptFsConfigWithDefaults instantiates a new CryptFsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptFsConfigWithDefaults() *CryptFsConfig {
	this := CryptFsConfig{}
	return &this
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *CryptFsConfig) GetPassphrase() Secret {
	if o == nil || o.Passphrase == nil {
		var ret Secret
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptFsConfig) GetPassphraseOk() (*Secret, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *CryptFsConfig) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given Secret and assigns it to the Passphrase field.
func (o *CryptFsConfig) SetPassphrase(v Secret) {
	o.Passphrase = &v
}

func (o CryptFsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullableCryptFsConfig struct {
	value *CryptFsConfig
	isSet bool
}

func (v NullableCryptFsConfig) Get() *CryptFsConfig {
	return v.value
}

func (v *NullableCryptFsConfig) Set(val *CryptFsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptFsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptFsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptFsConfig(val *CryptFsConfig) *NullableCryptFsConfig {
	return &NullableCryptFsConfig{value: val, isSet: true}
}

func (v NullableCryptFsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptFsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
