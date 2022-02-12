package types

import (
	"encoding/json"
)

// FilesystemConfig Storage filesystem details
type FilesystemConfig struct {
	Provider     *FsProviders       `json:"provider,omitempty"`
	S3config     *S3Config          `json:"s3config,omitempty"`
	Gcsconfig    *GCSConfig         `json:"gcsconfig,omitempty"`
	Azblobconfig *AzureBlobFsConfig `json:"azblobconfig,omitempty"`
	Cryptconfig  *CryptFsConfig     `json:"cryptconfig,omitempty"`
	Sftpconfig   *SFTPFsConfig      `json:"sftpconfig,omitempty"`
}

// NewFilesystemConfig instantiates a new FilesystemConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystemConfig() *FilesystemConfig {
	this := FilesystemConfig{}
	return &this
}

// NewFilesystemConfigWithDefaults instantiates a new FilesystemConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemConfigWithDefaults() *FilesystemConfig {
	this := FilesystemConfig{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *FilesystemConfig) GetProvider() FsProviders {
	if o == nil || o.Provider == nil {
		var ret FsProviders
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemConfig) GetProviderOk() (*FsProviders, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *FilesystemConfig) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given FsProviders and assigns it to the Provider field.
func (o *FilesystemConfig) SetProvider(v FsProviders) {
	o.Provider = &v
}

// GetS3config returns the S3config field value if set, zero value otherwise.
func (o *FilesystemConfig) GetS3config() S3Config {
	if o == nil || o.S3config == nil {
		var ret S3Config
		return ret
	}
	return *o.S3config
}

// GetS3configOk returns a tuple with the S3config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemConfig) GetS3configOk() (*S3Config, bool) {
	if o == nil || o.S3config == nil {
		return nil, false
	}
	return o.S3config, true
}

// HasS3config returns a boolean if a field has been set.
func (o *FilesystemConfig) HasS3config() bool {
	if o != nil && o.S3config != nil {
		return true
	}

	return false
}

// SetS3config gets a reference to the given S3Config and assigns it to the S3config field.
func (o *FilesystemConfig) SetS3config(v S3Config) {
	o.S3config = &v
}

// GetGcsconfig returns the Gcsconfig field value if set, zero value otherwise.
func (o *FilesystemConfig) GetGcsconfig() GCSConfig {
	if o == nil || o.Gcsconfig == nil {
		var ret GCSConfig
		return ret
	}
	return *o.Gcsconfig
}

// GetGcsconfigOk returns a tuple with the Gcsconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemConfig) GetGcsconfigOk() (*GCSConfig, bool) {
	if o == nil || o.Gcsconfig == nil {
		return nil, false
	}
	return o.Gcsconfig, true
}

// HasGcsconfig returns a boolean if a field has been set.
func (o *FilesystemConfig) HasGcsconfig() bool {
	if o != nil && o.Gcsconfig != nil {
		return true
	}

	return false
}

// SetGcsconfig gets a reference to the given GCSConfig and assigns it to the Gcsconfig field.
func (o *FilesystemConfig) SetGcsconfig(v GCSConfig) {
	o.Gcsconfig = &v
}

// GetAzblobconfig returns the Azblobconfig field value if set, zero value otherwise.
func (o *FilesystemConfig) GetAzblobconfig() AzureBlobFsConfig {
	if o == nil || o.Azblobconfig == nil {
		var ret AzureBlobFsConfig
		return ret
	}
	return *o.Azblobconfig
}

// GetAzblobconfigOk returns a tuple with the Azblobconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemConfig) GetAzblobconfigOk() (*AzureBlobFsConfig, bool) {
	if o == nil || o.Azblobconfig == nil {
		return nil, false
	}
	return o.Azblobconfig, true
}

// HasAzblobconfig returns a boolean if a field has been set.
func (o *FilesystemConfig) HasAzblobconfig() bool {
	if o != nil && o.Azblobconfig != nil {
		return true
	}

	return false
}

// SetAzblobconfig gets a reference to the given AzureBlobFsConfig and assigns it to the Azblobconfig field.
func (o *FilesystemConfig) SetAzblobconfig(v AzureBlobFsConfig) {
	o.Azblobconfig = &v
}

// GetCryptconfig returns the Cryptconfig field value if set, zero value otherwise.
func (o *FilesystemConfig) GetCryptconfig() CryptFsConfig {
	if o == nil || o.Cryptconfig == nil {
		var ret CryptFsConfig
		return ret
	}
	return *o.Cryptconfig
}

// GetCryptconfigOk returns a tuple with the Cryptconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemConfig) GetCryptconfigOk() (*CryptFsConfig, bool) {
	if o == nil || o.Cryptconfig == nil {
		return nil, false
	}
	return o.Cryptconfig, true
}

// HasCryptconfig returns a boolean if a field has been set.
func (o *FilesystemConfig) HasCryptconfig() bool {
	if o != nil && o.Cryptconfig != nil {
		return true
	}

	return false
}

// SetCryptconfig gets a reference to the given CryptFsConfig and assigns it to the Cryptconfig field.
func (o *FilesystemConfig) SetCryptconfig(v CryptFsConfig) {
	o.Cryptconfig = &v
}

// GetSftpconfig returns the Sftpconfig field value if set, zero value otherwise.
func (o *FilesystemConfig) GetSftpconfig() SFTPFsConfig {
	if o == nil || o.Sftpconfig == nil {
		var ret SFTPFsConfig
		return ret
	}
	return *o.Sftpconfig
}

// GetSftpconfigOk returns a tuple with the Sftpconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemConfig) GetSftpconfigOk() (*SFTPFsConfig, bool) {
	if o == nil || o.Sftpconfig == nil {
		return nil, false
	}
	return o.Sftpconfig, true
}

// HasSftpconfig returns a boolean if a field has been set.
func (o *FilesystemConfig) HasSftpconfig() bool {
	if o != nil && o.Sftpconfig != nil {
		return true
	}

	return false
}

// SetSftpconfig gets a reference to the given SFTPFsConfig and assigns it to the Sftpconfig field.
func (o *FilesystemConfig) SetSftpconfig(v SFTPFsConfig) {
	o.Sftpconfig = &v
}

func (o FilesystemConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.S3config != nil {
		toSerialize["s3config"] = o.S3config
	}
	if o.Gcsconfig != nil {
		toSerialize["gcsconfig"] = o.Gcsconfig
	}
	if o.Azblobconfig != nil {
		toSerialize["azblobconfig"] = o.Azblobconfig
	}
	if o.Cryptconfig != nil {
		toSerialize["cryptconfig"] = o.Cryptconfig
	}
	if o.Sftpconfig != nil {
		toSerialize["sftpconfig"] = o.Sftpconfig
	}
	return json.Marshal(toSerialize)
}

type NullableFilesystemConfig struct {
	value *FilesystemConfig
	isSet bool
}

func (v NullableFilesystemConfig) Get() *FilesystemConfig {
	return v.value
}

func (v *NullableFilesystemConfig) Set(val *FilesystemConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemConfig(val *FilesystemConfig) *NullableFilesystemConfig {
	return &NullableFilesystemConfig{value: val, isSet: true}
}

func (v NullableFilesystemConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
