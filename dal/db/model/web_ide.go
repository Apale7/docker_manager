package model

import (
	"gorm.io/gorm"
)

// Container [...]
type Container struct {
	gorm.Model
	ContainerID string `gorm:"unique;column:container_id;type:varchar(80);not null"`         // docker容器id
	Status      int8   `gorm:"column:status;type:tinyint(4);not null"`                       // Docker容器状态 1-Running 2-Paused 3-Restarting 4-OOMKilled 5-Dead 6-UNKNOWN
	ImageID     string `gorm:"index:idx_image_id;column:image_id;type:varchar(80);not null"` // docker镜像id
	Name        string `gorm:"column:name;type:varchar(32);not null"`                        // docker容器名字
}

// TableName get sql table name.获取数据库表名
func (m *Container) TableName() string {
	return "container"
}

// ContainerColumns get sql column name.获取数据库列名
var ContainerColumns = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	ContainerID string
	Status      string
	ImageID     string
	Name        string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	ContainerID: "container_id",
	Status:      "status",
	ImageID:     "image_id",
	Name:        "name",
}

// Image [...]
type Image struct {
	gorm.Model
	ImageID   string `gorm:"unique;column:image_id;type:varchar(80);not null"` // docker镜像id
	ImageSize int64  `gorm:"column:image_size;type:bigint(20);not null"`       // 镜像大小，单位byte
	Author    string `gorm:"column:author;type:varchar(32);not null"`          // 镜像作者
	RepoTags  string `gorm:"column:repo_tags;type:text"`                       // 仓库标签，以json数组的形式存储
}

// TableName get sql table name.获取数据库表名
func (m *Image) TableName() string {
	return "image"
}

// ImageColumns get sql column name.获取数据库列名
var ImageColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	ImageID   string
	ImageSize string
	Author    string
	RepoTags  string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	ImageID:   "image_id",
	ImageSize: "image_size",
	Author:    "author",
	RepoTags:  "repo_tags",
}

// Resource [...]
type Resource struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(16);not null"` // 资源名称
}

// TableName get sql table name.获取数据库表名
func (m *Resource) TableName() string {
	return "resource"
}

// ResourceColumns get sql column name.获取数据库列名
var ResourceColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Name:      "name",
}

// Role [...]
type Role struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(16);not null"` // 角色名称
}

// TableName get sql table name.获取数据库表名
func (m *Role) TableName() string {
	return "role"
}

// RoleColumns get sql column name.获取数据库列名
var RoleColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Name:      "name",
}

// RoleResource [...]
type RoleResource struct {
	gorm.Model
	RoleID     uint32 `gorm:"index:idx_user_id;column:role_id;type:int(11) unsigned;not null"` // role表的id
	ResourceID uint32 `gorm:"column:resource_id;type:int(11) unsigned;not null"`               // resource表的id
}

// TableName get sql table name.获取数据库表名
func (m *RoleResource) TableName() string {
	return "role_resource"
}

// RoleResourceColumns get sql column name.获取数据库列名
var RoleResourceColumns = struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	RoleID     string
	ResourceID string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	RoleID:     "role_id",
	ResourceID: "resource_id",
}

// User [...]
type User struct {
	gorm.Model
	Username string `gorm:"unique;unique;column:username;type:varchar(16);not null"` // 用户名
	Password string `gorm:"column:password;type:varchar(32);not null"`               // 密码
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Username  string
	Password  string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Username:  "username",
	Password:  "password",
}

// UserContainer [...]
type UserContainer struct {
	gorm.Model
	UserID      uint32 `gorm:"index:idx_user_id;column:user_id;type:int(11) unsigned;not null"` // user表的id
	ContainerID string `gorm:"column:container_id;type:varchar(80);not null"`                   // docker中的container_id
}

// TableName get sql table name.获取数据库表名
func (m *UserContainer) TableName() string {
	return "user_container"
}

// UserContainerColumns get sql column name.获取数据库列名
var UserContainerColumns = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	UserID      string
	ContainerID string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	UserID:      "user_id",
	ContainerID: "container_id",
}

// UserExtra [...]
type UserExtra struct {
	gorm.Model
	UserID      uint32 `gorm:"unique;column:user_id;type:int(11) unsigned;not null"` // user表的id
	Nickname    string `gorm:"column:nickname;type:varchar(16);not null"`            // 昵称
	PhoneNumber string `gorm:"column:phone_number;type:varchar(16)"`                 // 电话号码
	Email       string `gorm:"column:email;type:varchar(32)"`                        // 邮箱
	AvatarURL   string `gorm:"column:avatar_url;type:varchar(255)"`                  // 头像链接
}

// TableName get sql table name.获取数据库表名
func (m *UserExtra) TableName() string {
	return "user_extra"
}

// UserExtraColumns get sql column name.获取数据库列名
var UserExtraColumns = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	UserID      string
	Nickname    string
	PhoneNumber string
	Email       string
	AvatarURL   string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	UserID:      "user_id",
	Nickname:    "nickname",
	PhoneNumber: "phone_number",
	Email:       "email",
	AvatarURL:   "avatar_url",
}

// UserImage [...]
type UserImage struct {
	gorm.Model
	UserID  uint32 `gorm:"index:idx_user_id;column:user_id;type:int(11) unsigned;not null"` // user表的id
	ImageID string `gorm:"column:image_id;type:varchar(80);not null"`                       // docker中的image_id
}

// TableName get sql table name.获取数据库表名
func (m *UserImage) TableName() string {
	return "user_image"
}

// UserImageColumns get sql column name.获取数据库列名
var UserImageColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	UserID    string
	ImageID   string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	UserID:    "user_id",
	ImageID:   "image_id",
}

// UserRole [...]
type UserRole struct {
	gorm.Model
	UserID uint32 `gorm:"index:idx_user_id;column:user_id;type:int(11) unsigned;not null"` // user表的id
	RoleID uint32 `gorm:"column:role_id;type:int(11) unsigned;not null"`                   // role表的id
}

// TableName get sql table name.获取数据库表名
func (m *UserRole) TableName() string {
	return "user_role"
}

// UserRoleColumns get sql column name.获取数据库列名
var UserRoleColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	UserID    string
	RoleID    string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	UserID:    "user_id",
	RoleID:    "role_id",
}
