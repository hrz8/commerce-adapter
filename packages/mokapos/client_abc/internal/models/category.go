package models

// Category represents a category or group to which an item belongs.
type Category struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	BusinessID     int     `json:"business_id"`
	OutletID       int     `json:"outlet_id"`
	GUID           string  `json:"guid"`
	SynchronizedAt string  `json:"synchronized_at"`
	UpdatedAt      string  `json:"updated_at"`
	CreatedAt      string  `json:"created_at"`
	Description    *string `json:"description,omitempty"`
	IsDeleted      bool    `json:"is_deleted"`
	UniqID         *string `json:"uniq_id,omitempty"`
}
