package types

import (
	"encoding/json"
)

// UserFilters Additional user options
type UserFilters struct {
	// only clients connecting from these IP/Mask are allowed. IP/Mask must be in CIDR notation as defined in RFC 4632 and RFC 4291, for example \"192.0.2.0/24\" or \"2001:db8::/32\"
	AllowedIp []string `json:"allowed_ip,omitempty"`
	// clients connecting from these IP/Mask are not allowed. Denied rules are evaluated before allowed ones
	DeniedIp []string `json:"denied_ip,omitempty"`
	// if null or empty any available login method is allowed
	DeniedLoginMethods []LoginMethods `json:"denied_login_methods,omitempty"`
	// if null or empty any available protocol is allowed
	DeniedProtocols []SupportedProtocols `json:"denied_protocols,omitempty"`
	// filters based on shell like file patterns. These restrictions do not apply to files listing for performance reasons, so a denied file cannot be downloaded/overwritten/renamed but it will still be in the list of files. Please note that these restrictions can be easily bypassed
	FilePatterns []PatternsFilter `json:"file_patterns,omitempty"`
	// maximum allowed size, as bytes, for a single file upload. The upload will be aborted if/when the size of the file being sent exceeds this limit. 0 means unlimited. This restriction does not apply for SSH system commands such as `git` and `rsync`
	MaxUploadFileSize *int64 `json:"max_upload_file_size,omitempty"`
	// defines the TLS certificate field to use as username. For FTP clients it must match the name provided using the \"USER\" command. For WebDAV, if no username is provided, the CN will be used as username. For WebDAV clients it must match the implicit or provided username. Ignored if mutual TLS is disabled
	TlsUsername *string      `json:"tls_username,omitempty"`
	Hooks       *HooksFilter `json:"hooks,omitempty"`
	// Disable checks for existence and automatic creation of home directory and virtual folders. SFTPGo requires that the user's home directory, virtual folder root, and intermediate paths to virtual folders exist to work properly. If you already know that the required directories exist, disabling these checks will speed up login. You could, for example, disable these checks after the first login
	DisableFsChecks *bool `json:"disable_fs_checks,omitempty"`
	// WebClient/user REST API related configuration options
	WebClient []WebClientOptions `json:"web_client,omitempty"`
	// API key authentication allows to impersonate this user with an API key
	AllowApiKeyAuth    *bool               `json:"allow_api_key_auth,omitempty"`
	UserType           *UserType           `json:"user_type,omitempty"`
	TotpConfig         *UserTOTPConfig     `json:"totp_config,omitempty"`
	RecoveryCodes      []RecoveryCode      `json:"recovery_codes,omitempty"`
	BandwidthLimits    []BandwidthLimit    `json:"bandwidth_limits,omitempty"`
	DataTransferLimits []DataTransferLimit `json:"data_transfer_limits,omitempty"`
}

// NewUserFilters instantiates a new UserFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFilters() *UserFilters {
	this := UserFilters{}
	return &this
}

// NewUserFiltersWithDefaults instantiates a new UserFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFiltersWithDefaults() *UserFilters {
	this := UserFilters{}
	return &this
}

// GetAllowedIp returns the AllowedIp field value if set, zero value otherwise.
func (o *UserFilters) GetAllowedIp() []string {
	if o == nil || o.AllowedIp == nil {
		var ret []string
		return ret
	}
	return o.AllowedIp
}

// GetAllowedIpOk returns a tuple with the AllowedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetAllowedIpOk() ([]string, bool) {
	if o == nil || o.AllowedIp == nil {
		return nil, false
	}
	return o.AllowedIp, true
}

// HasAllowedIp returns a boolean if a field has been set.
func (o *UserFilters) HasAllowedIp() bool {
	if o != nil && o.AllowedIp != nil {
		return true
	}

	return false
}

// SetAllowedIp gets a reference to the given []string and assigns it to the AllowedIp field.
func (o *UserFilters) SetAllowedIp(v []string) {
	o.AllowedIp = v
}

