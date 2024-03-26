package constants

import "time"

type ModelUsers struct {
	UserID         uint                `json:"user_id" gorm:"primary_key;auto_increment"`
	Email          string              `json:"email" gorm:"type:varchar"`
	Password       string              `json:"password" gorm:"type:varchar"`
	AccessToken    string              `json:"access_token" gorm:"type:varchar"`
	OTP            uint                `json:"otp" gorm:"type:int"`
	OTPGeneratedOn time.Time           `json:"otp_on" gorm:"type:timestamp;null"`
	CreatedOn      time.Time           `json:"created_on" gorm:"type:timestamp; default:current_timestamp"`
	UpdatedAt      time.Time           `json:"updated_at" gorm:"type:timestamp; default:current_timestamp"`
	Status         StatusType          `gorm:"type:varchar" json:"status"`
	Cards          []ModelCards        `gorm:"foreignKey:UserID"`
	Transaction    []ModelTransaction  `gorm:"foreignKey:UserID"` // Add foreign key
	UserCategory   []ModelUserCategory `gorm:"foreignKey:UserID"`
}

type ModelCards struct {
	CardID      uint               `gorm:"primary_key;auto_increment" json:"card_id"`
	CardName    string             `gorm:"type:varchar" json:"card_name"`
	UserID      uint               `json:"user_id"`
	Limit       float64            `json:"limit"`
	LastDay     time.Time          `gorm:"type:timestamp" json:"last_day"`
	Expiry      time.Time          `gorm:"type:timestamp" json:"expiry"`
	Status      StatusType         `gorm:"type:varchar" json:"status"`
	Transaction []ModelTransaction `gorm:"foreignKey:CardID"` // Add foreign key
}

type ModelUserCategory struct {
	UCID         uint `gorm:"primary_key;auto_increment" json:"uc_id"`
	UserID       uint
	CategoryType TransactionType    `json:"category_type" gorm:"type:varchar"`
	CategoryName string             `gorm:"type:varchar" json:"category_name"`
	CreatedOn    time.Time          `gorm:"type:timestamp; default:current_timestamp" json:"created_on"`
	Transaction  []ModelTransaction `gorm:"foreignKey:UserID"` // Add foreign key
}

type ModelBill struct {
	BillID           uint               `gorm:"primary_key;auto_increment" json:"bill_id"`
	Frequency        FrequencyType      `json:"frequency" gorm:"type:varchar"`
	OtherFrequency   *time.Time         `gorm:"type:timestamp" json:"other_frequency"`
	ReminderAlert    bool               `json:"reminder_alert" gorm:"default:false"`
	AutoPay          bool               `json:"auto_pay" gorm:"default:false"`
	Description      *string            `gorm:"type:varchar" json:"description"`
	Status           StatusType         `gorm:"type:varchar" json:"status"`
	ModelTransaction []ModelTransaction `gorm:"foreignKey:BillID"`
}
type ModelInvestment struct {
	InvestmentID        uint               `gorm:"primary_key;auto_increment" json:"investment_id"`
	Type                InvestmentType     `json:"type" gorm:"type:varchar"`
	OtherType           *string            `gorm:"type:varchar" json:"other_type"`
	Description         *string            `gorm:"type:varchar" json:"description"`
	EndsOn              *time.Time         `gorm:"type:timestamp" json:"ends_on"`
	Returns             *float64           `gorm:"default:0" json:"returns"`
	Frequency           FrequencyType      `json:"frequency" gorm:"type:varchar"`
	OtherFrequency      *time.Time         `gorm:"type:timestamp" json:"other_frequency"`
	AutoPaid            bool               `json:"auto_paid" gorm:"default:false"`
	SkipNextInstallment bool               `json:"skip_next_installment" gorm:"default:false"`
	CreatedOn           time.Time          `gorm:"type:timestamp; default:current_timestamp" json:"created_on"`
	Status              StatusType         `gorm:"type:varchar" json:"status"`
	Transaction         []ModelTransaction `gorm:"foreignKey:InvestmentID"` // Add foreign key
}

type ModelTransaction struct {
	TransactionID uint            `gorm:"primary_key;auto_increment" json:"transaction_id"`
	Name          string          `gorm:"type:varchar" json:"name"`
	Type          TransactionType `gorm:"type:varchar" json:"type"`
	CategoryID    uint
	UserID        uint
	Amount        float64 `json:"amount"`
	CardID        *uint
	BillID        *uint
	InvestmentID  *uint
	TransactionOn time.Time
}
