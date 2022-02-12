package types

import "encoding/json"

// RecoveryCode Recovery codes to use if the user loses access to their second factor auth device. Each code can only be used once, you should use these codes to login and disable or reset 2FA for your account
type RecoveryCode struct {
	Secret *Secret `json:"secret,omitempty"`
	Used   *bool   `json:"used,omitempty"`
}

// NewRecoveryCode instantiates a new RecoveryCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryCode() *RecoveryCode {
	this := RecoveryCode{}
	return &this
}

// NewRecoveryCodeWithDefaults instantiates a new RecoveryCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryCodeWithDefaults() *RecoveryCode {
	this := RecoveryCode{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *RecoveryCode) GetSecret() Secret {
	if o == nil || o.Secret == nil {
		var ret Secret
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryCode) GetSecretOk() (*Secret, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *RecoveryCode) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given Secret and assigns it to the Secret field.
func (o *RecoveryCode) SetSecret(v Secret) {
	o.Secret = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *RecoveryCode) GetUsed() bool {
	if o == nil || o.Used == nil {
		var ret bool
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryCode) GetUsedOk() (*bool, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *RecoveryCode) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given bool and assigns it to the Used field.
func (o *RecoveryCode) SetUsed(v bool) {
	o.Used = &v
}

func (o RecoveryCode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Used != nil {
		toSerialize["used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableRecoveryCode struct {
	value *RecoveryCode
	isSet bool
}

func (v NullableRecoveryCode) Get() *RecoveryCode {
	return v.value
}

func (v *NullableRecoveryCode) Set(val *RecoveryCode) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryCode) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryCode(val *RecoveryCode) *NullableRecoveryCode {
	return &NullableRecoveryCode{value: val, isSet: true}
}

func (v NullableRecoveryCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
