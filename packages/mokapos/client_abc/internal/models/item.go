package models

// ItemImage represents an image associated with an item.
type ItemImage struct {
	URL    string `json:"url,omitempty"`
	ID     int    `json:"id,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

// Item represents a single product, menu item, or other entity returned by the API.
type Item struct {
	ID               int           `json:"id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Image            ItemImage     `json:"image"`
	BusinessID       int           `json:"business_id"`
	CategoryID       int           `json:"category_id"`
	CreatedAt        string        `json:"created_at"`
	UpdatedAt        string        `json:"updated_at"`
	OutletID         int           `json:"outlet_id"`
	GUID             string        `json:"guid"`
	SynchronizedAt   string        `json:"synchronized_at"`
	ItemVariants     []ItemVariant `json:"item_variants"`
	Category         Category      `json:"category"`
	IsEcommerce      bool          `json:"is_ecommerce"`
	BrandID          int           `json:"brand_id"`
	Condition        string        `json:"condition"`
	Weight           float64       `json:"weight"`
	Height           float64       `json:"height"`
	Width            float64       `json:"width"`
	Length           float64       `json:"length"`
	IsDeleted        bool          `json:"is_deleted"`
	IsRecipe         bool          `json:"is_recipe"`
	IsSalesTypePrice bool          `json:"is_sales_type_price"`
	Alert            bool          `json:"alert"`
	BackgroundColor  string        `json:"background_color"`
	UniqID           *string       `json:"uniq_id,omitempty"`
	ActiveModifiers  []interface{} `json:"active_modifiers"` // Placeholder - type depends on additional info
}

// ItemVariant represents different variations of an item (e.g., sizes, colors).
type ItemVariant struct {
	ID             int     `json:"id"`
	SKU            string  `json:"sku"`
	Price          float64 `json:"price"`
	ItemID         int     `json:"item_id"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	SynchronizedAt string  `json:"synchronized_at"`
	OutletID       int     `json:"outlet_id"`
	Name           string  `json:"name,omitempty"`
	Alert          bool    `json:"alert"`
	IsSaved        bool    `json:"is_saved"`
	AddInventory   int     `json:"add_inventory"`
	Guid           *string `json:"guid,omitempty"`
	Cogs           *int    `json:"cogs,omitempty"`
	IsDeleted      bool    `json:"is_deleted"`
	IsRecipe       bool    `json:"is_recipe"`
	UniqID         *string `json:"uniq_id,omitempty"`
	Position       int     `json:"position"`
	TrackStock     bool    `json:"track_stock"`
	StockAlert     int     `json:"stock_alert"`
	TrackCogs      bool    `json:"track_cogs"`
	InStock        int     `json:"in_stock"`
	LastModified   *string `json:"last_modified,omitempty"`
}
