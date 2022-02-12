package types

import (
	"encoding/json"
)

// S3Config S3 Compatible Object Storage configuration details
type S3Config struct {
	Bucket       *string `json:"bucket,omitempty"`
	Region       *string `json:"region,omitempty"`
	AccessKey    *string `json:"access_key,omitempty"`
	AccessSecret *Secret `json:"access_secret,omitempty"`
	// optional endpoint
	Endpoint     *string `json:"endpoint,omitempty"`
	StorageClass *string `json:"storage_class,omitempty"`
	// The canned ACL to apply to uploaded objects. Leave empty to use the default ACL. For more information and available ACLs, see here: https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl
	Acl *string `json:"acl,omitempty"`
	// the buffer size (in MB) to use for multipart uploads. The minimum allowed part size is 5MB, and if this value is set to zero, the default value (5MB) for the AWS SDK will be used. The minimum allowed value is 5.
	UploadPartSize *int32 `json:"upload_part_size,omitempty"`
	// the number of parts to upload in parallel. If this value is set to zero, the default value (5) will be used
	UploadConcurrency *int32 `json:"upload_concurrency,omitempty"`
	// the maximum time allowed, in seconds, to upload a single chunk (the chunk size is defined via \"upload_part_size\"). 0 means no timeout
	UploadPartMaxTime *int32 `json:"upload_part_max_time,omitempty"`
	// the buffer size (in MB) to use for multipart downloads. The minimum allowed part size is 5MB, and if this value is set to zero, the default value (5MB) for the AWS SDK will be used. The minimum allowed value is 5. Ignored for partial downloads
	DownloadPartSize *int32 `json:"download_part_size,omitempty"`
	// the number of parts to download in parallel. If this value is set to zero, the default value (5) will be used. Ignored for partial downloads
	DownloadConcurrency *int32 `json:"download_concurrency,omitempty"`
	// the maximum time allowed, in seconds, to download a single chunk (the chunk size is defined via \"download_part_size\"). 0 means no timeout. Ignored for partial downloads.
	DownloadPartMaxTime *int32 `json:"download_part_max_time,omitempty"`
	// Set this to \"true\" to force the request to use path-style addressing, i.e., \"http://s3.amazonaws.com/BUCKET/KEY\". By default, the S3 client will use virtual hosted bucket addressing when possible (\"http://BUCKET.s3.amazonaws.com/KEY\")
	ForcePathStyle *bool `json:"force_path_style,omitempty"`
	// key_prefix is similar to a chroot directory for a local filesystem. If specified the user will only see contents that starts with this prefix and so you can restrict access to a specific virtual folder. The prefix, if not empty, must not start with \"/\" and must end with \"/\". If empty the whole bucket contents will be available
	KeyPrefix *string `json:"key_prefix,omitempty"`
}

// NewS3Config instantiates a new S3Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3Config() *S3Config {
	this := S3Config{}
	return &this
}

// NewS3ConfigWithDefaults instantiates a new S3Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3ConfigWithDefaults() *S3Config {
	this := S3Config{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *S3Config) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *S3Config) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *S3Config) SetBucket(v string) {
	o.Bucket = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *S3Config) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *S3Config) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *S3Config) SetRegion(v string) {
	o.Region = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *S3Config) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *S3Config) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *S3Config) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetAccessSecret returns the AccessSecret field value if set, zero value otherwise.
func (o *S3Config) GetAccessSecret() Secret {
	if o == nil || o.AccessSecret == nil {
		var ret Secret
		return ret
	}
	return *o.AccessSecret
}

// GetAccessSecretOk returns a tuple with the AccessSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetAccessSecretOk() (*Secret, bool) {
	if o == nil || o.AccessSecret == nil {
		return nil, false
	}
	return o.AccessSecret, true
}

// HasAccessSecret returns a boolean if a field has been set.
func (o *S3Config) HasAccessSecret() bool {
	if o != nil && o.AccessSecret != nil {
		return true
	}

	return false
}

