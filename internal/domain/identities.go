package domain

type Identitie struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `json:"name"`
	Alias         string `json:"alias"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
	Email_Address string `gorm:"uniqueIndex" json:"email"`
}

func (Identitie) TableName() string {
	return "auth.users"
}
