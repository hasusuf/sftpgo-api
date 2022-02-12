package types

import (
	"encoding/json"
)

// DataTransferLimit struct for DataTransferLimit
type DataTransferLimit struct {
	// Source networks in CIDR notation as defined in RFC 4632 and RFC 4291 for example `192.0.2.0/24` or `2001:db8::/32`. The limit applies if the defined networks contain the client IP
	Sources []string `json:"sources,omitempty"`
	// Maximum data transfer allowed for uploads as MB. 0 means no limit
	UploadDataTransfer *int32 `json:"upload_data_transfer,omitempty"`
	// Maximum data transfer allowed for downloads as MB. 0 means no limit
	DownloadDataTransfer *int32 `json:"download_data_transfer,omitempty"`
	// Maximum total data transfer as MB. 0 means unlimited. You can set a total data transfer instead of the individual values for uploads and downloads
	TotalDataTransfer *int32 `json:"total_data_transfer,omitempty"`
}

// NewDataTransferLimit instantiates a new DataTransferLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataTransferLimit() *DataTransferLimit {
	this := DataTransferLimit{}
	return &this
}

// NewDataTransferLimitWithDefaults instantiates a new DataTransferLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataTransferLimitWithDefaults() *DataTransferLimit {
	this := DataTransferLimit{}
	return &this
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *DataTransferLimit) GetSources() []string {
	if o == nil || o.Sources == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransferLimit) GetSourcesOk() ([]string, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *DataTransferLimit) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *DataTransferLimit) SetSources(v []string) {
	o.Sources = v
}

// GetUploadDataTransfer returns the UploadDataTransfer field value if set, zero value otherwise.
func (o *DataTransferLimit) GetUploadDataTransfer() int32 {
	if o == nil || o.UploadDataTransfer == nil {
		var ret int32
		return ret
	}
	return *o.UploadDataTransfer
}

// GetUploadDataTransferOk returns a tuple with the UploadDataTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransferLimit) GetUploadDataTransferOk() (*int32, bool) {
	if o == nil || o.UploadDataTransfer == nil {
		return nil, false
	}
	return o.UploadDataTransfer, true
}

// HasUploadDataTransfer returns a boolean if a field has been set.
func (o *DataTransferLimit) HasUploadDataTransfer() bool {
	if o != nil && o.UploadDataTransfer != nil {
		return true
	}

	return false
}

// SetUploadDataTransfer gets a reference to the given int32 and assigns it to the UploadDataTransfer field.
func (o *DataTransferLimit) SetUploadDataTransfer(v int32) {
	o.UploadDataTransfer = &v
}

// GetDownloadDataTransfer returns the DownloadDataTransfer field value if set, zero value otherwise.
func (o *DataTransferLimit) GetDownloadDataTransfer() int32 {
	if o == nil || o.DownloadDataTransfer == nil {
		var ret int32
		return ret
	}
	return *o.DownloadDataTransfer
}

// GetDownloadDataTransferOk returns a tuple with the DownloadDataTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransferLimit) GetDownloadDataTransferOk() (*int32, bool) {
	if o == nil || o.DownloadDataTransfer == nil {
		return nil, false
	}
	return o.DownloadDataTransfer, true
}

// HasDownloadDataTransfer returns a boolean if a field has been set.
func (o *DataTransferLimit) HasDownloadDataTransfer() bool {
	if o != nil && o.DownloadDataTransfer != nil {
		return true
	}

	return false
}

// SetDownloadDataTransfer gets a reference to the given int32 and assigns it to the DownloadDataTransfer field.
func (o *DataTransferLimit) SetDownloadDataTransfer(v int32) {
	o.DownloadDataTransfer = &v
}

// GetTotalDataTransfer returns the TotalDataTransfer field value if set, zero value otherwise.
func (o *DataTransferLimit) GetTotalDataTransfer() int32 {
	if o == nil || o.TotalDataTransfer == nil {
		var ret int32
		return ret
	}
	return *o.TotalDataTransfer
}

// GetTotalDataTransferOk returns a tuple with the TotalDataTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransferLimit) GetTotalDataTransferOk() (*int32, bool) {
	if o == nil || o.TotalDataTransfer == nil {
		return nil, false
	}
	return o.TotalDataTransfer, true
}

// HasTotalDataTransfer returns a boolean if a field has been set.
func (o *DataTransferLimit) HasTotalDataTransfer() bool {
	if o != nil && o.TotalDataTransfer != nil {
		return true
	}

	return false
}

// SetTotalDataTransfer gets a reference to the given int32 and assigns it to the TotalDataTransfer field.
func (o *DataTransferLimit) SetTotalDataTransfer(v int32) {
	o.TotalDataTransfer = &v
}

func (o DataTransferLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.UploadDataTransfer != nil {
		toSerialize["upload_data_transfer"] = o.UploadDataTransfer
	}
	if o.DownloadDataTransfer != nil {
		toSerialize["download_data_transfer"] = o.DownloadDataTransfer
	}
	if o.TotalDataTransfer != nil {
		toSerialize["total_data_transfer"] = o.TotalDataTransfer
	}
	return json.Marshal(toSerialize)
}

type NullableDataTransferLimit struct {
	value *DataTransferLimit
	isSet bool
}

func (v NullableDataTransferLimit) Get() *DataTransferLimit {
	return v.value
}

func (v *NullableDataTransferLimit) Set(val *DataTransferLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTransferLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTransferLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTransferLimit(val *DataTransferLimit) *NullableDataTransferLimit {
	return &NullableDataTransferLimit{value: val, isSet: true}
}

func (v NullableDataTransferLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTransferLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
