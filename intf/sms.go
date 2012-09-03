package intf

import "net/url"

type Sms interface {
	Send(url.Values) string
	BulkSend(url.Values) string
	GetHistory(url.Values) string
	GetRate(url.Values) string
	GetStatus(url.Values) string
}