// GetDeniedIp returns the DeniedIp field value if set, zero value otherwise.
func (o *UserFilters) GetDeniedIp() []string {
	if o == nil || o.DeniedIp == nil {
		var ret []string
		return ret
	}
	return o.DeniedIp
}

// GetDeniedIpOk returns a tuple with the DeniedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetDeniedIpOk() ([]string, bool) {
	if o == nil || o.DeniedIp == nil {
		return nil, false
	}
	return o.DeniedIp, true
}

// HasDeniedIp returns a boolean if a field has been set.
func (o *UserFilters) HasDeniedIp() bool {
	if o != nil && o.DeniedIp != nil {
		return true
	}

	return false
}

// SetDeniedIp gets a reference to the given []string and assigns it to the DeniedIp field.
func (o *UserFilters) SetDeniedIp(v []string) {
	o.DeniedIp = v
}

// GetDeniedLoginMethods returns the DeniedLoginMethods field value if set, zero value otherwise.
func (o *UserFilters) GetDeniedLoginMethods() []LoginMethods {
	if o == nil || o.DeniedLoginMethods == nil {
		var ret []LoginMethods
		return ret
	}
	return o.DeniedLoginMethods
}

// GetDeniedLoginMethodsOk returns a tuple with the DeniedLoginMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetDeniedLoginMethodsOk() ([]LoginMethods, bool) {
	if o == nil || o.DeniedLoginMethods == nil {
		return nil, false
	}
	return o.DeniedLoginMethods, true
}

// HasDeniedLoginMethods returns a boolean if a field has been set.
func (o *UserFilters) HasDeniedLoginMethods() bool {
	if o != nil && o.DeniedLoginMethods != nil {
		return true
	}

	return false
}

// SetDeniedLoginMethods gets a reference to the given []LoginMethods and assigns it to the DeniedLoginMethods field.
func (o *UserFilters) SetDeniedLoginMethods(v []LoginMethods) {
	o.DeniedLoginMethods = v
}

// GetDeniedProtocols returns the DeniedProtocols field value if set, zero value otherwise.
func (o *UserFilters) GetDeniedProtocols() []SupportedProtocols {
	if o == nil || o.DeniedProtocols == nil {
		var ret []SupportedProtocols
		return ret
	}
	return o.DeniedProtocols
}

// GetDeniedProtocolsOk returns a tuple with the DeniedProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetDeniedProtocolsOk() ([]SupportedProtocols, bool) {
	if o == nil || o.DeniedProtocols == nil {
		return nil, false
	}
	return o.DeniedProtocols, true
}

// HasDeniedProtocols returns a boolean if a field has been set.
func (o *UserFilters) HasDeniedProtocols() bool {
	if o != nil && o.DeniedProtocols != nil {
		return true
	}

	return false
}

// SetDeniedProtocols gets a reference to the given []SupportedProtocols and assigns it to the DeniedProtocols field.
func (o *UserFilters) SetDeniedProtocols(v []SupportedProtocols) {
	o.DeniedProtocols = v
}

// GetFilePatterns returns the FilePatterns field value if set, zero value otherwise.
func (o *UserFilters) GetFilePatterns() []PatternsFilter {
	if o == nil || o.FilePatterns == nil {
		var ret []PatternsFilter
		return ret
	}
	return o.FilePatterns
}

// GetFilePatternsOk returns a tuple with the FilePatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetFilePatternsOk() ([]PatternsFilter, bool) {
	if o == nil || o.FilePatterns == nil {
		return nil, false
	}
	return o.FilePatterns, true
}

// HasFilePatterns returns a boolean if a field has been set.
func (o *UserFilters) HasFilePatterns() bool {
	if o != nil && o.FilePatterns != nil {
		return true
	}

	return false
}

// SetFilePatterns gets a reference to the given []PatternsFilter and assigns it to the FilePatterns field.
func (o *UserFilters) SetFilePatterns(v []PatternsFilter) {
	o.FilePatterns = v
}

