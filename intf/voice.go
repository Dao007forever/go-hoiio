package intf

import "net/url"

type Voice interface {
	Call(url.Values) string
	Conference(url.Values) string
	HangUp(url.Values) string
	GetHistory(url.Values) string
	GetRate(url.Values) string
	GetStatus(url.Values) string
}