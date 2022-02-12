package types

import (
	"encoding/json"
)

// ApiResponse struct for ApiResponse
type ApiResponse struct {
	// message, can be empty
	Message *string `json:"message,omitempty"`
	// error description if any
	Error *string `json:"error,omitempty"`
}

// NewApiResponse instantiates a new ApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResponse() *ApiResponse {
	this := ApiResponse{}
	return &this
}

// NewApiResponseWithDefaults instantiates a new ApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResponseWithDefaults() *ApiResponse {
	this := ApiResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiResponse) SetMessage(v string) {
	o.Message = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ApiResponse) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ApiResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ApiResponse) SetError(v string) {
	o.Error = &v
}

func (o ApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableApiResponse struct {
	value *ApiResponse
	isSet bool
}

func (v NullableApiResponse) Get() *ApiResponse {
	return v.value
}

func (v *NullableApiResponse) Set(val *ApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResponse(val *ApiResponse) *NullableApiResponse {
	return &NullableApiResponse{value: val, isSet: true}
}

func (v NullableApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
