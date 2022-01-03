package coinbasepro

import (
	"context"
	"encoding/json"
	"fmt"
)

// Account holds funds for trading on coinbasepro.
// Coinbasepro Accounts are separate from Coinbase accounts. You Deposit funds to begin trading.
type Account struct {
	// Funds available for withdrawal or trade
	Available float64 `json:"available,string"`
	Balance   float64 `json:"balance,string"`
	// Currency is the native currency of the account
	Currency CurrencyName `json:"currency"`
	Hold     float64      `json:"hold,string"`
	// ID of the account
	ID string `json:"id"`
	// ProfileID is the id of the profile to which the account belongs
	ProfileID      string `json:"profile_id"`
	TradingEnabled bool   `json:"trading_enabled"`
}

// ListAccounts retrieves the list of trading accounts belonging to the Profile of the API key. The list is not paginated.
func (c *Client) ListAccounts(ctx context.Context) ([]Account, error) {
	var accounts []Account
	return accounts, c.Get(ctx, fmt.Sprintf("/%s/", coinbaseproAccounts), &accounts)
}

// GetAccount retrieves the detailed representation of a trading Account. The requested Account must belong to the current Profile.
func (c *Client) GetAccount(ctx context.Context, accountID string) (Account, error) {
	var account Account
	return account, c.Get(ctx, fmt.Sprintf("/%s/%s", coinbaseproAccounts, accountID), &account)
}

// Ledger holds the detailed activity of the profile associated with the current API key.
// Ledger is paginated and sorted newest first.
type Ledger struct {
	Entries []*LedgerEntry `json:"entries"`
	Page    *Pagination    `json:"page"`
}

// LedgerEntry represents an instance of account activity.
// A LedgerEntry will either increase or decrease the Account balance.
type LedgerEntry struct {
	// Amount of the transaction
	Amount float64 `json:"amount,string"`
	// Balance after transaction applied
	Balance float64 `json:"balance,string"`
	// CreatedAt is the timestamp of the transaction time
	CreatedAt Time `json:"created_at"`
	// Details will contain additional information if an entry is the result of a trade ('match', 'fee')
	Details LedgerDetails `json:"details"`
	// ID of the transaction
	ID string `json:"id"`
	// Type of transaction ('conversion', 'fee', 'match', 'rebate')
	Type LedgerEntryType `json:"type"`
}

// LedgerEntryType describes the reason for the account balance change.
type LedgerEntryType string

const (
	// LedgerEntryTypeConversion funds converted between fiat currency and a stablecoin
	LedgerEntryTypeConversion LedgerEntryType = "conversion"
	// LedgerEntryTypeFee funds moved to/from Coinbase to Coinbase Pro
	LedgerEntryTypeFee LedgerEntryType = "fee"
	// LedgerEntryTypeMatch funds moved as a result of a trade
	LedgerEntryTypeMatch LedgerEntryType = "match"
	// LedgerEntryTypeRebate fee rebate as per coinbasepro fee schedule (see https://pro.coinbase.com/fees)
	LedgerEntryTypeRebate LedgerEntryType = "rebate"
)

// LedgerDetails contains additional details for LedgerEntryTypeFee and LedgerEntryTypeMatch trades.
type LedgerDetails struct {
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
	TradeID   string `json:"trade_id"`
}

// UnmarshalJSON allows the raw slice of entries to be mapped to a named field on the struct.
// Pagination details are added post-unmarshal.
func (l *Ledger) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &l.Entries)
}

// GetLedger retrieves a paginated list of Account activity for the current Profile.
func (c *Client) GetLedger(ctx context.Context, accountID string, pagination PaginationParams) (Ledger, error) {
	if err := pagination.Validate(); err != nil {
		return Ledger{}, err
	}
	var ledger Ledger
	query := query(pagination.Params())
	return ledger, c.Get(ctx, fmt.Sprintf("/%s/%s/%s/%s", coinbaseproAccounts, accountID, coinbaseproLedger, query), &ledger)
}

// Holds are placed on an account for any active orders or pending withdraw requests.
// For `limit buy` orders, Price x Size x (1 + fee-percent) USD will be held.
// For sell orders, the number of base currency to be sold is held. Actual Fees are assessed at time of trade.
// If a partially filled or unfilled Order is canceled, any remaining funds will be released from hold.
// For a MarketOrder `buy` with Order.Funds, the Order.Funds amount will be put on hold. If only Order.Size is specified,
// the total Account.Balance (in the quote account) will be put on hold for the duration of the MarketOrder
// (usually a trivially short time).
// For a 'sell' Order, the Order.Size in base currency will be put on hold.
// If Order.Size is not specified (and only Order.Funds is specified), the entire base currency balance will be on
// hold for the duration of the MarketOrder.
type Holds struct {
	Holds []*Hold     `json:"holds"`
	Page  *Pagination `json:"page,omitempty"`
}

// A Hold will make the Amount of funds unavailable for trade or withdrawal.
type Hold struct {
	// Account identifies the Account to which the Hold applies
	AccountID string `json:"account_id"`
	// Amount of hold
	Amount float64 `json:"amount,string"`
	// Time hold was created
	CreatedAt Time `json:"created_at"`
	// Ref contains the id of the order or transfer which created the hold.
	Ref string `json:"ref"`
	// Type of indicates whether the Hold is the result of open orders or withdrawals.
	Type HoldType `json:"type"`
	// Time order was filled
	UpdatedAt Time `json:"updated_at"`
}

// HoldType indicates why the hold exists.
type HoldType string

const (
	// HoldTypeOpenOrders type holds are related to open orders.
	HoldTypeOpenOrders HoldType = "order"
	// HoldTypeWithdrawal type holds are related to a withdrawal.
	HoldTypeWithdrawal HoldType = "transfer"
)

// UnmarshalJSON allows the raw slice of entries to be mapped to a named field on the struct.
// Pagination details are added post-unmarshal.
func (h *Holds) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &h.Holds)
}

// GetHolds retrieves the list of Holds for the Account. The requested Account must belong to the current Profile.
func (c *Client) GetHolds(ctx context.Context, accountID string, pagination PaginationParams) (Holds, error) {
	if err := pagination.Validate(); err != nil {
		return Holds{}, err
	}
	var holds Holds
	query := query(pagination.Params())
	return holds, c.Get(ctx, fmt.Sprintf("/%s/%s/%s/%s", coinbaseproAccounts, accountID, coinbaseproHolds, query), &holds)
}
