package campaign

import (
	"moyu/user"
	"time"
)

type Campaign struct {
	ID               int       `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	UserID           int       `gorm:"column:user_id;type:int"`
	Name             string    `gorm:"column:name;type:varchar;size:255"`
	ShortDescription string    `gorm:"column:short_description;type:varchar;size:255"`
	Description      string    `gorm:"column:description;type:varchar;size:255"`
	GoalAmount       int       `gorm:"column:goal_amount;type:int"`
	CurrentAmount    int       `gorm:"column:current_amount;type:int"`
	BackerCount      int       `gorm:"column:backer_count;type:int"`
	Perks            string    `gorm:"column:perks;type:varchar;size:255"`
	Slug             string    `gorm:"column:slug;type:varchar;size:255"`
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt        time.Time `gorm:"column:updated_at;type:timestamp"`
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	ID         int       `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	CampaignID int       `gorm:"column:campaign_id;type:int"`
	FileName   string    `gorm:"column:file_name;type:varchar;size:255"`
	IsPrimary  int       `gorm:"column:is_primary;type:int"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp"`
}