// GetMaxUploadFileSize returns the MaxUploadFileSize field value if set, zero value otherwise.
func (o *UserFilters) GetMaxUploadFileSize() int64 {
	if o == nil || o.MaxUploadFileSize == nil {
		var ret int64
		return ret
	}
	return *o.MaxUploadFileSize
}

// GetMaxUploadFileSizeOk returns a tuple with the MaxUploadFileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetMaxUploadFileSizeOk() (*int64, bool) {
	if o == nil || o.MaxUploadFileSize == nil {
		return nil, false
	}
	return o.MaxUploadFileSize, true
}

// HasMaxUploadFileSize returns a boolean if a field has been set.
func (o *UserFilters) HasMaxUploadFileSize() bool {
	if o != nil && o.MaxUploadFileSize != nil {
		return true
	}

	return false
}

// SetMaxUploadFileSize gets a reference to the given int64 and assigns it to the MaxUploadFileSize field.
func (o *UserFilters) SetMaxUploadFileSize(v int64) {
	o.MaxUploadFileSize = &v
}

// GetTlsUsername returns the TlsUsername field value if set, zero value otherwise.
func (o *UserFilters) GetTlsUsername() string {
	if o == nil || o.TlsUsername == nil {
		var ret string
		return ret
	}
	return *o.TlsUsername
}

// GetTlsUsernameOk returns a tuple with the TlsUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetTlsUsernameOk() (*string, bool) {
	if o == nil || o.TlsUsername == nil {
		return nil, false
	}
	return o.TlsUsername, true
}

// HasTlsUsername returns a boolean if a field has been set.
func (o *UserFilters) HasTlsUsername() bool {
	if o != nil && o.TlsUsername != nil {
		return true
	}

	return false
}

// SetTlsUsername gets a reference to the given string and assigns it to the TlsUsername field.
func (o *UserFilters) SetTlsUsername(v string) {
	o.TlsUsername = &v
}

// GetHooks returns the Hooks field value if set, zero value otherwise.
func (o *UserFilters) GetHooks() HooksFilter {
	if o == nil || o.Hooks == nil {
		var ret HooksFilter
		return ret
	}
	return *o.Hooks
}

// GetHooksOk returns a tuple with the Hooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetHooksOk() (*HooksFilter, bool) {
	if o == nil || o.Hooks == nil {
		return nil, false
	}
	return o.Hooks, true
}

// HasHooks returns a boolean if a field has been set.
func (o *UserFilters) HasHooks() bool {
	if o != nil && o.Hooks != nil {
		return true
	}

	return false
}

// SetHooks gets a reference to the given HooksFilter and assigns it to the Hooks field.
func (o *UserFilters) SetHooks(v HooksFilter) {
	o.Hooks = &v
}

// GetDisableFsChecks returns the DisableFsChecks field value if set, zero value otherwise.
func (o *UserFilters) GetDisableFsChecks() bool {
	if o == nil || o.DisableFsChecks == nil {
		var ret bool
		return ret
	}
	return *o.DisableFsChecks
}

// GetDisableFsChecksOk returns a tuple with the DisableFsChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetDisableFsChecksOk() (*bool, bool) {
	if o == nil || o.DisableFsChecks == nil {
		return nil, false
	}
	return o.DisableFsChecks, true
}

// HasDisableFsChecks returns a boolean if a field has been set.
func (o *UserFilters) HasDisableFsChecks() bool {
	if o != nil && o.DisableFsChecks != nil {
		return true
	}

	return false
}

// SetDisableFsChecks gets a reference to the given bool and assigns it to the DisableFsChecks field.
func (o *UserFilters) SetDisableFsChecks(v bool) {
	o.DisableFsChecks = &v
}

// GetWebClient returns the WebClient field value if set, zero value otherwise.
func (o *UserFilters) GetWebClient() []WebClientOptions {
	if o == nil || o.WebClient == nil {
		var ret []WebClientOptions
		return ret
	}
	return o.WebClient
}

