package types

import (
	"encoding/json"
	"fmt"
)

// Permission Permissions:   * `*` - all permissions are granted   * `list` - list items is allowed   * `download` - download files is allowed   * `upload` - upload files is allowed   * `overwrite` - overwrite an existing file, while uploading, is allowed. upload permission is required to allow file overwrite   * `delete` - delete files or directories is allowed   * `delete_files` - delete files is allowed   * `delete_dirs` - delete directories is allowed   * `rename` - rename files or directories is allowed   * `rename_files` - rename files is allowed   * `rename_dirs` - rename directories is allowed   * `create_dirs` - create directories is allowed   * `create_symlinks` - create links is allowed   * `chmod` changing file or directory permissions is allowed   * `chown` changing file or directory owner and group is allowed   * `chtimes` changing file or directory access and modification time is allowed
type Permission string

// List of Permission
const (
	PERMISSION_STAR            Permission = "*"
	PERMISSION_LIST            Permission = "list"
	PERMISSION_DOWNLOAD        Permission = "download"
	PERMISSION_UPLOAD          Permission = "upload"
	PERMISSION_OVERWRITE       Permission = "overwrite"
	PERMISSION_DELETE          Permission = "delete"
	PERMISSION_DELETE_FILES    Permission = "delete_files"
	PERMISSION_DELETE_DIRS     Permission = "delete_dirs"
	PERMISSION_RENAME          Permission = "rename"
	PERMISSION_RENAME_FILES    Permission = "rename_files"
	PERMISSION_RENAME_DIRS     Permission = "rename_dirs"
	PERMISSION_CREATE_DIRS     Permission = "create_dirs"
	PERMISSION_CREATE_SYMLINKS Permission = "create_symlinks"
	PERMISSION_CHMOD           Permission = "chmod"
	PERMISSION_CHOWN           Permission = "chown"
	PERMISSION_CHTIMES         Permission = "chtimes"
)

// AllowedPermissionEnumValues All allowed values of Permission enum
var AllowedPermissionEnumValues = []Permission{
	"*",
	"list",
	"download",
	"upload",
	"overwrite",
	"delete",
	"delete_files",
	"delete_dirs",
	"rename",
	"rename_files",
	"rename_dirs",
	"create_dirs",
	"create_symlinks",
	"chmod",
	"chown",
	"chtimes",
}

func (v *Permission) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Permission(value)
	for _, existing := range AllowedPermissionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Permission", value)
}

// NewPermissionFromValue returns a pointer to a valid Permission
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPermissionFromValue(v string) (*Permission, error) {
	ev := Permission(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Permission: valid values are %v", v, AllowedPermissionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Permission) IsValid() bool {
	for _, existing := range AllowedPermissionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Permission value
func (v Permission) Ptr() *Permission {
	return &v
}

type NullablePermission struct {
	value *Permission
	isSet bool
}

func (v NullablePermission) Get() *Permission {
	return v.value
}

func (v *NullablePermission) Set(val *Permission) {
	v.value = val
	v.isSet = true
}

func (v NullablePermission) IsSet() bool {
	return v.isSet
}

func (v *NullablePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermission(val *Permission) *NullablePermission {
	return &NullablePermission{value: val, isSet: true}
}

func (v NullablePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