// SetAccessSecret gets a reference to the given Secret and assigns it to the AccessSecret field.
func (o *S3Config) SetAccessSecret(v Secret) {
	o.AccessSecret = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *S3Config) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *S3Config) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *S3Config) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *S3Config) GetStorageClass() string {
	if o == nil || o.StorageClass == nil {
		var ret string
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetStorageClassOk() (*string, bool) {
	if o == nil || o.StorageClass == nil {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *S3Config) HasStorageClass() bool {
	if o != nil && o.StorageClass != nil {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given string and assigns it to the StorageClass field.
func (o *S3Config) SetStorageClass(v string) {
	o.StorageClass = &v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *S3Config) GetAcl() string {
	if o == nil || o.Acl == nil {
		var ret string
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetAclOk() (*string, bool) {
	if o == nil || o.Acl == nil {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *S3Config) HasAcl() bool {
	if o != nil && o.Acl != nil {
		return true
	}

	return false
}

// SetAcl gets a reference to the given string and assigns it to the Acl field.
func (o *S3Config) SetAcl(v string) {
	o.Acl = &v
}

// GetUploadPartSize returns the UploadPartSize field value if set, zero value otherwise.
func (o *S3Config) GetUploadPartSize() int32 {
	if o == nil || o.UploadPartSize == nil {
		var ret int32
		return ret
	}
	return *o.UploadPartSize
}

// GetUploadPartSizeOk returns a tuple with the UploadPartSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetUploadPartSizeOk() (*int32, bool) {
	if o == nil || o.UploadPartSize == nil {
		return nil, false
	}
	return o.UploadPartSize, true
}

// HasUploadPartSize returns a boolean if a field has been set.
func (o *S3Config) HasUploadPartSize() bool {
	if o != nil && o.UploadPartSize != nil {
		return true
	}

	return false
}

// SetUploadPartSize gets a reference to the given int32 and assigns it to the UploadPartSize field.
func (o *S3Config) SetUploadPartSize(v int32) {
	o.UploadPartSize = &v
}

// GetUploadConcurrency returns the UploadConcurrency field value if set, zero value otherwise.
func (o *S3Config) GetUploadConcurrency() int32 {
	if o == nil || o.UploadConcurrency == nil {
		var ret int32
		return ret
	}
	return *o.UploadConcurrency
}

// GetUploadConcurrencyOk returns a tuple with the UploadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetUploadConcurrencyOk() (*int32, bool) {
	if o == nil || o.UploadConcurrency == nil {
		return nil, false
	}
	return o.UploadConcurrency, true
}

// HasUploadConcurrency returns a boolean if a field has been set.
func (o *S3Config) HasUploadConcurrency() bool {
	if o != nil && o.UploadConcurrency != nil {
		return true
	}

	return false
}

// SetUploadConcurrency gets a reference to the given int32 and assigns it to the UploadConcurrency field.
func (o *S3Config) SetUploadConcurrency(v int32) {
	o.UploadConcurrency = &v
}

// GetUploadPartMaxTime returns the UploadPartMaxTime field value if set, zero value otherwise.
func (o *S3Config) GetUploadPartMaxTime() int32 {
	if o == nil || o.UploadPartMaxTime == nil {
		var ret int32
		return ret
	}
	return *o.UploadPartMaxTime
}

// GetUploadPartMaxTimeOk returns a tuple with the UploadPartMaxTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetUploadPartMaxTimeOk() (*int32, bool) {
	if o == nil || o.UploadPartMaxTime == nil {
		return nil, false
	}
	return o.UploadPartMaxTime, true
}

// HasUploadPartMaxTime returns a boolean if a field has been set.
func (o *S3Config) HasUploadPartMaxTime() bool {
	if o != nil && o.UploadPartMaxTime != nil {
		return true
	}

	return false
}

// SetUploadPartMaxTime gets a reference to the given int32 and assigns it to the UploadPartMaxTime field.
func (o *S3Config) SetUploadPartMaxTime(v int32) {
	o.UploadPartMaxTime = &v
}

// GetDownloadPartSize returns the DownloadPartSize field value if set, zero value otherwise.
func (o *S3Config) GetDownloadPartSize() int32 {
	if o == nil || o.DownloadPartSize == nil {
		var ret int32
		return ret
	}
	return *o.DownloadPartSize
}

// GetDownloadPartSizeOk returns a tuple with the DownloadPartSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetDownloadPartSizeOk() (*int32, bool) {
	if o == nil || o.DownloadPartSize == nil {
		return nil, false
	}
	return o.DownloadPartSize, true
}

// HasDownloadPartSize returns a boolean if a field has been set.
func (o *S3Config) HasDownloadPartSize() bool {
	if o != nil && o.DownloadPartSize != nil {
		return true
	}

	return false
}

// SetDownloadPartSize gets a reference to the given int32 and assigns it to the DownloadPartSize field.
func (o *S3Config) SetDownloadPartSize(v int32) {
	o.DownloadPartSize = &v
}

// GetDownloadConcurrency returns the DownloadConcurrency field value if set, zero value otherwise.
func (o *S3Config) GetDownloadConcurrency() int32 {
	if o == nil || o.DownloadConcurrency == nil {
		var ret int32
		return ret
	}
	return *o.DownloadConcurrency
}

// GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetDownloadConcurrencyOk() (*int32, bool) {
	if o == nil || o.DownloadConcurrency == nil {
		return nil, false
	}
	return o.DownloadConcurrency, true
}

// HasDownloadConcurrency returns a boolean if a field has been set.
func (o *S3Config) HasDownloadConcurrency() bool {
	if o != nil && o.DownloadConcurrency != nil {
		return true
	}

	return false
}

// SetDownloadConcurrency gets a reference to the given int32 and assigns it to the DownloadConcurrency field.
func (o *S3Config) SetDownloadConcurrency(v int32) {
	o.DownloadConcurrency = &v
}

// GetDownloadPartMaxTime returns the DownloadPartMaxTime field value if set, zero value otherwise.
func (o *S3Config) GetDownloadPartMaxTime() int32 {
	if o == nil || o.DownloadPartMaxTime == nil {
		var ret int32
		return ret
	}
	return *o.DownloadPartMaxTime
}

// GetDownloadPartMaxTimeOk returns a tuple with the DownloadPartMaxTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetDownloadPartMaxTimeOk() (*int32, bool) {
	if o == nil || o.DownloadPartMaxTime == nil {
		return nil, false
	}
	return o.DownloadPartMaxTime, true
}

// HasDownloadPartMaxTime returns a boolean if a field has been set.
func (o *S3Config) HasDownloadPartMaxTime() bool {
	if o != nil && o.DownloadPartMaxTime != nil {
		return true
	}

	return false
}

// SetDownloadPartMaxTime gets a reference to the given int32 and assigns it to the DownloadPartMaxTime field.
func (o *S3Config) SetDownloadPartMaxTime(v int32) {
	o.DownloadPartMaxTime = &v
}

// GetForcePathStyle returns the ForcePathStyle field value if set, zero value otherwise.
func (o *S3Config) GetForcePathStyle() bool {
	if o == nil || o.ForcePathStyle == nil {
		var ret bool
		return ret
	}
	return *o.ForcePathStyle
}

// GetForcePathStyleOk returns a tuple with the ForcePathStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetForcePathStyleOk() (*bool, bool) {
	if o == nil || o.ForcePathStyle == nil {
		return nil, false
	}
	return o.ForcePathStyle, true
}

// HasForcePathStyle returns a boolean if a field has been set.
func (o *S3Config) HasForcePathStyle() bool {
	if o != nil && o.ForcePathStyle != nil {
		return true
	}

	return false
}

// SetForcePathStyle gets a reference to the given bool and assigns it to the ForcePathStyle field.
func (o *S3Config) SetForcePathStyle(v bool) {
	o.ForcePathStyle = &v
}

// GetKeyPrefix returns the KeyPrefix field value if set, zero value otherwise.
func (o *S3Config) GetKeyPrefix() string {
	if o == nil || o.KeyPrefix == nil {
		var ret string
		return ret
	}
	return *o.KeyPrefix
}

// GetKeyPrefixOk returns a tuple with the KeyPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetKeyPrefixOk() (*string, bool) {
	if o == nil || o.KeyPrefix == nil {
		return nil, false
	}
	return o.KeyPrefix, true
}

