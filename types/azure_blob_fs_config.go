package types

import (
	"encoding/json"
)

// AzureBlobFsConfig Azure Blob Storage configuration details
type AzureBlobFsConfig struct {
	Container *string `json:"container,omitempty"`
	// Storage Account Name, leave blank to use SAS URL
	AccountName *string `json:"account_name,omitempty"`
	AccountKey  *Secret `json:"account_key,omitempty"`
	SasUrl      *Secret `json:"sas_url,omitempty"`
	// optional endpoint. Default is \"blob.core.windows.net\". If you use the emulator the endpoint must include the protocol, for example \"http://127.0.0.1:10000\"
	Endpoint *string `json:"endpoint,omitempty"`
	// the buffer size (in MB) to use for multipart uploads. If this value is set to zero, the default value (4MB) will be used.
	UploadPartSize *int32 `json:"upload_part_size,omitempty"`
	// the number of parts to upload in parallel. If this value is set to zero, the default value (2) will be used
	UploadConcurrency *int32  `json:"upload_concurrency,omitempty"`
	AccessTier        *string `json:"access_tier,omitempty"`
	// key_prefix is similar to a chroot directory for a local filesystem. If specified the user will only see contents that starts with this prefix and so you can restrict access to a specific virtual folder. The prefix, if not empty, must not start with \"/\" and must end with \"/\". If empty the whole container contents will be available
	KeyPrefix   *string `json:"key_prefix,omitempty"`
	UseEmulator *bool   `json:"use_emulator,omitempty"`
}

// NewAzureBlobFsConfig instantiates a new AzureBlobFsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobFsConfig() *AzureBlobFsConfig {
	this := AzureBlobFsConfig{}
	return &this
}

// NewAzureBlobFsConfigWithDefaults instantiates a new AzureBlobFsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobFsConfigWithDefaults() *AzureBlobFsConfig {
	this := AzureBlobFsConfig{}
	return &this
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *AzureBlobFsConfig) GetContainer() string {
	if o == nil || o.Container == nil {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobFsConfig) GetContainerOk() (*string, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *AzureBlobFsConfig) HasContainer() bool {
	if o != nil && o.Container != nil {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *AzureBlobFsConfig) SetContainer(v string) {
	o.Container = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *AzureBlobFsConfig) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobFsConfig) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *AzureBlobFsConfig) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *AzureBlobFsConfig) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountKey returns the AccountKey field value if set, zero value otherwise.
func (o *AzureBlobFsConfig) GetAccountKey() Secret {
	if o == nil || o.AccountKey == nil {
		var ret Secret
		return ret
	}
	return *o.AccountKey
}

// GetAccountKeyOk returns a tuple with the AccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobFsConfig) GetAccountKeyOk() (*Secret, bool) {
	if o == nil || o.AccountKey == nil {
		return nil, false
	}
	return o.AccountKey, true
}

// HasAccountKey returns a boolean if a field has been set.
func (o *AzureBlobFsConfig) HasAccountKey() bool {
	if o != nil && o.AccountKey != nil {
		return true
	}

	return false
}

// SetAccountKey gets a reference to the given Secret and assigns it to the AccountKey field.
func (o *AzureBlobFsConfig) SetAccountKey(v Secret) {
	o.AccountKey = &v
}

// GetSasUrl returns the SasUrl field value if set, zero value otherwise.
func (o *AzureBlobFsConfig) GetSasUrl() Secret {
	if o == nil || o.SasUrl == nil {
		var ret Secret
		return ret
	}
	return *o.SasUrl
}

// GetSasUrlOk returns a tuple with the SasUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobFsConfig) GetSasUrlOk() (*Secret, bool) {
	if o == nil || o.SasUrl == nil {
		return nil, false
	}
	return o.SasUrl, true
}

// HasSasUrl returns a boolean if a field has been set.
func (o *AzureBlobFsConfig) HasSasUrl() bool {
	if o != nil && o.SasUrl != nil {
		return true
	}

	return false
}

// SetSasUrl gets a reference to the given Secret and assigns it to the SasUrl field.
func (o *AzureBlobFsConfig) SetSasUrl(v Secret) {
	o.SasUrl = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *AzureBlobFsConfig) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobFsConfig) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *AzureBlobFsConfig) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *AzureBlobFsConfig) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetUploadPartSize returns the UploadPartSize field value if set, zero value otherwise.
func (o *AzureBlobFsConfig) GetUploadPartSize() int32 {
	if o == nil || o.UploadPartSize == nil {
		var ret int32
		return ret
	}
	return *o.UploadPartSize
}

// GetUploadPartSizeOk returns a tuple with the UploadPartSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobFsConfig) GetUploadPartSizeOk() (*int32, bool) {
	if o == nil || o.UploadPartSize == nil {
		return nil, false
	}
	return o.UploadPartSize, true
}

// HasUploadPartSize returns a boolean if a field has been set.
func (o *AzureBlobFsConfig) HasUploadPartSize() bool {
	if o != nil && o.UploadPartSize != nil {
		return true
	}

	return false
}

// SetUploadPartSize gets a reference to the given int32 and assigns it to the UploadPartSize field.
func (o *AzureBlobFsConfig) SetUploadPartSize(v int32) {
	o.UploadPartSize = &v
}

