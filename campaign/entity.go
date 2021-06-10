package campaign

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	BackerCount      int
	Perks            string
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
