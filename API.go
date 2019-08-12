package gozaim

import (
	"encoding/json"
	"net/url"
)

func (c *Client) FetchMe() (Me, error) {
	body, err := c.get("home/user/verify", nil)
	if err != nil {
		return Me{}, err
	}

	var raw struct {
		Me        Me
		Requested string
	}
	json.Unmarshal(body, &raw)

	return raw.Me, nil
}

func (c *Client) FetchMoney(params url.Values) ([]Money, error) {
	body, err := c.get("home/money", params)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Money []Money
	}
	json.Unmarshal(body, &raw)

	return raw.Money, nil
}

func (c *Client) FetchCategories() ([]Category, error) {
	body, err := c.get("home/category", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Categories []Category
		Requested  int
	}
	json.Unmarshal(body, &raw)

	return raw.Categories, nil
}

func (c *Client) FetchDefaultCategories() ([]Category, error) {
	body, err := c.get("category", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Categories []Category
		Requested  int
	}
	json.Unmarshal(body, &raw)

	return raw.Categories, nil
}

func (c *Client) FetchGenres() ([]Genre, error) {
	body, err := c.get("home/genre", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Genres    []Genre
		Requested int
	}
	json.Unmarshal(body, &raw)

	return raw.Genres, nil
}

func (c *Client) FetchDefaultGenres() ([]Genre, error) {
	body, err := c.get("genre", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Genres    []Genre
		Requested int
	}
	json.Unmarshal(body, &raw)

	return raw.Genres, nil
}

func (c *Client) FetchAccounts() ([]Account, error) {
	body, err := c.get("home/account", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Accounts  []Account
		Requested int
	}
	json.Unmarshal(body, &raw)

	return raw.Accounts, nil
}

func (c *Client) FetchDefaultAccounts() ([]Account, error) {
	body, err := c.get("account", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Accounts  []Account
		Requested int
	}
	json.Unmarshal(body, &raw)

	return raw.Accounts, nil
}

func (c *Client) FetchCurrencies() ([]Currency, error) {
	body, err := c.get("currency", nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Currencies []Currency
		Requested  int
	}
	json.Unmarshal(body, &raw)

	return raw.Currencies, nil
}

func (c *Client) CreatePayment(params url.Values) (bool, error) {
	_, err := c.post("home/money/payment", params)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) CreateIncome(params url.Values) (bool, error) {
	_, err := c.post("home/money/income", params)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) UpdatePayment(id string, params url.Values) (bool, error) {
	_, err := c.put("home/money/payment/"+id, params)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) UpdateIncome(id string, params url.Values) (bool, error) {
	_, err := c.put("home/money/income/"+id, params)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) DeletePayment(id string) (bool, error) {
	_, err := c.put("home/money/payment/"+id, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) DeleteIncome(id string) (bool, error) {
	_, err := c.put("home/money/income/"+id, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
