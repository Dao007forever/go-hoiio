package impl

import (
	"net/url"
	"resource"
	. "intf"
)

type NumberImpl struct {
	client Client
}

func NewNumber(client Client) Number {
	return &NumberImpl{client}
}

func (num *NumberImpl) GetCountries() string {
	return num.client.Execute(resource.NumberService, resource.Number.GetCountries, nil)
}

func (num *NumberImpl) GetChoices(params url.Values) string {
	return num.client.Execute(resource.NumberService, resource.Number.GetChoices, params)
}

func (num *NumberImpl) GetRate(params url.Values) string {
	return num.client.Execute(resource.NumberService, resource.Number.GetRate, params)
}

func (num *NumberImpl) Subscribe(params url.Values) string {
	return num.client.Execute(resource.NumberService, resource.Number.Subscribe, params)
}

func (num *NumberImpl) UpdateForwarding(params url.Values) string {
	return num.client.Execute(resource.NumberService, resource.Number.UpdateForwarding, params)
}

func (num *NumberImpl) ActiveNumbers() string {
	return num.client.Execute(resource.NumberService, resource.Number.GetActive, nil)
}