package impl

import (
	"net/url"
	"resource"
	. "intf"
)

type FaxImpl struct {
	client Client
}

func NewFax(client Client) Fax {
	return &FaxImpl{client}
}

func (fax *FaxImpl) Send(params url.Values) string {
	return fax.client.Execute(resource.FaxService, resource.Fax.Send, params)
}

func (fax *FaxImpl) GetHistory(params url.Values) string {
	return fax.client.Execute(resource.FaxService, resource.Fax.GetHistory, params)
}

func (fax *FaxImpl) GetRate(params url.Values) string {
	return fax.client.Execute(resource.FaxService, resource.Fax.GetRate, params)
}

func (fax *FaxImpl) GetStatus(params url.Values) string {
	return fax.client.Execute(resource.FaxService, resource.Fax.GetStatus, params)
}
