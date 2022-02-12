package types

import (
	"encoding/json"
)

// PatternsFilter struct for PatternsFilter
type PatternsFilter struct {
	// exposed virtual path, if no other specific filter is defined, the filter applies for sub directories too. For example if filters are defined for the paths \"/\" and \"/sub\" then the filters for \"/\" are applied for any file outside the \"/sub\" directory
	Path *string `json:"path,omitempty"`
	// list of, case insensitive, allowed shell like patterns.
	AllowedPatterns []string `json:"allowed_patterns,omitempty"`
	// list of, case insensitive, denied shell like patterns. Denied patterns are evaluated before the allowed ones
	DeniedPatterns []string `json:"denied_patterns,omitempty"`
	// Deny policies   * `0` - default policy. Denied files/directories matching the filters are visible in directory listing but cannot be uploaded/downloaded/overwritten/renamed   * `1` - deny policy hide. This policy applies the same restrictions as the default one and denied files/directories matching the filters will also be hidden in directory listing. This mode may cause performance issues for large directories
	DenyPolicy *int32 `json:"deny_policy,omitempty"`
}

// NewPatternsFilter instantiates a new PatternsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatternsFilter() *PatternsFilter {
	this := PatternsFilter{}
	return &this
}

// NewPatternsFilterWithDefaults instantiates a new PatternsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatternsFilterWithDefaults() *PatternsFilter {
	this := PatternsFilter{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PatternsFilter) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatternsFilter) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PatternsFilter) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PatternsFilter) SetPath(v string) {
	o.Path = &v
}

// GetAllowedPatterns returns the AllowedPatterns field value if set, zero value otherwise.
func (o *PatternsFilter) GetAllowedPatterns() []string {
	if o == nil || o.AllowedPatterns == nil {
		var ret []string
		return ret
	}
	return o.AllowedPatterns
}

// GetAllowedPatternsOk returns a tuple with the AllowedPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatternsFilter) GetAllowedPatternsOk() ([]string, bool) {
	if o == nil || o.AllowedPatterns == nil {
		return nil, false
	}
	return o.AllowedPatterns, true
}

// HasAllowedPatterns returns a boolean if a field has been set.
func (o *PatternsFilter) HasAllowedPatterns() bool {
	if o != nil && o.AllowedPatterns != nil {
		return true
	}

	return false
}

// SetAllowedPatterns gets a reference to the given []string and assigns it to the AllowedPatterns field.
func (o *PatternsFilter) SetAllowedPatterns(v []string) {
	o.AllowedPatterns = v
}

// GetDeniedPatterns returns the DeniedPatterns field value if set, zero value otherwise.
func (o *PatternsFilter) GetDeniedPatterns() []string {
	if o == nil || o.DeniedPatterns == nil {
		var ret []string
		return ret
	}
	return o.DeniedPatterns
}

// GetDeniedPatternsOk returns a tuple with the DeniedPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatternsFilter) GetDeniedPatternsOk() ([]string, bool) {
	if o == nil || o.DeniedPatterns == nil {
		return nil, false
	}
	return o.DeniedPatterns, true
}

// HasDeniedPatterns returns a boolean if a field has been set.
func (o *PatternsFilter) HasDeniedPatterns() bool {
	if o != nil && o.DeniedPatterns != nil {
		return true
	}

	return false
}

// SetDeniedPatterns gets a reference to the given []string and assigns it to the DeniedPatterns field.
func (o *PatternsFilter) SetDeniedPatterns(v []string) {
	o.DeniedPatterns = v
}

// GetDenyPolicy returns the DenyPolicy field value if set, zero value otherwise.
func (o *PatternsFilter) GetDenyPolicy() int32 {
	if o == nil || o.DenyPolicy == nil {
		var ret int32
		return ret
	}
	return *o.DenyPolicy
}

// GetDenyPolicyOk returns a tuple with the DenyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatternsFilter) GetDenyPolicyOk() (*int32, bool) {
	if o == nil || o.DenyPolicy == nil {
		return nil, false
	}
	return o.DenyPolicy, true
}

// HasDenyPolicy returns a boolean if a field has been set.
func (o *PatternsFilter) HasDenyPolicy() bool {
	if o != nil && o.DenyPolicy != nil {
		return true
	}

	return false
}

// SetDenyPolicy gets a reference to the given int32 and assigns it to the DenyPolicy field.
func (o *PatternsFilter) SetDenyPolicy(v int32) {
	o.DenyPolicy = &v
}

func (o PatternsFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.AllowedPatterns != nil {
		toSerialize["allowed_patterns"] = o.AllowedPatterns
	}
	if o.DeniedPatterns != nil {
		toSerialize["denied_patterns"] = o.DeniedPatterns
	}
	if o.DenyPolicy != nil {
		toSerialize["deny_policy"] = o.DenyPolicy
	}
	return json.Marshal(toSerialize)
}

type NullablePatternsFilter struct {
	value *PatternsFilter
	isSet bool
}

func (v NullablePatternsFilter) Get() *PatternsFilter {
	return v.value
}

func (v *NullablePatternsFilter) Set(val *PatternsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePatternsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePatternsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatternsFilter(val *PatternsFilter) *NullablePatternsFilter {
	return &NullablePatternsFilter{value: val, isSet: true}
}

func (v NullablePatternsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatternsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