// GetWebClientOk returns a tuple with the WebClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetWebClientOk() ([]WebClientOptions, bool) {
	if o == nil || o.WebClient == nil {
		return nil, false
	}
	return o.WebClient, true
}

// HasWebClient returns a boolean if a field has been set.
func (o *UserFilters) HasWebClient() bool {
	if o != nil && o.WebClient != nil {
		return true
	}

	return false
}

// SetWebClient gets a reference to the given []WebClientOptions and assigns it to the WebClient field.
func (o *UserFilters) SetWebClient(v []WebClientOptions) {
	o.WebClient = v
}

// GetAllowApiKeyAuth returns the AllowApiKeyAuth field value if set, zero value otherwise.
func (o *UserFilters) GetAllowApiKeyAuth() bool {
	if o == nil || o.AllowApiKeyAuth == nil {
		var ret bool
		return ret
	}
	return *o.AllowApiKeyAuth
}

// GetAllowApiKeyAuthOk returns a tuple with the AllowApiKeyAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetAllowApiKeyAuthOk() (*bool, bool) {
	if o == nil || o.AllowApiKeyAuth == nil {
		return nil, false
	}
	return o.AllowApiKeyAuth, true
}

// HasAllowApiKeyAuth returns a boolean if a field has been set.
func (o *UserFilters) HasAllowApiKeyAuth() bool {
	if o != nil && o.AllowApiKeyAuth != nil {
		return true
	}

	return false
}

// SetAllowApiKeyAuth gets a reference to the given bool and assigns it to the AllowApiKeyAuth field.
func (o *UserFilters) SetAllowApiKeyAuth(v bool) {
	o.AllowApiKeyAuth = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *UserFilters) GetUserType() UserType {
	if o == nil || o.UserType == nil {
		var ret UserType
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetUserTypeOk() (*UserType, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *UserFilters) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given UserType and assigns it to the UserType field.
func (o *UserFilters) SetUserType(v UserType) {
	o.UserType = &v
}

// GetTotpConfig returns the TotpConfig field value if set, zero value otherwise.
func (o *UserFilters) GetTotpConfig() UserTOTPConfig {
	if o == nil || o.TotpConfig == nil {
		var ret UserTOTPConfig
		return ret
	}
	return *o.TotpConfig
}

// GetTotpConfigOk returns a tuple with the TotpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetTotpConfigOk() (*UserTOTPConfig, bool) {
	if o == nil || o.TotpConfig == nil {
		return nil, false
	}
	return o.TotpConfig, true
}

// HasTotpConfig returns a boolean if a field has been set.
func (o *UserFilters) HasTotpConfig() bool {
	if o != nil && o.TotpConfig != nil {
		return true
	}

	return false
}

// SetTotpConfig gets a reference to the given UserTOTPConfig and assigns it to the TotpConfig field.
func (o *UserFilters) SetTotpConfig(v UserTOTPConfig) {
	o.TotpConfig = &v
}

// GetRecoveryCodes returns the RecoveryCodes field value if set, zero value otherwise.
func (o *UserFilters) GetRecoveryCodes() []RecoveryCode {
	if o == nil || o.RecoveryCodes == nil {
		var ret []RecoveryCode
		return ret
	}
	return o.RecoveryCodes
}

// GetRecoveryCodesOk returns a tuple with the RecoveryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetRecoveryCodesOk() ([]RecoveryCode, bool) {
	if o == nil || o.RecoveryCodes == nil {
		return nil, false
	}
	return o.RecoveryCodes, true
}

// HasRecoveryCodes returns a boolean if a field has been set.
func (o *UserFilters) HasRecoveryCodes() bool {
	if o != nil && o.RecoveryCodes != nil {
		return true
	}

	return false
}

