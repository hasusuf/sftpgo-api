package types

import (
	"encoding/json"
)

// BandwidthLimit struct for BandwidthLimit
type BandwidthLimit struct {
	// Source networks in CIDR notation as defined in RFC 4632 and RFC 4291 for example `192.0.2.0/24` or `2001:db8::/32`. The limit applies if the defined networks contain the client IP
	Sources []string `json:"sources,omitempty"`
	// Maximum upload bandwidth as KB/s, 0 means unlimited
	UploadBandwidth *int32 `json:"upload_bandwidth,omitempty"`
	// Maximum download bandwidth as KB/s, 0 means unlimited
	DownloadBandwidth *int32 `json:"download_bandwidth,omitempty"`
}

// NewBandwidthLimit instantiates a new BandwidthLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBandwidthLimit() *BandwidthLimit {
	this := BandwidthLimit{}
	return &this
}

// NewBandwidthLimitWithDefaults instantiates a new BandwidthLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBandwidthLimitWithDefaults() *BandwidthLimit {
	this := BandwidthLimit{}
	return &this
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *BandwidthLimit) GetSources() []string {
	if o == nil || o.Sources == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BandwidthLimit) GetSourcesOk() ([]string, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *BandwidthLimit) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *BandwidthLimit) SetSources(v []string) {
	o.Sources = v
}

// GetUploadBandwidth returns the UploadBandwidth field value if set, zero value otherwise.
func (o *BandwidthLimit) GetUploadBandwidth() int32 {
	if o == nil || o.UploadBandwidth == nil {
		var ret int32
		return ret
	}
	return *o.UploadBandwidth
}

// GetUploadBandwidthOk returns a tuple with the UploadBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BandwidthLimit) GetUploadBandwidthOk() (*int32, bool) {
	if o == nil || o.UploadBandwidth == nil {
		return nil, false
	}
	return o.UploadBandwidth, true
}

// HasUploadBandwidth returns a boolean if a field has been set.
func (o *BandwidthLimit) HasUploadBandwidth() bool {
	if o != nil && o.UploadBandwidth != nil {
		return true
	}

	return false
}

// SetUploadBandwidth gets a reference to the given int32 and assigns it to the UploadBandwidth field.
func (o *BandwidthLimit) SetUploadBandwidth(v int32) {
	o.UploadBandwidth = &v
}

// GetDownloadBandwidth returns the DownloadBandwidth field value if set, zero value otherwise.
func (o *BandwidthLimit) GetDownloadBandwidth() int32 {
	if o == nil || o.DownloadBandwidth == nil {
		var ret int32
		return ret
	}
	return *o.DownloadBandwidth
}

// GetDownloadBandwidthOk returns a tuple with the DownloadBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BandwidthLimit) GetDownloadBandwidthOk() (*int32, bool) {
	if o == nil || o.DownloadBandwidth == nil {
		return nil, false
	}
	return o.DownloadBandwidth, true
}

// HasDownloadBandwidth returns a boolean if a field has been set.
func (o *BandwidthLimit) HasDownloadBandwidth() bool {
	if o != nil && o.DownloadBandwidth != nil {
		return true
	}

	return false
}

// SetDownloadBandwidth gets a reference to the given int32 and assigns it to the DownloadBandwidth field.
func (o *BandwidthLimit) SetDownloadBandwidth(v int32) {
	o.DownloadBandwidth = &v
}

func (o BandwidthLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.UploadBandwidth != nil {
		toSerialize["upload_bandwidth"] = o.UploadBandwidth
	}
	if o.DownloadBandwidth != nil {
		toSerialize["download_bandwidth"] = o.DownloadBandwidth
	}
	return json.Marshal(toSerialize)
}

type NullableBandwidthLimit struct {
	value *BandwidthLimit
	isSet bool
}

func (v NullableBandwidthLimit) Get() *BandwidthLimit {
	return v.value
}

func (v *NullableBandwidthLimit) Set(val *BandwidthLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableBandwidthLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableBandwidthLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBandwidthLimit(val *BandwidthLimit) *NullableBandwidthLimit {
	return &NullableBandwidthLimit{value: val, isSet: true}
}

func (v NullableBandwidthLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBandwidthLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
