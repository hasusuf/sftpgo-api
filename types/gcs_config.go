package types

import (
	"encoding/json"
)

// GCSConfig Google Cloud Storage configuration details. The \"credentials\" field must be populated only when adding/updating a user. It will be always omitted, since there are sensitive data, when you search/get users
type GCSConfig struct {
	Bucket      *string `json:"bucket,omitempty"`
	Credentials *Secret `json:"credentials,omitempty"`
	// Automatic credentials:   * `0` - disabled, explicit credentials, using a JSON credentials file, must be provided. This is the default value if the field is null   * `1` - enabled, we try to use the Application Default Credentials (ADC) strategy to find your application's credentials
	AutomaticCredentials *int32  `json:"automatic_credentials,omitempty"`
	StorageClass         *string `json:"storage_class,omitempty"`
	// The ACL to apply to uploaded objects. Leave empty to use the default ACL. For more information and available ACLs, refer to the JSON API here: https://cloud.google.com/storage/docs/access-control/lists#predefined-acl
	Acl *string `json:"acl,omitempty"`
	// key_prefix is similar to a chroot directory for a local filesystem. If specified the user will only see contents that starts with this prefix and so you can restrict access to a specific virtual folder. The prefix, if not empty, must not start with \"/\" and must end with \"/\". If empty the whole bucket contents will be available
	KeyPrefix *string `json:"key_prefix,omitempty"`
}

// NewGCSConfig instantiates a new GCSConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCSConfig() *GCSConfig {
	this := GCSConfig{}
	return &this
}

// NewGCSConfigWithDefaults instantiates a new GCSConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCSConfigWithDefaults() *GCSConfig {
	this := GCSConfig{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *GCSConfig) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSConfig) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *GCSConfig) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *GCSConfig) SetBucket(v string) {
	o.Bucket = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *GCSConfig) GetCredentials() Secret {
	if o == nil || o.Credentials == nil {
		var ret Secret
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSConfig) GetCredentialsOk() (*Secret, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *GCSConfig) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given Secret and assigns it to the Credentials field.
func (o *GCSConfig) SetCredentials(v Secret) {
	o.Credentials = &v
}

// GetAutomaticCredentials returns the AutomaticCredentials field value if set, zero value otherwise.
func (o *GCSConfig) GetAutomaticCredentials() int32 {
	if o == nil || o.AutomaticCredentials == nil {
		var ret int32
		return ret
	}
	return *o.AutomaticCredentials
}

// GetAutomaticCredentialsOk returns a tuple with the AutomaticCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSConfig) GetAutomaticCredentialsOk() (*int32, bool) {
	if o == nil || o.AutomaticCredentials == nil {
		return nil, false
	}
	return o.AutomaticCredentials, true
}

// HasAutomaticCredentials returns a boolean if a field has been set.
func (o *GCSConfig) HasAutomaticCredentials() bool {
	if o != nil && o.AutomaticCredentials != nil {
		return true
	}

	return false
}

// SetAutomaticCredentials gets a reference to the given int32 and assigns it to the AutomaticCredentials field.
func (o *GCSConfig) SetAutomaticCredentials(v int32) {
	o.AutomaticCredentials = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *GCSConfig) GetStorageClass() string {
	if o == nil || o.StorageClass == nil {
		var ret string
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSConfig) GetStorageClassOk() (*string, bool) {
	if o == nil || o.StorageClass == nil {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *GCSConfig) HasStorageClass() bool {
	if o != nil && o.StorageClass != nil {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given string and assigns it to the StorageClass field.
func (o *GCSConfig) SetStorageClass(v string) {
	o.StorageClass = &v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *GCSConfig) GetAcl() string {
	if o == nil || o.Acl == nil {
		var ret string
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSConfig) GetAclOk() (*string, bool) {
	if o == nil || o.Acl == nil {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *GCSConfig) HasAcl() bool {
	if o != nil && o.Acl != nil {
		return true
	}

	return false
}

// SetAcl gets a reference to the given string and assigns it to the Acl field.
func (o *GCSConfig) SetAcl(v string) {
	o.Acl = &v
}

// GetKeyPrefix returns the KeyPrefix field value if set, zero value otherwise.
func (o *GCSConfig) GetKeyPrefix() string {
	if o == nil || o.KeyPrefix == nil {
		var ret string
		return ret
	}
	return *o.KeyPrefix
}

// GetKeyPrefixOk returns a tuple with the KeyPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSConfig) GetKeyPrefixOk() (*string, bool) {
	if o == nil || o.KeyPrefix == nil {
		return nil, false
	}
	return o.KeyPrefix, true
}

// HasKeyPrefix returns a boolean if a field has been set.
func (o *GCSConfig) HasKeyPrefix() bool {
	if o != nil && o.KeyPrefix != nil {
		return true
	}

	return false
}

// SetKeyPrefix gets a reference to the given string and assigns it to the KeyPrefix field.
func (o *GCSConfig) SetKeyPrefix(v string) {
	o.KeyPrefix = &v
}

func (o GCSConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.AutomaticCredentials != nil {
		toSerialize["automatic_credentials"] = o.AutomaticCredentials
	}
	if o.StorageClass != nil {
		toSerialize["storage_class"] = o.StorageClass
	}
	if o.Acl != nil {
		toSerialize["acl"] = o.Acl
	}
	if o.KeyPrefix != nil {
		toSerialize["key_prefix"] = o.KeyPrefix
	}
	return json.Marshal(toSerialize)
}

type NullableGCSConfig struct {
	value *GCSConfig
	isSet bool
}

func (v NullableGCSConfig) Get() *GCSConfig {
	return v.value
}

func (v *NullableGCSConfig) Set(val *GCSConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGCSConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGCSConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCSConfig(val *GCSConfig) *NullableGCSConfig {
	return &NullableGCSConfig{value: val, isSet: true}
}

func (v NullableGCSConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCSConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