// GetUploadConcurrency returns the UploadConcurrency field value if set, zero value otherwise.
func (o *AzureBlobFsConfig) GetUploadConcurrency() int32 {
	if o == nil || o.UploadConcurrency == nil {
		var ret int32
		return ret
	}
	return *o.UploadConcurrency
}

// GetUploadConcurrencyOk returns a tuple with the UploadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobFsConfig) GetUploadConcurrencyOk() (*int32, bool) {
	if o == nil || o.UploadConcurrency == nil {
		return nil, false
	}
	return o.UploadConcurrency, true
}

// HasUploadConcurrency returns a boolean if a field has been set.
func (o *AzureBlobFsConfig) HasUploadConcurrency() bool {
	if o != nil && o.UploadConcurrency != nil {
		return true
	}

	return false
}

// SetUploadConcurrency gets a reference to the given int32 and assigns it to the UploadConcurrency field.
func (o *AzureBlobFsConfig) SetUploadConcurrency(v int32) {
	o.UploadConcurrency = &v
}

// GetAccessTier returns the AccessTier field value if set, zero value otherwise.
func (o *AzureBlobFsConfig) GetAccessTier() string {
	if o == nil || o.AccessTier == nil {
		var ret string
		return ret
	}
	return *o.AccessTier
}

// GetAccessTierOk returns a tuple with the AccessTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobFsConfig) GetAccessTierOk() (*string, bool) {
	if o == nil || o.AccessTier == nil {
		return nil, false
	}
	return o.AccessTier, true
}

// HasAccessTier returns a boolean if a field has been set.
func (o *AzureBlobFsConfig) HasAccessTier() bool {
	if o != nil && o.AccessTier != nil {
		return true
	}

	return false
}

// SetAccessTier gets a reference to the given string and assigns it to the AccessTier field.
func (o *AzureBlobFsConfig) SetAccessTier(v string) {
	o.AccessTier = &v
}

// GetKeyPrefix returns the KeyPrefix field value if set, zero value otherwise.
func (o *AzureBlobFsConfig) GetKeyPrefix() string {
	if o == nil || o.KeyPrefix == nil {
		var ret string
		return ret
	}
	return *o.KeyPrefix
}

// GetKeyPrefixOk returns a tuple with the KeyPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobFsConfig) GetKeyPrefixOk() (*string, bool) {
	if o == nil || o.KeyPrefix == nil {
		return nil, false
	}
	return o.KeyPrefix, true
}

// HasKeyPrefix returns a boolean if a field has been set.
func (o *AzureBlobFsConfig) HasKeyPrefix() bool {
	if o != nil && o.KeyPrefix != nil {
		return true
	}

	return false
}

// SetKeyPrefix gets a reference to the given string and assigns it to the KeyPrefix field.
func (o *AzureBlobFsConfig) SetKeyPrefix(v string) {
	o.KeyPrefix = &v
}

// GetUseEmulator returns the UseEmulator field value if set, zero value otherwise.
func (o *AzureBlobFsConfig) GetUseEmulator() bool {
	if o == nil || o.UseEmulator == nil {
		var ret bool
		return ret
	}
	return *o.UseEmulator
}

// GetUseEmulatorOk returns a tuple with the UseEmulator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobFsConfig) GetUseEmulatorOk() (*bool, bool) {
	if o == nil || o.UseEmulator == nil {
		return nil, false
	}
	return o.UseEmulator, true
}

// HasUseEmulator returns a boolean if a field has been set.
func (o *AzureBlobFsConfig) HasUseEmulator() bool {
	if o != nil && o.UseEmulator != nil {
		return true
	}

	return false
}

// SetUseEmulator gets a reference to the given bool and assigns it to the UseEmulator field.
func (o *AzureBlobFsConfig) SetUseEmulator(v bool) {
	o.UseEmulator = &v
}

func (o AzureBlobFsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Container != nil {
		toSerialize["container"] = o.Container
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.AccountKey != nil {
		toSerialize["account_key"] = o.AccountKey
	}
	if o.SasUrl != nil {
		toSerialize["sas_url"] = o.SasUrl
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.UploadPartSize != nil {
		toSerialize["upload_part_size"] = o.UploadPartSize
	}
	if o.UploadConcurrency != nil {
		toSerialize["upload_concurrency"] = o.UploadConcurrency
	}
	if o.AccessTier != nil {
		toSerialize["access_tier"] = o.AccessTier
	}
	if o.KeyPrefix != nil {
		toSerialize["key_prefix"] = o.KeyPrefix
	}
	if o.UseEmulator != nil {
		toSerialize["use_emulator"] = o.UseEmulator
	}
	return json.Marshal(toSerialize)
}

type NullableAzureBlobFsConfig struct {
	value *AzureBlobFsConfig
	isSet bool
}

func (v NullableAzureBlobFsConfig) Get() *AzureBlobFsConfig {
	return v.value
}

func (v *NullableAzureBlobFsConfig) Set(val *AzureBlobFsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobFsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobFsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobFsConfig(val *AzureBlobFsConfig) *NullableAzureBlobFsConfig {
	return &NullableAzureBlobFsConfig{value: val, isSet: true}
}

func (v NullableAzureBlobFsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobFsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
