package intf

import "net/url"

type Client interface {
	GetMethodURI(path ...string) string
	Request(uri string, data url.Values) string
	Execute(service string, use string, data url.Values) string
}
