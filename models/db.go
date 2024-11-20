package models

import "time"

type UserCredentialTable struct {
	UserID    string    `gorm:"primaryKey;column:user_id;"`
	Salt      string    `gorm:"column:salt;"`
	Hash      string    `gorm:"column:hash;"`
	Retries   int       `gorm:"column:retries;"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (UserCredentialTable) TableName() string {
	return "user_credential"
}

type UserTable struct {
	ID            string    `gorm:"primaryKey;column:id;"`
	VerifiedEmail bool      `gorm:"column:is_email_verified;"`
	FirstName     string    `gorm:"column:first_name;"`
	LastName      string    `gorm:"column:last_name;"`
	ContactNo     string    `gorm:"column:contact_no;"`
	CountryCode   string    `gorm:"column:country_code;"`
	Email         string    `gorm:"column:email;"`
	Locale        string    `gorm:"column:locale;"`
	Status        string    `gorm:"column:status;"`
	IsDeleted     bool      `gorm:"column:is_deleted;"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func (UserTable) TableName() string {
	return "user"
}

type UserWithCredentials struct {
	UserTable
	C UserCredentialTable `gorm:"foreignKey:UserID;references:ID"`
}
