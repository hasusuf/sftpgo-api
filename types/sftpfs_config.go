package types

import (
	"encoding/json"
)

// SFTPFsConfig struct for SFTPFsConfig
type SFTPFsConfig struct {
	// remote SFTP endpoint as host:port
	Endpoint *string `json:"endpoint,omitempty"`
	// you can specify a password or private key or both. In the latter case the private key will be tried first.
	Username   *string `json:"username,omitempty"`
	Password   *Secret `json:"password,omitempty"`
	PrivateKey *Secret `json:"private_key,omitempty"`
	// SHA256 fingerprints to use for host key verification. If you don't provide any fingerprint the remote host key will not be verified, this is a security risk
	Fingerprints []string `json:"fingerprints,omitempty"`
	// Specifying a prefix you can restrict all operations to a given path within the remote SFTP server.
	Prefix *string `json:"prefix,omitempty"`
	// Concurrent reads are safe to use and disabling them will degrade performance. Some servers automatically delete files once they are downloaded. Using concurrent reads is problematic with such servers.
	DisableConcurrentReads *bool `json:"disable_concurrent_reads,omitempty"`
	// The size of the buffer (in MB) to use for transfers. By enabling buffering, the reads and writes, from/to the remote SFTP server, are split in multiple concurrent requests and this allows data to be transferred at a faster rate, over high latency networks, by overlapping round-trip times. With buffering enabled, resuming uploads is not supported and a file cannot be opened for both reading and writing at the same time. 0 means disabled.
	BufferSize *int32 `json:"buffer_size,omitempty"`
}

// NewSFTPFsConfig instantiates a new SFTPFsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSFTPFsConfig() *SFTPFsConfig {
	this := SFTPFsConfig{}
	return &this
}

// NewSFTPFsConfigWithDefaults instantiates a new SFTPFsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSFTPFsConfigWithDefaults() *SFTPFsConfig {
	this := SFTPFsConfig{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *SFTPFsConfig) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SFTPFsConfig) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *SFTPFsConfig) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *SFTPFsConfig) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SFTPFsConfig) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SFTPFsConfig) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SFTPFsConfig) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SFTPFsConfig) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SFTPFsConfig) GetPassword() Secret {
	if o == nil || o.Password == nil {
		var ret Secret
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SFTPFsConfig) GetPasswordOk() (*Secret, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SFTPFsConfig) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given Secret and assigns it to the Password field.
func (o *SFTPFsConfig) SetPassword(v Secret) {
	o.Password = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *SFTPFsConfig) GetPrivateKey() Secret {
	if o == nil || o.PrivateKey == nil {
		var ret Secret
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SFTPFsConfig) GetPrivateKeyOk() (*Secret, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *SFTPFsConfig) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given Secret and assigns it to the PrivateKey field.
func (o *SFTPFsConfig) SetPrivateKey(v Secret) {
	o.PrivateKey = &v
}

// GetFingerprints returns the Fingerprints field value if set, zero value otherwise.
func (o *SFTPFsConfig) GetFingerprints() []string {
	if o == nil || o.Fingerprints == nil {
		var ret []string
		return ret
	}
	return o.Fingerprints
}

// GetFingerprintsOk returns a tuple with the Fingerprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SFTPFsConfig) GetFingerprintsOk() ([]string, bool) {
	if o == nil || o.Fingerprints == nil {
		return nil, false
	}
	return o.Fingerprints, true
}

// HasFingerprints returns a boolean if a field has been set.
func (o *SFTPFsConfig) HasFingerprints() bool {
	if o != nil && o.Fingerprints != nil {
		return true
	}

	return false
}

// SetFingerprints gets a reference to the given []string and assigns it to the Fingerprints field.
func (o *SFTPFsConfig) SetFingerprints(v []string) {
	o.Fingerprints = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *SFTPFsConfig) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SFTPFsConfig) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *SFTPFsConfig) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *SFTPFsConfig) SetPrefix(v string) {
	o.Prefix = &v
}

// GetDisableConcurrentReads returns the DisableConcurrentReads field value if set, zero value otherwise.
func (o *SFTPFsConfig) GetDisableConcurrentReads() bool {
	if o == nil || o.DisableConcurrentReads == nil {
		var ret bool
		return ret
	}
	return *o.DisableConcurrentReads
}

// GetDisableConcurrentReadsOk returns a tuple with the DisableConcurrentReads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SFTPFsConfig) GetDisableConcurrentReadsOk() (*bool, bool) {
	if o == nil || o.DisableConcurrentReads == nil {
		return nil, false
	}
	return o.DisableConcurrentReads, true
}

// HasDisableConcurrentReads returns a boolean if a field has been set.
func (o *SFTPFsConfig) HasDisableConcurrentReads() bool {
	if o != nil && o.DisableConcurrentReads != nil {
		return true
	}

	return false
}

// SetDisableConcurrentReads gets a reference to the given bool and assigns it to the DisableConcurrentReads field.
func (o *SFTPFsConfig) SetDisableConcurrentReads(v bool) {
	o.DisableConcurrentReads = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *SFTPFsConfig) GetBufferSize() int32 {
	if o == nil || o.BufferSize == nil {
		var ret int32
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SFTPFsConfig) GetBufferSizeOk() (*int32, bool) {
	if o == nil || o.BufferSize == nil {
		return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *SFTPFsConfig) HasBufferSize() bool {
	if o != nil && o.BufferSize != nil {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given int32 and assigns it to the BufferSize field.
func (o *SFTPFsConfig) SetBufferSize(v int32) {
	o.BufferSize = &v
}

func (o SFTPFsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	if o.Fingerprints != nil {
		toSerialize["fingerprints"] = o.Fingerprints
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	if o.DisableConcurrentReads != nil {
		toSerialize["disable_concurrent_reads"] = o.DisableConcurrentReads
	}
	if o.BufferSize != nil {
		toSerialize["buffer_size"] = o.BufferSize
	}
	return json.Marshal(toSerialize)
}

type NullableSFTPFsConfig struct {
	value *SFTPFsConfig
	isSet bool
}

func (v NullableSFTPFsConfig) Get() *SFTPFsConfig {
	return v.value
}

func (v *NullableSFTPFsConfig) Set(val *SFTPFsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSFTPFsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSFTPFsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSFTPFsConfig(val *SFTPFsConfig) *NullableSFTPFsConfig {
	return &NullableSFTPFsConfig{value: val, isSet: true}
}

func (v NullableSFTPFsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSFTPFsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
