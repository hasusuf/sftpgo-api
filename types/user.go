package types

import "encoding/json"

type GetUsersRequest struct {
	Offset int
	Limit  int
	Order  string
}

// User struct for User
type User struct {
	Id *int32 `json:"id,omitempty"`
	// status:   * `0` user is disabled, login is not allowed   * `1` user is enabled
	Status *int32 `json:"status,omitempty"`
	// username is unique
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	// optional description, for example the user full name
	Description *string `json:"description,omitempty"`
	// expiration date as unix timestamp in milliseconds. An expired account cannot login. 0 means no expiration
	ExpirationDate *int64 `json:"expiration_date,omitempty"`
	// password or public key/SSH user certificate are mandatory. If the password has no known hashing algo prefix it will be stored, by default, using bcrypt, argon2id is supported too. You can send a password hashed as bcrypt ($2a$ prefix), argon2id, pbkdf2 or unix crypt and it will be stored as is. For security reasons this field is omitted when you search/get users
	Password *string `json:"password,omitempty"`
	// Public keys in OpenSSH format. A password or at least one public key/SSH user certificate are mandatory.
	PublicKeys []string `json:"public_keys,omitempty"`
	// path to the user home directory. The user cannot upload or download files outside this directory. SFTPGo tries to automatically create this folder if missing. Must be an absolute path
	HomeDir *string `json:"home_dir,omitempty"`
	// mapping between virtual SFTPGo paths and filesystem paths outside the user home directory. Supported for local filesystem only. If one or more of the specified folders are not inside the dataprovider they will be automatically created. You have to create the folder on the filesystem yourself
	VirtualFolders []VirtualFolder `json:"virtual_folders,omitempty"`
	// if you run SFTPGo as root user, the created files and directories will be assigned to this uid. 0 means no change, the owner will be the user that runs SFTPGo. Ignored on windows
	Uid *int32 `json:"uid,omitempty"`
	// if you run SFTPGo as root user, the created files and directories will be assigned to this gid. 0 means no change, the group will be the one of the user that runs SFTPGo. Ignored on windows
	Gid *int32 `json:"gid,omitempty"`
	// Limit the sessions that a user can open. 0 means unlimited
	MaxSessions *int32 `json:"max_sessions,omitempty"`
	// Quota as size in bytes. 0 menas unlimited. Please note that quota is updated if files are added/removed via SFTPGo otherwise a quota scan or a manual quota update is needed
	QuotaSize *int64 `json:"quota_size,omitempty"`
	// Quota as number of files. 0 menas unlimited. Please note that quota is updated if files are added/removed via SFTPGo otherwise a quota scan or a manual quota update is needed
	QuotaFiles     *int32                  `json:"quota_files,omitempty"`
	Permissions    map[string][]Permission `json:"permissions,omitempty"`
	UsedQuotaSize  *int64                  `json:"used_quota_size,omitempty"`
	UsedQuotaFiles *int32                  `json:"used_quota_files,omitempty"`
	// Last quota update as unix timestamp in milliseconds
	LastQuotaUpdate *int64 `json:"last_quota_update,omitempty"`
	// Maximum upload bandwidth as KB/s, 0 means unlimited
	UploadBandwidth *int32 `json:"upload_bandwidth,omitempty"`
	// Maximum download bandwidth as KB/s, 0 means unlimited
	DownloadBandwidth *int32 `json:"download_bandwidth,omitempty"`
	// Maximum data transfer allowed for uploads as MB. 0 means no limit
	UploadDataTransfer *int32 `json:"upload_data_transfer,omitempty"`
	// Maximum data transfer allowed for downloads as MB. 0 means no limit
	DownloadDataTransfer *int32 `json:"download_data_transfer,omitempty"`
	// Maximum total data transfer as MB. 0 means unlimited. You can set a total data transfer instead of the individual values for uploads and downloads
	TotalDataTransfer *int32 `json:"total_data_transfer,omitempty"`
	// Uploaded size, as bytes, since the last reset
	UsedUploadDataTransfer *int32 `json:"used_upload_data_transfer,omitempty"`
	// Downloaded size, as bytes, since the last reset
	UsedDownloadDataTransfer *int32 `json:"used_download_data_transfer,omitempty"`
	// creation time as unix timestamp in milliseconds. It will be 0 for users created before v2.2.0
	CreatedAt *int64 `json:"created_at,omitempty"`
	// last update time as unix timestamp in milliseconds
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// Last user login as unix timestamp in milliseconds. It is saved at most once every 10 minutes
	LastLogin  *int64            `json:"last_login,omitempty"`
	Filters    *UserFilters      `json:"filters,omitempty"`
	Filesystem *FilesystemConfig `json:"filesystem,omitempty"`
	// Free form text field for external systems
	AdditionalInfo *string `json:"additional_info,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *User) SetId(v int32) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *User) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *User) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *User) SetStatus(v int32) {
	o.Status = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *User) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *User) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *User) SetDescription(v string) {
	o.Description = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *User) GetExpirationDate() int64 {
	if o == nil || o.ExpirationDate == nil {
		var ret int64
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetExpirationDateOk() (*int64, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *User) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *User) SetExpirationDate(v int64) {
	o.ExpirationDate = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *User) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *User) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *User) SetPassword(v string) {
	o.Password = &v
}

// GetPublicKeys returns the PublicKeys field value if set, zero value otherwise.
func (o *User) GetPublicKeys() []string {
	if o == nil || o.PublicKeys == nil {
		var ret []string
		return ret
	}
	return o.PublicKeys
}

// GetPublicKeysOk returns a tuple with the PublicKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPublicKeysOk() ([]string, bool) {
	if o == nil || o.PublicKeys == nil {
		return nil, false
	}
	return o.PublicKeys, true
}

// HasPublicKeys returns a boolean if a field has been set.
func (o *User) HasPublicKeys() bool {
	if o != nil && o.PublicKeys != nil {
		return true
	}

	return false
}

// SetPublicKeys gets a reference to the given []string and assigns it to the PublicKeys field.
func (o *User) SetPublicKeys(v []string) {
	o.PublicKeys = v
}

// GetHomeDir returns the HomeDir field value if set, zero value otherwise.
func (o *User) GetHomeDir() string {
	if o == nil || o.HomeDir == nil {
		var ret string
		return ret
	}
	return *o.HomeDir
}

// GetHomeDirOk returns a tuple with the HomeDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetHomeDirOk() (*string, bool) {
	if o == nil || o.HomeDir == nil {
		return nil, false
	}
	return o.HomeDir, true
}

// HasHomeDir returns a boolean if a field has been set.
func (o *User) HasHomeDir() bool {
	if o != nil && o.HomeDir != nil {
		return true
	}

	return false
}

// SetHomeDir gets a reference to the given string and assigns it to the HomeDir field.
func (o *User) SetHomeDir(v string) {
	o.HomeDir = &v
}

// GetVirtualFolders returns the VirtualFolders field value if set, zero value otherwise.
func (o *User) GetVirtualFolders() []VirtualFolder {
	if o == nil || o.VirtualFolders == nil {
		var ret []VirtualFolder
		return ret
	}
	return o.VirtualFolders
}

// GetVirtualFoldersOk returns a tuple with the VirtualFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVirtualFoldersOk() ([]VirtualFolder, bool) {
	if o == nil || o.VirtualFolders == nil {
		return nil, false
	}
	return o.VirtualFolders, true
}

// HasVirtualFolders returns a boolean if a field has been set.
func (o *User) HasVirtualFolders() bool {
	if o != nil && o.VirtualFolders != nil {
		return true
	}

	return false
}

// SetVirtualFolders gets a reference to the given []VirtualFolder and assigns it to the VirtualFolders field.
func (o *User) SetVirtualFolders(v []VirtualFolder) {
	o.VirtualFolders = v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *User) GetUid() int32 {
	if o == nil || o.Uid == nil {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUidOk() (*int32, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *User) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *User) SetUid(v int32) {
	o.Uid = &v
}

// GetGid returns the Gid field value if set, zero value otherwise.
func (o *User) GetGid() int32 {
	if o == nil || o.Gid == nil {
		var ret int32
		return ret
	}
	return *o.Gid
}

// GetGidOk returns a tuple with the Gid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGidOk() (*int32, bool) {
	if o == nil || o.Gid == nil {
		return nil, false
	}
	return o.Gid, true
}

// HasGid returns a boolean if a field has been set.
func (o *User) HasGid() bool {
	if o != nil && o.Gid != nil {
		return true
	}

	return false
}

// SetGid gets a reference to the given int32 and assigns it to the Gid field.
func (o *User) SetGid(v int32) {
	o.Gid = &v
}

// GetMaxSessions returns the MaxSessions field value if set, zero value otherwise.
func (o *User) GetMaxSessions() int32 {
	if o == nil || o.MaxSessions == nil {
		var ret int32
		return ret
	}
	return *o.MaxSessions
}

// GetMaxSessionsOk returns a tuple with the MaxSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetMaxSessionsOk() (*int32, bool) {
	if o == nil || o.MaxSessions == nil {
		return nil, false
	}
	return o.MaxSessions, true
}

// HasMaxSessions returns a boolean if a field has been set.
func (o *User) HasMaxSessions() bool {
	if o != nil && o.MaxSessions != nil {
		return true
	}

	return false
}

// SetMaxSessions gets a reference to the given int32 and assigns it to the MaxSessions field.
func (o *User) SetMaxSessions(v int32) {
	o.MaxSessions = &v
}

// GetQuotaSize returns the QuotaSize field value if set, zero value otherwise.
func (o *User) GetQuotaSize() int64 {
	if o == nil || o.QuotaSize == nil {
		var ret int64
		return ret
	}
	return *o.QuotaSize
}

// GetQuotaSizeOk returns a tuple with the QuotaSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetQuotaSizeOk() (*int64, bool) {
	if o == nil || o.QuotaSize == nil {
		return nil, false
	}
	return o.QuotaSize, true
}

// HasQuotaSize returns a boolean if a field has been set.
func (o *User) HasQuotaSize() bool {
	if o != nil && o.QuotaSize != nil {
		return true
	}

	return false
}

// SetQuotaSize gets a reference to the given int64 and assigns it to the QuotaSize field.
func (o *User) SetQuotaSize(v int64) {
	o.QuotaSize = &v
}

// GetQuotaFiles returns the QuotaFiles field value if set, zero value otherwise.
func (o *User) GetQuotaFiles() int32 {
	if o == nil || o.QuotaFiles == nil {
		var ret int32
		return ret
	}
	return *o.QuotaFiles
}

// GetQuotaFilesOk returns a tuple with the QuotaFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetQuotaFilesOk() (*int32, bool) {
	if o == nil || o.QuotaFiles == nil {
		return nil, false
	}
	return o.QuotaFiles, true
}

// HasQuotaFiles returns a boolean if a field has been set.
func (o *User) HasQuotaFiles() bool {
	if o != nil && o.QuotaFiles != nil {
		return true
	}

	return false
}

// SetQuotaFiles gets a reference to the given int32 and assigns it to the QuotaFiles field.
func (o *User) SetQuotaFiles(v int32) {
	o.QuotaFiles = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *User) GetPermissions() map[string][]Permission {
	if o == nil || o.Permissions == nil {
		var ret map[string][]Permission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPermissionsOk() (map[string][]Permission, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *User) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []map[string][]Permission and assigns it to the Permissions field.
func (o *User) SetPermissions(v map[string][]Permission) {
	o.Permissions = v
}

// GetUsedQuotaSize returns the UsedQuotaSize field value if set, zero value otherwise.
func (o *User) GetUsedQuotaSize() int64 {
	if o == nil || o.UsedQuotaSize == nil {
		var ret int64
		return ret
	}
	return *o.UsedQuotaSize
}

// GetUsedQuotaSizeOk returns a tuple with the UsedQuotaSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsedQuotaSizeOk() (*int64, bool) {
	if o == nil || o.UsedQuotaSize == nil {
		return nil, false
	}
	return o.UsedQuotaSize, true
}

// HasUsedQuotaSize returns a boolean if a field has been set.
func (o *User) HasUsedQuotaSize() bool {
	if o != nil && o.UsedQuotaSize != nil {
		return true
	}

	return false
}

// SetUsedQuotaSize gets a reference to the given int64 and assigns it to the UsedQuotaSize field.
func (o *User) SetUsedQuotaSize(v int64) {
	o.UsedQuotaSize = &v
}

// GetUsedQuotaFiles returns the UsedQuotaFiles field value if set, zero value otherwise.
func (o *User) GetUsedQuotaFiles() int32 {
	if o == nil || o.UsedQuotaFiles == nil {
		var ret int32
		return ret
	}
	return *o.UsedQuotaFiles
}

// GetUsedQuotaFilesOk returns a tuple with the UsedQuotaFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsedQuotaFilesOk() (*int32, bool) {
	if o == nil || o.UsedQuotaFiles == nil {
		return nil, false
	}
	return o.UsedQuotaFiles, true
}

// HasUsedQuotaFiles returns a boolean if a field has been set.
func (o *User) HasUsedQuotaFiles() bool {
	if o != nil && o.UsedQuotaFiles != nil {
		return true
	}

	return false
}

// SetUsedQuotaFiles gets a reference to the given int32 and assigns it to the UsedQuotaFiles field.
func (o *User) SetUsedQuotaFiles(v int32) {
	o.UsedQuotaFiles = &v
}

// GetLastQuotaUpdate returns the LastQuotaUpdate field value if set, zero value otherwise.
func (o *User) GetLastQuotaUpdate() int64 {
	if o == nil || o.LastQuotaUpdate == nil {
		var ret int64
		return ret
	}
	return *o.LastQuotaUpdate
}

// GetLastQuotaUpdateOk returns a tuple with the LastQuotaUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastQuotaUpdateOk() (*int64, bool) {
	if o == nil || o.LastQuotaUpdate == nil {
		return nil, false
	}
	return o.LastQuotaUpdate, true
}

// HasLastQuotaUpdate returns a boolean if a field has been set.
func (o *User) HasLastQuotaUpdate() bool {
	if o != nil && o.LastQuotaUpdate != nil {
		return true
	}

	return false
}

// SetLastQuotaUpdate gets a reference to the given int64 and assigns it to the LastQuotaUpdate field.
func (o *User) SetLastQuotaUpdate(v int64) {
	o.LastQuotaUpdate = &v
}

// GetUploadBandwidth returns the UploadBandwidth field value if set, zero value otherwise.
func (o *User) GetUploadBandwidth() int32 {
	if o == nil || o.UploadBandwidth == nil {
		var ret int32
		return ret
	}
	return *o.UploadBandwidth
}

// GetUploadBandwidthOk returns a tuple with the UploadBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUploadBandwidthOk() (*int32, bool) {
	if o == nil || o.UploadBandwidth == nil {
		return nil, false
	}
	return o.UploadBandwidth, true
}

// HasUploadBandwidth returns a boolean if a field has been set.
func (o *User) HasUploadBandwidth() bool {
	if o != nil && o.UploadBandwidth != nil {
		return true
	}

	return false
}

// SetUploadBandwidth gets a reference to the given int32 and assigns it to the UploadBandwidth field.
func (o *User) SetUploadBandwidth(v int32) {
	o.UploadBandwidth = &v
}

// GetDownloadBandwidth returns the DownloadBandwidth field value if set, zero value otherwise.
func (o *User) GetDownloadBandwidth() int32 {
	if o == nil || o.DownloadBandwidth == nil {
		var ret int32
		return ret
	}
	return *o.DownloadBandwidth
}

// GetDownloadBandwidthOk returns a tuple with the DownloadBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDownloadBandwidthOk() (*int32, bool) {
	if o == nil || o.DownloadBandwidth == nil {
		return nil, false
	}
	return o.DownloadBandwidth, true
}

// HasDownloadBandwidth returns a boolean if a field has been set.
func (o *User) HasDownloadBandwidth() bool {
	if o != nil && o.DownloadBandwidth != nil {
		return true
	}

	return false
}

// SetDownloadBandwidth gets a reference to the given int32 and assigns it to the DownloadBandwidth field.
func (o *User) SetDownloadBandwidth(v int32) {
	o.DownloadBandwidth = &v
}

// GetUploadDataTransfer returns the UploadDataTransfer field value if set, zero value otherwise.
func (o *User) GetUploadDataTransfer() int32 {
	if o == nil || o.UploadDataTransfer == nil {
		var ret int32
		return ret
	}
	return *o.UploadDataTransfer
}

// GetUploadDataTransferOk returns a tuple with the UploadDataTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUploadDataTransferOk() (*int32, bool) {
	if o == nil || o.UploadDataTransfer == nil {
		return nil, false
	}
	return o.UploadDataTransfer, true
}

// HasUploadDataTransfer returns a boolean if a field has been set.
func (o *User) HasUploadDataTransfer() bool {
	if o != nil && o.UploadDataTransfer != nil {
		return true
	}

	return false
}

// SetUploadDataTransfer gets a reference to the given int32 and assigns it to the UploadDataTransfer field.
func (o *User) SetUploadDataTransfer(v int32) {
	o.UploadDataTransfer = &v
}

// GetDownloadDataTransfer returns the DownloadDataTransfer field value if set, zero value otherwise.
func (o *User) GetDownloadDataTransfer() int32 {
	if o == nil || o.DownloadDataTransfer == nil {
		var ret int32
		return ret
	}
	return *o.DownloadDataTransfer
}

// GetDownloadDataTransferOk returns a tuple with the DownloadDataTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDownloadDataTransferOk() (*int32, bool) {
	if o == nil || o.DownloadDataTransfer == nil {
		return nil, false
	}
	return o.DownloadDataTransfer, true
}

// HasDownloadDataTransfer returns a boolean if a field has been set.
func (o *User) HasDownloadDataTransfer() bool {
	if o != nil && o.DownloadDataTransfer != nil {
		return true
	}

	return false
}

// SetDownloadDataTransfer gets a reference to the given int32 and assigns it to the DownloadDataTransfer field.
func (o *User) SetDownloadDataTransfer(v int32) {
	o.DownloadDataTransfer = &v
}

// GetTotalDataTransfer returns the TotalDataTransfer field value if set, zero value otherwise.
func (o *User) GetTotalDataTransfer() int32 {
	if o == nil || o.TotalDataTransfer == nil {
		var ret int32
		return ret
	}
	return *o.TotalDataTransfer
}

// GetTotalDataTransferOk returns a tuple with the TotalDataTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTotalDataTransferOk() (*int32, bool) {
	if o == nil || o.TotalDataTransfer == nil {
		return nil, false
	}
	return o.TotalDataTransfer, true
}

// HasTotalDataTransfer returns a boolean if a field has been set.
func (o *User) HasTotalDataTransfer() bool {
	if o != nil && o.TotalDataTransfer != nil {
		return true
	}

	return false
}

// SetTotalDataTransfer gets a reference to the given int32 and assigns it to the TotalDataTransfer field.
func (o *User) SetTotalDataTransfer(v int32) {
	o.TotalDataTransfer = &v
}

// GetUsedUploadDataTransfer returns the UsedUploadDataTransfer field value if set, zero value otherwise.
func (o *User) GetUsedUploadDataTransfer() int32 {
	if o == nil || o.UsedUploadDataTransfer == nil {
		var ret int32
		return ret
	}
	return *o.UsedUploadDataTransfer
}

// GetUsedUploadDataTransferOk returns a tuple with the UsedUploadDataTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsedUploadDataTransferOk() (*int32, bool) {
	if o == nil || o.UsedUploadDataTransfer == nil {
		return nil, false
	}
	return o.UsedUploadDataTransfer, true
}

// HasUsedUploadDataTransfer returns a boolean if a field has been set.
func (o *User) HasUsedUploadDataTransfer() bool {
	if o != nil && o.UsedUploadDataTransfer != nil {
		return true
	}

	return false
}

// SetUsedUploadDataTransfer gets a reference to the given int32 and assigns it to the UsedUploadDataTransfer field.
func (o *User) SetUsedUploadDataTransfer(v int32) {
	o.UsedUploadDataTransfer = &v
}

// GetUsedDownloadDataTransfer returns the UsedDownloadDataTransfer field value if set, zero value otherwise.
func (o *User) GetUsedDownloadDataTransfer() int32 {
	if o == nil || o.UsedDownloadDataTransfer == nil {
		var ret int32
		return ret
	}
	return *o.UsedDownloadDataTransfer
}

// GetUsedDownloadDataTransferOk returns a tuple with the UsedDownloadDataTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsedDownloadDataTransferOk() (*int32, bool) {
	if o == nil || o.UsedDownloadDataTransfer == nil {
		return nil, false
	}
	return o.UsedDownloadDataTransfer, true
}

// HasUsedDownloadDataTransfer returns a boolean if a field has been set.
func (o *User) HasUsedDownloadDataTransfer() bool {
	if o != nil && o.UsedDownloadDataTransfer != nil {
		return true
	}

	return false
}

// SetUsedDownloadDataTransfer gets a reference to the given int32 and assigns it to the UsedDownloadDataTransfer field.
func (o *User) SetUsedDownloadDataTransfer(v int32) {
	o.UsedDownloadDataTransfer = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *User) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *User) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *User) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *User) GetLastLogin() int64 {
	if o == nil || o.LastLogin == nil {
		var ret int64
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginOk() (*int64, bool) {
	if o == nil || o.LastLogin == nil {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *User) HasLastLogin() bool {
	if o != nil && o.LastLogin != nil {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given int64 and assigns it to the LastLogin field.
func (o *User) SetLastLogin(v int64) {
	o.LastLogin = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *User) GetFilters() UserFilters {
	if o == nil || o.Filters == nil {
		var ret UserFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFiltersOk() (*UserFilters, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *User) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given UserFilters and assigns it to the Filters field.
func (o *User) SetFilters(v UserFilters) {
	o.Filters = &v
}

// GetFilesystem returns the Filesystem field value if set, zero value otherwise.
func (o *User) GetFilesystem() FilesystemConfig {
	if o == nil || o.Filesystem == nil {
		var ret FilesystemConfig
		return ret
	}
	return *o.Filesystem
}

// GetFilesystemOk returns a tuple with the Filesystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFilesystemOk() (*FilesystemConfig, bool) {
	if o == nil || o.Filesystem == nil {
		return nil, false
	}
	return o.Filesystem, true
}

// HasFilesystem returns a boolean if a field has been set.
func (o *User) HasFilesystem() bool {
	if o != nil && o.Filesystem != nil {
		return true
	}

	return false
}

// SetFilesystem gets a reference to the given FilesystemConfig and assigns it to the Filesystem field.
func (o *User) SetFilesystem(v FilesystemConfig) {
	o.Filesystem = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *User) GetAdditionalInfo() string {
	if o == nil || o.AdditionalInfo == nil {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || o.AdditionalInfo == nil {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *User) HasAdditionalInfo() bool {
	if o != nil && o.AdditionalInfo != nil {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *User) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PublicKeys != nil {
		toSerialize["public_keys"] = o.PublicKeys
	}
	if o.HomeDir != nil {
		toSerialize["home_dir"] = o.HomeDir
	}
	if o.VirtualFolders != nil {
		toSerialize["virtual_folders"] = o.VirtualFolders
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if o.Gid != nil {
		toSerialize["gid"] = o.Gid
	}
	if o.MaxSessions != nil {
		toSerialize["max_sessions"] = o.MaxSessions
	}
	if o.QuotaSize != nil {
		toSerialize["quota_size"] = o.QuotaSize
	}
	if o.QuotaFiles != nil {
		toSerialize["quota_files"] = o.QuotaFiles
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
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
	if o.UploadBandwidth != nil {
		toSerialize["upload_bandwidth"] = o.UploadBandwidth
	}
	if o.DownloadBandwidth != nil {
		toSerialize["download_bandwidth"] = o.DownloadBandwidth
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
	if o.UsedUploadDataTransfer != nil {
		toSerialize["used_upload_data_transfer"] = o.UsedUploadDataTransfer
	}
	if o.UsedDownloadDataTransfer != nil {
		toSerialize["used_download_data_transfer"] = o.UsedDownloadDataTransfer
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.LastLogin != nil {
		toSerialize["last_login"] = o.LastLogin
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Filesystem != nil {
		toSerialize["filesystem"] = o.Filesystem
	}
	if o.AdditionalInfo != nil {
		toSerialize["additional_info"] = o.AdditionalInfo
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
