package types

import "encoding/json"

// VirtualFolder A virtual folder is a mapping between a SFTPGo virtual path and a filesystem path outside the user home directory. The specified paths must be absolute and the virtual path cannot be \"/\", it must be a sub directory. The parent directory for the specified virtual path must exist. SFTPGo will try to automatically create any missing parent directory for the configured virtual folders at user login.
type VirtualFolder struct {
	Id *int32 `json:"id,omitempty"`
	// unique name for this virtual folder
	Name *string `json:"name,omitempty"`
	// absolute filesystem path to use as virtual folder
	MappedPath *string `json:"mapped_path,omitempty"`
	// optional description
	Description    *string `json:"description,omitempty"`
	UsedQuotaSize  *int64  `json:"used_quota_size,omitempty"`
	UsedQuotaFiles *int32  `json:"used_quota_files,omitempty"`
	// Last quota update as unix timestamp in milliseconds
	LastQuotaUpdate *int64 `json:"last_quota_update,omitempty"`
	// list of usernames associated with this virtual folder
	Users       []string          `json:"users,omitempty"`
	Filesystem  *FilesystemConfig `json:"filesystem,omitempty"`
	VirtualPath string            `json:"virtual_path"`
	// Quota as size in bytes. 0 menas unlimited, -1 means included in user quota. Please note that quota is updated if files are added/removed via SFTPGo otherwise a quota scan or a manual quota update is needed
	QuotaSize *int64 `json:"quota_size,omitempty"`
	// Quota as number of files. 0 menas unlimited, , -1 means included in user quota. Please note that quota is updated if files are added/removed via SFTPGo otherwise a quota scan or a manual quota update is needed
	QuotaFiles *int32 `json:"quota_files,omitempty"`
}

// NewVirtualFolder instantiates a new VirtualFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualFolder(virtualPath string) *VirtualFolder {
	this := VirtualFolder{}
	this.VirtualPath = virtualPath
	return &this
}

// NewVirtualFolderWithDefaults instantiates a new VirtualFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualFolderWithDefaults() *VirtualFolder {
	this := VirtualFolder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VirtualFolder) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VirtualFolder) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VirtualFolder) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualFolder) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualFolder) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualFolder) SetName(v string) {
	o.Name = &v
}

// GetMappedPath returns the MappedPath field value if set, zero value otherwise.
func (o *VirtualFolder) GetMappedPath() string {
	if o == nil || o.MappedPath == nil {
		var ret string
		return ret
	}
	return *o.MappedPath
}

// GetMappedPathOk returns a tuple with the MappedPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetMappedPathOk() (*string, bool) {
	if o == nil || o.MappedPath == nil {
		return nil, false
	}
	return o.MappedPath, true
}

// HasMappedPath returns a boolean if a field has been set.
func (o *VirtualFolder) HasMappedPath() bool {
	if o != nil && o.MappedPath != nil {
		return true
	}

	return false
}

// SetMappedPath gets a reference to the given string and assigns it to the MappedPath field.
func (o *VirtualFolder) SetMappedPath(v string) {
	o.MappedPath = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualFolder) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualFolder) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualFolder) SetDescription(v string) {
	o.Description = &v
}

// GetUsedQuotaSize returns the UsedQuotaSize field value if set, zero value otherwise.
func (o *VirtualFolder) GetUsedQuotaSize() int64 {
	if o == nil || o.UsedQuotaSize == nil {
		var ret int64
		return ret
	}
	return *o.UsedQuotaSize
}

// GetUsedQuotaSizeOk returns a tuple with the UsedQuotaSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetUsedQuotaSizeOk() (*int64, bool) {
	if o == nil || o.UsedQuotaSize == nil {
		return nil, false
	}
	return o.UsedQuotaSize, true
}

// HasUsedQuotaSize returns a boolean if a field has been set.
func (o *VirtualFolder) HasUsedQuotaSize() bool {
	if o != nil && o.UsedQuotaSize != nil {
		return true
	}

	return false
}

// SetUsedQuotaSize gets a reference to the given int64 and assigns it to the UsedQuotaSize field.
func (o *VirtualFolder) SetUsedQuotaSize(v int64) {
	o.UsedQuotaSize = &v
}

// GetUsedQuotaFiles returns the UsedQuotaFiles field value if set, zero value otherwise.
func (o *VirtualFolder) GetUsedQuotaFiles() int32 {
	if o == nil || o.UsedQuotaFiles == nil {
		var ret int32
		return ret
	}
	return *o.UsedQuotaFiles
}

// GetUsedQuotaFilesOk returns a tuple with the UsedQuotaFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetUsedQuotaFilesOk() (*int32, bool) {
	if o == nil || o.UsedQuotaFiles == nil {
		return nil, false
	}
	return o.UsedQuotaFiles, true
}

// HasUsedQuotaFiles returns a boolean if a field has been set.
func (o *VirtualFolder) HasUsedQuotaFiles() bool {
	if o != nil && o.UsedQuotaFiles != nil {
		return true
	}

	return false
}

// SetUsedQuotaFiles gets a reference to the given int32 and assigns it to the UsedQuotaFiles field.
func (o *VirtualFolder) SetUsedQuotaFiles(v int32) {
	o.UsedQuotaFiles = &v
}

// GetLastQuotaUpdate returns the LastQuotaUpdate field value if set, zero value otherwise.
func (o *VirtualFolder) GetLastQuotaUpdate() int64 {
	if o == nil || o.LastQuotaUpdate == nil {
		var ret int64
		return ret
	}
	return *o.LastQuotaUpdate
}

