package models

import (
	"encoding/json"
)

type User struct {
	UID    uint     `gorm:"primaryKey" json:"uid"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Status int      `json:"status"`
	Groups []*Group `gorm:"many2many:user_groups;joinForeignKey:UserUID;joinReferences:GroupUID" json:"groups"`
	RoleID uint     `json:"role_id"`
	Role   Role     `json:"role"`
}

type Group struct {
	UID   uint    `gorm:"primaryKey" json:"uid"`
	Name  string  `json:"name"`
	Users []*User `gorm:"many2many:user_groups;joinForeignKey:GroupUID;joinReferences:UserUID" json:"users"`
}

type Role struct {
	UID    uint            `gorm:"primaryKey" json:"uid"`
	Name   string          `json:"name"`
	Rights json.RawMessage `json:"rights"`
}

type UserGroup struct {
	UserUID  uint `gorm:"primaryKey" json:"user_uid"`
	GroupUID uint `gorm:"primaryKey" json:"group_uid"`
}
