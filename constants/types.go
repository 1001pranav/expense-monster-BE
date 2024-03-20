package constants

const (
	STATUS_ACTIVE   StatusType = "Active"
	STATUS_BLOCKED  StatusType = "Blocked"
	STATUS_INACTIVE StatusType = "Inactive"
	STATUS_DELETED  StatusType = "Deleted"
	STATUS_PAUSED   StatusType = "Paused"
	STATUS_PENDING  StatusType = "Pending"
)

type StatusType string

type TransactionType string

const (
	TRANSACTION_INCOME     TransactionType = "Income"
	TRANSACTION_EXPENSE    TransactionType = "Expense"
	TRANSACTION_INVESTMENT TransactionType = "Investment"
)

type InvestmentType string

const (
	INVESTMENT_RD     InvestmentType = "RD"
	INVESTMENT_FD     InvestmentType = "FD"
	INVESTMENT_STOCKS InvestmentType = "Stocks"
	INVESTMENT_MF     InvestmentType = "MF"
	INVESTMENT_GOLD   InvestmentType = "Gold"
	INVESTMENT_OTHERS InvestmentType = "Others"
)

type FrequencyType string

const (
	FREQUENCY_DAILY         FrequencyType = "Daily"
	FREQUENCY_WEEKLY        FrequencyType = "Weekly"
	FREQUENCY_MONTHLY       FrequencyType = "Monthly"
	FREQUENCY_YEARLY        FrequencyType = "Annually"
	FREQUENCY_QUARTERLY     FrequencyType = "Quarterly"
	FREQUENCY_SEMI_ANNUALLY FrequencyType = "Semi-Annually"
)

type RestPasswordType int

const (
	FORGOT_RESET_PASSWORD RestPasswordType = 1
	CHANGE_RESET_PASSWORD RestPasswordType = 2
)