// GetLastQuotaUpdateOk returns a tuple with the LastQuotaUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetLastQuotaUpdateOk() (*int64, bool) {
	if o == nil || o.LastQuotaUpdate == nil {
		return nil, false
	}
	return o.LastQuotaUpdate, true
}

// HasLastQuotaUpdate returns a boolean if a field has been set.
func (o *VirtualFolder) HasLastQuotaUpdate() bool {
	if o != nil && o.LastQuotaUpdate != nil {
		return true
	}

	return false
}

// SetLastQuotaUpdate gets a reference to the given int64 and assigns it to the LastQuotaUpdate field.
func (o *VirtualFolder) SetLastQuotaUpdate(v int64) {
	o.LastQuotaUpdate = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *VirtualFolder) GetUsers() []string {
	if o == nil || o.Users == nil {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetUsersOk() ([]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *VirtualFolder) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *VirtualFolder) SetUsers(v []string) {
	o.Users = v
}

// GetFilesystem returns the Filesystem field value if set, zero value otherwise.
func (o *VirtualFolder) GetFilesystem() FilesystemConfig {
	if o == nil || o.Filesystem == nil {
		var ret FilesystemConfig
		return ret
	}
	return *o.Filesystem
}

// GetFilesystemOk returns a tuple with the Filesystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetFilesystemOk() (*FilesystemConfig, bool) {
	if o == nil || o.Filesystem == nil {
		return nil, false
	}
	return o.Filesystem, true
}

// HasFilesystem returns a boolean if a field has been set.
func (o *VirtualFolder) HasFilesystem() bool {
	if o != nil && o.Filesystem != nil {
		return true
	}

	return false
}

// SetFilesystem gets a reference to the given FilesystemConfig and assigns it to the Filesystem field.
func (o *VirtualFolder) SetFilesystem(v FilesystemConfig) {
	o.Filesystem = &v
}

// GetVirtualPath returns the VirtualPath field value
func (o *VirtualFolder) GetVirtualPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualPath
}

// GetVirtualPathOk returns a tuple with the VirtualPath field value
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetVirtualPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualPath, true
}

// SetVirtualPath sets field value
func (o *VirtualFolder) SetVirtualPath(v string) {
	o.VirtualPath = v
}

// GetQuotaSize returns the QuotaSize field value if set, zero value otherwise.
func (o *VirtualFolder) GetQuotaSize() int64 {
	if o == nil || o.QuotaSize == nil {
		var ret int64
		return ret
	}
	return *o.QuotaSize
}

// GetQuotaSizeOk returns a tuple with the QuotaSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetQuotaSizeOk() (*int64, bool) {
	if o == nil || o.QuotaSize == nil {
		return nil, false
	}
	return o.QuotaSize, true
}

// HasQuotaSize returns a boolean if a field has been set.
func (o *VirtualFolder) HasQuotaSize() bool {
	if o != nil && o.QuotaSize != nil {
		return true
	}

	return false
}

// SetQuotaSize gets a reference to the given int64 and assigns it to the QuotaSize field.
func (o *VirtualFolder) SetQuotaSize(v int64) {
	o.QuotaSize = &v
}

// GetQuotaFiles returns the QuotaFiles field value if set, zero value otherwise.
func (o *VirtualFolder) GetQuotaFiles() int32 {
	if o == nil || o.QuotaFiles == nil {
		var ret int32
		return ret
	}
	return *o.QuotaFiles
}

// GetQuotaFilesOk returns a tuple with the QuotaFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFolder) GetQuotaFilesOk() (*int32, bool) {
	if o == nil || o.QuotaFiles == nil {
		return nil, false
	}
	return o.QuotaFiles, true
}

// HasQuotaFiles returns a boolean if a field has been set.
func (o *VirtualFolder) HasQuotaFiles() bool {
	if o != nil && o.QuotaFiles != nil {
		return true
	}

	return false
}

// SetQuotaFiles gets a reference to the given int32 and assigns it to the QuotaFiles field.
func (o *VirtualFolder) SetQuotaFiles(v int32) {
	o.QuotaFiles = &v
}

func (o VirtualFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.MappedPath != nil {
		toSerialize["mapped_path"] = o.MappedPath
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.UsedQuotaSize != nil {
		toSerialize["used_quota_size"] = o.UsedQuotaSize
	}
	if o.UsedQuotaFiles != nil {
		toSerialize["used_quota_files"] = o.UsedQuotaFiles
	}
	if o.LastQuotaUpdate != nil {
		toSerialize["last_quota_update"] = o.LastQuotaUpdate
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Filesystem != nil {
		toSerialize["filesystem"] = o.Filesystem
	}
	if true {
		toSerialize["virtual_path"] = o.VirtualPath
	}
	if o.QuotaSize != nil {
		toSerialize["quota_size"] = o.QuotaSize
	}
	if o.QuotaFiles != nil {
		toSerialize["quota_files"] = o.QuotaFiles
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualFolder struct {
	value *VirtualFolder
	isSet bool
}

func (v NullableVirtualFolder) Get() *VirtualFolder {
	return v.value
}

func (v *NullableVirtualFolder) Set(val *VirtualFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualFolder(val *VirtualFolder) *NullableVirtualFolder {
	return &NullableVirtualFolder{value: val, isSet: true}
}

func (v NullableVirtualFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
