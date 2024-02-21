package constants

import (
	"time"
)

type LoginAPIData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StatusType string
type TransactionType string
type InvestmentType string
type FrequencyType string

const (
	STATUS_ACTIVE   StatusType = "Active"
	STATUS_BLOCKED  StatusType = "Blocked"
	STATUS_INACTIVE StatusType = "Inactive"
	STATUS_DELETED  StatusType = "Deleted"
	STATUS_PAUSED   StatusType = "Paused"
	STATUS_PENDING  StatusType = "Pending"
)

const (
	TRANSACTION_INCOME     TransactionType = "Income"
	TRANSACTION_EXPENSE    TransactionType = "Expense"
	TRANSACTION_INVESTMENT TransactionType = "Expense"
)

const (
	INVESTMENT_RD     InvestmentType = "RD"
	INVESTMENT_FD     InvestmentType = "FD"
	INVESTMENT_STOCKS InvestmentType = "Stocks"
	INVESTMENT_MF     InvestmentType = "MF"
	INVESTMENT_GOLD   InvestmentType = "Gold"
	INVESTMENT_OTHERS InvestmentType = "Others"
)

const (
	FREQUENCY_DAILY         FrequencyType = "Daily"
	FREQUENCY_WEEKLY        FrequencyType = "Weekly"
	FREQUENCY_MONTHLY       FrequencyType = "Monthly"
	FREQUENCY_YEARLY        FrequencyType = "Annually"
	FREQUENCY_QUARTERLY     FrequencyType = "Quarterly"
	FREQUENCY_SEMI_ANNUALLY FrequencyType = "Semi-Annually"
)

type ModelUser struct {
	UserID    int        `json:"user_id" gorm:"primary_key;auto_increment"`
	Email     string     `json:"email" gorm:"type:varchar"`
	Password  string     `json:"password" gorm:"type:varchar"`
	CreatedOn time.Time  `json:"created_on" gorm:"type:timestamp; default:current_timestamp"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:timestamp; default:current_timestamp"`
	Status    StatusType `gorm:"type:varchar" json:"status"`
}
type ModelCards struct {
	CardID   int         `gorm:"primary_key;auto_increment" json:"card_id"`
	CardName string      `gorm:"type:varchar" json:"card_name"`
	UserID   []ModelUser `json:"user_id" gorm:"foreignkey:UserID"`
	Limit    float64     `json:"limit"`
	LastDay  time.Time   `gorm:"type:timestamp" json:"last_day"`
	Expiry   time.Time   `gorm:"type:timestamp" json:"expiry"`
	Status   StatusType  `gorm:"type:varchar" json:"status"`
}
type ModelUserCategory struct {
	UCID         int             `gorm:"primary_key;auto_increment" json:"uc_id"`
	UserID       []ModelUser     `gorm:"foreignkey:UserID" json:"user"`
	CategoryType TransactionType `json:"category_type" gorm:"type:varchar"`
	CategoryName string          `gorm:"type:varchar" json:"category_name"`
	CreatedOn    time.Time       `gorm:"type:timestamp; default:current_timestamp" json:"created_on"`
}

type ModelBill struct {
	BillID         int           `gorm:"primary_key;auto_increment" json:"bill_id"`
	Frequency      FrequencyType `json:"frequency" gorm:"type:varchar"`
	OtherFrequency *time.Time    `gorm:"type:timestamp" json:"other_frequency"`
	ReminderAlert  bool          `json:"reminder_alert" gorm:"default:false"`
	AutoPay        bool          `json:"auto_pay" gorm:"default:false"`
	Description    *string       `gorm:"type:varchar" json:"description"`
	Status         StatusType    `gorm:"type:varchar" json:"status"`
}

type ModelInvestment struct {
	InvestmentID        int            `gorm:"primary_key;auto_increment" json:"investment_id"`
	Type                InvestmentType `json:"type" gorm:"type:varchar"`
	OtherType           *string        `gorm:"type:varchar" json:"other_type"`
	Description         *string        `gorm:"type:varchar" json:"description"`
	EndsOn              *time.Time     `gorm:"type:timestamp" json:"ends_on"`
	Returns             *float64       `gorm:"default:0" json:"returns"`
	Frequency           FrequencyType  `json:"frequency" gorm:"type:varchar"`
	OtherFrequency      *time.Time     `gorm:"type:timestamp" json:"other_frequency"`
	AutoPaid            bool           `json:"auto_paid" gorm:"default:false"`
	SkipNextInstallment bool           `json:"skip_next_installment" gorm:"default:false"`
	CreatedOn           time.Time      `gorm:"type:timestamp; default:current_timestamp" json:"created_on"`
	Status              StatusType     `gorm:"type:varchar" json:"status"`
}

type ModelTransaction struct {
	TransactionID int                 `gorm:"primary_key;auto_increment" json:"transaction_id"`
	Name          string              `gorm:"type:varchar" json:"name"`
	Type          TransactionType     `gorm:"type:varchar" json:"type"`
	CategoryID    []ModelUserCategory `gorm:"foreignkey:CategoryID" json:"category_id"`
	UserID        []ModelUser         `gorm:"foreignkey:UserID" json:"user_id"`
	Amount        float64             `json:"amount"`
	CreditCardID  *[]ModelCards       `json:"credit_card_id" gorm:"foreignkey: CardID"`
	BillID        *[]ModelBill        `json:"bill_id" gorm:"foreignkey: BillID"`
	InvestmentID  *[]ModelInvestment  `json:"investment_id" gorm:"foreignkey: InvestmentID"`
	TransactionOn time.Time           `gorm:"type:timestamp" json:"transaction_on"`
}
