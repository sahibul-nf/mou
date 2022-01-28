package transaction

import (
	"moyu/campaign"
	"moyu/user"
	"time"
)

type Transaction struct {
	ID         int    `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	CampaignID int    `gorm:"column:campaign_id;type:int"`
	UserID     int    `gorm:"column:user_id;type:int"`
	Amount     int    `gorm:"column:amount;type:int"`
	Status     string `gorm:"column:status;type:varchar;size:255"`
	Code       string `gorm:"column:code;type:varchar;size:255"`
	PaymentURL string `gorm:"column:payment_url;type:varchar;size:255"`
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp"`
}
