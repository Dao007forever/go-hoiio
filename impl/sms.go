package impl

import (
	"net/url"
	"resource"
	. "intf"
)

type SmsImpl struct {
	client Client
}

func NewSms(client Client) Sms {
	return &SmsImpl{client}
}

func (sms *SmsImpl) Send(params url.Values) string {
	return sms.client.Execute(resource.SmsService, resource.Sms.Send, params)
}

func (sms *SmsImpl) BulkSend(params url.Values) string {
	return sms.client.Execute(resource.SmsService, resource.Sms.BulkSend, params)
}

func (sms *SmsImpl) GetHistory(params url.Values) string {
	return sms.client.Execute(resource.SmsService, resource.Sms.GetHistory, params)
}

func (sms *SmsImpl) GetRate(params url.Values) string {
	return sms.client.Execute(resource.SmsService, resource.Sms.GetRate, params)
}

func (sms *SmsImpl) GetStatus(params url.Values) string {
	return sms.client.Execute(resource.SmsService, resource.Sms.GetStatus, params)
}