package coinbasepro

import (
	"context"
	"encoding/json"
	"fmt"
)

type Profile struct {
	Active    bool   `json:"active"`
	CreatedAt Time   `json:"created_at"`
	ID        string `json:"id"`
	IsDefault bool   `json:"is_default"`
	Name      string `json:"name"`
	UserID    string `json:"user_id"`
}

type ProfileFilter struct {
	Active bool `json:"active"`
}

func (p ProfileFilter) Params() []string {
	var params []string
	if p.Active {
		params = append(params, "active")
	}
	return params
}

type ProfileTransferSpec struct {
	Amount   float64      `json:"amount,string"`
	Currency CurrencyName `json:"currency"`
	From     string       `json:"from"`
	To       string       `json:"to"`
}

func (p *ProfileTransferSpec) UnmarshalJSON(b []byte) error {
	type Alias ProfileTransferSpec
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	return json.Unmarshal(b, &aux)
}

type ProfileTransfer struct {
	Amount   float64      `json:"amount"`
	Currency CurrencyName `json:"currency"`
	From     string       `json:"from"`
	To       string       `json:"to"`
}

// ListProfiles retrieves a list of Profiles (portfolio equivalents). A given user can have a maximum of 10 profiles.
// The list is not paginated.
func (c *Client) ListProfiles(ctx context.Context, filter ProfileFilter) ([]Profile, error) {
	var profiles []Profile
	return profiles, c.Get(ctx, fmt.Sprintf("/profiles/%s", query(filter.Params())), &profiles)
}

// GetProfile retrieves the details of a single Profile.
func (c *Client) GetProfile(ctx context.Context, profileID string) (Profile, error) {
	var profile Profile
	return profile, c.Get(ctx, fmt.Sprintf("/profiles/%s", profileID), &profile)
}

// CreateProfileTransfer transfers funds between user Profiles.
func (c *Client) CreateProfileTransfer(ctx context.Context, transferSpec ProfileTransferSpec) (ProfileTransfer, error) {
	var transfer ProfileTransfer
	return transfer, c.Post(ctx, "/profiles/transfer", transferSpec, &transfer)
}
