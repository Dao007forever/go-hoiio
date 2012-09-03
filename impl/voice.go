package impl

import (
	"net/url"
	"resource"
	. "intf"
)

type VoiceImpl struct {
	client Client
}

func NewVoice(client Client) Voice {
	return &VoiceImpl{client}
}

func (voice *VoiceImpl) Call(params url.Values) string {
	return voice.client.Execute(resource.VoiceService, resource.Voice.Call, params)
}

func (voice *VoiceImpl) Conference(params url.Values) string{
	return voice.client.Execute(resource.VoiceService, resource.Voice.Conference, params)
}

func (voice *VoiceImpl) HangUp(params url.Values) string{
	return voice.client.Execute(resource.VoiceService, resource.Voice.HangUp, params)
}

func (voice *VoiceImpl) GetHistory(params url.Values) string {
	return voice.client.Execute(resource.VoiceService, resource.Voice.GetHistory, params)
}

func (voice *VoiceImpl) GetRate(params url.Values) string {
	return voice.client.Execute(resource.VoiceService, resource.Voice.GetRate, params)
}

func (voice *VoiceImpl) GetStatus(params url.Values) string {
	return voice.client.Execute(resource.VoiceService, resource.Voice.GetStatus, params)
}
