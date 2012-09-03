package impl

import (
	"resource"
	. "intf"
)

type AccountImpl struct {
	client Client
}

func NewAccount(client Client) Account {
	return &AccountImpl{client}
}

func (acc *AccountImpl) GetBalance() string {
	return acc.client.Execute(resource.AccountService, resource.Account.GetBalance, nil)
}

func (acc *AccountImpl) GetInfo() string {
	return acc.client.Execute(resource.AccountService, resource.Account.GetInfo, nil)
}