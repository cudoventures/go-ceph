package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Account is Go representation of the JSON output of an account creation
type Account struct {
	ID            string    `json:"id" url:"id"`
	Tenant        string    `json:"tenant" url:"tenant"`
	Name          string    `json:"name" url:"name"`
	Email         string    `json:"email" url:"email"`
	Quota         QuotaSpec `json:"quota"`
	BucketQuota   QuotaSpec `json:"bucket_quota"`
	MaxUsers      *int      `json:"max_users" url:"max-users"`
	MaxRoles      *int      `json:"max_roles" url:"max-roles"`
	MaxGroups     *int      `json:"max_groups" url:"max-groups"`
	MaxBuckets    *int      `json:"max_buckets" url:"max-buckets"`
	MaxAccessKeys *int      `json:"max_access_keys" url:"max-access_keys"`
}

// GetAccount retrieves a given object store account
func (api *API) GetAccount(ctx context.Context, account Account) (Account, error) {
	if account.ID == "" {
		return Account{}, errMissingAccountID
	}

	body, err := api.call(ctx, http.MethodGet, "/account", valueToURLParams(account, []string{"id", "tenant", "name"}))
	if err != nil {
		return Account{}, err
	}

	a := Account{}
	err = json.Unmarshal(body, &a)
	if err != nil {
		return Account{}, fmt.Errorf("%s. %s. %w", unmarshalError, string(body), err)
	}

	return a, nil
}

// CreateAccount creates an account in the object store
func (api *API) CreateAccount(ctx context.Context, account Account) (Account, error) {
	body, err := api.call(ctx, http.MethodPost, "/account", valueToURLParams(account, []string{"id", "tenant", "name", "email", "quota", "bucket-quota", "max-users", "max-roles", "max-groups", "max-buckets", "max-access-keys"}))
	if err != nil {
		return Account{}, err
	}

	// Unmarshal response into Go type
	a := Account{}
	err = json.Unmarshal(body, &a)
	if err != nil {
		return Account{}, fmt.Errorf("%s. %s. %w", unmarshalError, string(body), err)
	}

	return a, nil
}

// RemoveAccount removes an account from the object store
func (api *API) RemoveAccount(ctx context.Context, account Account) error {
	if account.ID == "" {
		return errMissingAccountID
	}

	_, err := api.call(ctx, http.MethodDelete, "/account", valueToURLParams(account, []string{"id"}))
	if err != nil {
		return err
	}

	return nil
}

// ModifyAccount updates an account in the object store
func (api *API) ModifyAccount(ctx context.Context, account Account) (Account, error) {
	if account.ID == "" {
		return Account{}, errMissingAccountID
	}

	body, err := api.call(ctx, http.MethodPost, "/account", valueToURLParams(account, []string{"id", "tenant", "name", "email", "quota", "bucket-quota", "max-users", "max-roles", "max-groups", "max-buckets", "max-access-keys"}))
	if err != nil {
		return Account{}, err
	}

	a := Account{}
	err = json.Unmarshal(body, &a)
	if err != nil {
		return Account{}, fmt.Errorf("%s. %s. %w", unmarshalError, string(body), err)
	}

	return a, nil
}
