package types

import (
	"encoding/json"
)

// HooksFilter User specific hook overrides
type HooksFilter struct {
	// If true, the external auth hook, if defined, will not be executed
	ExternalAuthDisabled *bool `json:"external_auth_disabled,omitempty"`
	// If true, the pre-login hook, if defined, will not be executed
	PreLoginDisabled *bool `json:"pre_login_disabled,omitempty"`
	// If true, the check password hook, if defined, will not be executed
	CheckPasswordDisabled *bool `json:"check_password_disabled,omitempty"`
}

// NewHooksFilter instantiates a new HooksFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHooksFilter() *HooksFilter {
	this := HooksFilter{}
	return &this
}

// NewHooksFilterWithDefaults instantiates a new HooksFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHooksFilterWithDefaults() *HooksFilter {
	this := HooksFilter{}
	return &this
}

// GetExternalAuthDisabled returns the ExternalAuthDisabled field value if set, zero value otherwise.
func (o *HooksFilter) GetExternalAuthDisabled() bool {
	if o == nil || o.ExternalAuthDisabled == nil {
		var ret bool
		return ret
	}
	return *o.ExternalAuthDisabled
}

// GetExternalAuthDisabledOk returns a tuple with the ExternalAuthDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HooksFilter) GetExternalAuthDisabledOk() (*bool, bool) {
	if o == nil || o.ExternalAuthDisabled == nil {
		return nil, false
	}
	return o.ExternalAuthDisabled, true
}

// HasExternalAuthDisabled returns a boolean if a field has been set.
func (o *HooksFilter) HasExternalAuthDisabled() bool {
	if o != nil && o.ExternalAuthDisabled != nil {
		return true
	}

	return false
}

// SetExternalAuthDisabled gets a reference to the given bool and assigns it to the ExternalAuthDisabled field.
func (o *HooksFilter) SetExternalAuthDisabled(v bool) {
	o.ExternalAuthDisabled = &v
}

// GetPreLoginDisabled returns the PreLoginDisabled field value if set, zero value otherwise.
func (o *HooksFilter) GetPreLoginDisabled() bool {
	if o == nil || o.PreLoginDisabled == nil {
		var ret bool
		return ret
	}
	return *o.PreLoginDisabled
}

// GetPreLoginDisabledOk returns a tuple with the PreLoginDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HooksFilter) GetPreLoginDisabledOk() (*bool, bool) {
	if o == nil || o.PreLoginDisabled == nil {
		return nil, false
	}
	return o.PreLoginDisabled, true
}

// HasPreLoginDisabled returns a boolean if a field has been set.
func (o *HooksFilter) HasPreLoginDisabled() bool {
	if o != nil && o.PreLoginDisabled != nil {
		return true
	}

	return false
}

// SetPreLoginDisabled gets a reference to the given bool and assigns it to the PreLoginDisabled field.
func (o *HooksFilter) SetPreLoginDisabled(v bool) {
	o.PreLoginDisabled = &v
}

// GetCheckPasswordDisabled returns the CheckPasswordDisabled field value if set, zero value otherwise.
func (o *HooksFilter) GetCheckPasswordDisabled() bool {
	if o == nil || o.CheckPasswordDisabled == nil {
		var ret bool
		return ret
	}
	return *o.CheckPasswordDisabled
}

// GetCheckPasswordDisabledOk returns a tuple with the CheckPasswordDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HooksFilter) GetCheckPasswordDisabledOk() (*bool, bool) {
	if o == nil || o.CheckPasswordDisabled == nil {
		return nil, false
	}
	return o.CheckPasswordDisabled, true
}

// HasCheckPasswordDisabled returns a boolean if a field has been set.
func (o *HooksFilter) HasCheckPasswordDisabled() bool {
	if o != nil && o.CheckPasswordDisabled != nil {
		return true
	}

	return false
}

// SetCheckPasswordDisabled gets a reference to the given bool and assigns it to the CheckPasswordDisabled field.
func (o *HooksFilter) SetCheckPasswordDisabled(v bool) {
	o.CheckPasswordDisabled = &v
}

func (o HooksFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalAuthDisabled != nil {
		toSerialize["external_auth_disabled"] = o.ExternalAuthDisabled
	}
	if o.PreLoginDisabled != nil {
		toSerialize["pre_login_disabled"] = o.PreLoginDisabled
	}
	if o.CheckPasswordDisabled != nil {
		toSerialize["check_password_disabled"] = o.CheckPasswordDisabled
	}
	return json.Marshal(toSerialize)
}

type NullableHooksFilter struct {
	value *HooksFilter
	isSet bool
}

func (v NullableHooksFilter) Get() *HooksFilter {
	return v.value
}

func (v *NullableHooksFilter) Set(val *HooksFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableHooksFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableHooksFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHooksFilter(val *HooksFilter) *NullableHooksFilter {
	return &NullableHooksFilter{value: val, isSet: true}
}

func (v NullableHooksFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHooksFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