// HasKeyPrefix returns a boolean if a field has been set.
func (o *S3Config) HasKeyPrefix() bool {
	if o != nil && o.KeyPrefix != nil {
		return true
	}

	return false
}

// SetKeyPrefix gets a reference to the given string and assigns it to the KeyPrefix field.
func (o *S3Config) SetKeyPrefix(v string) {
	o.KeyPrefix = &v
}

func (o S3Config) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.AccessKey != nil {
		toSerialize["access_key"] = o.AccessKey
	}
	if o.AccessSecret != nil {
		toSerialize["access_secret"] = o.AccessSecret
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.StorageClass != nil {
		toSerialize["storage_class"] = o.StorageClass
	}
	if o.Acl != nil {
		toSerialize["acl"] = o.Acl
	}
	if o.UploadPartSize != nil {
		toSerialize["upload_part_size"] = o.UploadPartSize
	}
	if o.UploadConcurrency != nil {
		toSerialize["upload_concurrency"] = o.UploadConcurrency
	}
	if o.UploadPartMaxTime != nil {
		toSerialize["upload_part_max_time"] = o.UploadPartMaxTime
	}
	if o.DownloadPartSize != nil {
		toSerialize["download_part_size"] = o.DownloadPartSize
	}
	if o.DownloadConcurrency != nil {
		toSerialize["download_concurrency"] = o.DownloadConcurrency
	}
	if o.DownloadPartMaxTime != nil {
		toSerialize["download_part_max_time"] = o.DownloadPartMaxTime
	}
	if o.ForcePathStyle != nil {
		toSerialize["force_path_style"] = o.ForcePathStyle
	}
	if o.KeyPrefix != nil {
		toSerialize["key_prefix"] = o.KeyPrefix
	}
	return json.Marshal(toSerialize)
}

type NullableS3Config struct {
	value *S3Config
	isSet bool
}

func (v NullableS3Config) Get() *S3Config {
	return v.value
}

func (v *NullableS3Config) Set(val *S3Config) {
	v.value = val
	v.isSet = true
}

func (v NullableS3Config) IsSet() bool {
	return v.isSet
}

func (v *NullableS3Config) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3Config(val *S3Config) *NullableS3Config {
	return &NullableS3Config{value: val, isSet: true}
}

func (v NullableS3Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3Config) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
