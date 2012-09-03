package intf

import "net/url"

type Fax interface {
	Send(url.Values) string
	GetHistory(url.Values) string
	GetRate(url.Values) string
	GetStatus(url.Values) string
}