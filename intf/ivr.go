package intf

import "net/url"

type Ivr interface {
	Dial(url.Values) string
	Play(url.Values) string
	Gather(url.Values) string
	Record(url.Values) string
	Transfer(url.Values) string
	HangUp(url.Values) string
}