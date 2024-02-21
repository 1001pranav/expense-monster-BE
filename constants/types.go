package constants

type LoginAPIData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StatusType string

const (
	ACTIVE   StatusType = "Active"
	BLOCKED  StatusType = "Blocked"
	INACTIVE StatusType = "Inactive"
	DELETED  StatusType = "Deleted"
)

type ModelUsers struct {
	UserID    uint   `json:"user_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedOn uint64 `json:"created_on"`
	UpdatedAt uint64 `json:"updated_at"`
	Status    string `json:"status"`
}

type ModelCards struct {
	CardId   uint   `json:"card_id"`
	CardName string `json:"card_name"`
	UserID   uint   `json:"user_id"`
	Limit    uint   `json:"limit"`
	LastDay  uint   `json:"last_day"`
	Expiry   uint   `json:"expiry"`
	Status   string `json:"status"`
}

type ModelUserCategory struct {
	UCID         uint   `json:"uc_id"`
	UserID       uint   `json:"user_id"`
	CategoryType string `json:"category_type"`
	CategoryName string `json:"category_name"`
	CreatedOn    int    `json:"created_on"`
}

type ModelTransaction struct {
	TransactionID uint    `json:"transaction_id"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	CategoryID    uint    `json:"category_id"`
	UserID        uint    `json:"user_id"`
	Amount        float64 `json:"amount"`
	CreditCardID  uint    `json:"credit_card_card_id"`
	BillID        uint    `json:"bill_id"`
	InvestmentID  uint    `json:"investment_id"`
	TransactionOn int     `json:"transaction_on"`
}
