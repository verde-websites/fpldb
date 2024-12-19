package models

import "time"

type FplManagerSession struct {
	ID                 int64      `bun:"id,pk,autoincrement"`
	Email              string     `bun:"email"`
	Password           string     `bun:"password"`
	ManagerID          string     `bun:"manager_id"`
	Cookies            string     `bun:"cookies"`
	InUse              *bool      `bun:"in_use"`
	LastUsed           time.Time  `bun:"last_used"`
	CookiesLastUpdated *time.Time `bun:"cookies_last_updated"`
	Active             bool       `bun:"active"`
}
