package intf

import "net/url"

type Number interface {
	GetCountries() string
	GetChoices(url.Values) string
	GetRate(url.Values) string
	Subscribe(url.Values) string
	UpdateForwarding(url.Values) string
	ActiveNumbers() string
}
