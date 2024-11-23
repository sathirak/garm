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

type MailTemplateTable struct {
	ID        int       `gorm:"primaryKey;column:id;autoIncrement"`
	Name      string    `gorm:"column:name;size:256;not null"`
	Subject   string    `gorm:"column:subject;size:256;not null"`
	Body      string    `gorm:"column:body;not null"`
	HTMLBody  string    `gorm:"column:html_body;not null"`
	IsDeleted bool      `gorm:"column:is_deleted"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (MailTemplateTable) TableName() string {
	return "mail_template"
}

type MailLogTable struct {
	ID             string    `gorm:"primaryKey;type:uuid;column:id"`
	SentAt         time.Time `gorm:"column:sent_at"`
	RecepientEmail string    `gorm:"column:recepient_email;size:256;not null"`
	RecepientID    string    `gorm:"column:recepient_id;type:uuid"`
	Data           string    `gorm:"column:data;type:jsonb;not null"`
	TemplateID     int       `gorm:"column:template_id;not null"`
	Status         string    `gorm:"column:status;not null"`
}

func (MailLogTable) TableName() string {
	return "mail_log"
}
