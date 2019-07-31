package viewmodels

// ViewGift font form end database
type (
	// Accounts ...
	Accounts struct {
		Email       string `json:"email"`
		IdentitiyID string `json:"identity_id"`
		UserName    string `json:"username"`
	}

	// Content ...
	Content struct {
		Title       string `json:"title"`
		Content     string `json:"content"`
		ContentHash string `json:"content_hash"`
		Price       int64  `json:"price"`
		Weight      int64  `json:"weight"`
	}
	// Auction ...
	Auction struct {
		ContentHash string `json:"content_hash"`
		Address     string `json:"address"`
		TokenID     int64  `json:"token_id"`
		Percent     int64  `json:"percent"`
		Price       int64  `json:"price"`
	}
	// BidPerson ...
	BidPerson struct {
		TokenID int64  `json:"token_id"`
		Price   int64  `json:"price"`
		Address string `json:"address"`
	}
)