// SetRecoveryCodes gets a reference to the given []RecoveryCode and assigns it to the RecoveryCodes field.
func (o *UserFilters) SetRecoveryCodes(v []RecoveryCode) {
	o.RecoveryCodes = v
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *UserFilters) GetBandwidthLimits() []BandwidthLimit {
	if o == nil || o.BandwidthLimits == nil {
		var ret []BandwidthLimit
		return ret
	}
	return o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetBandwidthLimitsOk() ([]BandwidthLimit, bool) {
	if o == nil || o.BandwidthLimits == nil {
		return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *UserFilters) HasBandwidthLimits() bool {
	if o != nil && o.BandwidthLimits != nil {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given []BandwidthLimit and assigns it to the BandwidthLimits field.
func (o *UserFilters) SetBandwidthLimits(v []BandwidthLimit) {
	o.BandwidthLimits = v
}

// GetDataTransferLimits returns the DataTransferLimits field value if set, zero value otherwise.
func (o *UserFilters) GetDataTransferLimits() []DataTransferLimit {
	if o == nil || o.DataTransferLimits == nil {
		var ret []DataTransferLimit
		return ret
	}
	return o.DataTransferLimits
}

// GetDataTransferLimitsOk returns a tuple with the DataTransferLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFilters) GetDataTransferLimitsOk() ([]DataTransferLimit, bool) {
	if o == nil || o.DataTransferLimits == nil {
		return nil, false
	}
	return o.DataTransferLimits, true
}

// HasDataTransferLimits returns a boolean if a field has been set.
func (o *UserFilters) HasDataTransferLimits() bool {
	if o != nil && o.DataTransferLimits != nil {
		return true
	}

	return false
}

// SetDataTransferLimits gets a reference to the given []DataTransferLimit and assigns it to the DataTransferLimits field.
func (o *UserFilters) SetDataTransferLimits(v []DataTransferLimit) {
	o.DataTransferLimits = v
}

func (o UserFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedIp != nil {
		toSerialize["allowed_ip"] = o.AllowedIp
	}
	if o.DeniedIp != nil {
		toSerialize["denied_ip"] = o.DeniedIp
	}
	if o.DeniedLoginMethods != nil {
		toSerialize["denied_login_methods"] = o.DeniedLoginMethods
	}
	if o.DeniedProtocols != nil {
		toSerialize["denied_protocols"] = o.DeniedProtocols
	}
	if o.FilePatterns != nil {
		toSerialize["file_patterns"] = o.FilePatterns
	}
	if o.MaxUploadFileSize != nil {
		toSerialize["max_upload_file_size"] = o.MaxUploadFileSize
	}
	if o.TlsUsername != nil {
		toSerialize["tls_username"] = o.TlsUsername
	}
	if o.Hooks != nil {
		toSerialize["hooks"] = o.Hooks
	}
	if o.DisableFsChecks != nil {
		toSerialize["disable_fs_checks"] = o.DisableFsChecks
	}
	if o.WebClient != nil {
		toSerialize["web_client"] = o.WebClient
	}
	if o.AllowApiKeyAuth != nil {
		toSerialize["allow_api_key_auth"] = o.AllowApiKeyAuth
	}
	if o.UserType != nil {
		toSerialize["user_type"] = o.UserType
	}
	if o.TotpConfig != nil {
		toSerialize["totp_config"] = o.TotpConfig
	}
	if o.RecoveryCodes != nil {
		toSerialize["recovery_codes"] = o.RecoveryCodes
	}
	if o.BandwidthLimits != nil {
		toSerialize["bandwidth_limits"] = o.BandwidthLimits
	}
	if o.DataTransferLimits != nil {
		toSerialize["data_transfer_limits"] = o.DataTransferLimits
	}
	return json.Marshal(toSerialize)
}

type NullableUserFilters struct {
	value *UserFilters
	isSet bool
}

func (v NullableUserFilters) Get() *UserFilters {
	return v.value
}

func (v *NullableUserFilters) Set(val *UserFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFilters(val *UserFilters) *NullableUserFilters {
	return &NullableUserFilters{value: val, isSet: true}
}

func (v NullableUserFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
