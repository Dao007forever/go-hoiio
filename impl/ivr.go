package impl

import (
	"net/url"
	"resource"
	. "intf"
)

type IvrImpl struct {
	client Client
}

func NewIvr(client Client) Ivr {
	return &IvrImpl{client}
}

func (ivr *IvrImpl) Dial(params url.Values) string {
	return ivr.client.Execute(resource.IvrService, resource.Ivr.Dial, params)
}

func (ivr *IvrImpl) Play(params url.Values) string {
	return ivr.client.Execute(resource.IvrService, resource.Ivr.Play, params)
}

func (ivr *IvrImpl) Gather(params url.Values) string {
	return ivr.client.Execute(resource.IvrService, resource.Ivr.Gather, params)
}

func (ivr *IvrImpl) Record(params url.Values) string {
	return ivr.client.Execute(resource.IvrService, resource.Ivr.Record, params)
}

func (ivr *IvrImpl) Transfer(params url.Values) string {
	return ivr.client.Execute(resource.IvrService, resource.Ivr.Transfer, params)
}

func (ivr *IvrImpl) HangUp(params url.Values) string {
	return ivr.client.Execute(resource.IvrService, resource.Ivr.HangUp, params)
}