package entity

import "time"

type Item struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"size:100;not null" json:"name"`
	Category      string    `gorm:"not null" json:"category"`
	Brand         string    `gorm:"size:100;not null" json:"brand"`
	PurchasePrice int       `gorm:"not null" json:"purchase_price"`
	PurchaseDate  time.Time `gorm:"type:date;not null" json:"purchase_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// UpdateItemRequest: タスク1要件のフィールドのみ
type UpdateItemRequest struct {
	Name          *string `json:"name"`
	Brand         *string `json:"brand"`
	PurchasePrice *int    `json:"purchase_price"`
}

// バリデーションロジック
func (r *UpdateItemRequest) Validate() []string {
	var details []string
	if r.Name != nil && (len(*r.Name) == 0 || len(*r.Name) > 100) {
		details = append(details, "name must be between 1 and 100 characters")
	}
	if r.Brand != nil && (len(*r.Brand) == 0 || len(*r.Brand) > 100) {
		details = append(details, "brand must be between 1 and 100 characters")
	}
	if r.PurchasePrice != nil && *r.PurchasePrice < 0 {
		details = append(details, "purchase_price must be 0 or greater")
	}
	return details
}
