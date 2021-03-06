// Code generated by go-queryset. DO NOT EDIT.
package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set UserQuerySet

// UserQuerySet is an queryset type for User
type UserQuerySet struct {
	db *gorm.DB
}

// NewUserQuerySet constructs new UserQuerySet
func NewUserQuerySet(db *gorm.DB) UserQuerySet {
	return UserQuerySet{
		db: db.Model(&User{}),
	}
}

func (qs UserQuerySet) w(db *gorm.DB) UserQuerySet {
	return NewUserQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) All(ret *[]User) error {
	return qs.db.Find(ret).Error
}

// AvatarEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) AvatarEq(avatar string) UserQuerySet {
	return qs.w(qs.db.Where("avatar = ?", avatar))
}

// AvatarIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) AvatarIn(avatar ...string) UserQuerySet {
	if len(avatar) == 0 {
		qs.db.AddError(errors.New("must at least pass one avatar in AvatarIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("avatar IN (?)", avatar))
}

// AvatarNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) AvatarNe(avatar string) UserQuerySet {
	return qs.w(qs.db.Where("avatar != ?", avatar))
}

// AvatarNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) AvatarNotIn(avatar ...string) UserQuerySet {
	if len(avatar) == 0 {
		qs.db.AddError(errors.New("must at least pass one avatar in AvatarNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("avatar NOT IN (?)", avatar))
}

// Count is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *User) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtEq(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtGt(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtGte(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtLt(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtLte(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtNe(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *User) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Delete() error {
	return qs.db.Delete(User{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(User{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(User{})
	return db.RowsAffected, db.Error
}

// EmailEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) EmailEq(email string) UserQuerySet {
	return qs.w(qs.db.Where("email = ?", email))
}

// EmailIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) EmailIn(email ...string) UserQuerySet {
	if len(email) == 0 {
		qs.db.AddError(errors.New("must at least pass one email in EmailIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("email IN (?)", email))
}

// EmailNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) EmailNe(email string) UserQuerySet {
	return qs.w(qs.db.Where("email != ?", email))
}

// EmailNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) EmailNotIn(email ...string) UserQuerySet {
	if len(email) == 0 {
		qs.db.AddError(errors.New("must at least pass one email in EmailNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("email NOT IN (?)", email))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) GetUpdater() UserUpdater {
	return NewUserUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDEq(ID uint64) UserQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDGt(ID uint64) UserQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDGte(ID uint64) UserQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDIn(ID ...uint64) UserQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDLt(ID uint64) UserQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDLte(ID uint64) UserQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDNe(ID uint64) UserQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDNotIn(ID ...uint64) UserQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// IntroductionEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IntroductionEq(introduction string) UserQuerySet {
	return qs.w(qs.db.Where("introduction = ?", introduction))
}

// IntroductionIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IntroductionIn(introduction ...string) UserQuerySet {
	if len(introduction) == 0 {
		qs.db.AddError(errors.New("must at least pass one introduction in IntroductionIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("introduction IN (?)", introduction))
}

// IntroductionNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IntroductionNe(introduction string) UserQuerySet {
	return qs.w(qs.db.Where("introduction != ?", introduction))
}

// IntroductionNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IntroductionNotIn(introduction ...string) UserQuerySet {
	if len(introduction) == 0 {
		qs.db.AddError(errors.New("must at least pass one introduction in IntroductionNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("introduction NOT IN (?)", introduction))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Limit(limit int) UserQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Offset(offset int) UserQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs UserQuerySet) One(ret *User) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByCreatedAt() UserQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByID() UserQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByUpdatedAt() UserQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByCreatedAt() UserQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByID() UserQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByUpdatedAt() UserQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// PasswordEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) PasswordEq(password string) UserQuerySet {
	return qs.w(qs.db.Where("password = ?", password))
}

// PasswordIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) PasswordIn(password ...string) UserQuerySet {
	if len(password) == 0 {
		qs.db.AddError(errors.New("must at least pass one password in PasswordIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("password IN (?)", password))
}

// PasswordNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) PasswordNe(password string) UserQuerySet {
	return qs.w(qs.db.Where("password != ?", password))
}

// PasswordNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) PasswordNotIn(password ...string) UserQuerySet {
	if len(password) == 0 {
		qs.db.AddError(errors.New("must at least pass one password in PasswordNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("password NOT IN (?)", password))
}

// RolesEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RolesEq(roles string) UserQuerySet {
	return qs.w(qs.db.Where("roles = ?", roles))
}

// RolesIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RolesIn(roles ...string) UserQuerySet {
	if len(roles) == 0 {
		qs.db.AddError(errors.New("must at least pass one roles in RolesIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("roles IN (?)", roles))
}

// RolesNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RolesNe(roles string) UserQuerySet {
	return qs.w(qs.db.Where("roles != ?", roles))
}

// RolesNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RolesNotIn(roles ...string) UserQuerySet {
	if len(roles) == 0 {
		qs.db.AddError(errors.New("must at least pass one roles in RolesNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("roles NOT IN (?)", roles))
}

// SetAvatar is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetAvatar(avatar string) UserUpdater {
	u.fields[string(UserDBSchema.Avatar)] = avatar
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetCreatedAt(createdAt time.Time) UserUpdater {
	u.fields[string(UserDBSchema.CreatedAt)] = createdAt
	return u
}

// SetEmail is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetEmail(email string) UserUpdater {
	u.fields[string(UserDBSchema.Email)] = email
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetID(ID uint64) UserUpdater {
	u.fields[string(UserDBSchema.ID)] = ID
	return u
}

// SetIntroduction is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetIntroduction(introduction string) UserUpdater {
	u.fields[string(UserDBSchema.Introduction)] = introduction
	return u
}

// SetPassword is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetPassword(password string) UserUpdater {
	u.fields[string(UserDBSchema.Password)] = password
	return u
}

// SetRoles is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetRoles(roles string) UserUpdater {
	u.fields[string(UserDBSchema.Roles)] = roles
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetUpdatedAt(updatedAt time.Time) UserUpdater {
	u.fields[string(UserDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUserName is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetUserName(userName string) UserUpdater {
	u.fields[string(UserDBSchema.UserName)] = userName
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u UserUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u UserUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtEq(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtGt(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtGte(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtLt(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtLte(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtNe(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// UserNameEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UserNameEq(userName string) UserQuerySet {
	return qs.w(qs.db.Where("user_name = ?", userName))
}

// UserNameIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UserNameIn(userName ...string) UserQuerySet {
	if len(userName) == 0 {
		qs.db.AddError(errors.New("must at least pass one userName in UserNameIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("user_name IN (?)", userName))
}

// UserNameNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UserNameNe(userName string) UserQuerySet {
	return qs.w(qs.db.Where("user_name != ?", userName))
}

// UserNameNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UserNameNotIn(userName ...string) UserQuerySet {
	if len(userName) == 0 {
		qs.db.AddError(errors.New("must at least pass one userName in UserNameNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("user_name NOT IN (?)", userName))
}

// ===== END of query set UserQuerySet

// ===== BEGIN of User modifiers

// UserDBSchemaField describes database schema field. It requires for method 'Update'
type UserDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f UserDBSchemaField) String() string {
	return string(f)
}

// UserDBSchema stores db field names of User
var UserDBSchema = struct {
	ID           UserDBSchemaField
	UserName     UserDBSchemaField
	Password     UserDBSchemaField
	Email        UserDBSchemaField
	Roles        UserDBSchemaField
	Introduction UserDBSchemaField
	Avatar       UserDBSchemaField
	CreatedAt    UserDBSchemaField
	UpdatedAt    UserDBSchemaField
}{

	ID:           UserDBSchemaField("id"),
	UserName:     UserDBSchemaField("user_name"),
	Password:     UserDBSchemaField("password"),
	Email:        UserDBSchemaField("email"),
	Roles:        UserDBSchemaField("roles"),
	Introduction: UserDBSchemaField("introduction"),
	Avatar:       UserDBSchemaField("avatar"),
	CreatedAt:    UserDBSchemaField("created_at"),
	UpdatedAt:    UserDBSchemaField("updated_at"),
}

// Update updates User fields by primary key
// nolint: dupl
func (o *User) Update(db *gorm.DB, fields ...UserDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":           o.ID,
		"user_name":    o.UserName,
		"password":     o.Password,
		"email":        o.Email,
		"roles":        o.Roles,
		"introduction": o.Introduction,
		"avatar":       o.Avatar,
		"created_at":   o.CreatedAt,
		"updated_at":   o.UpdatedAt,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update User %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// UserUpdater is an User updates manager
type UserUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewUserUpdater creates new User updater
// nolint: dupl
func NewUserUpdater(db *gorm.DB) UserUpdater {
	return UserUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&User{}),
	}
}

// ===== END of User modifiers

// ===== END of all query sets
