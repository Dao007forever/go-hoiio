package resource

// Account
var AccountService string = "user"

type AccStruct struct {
	GetBalance string
	GetInfo string
}

var Account AccStruct = AccStruct{
	"get_balance",
	"get_info",
}

// Number
var NumberService string = "number"

type NumStruct struct {
	GetCountries string
	GetChoices string
	GetRate string
	Subscribe string
	UpdateForwarding string
	GetActive string
}

var Number NumStruct = NumStruct{
	"get_countries",
	"get_choices",
	"get_rates",
	"subscribe",
	"update_forwarding",
	"get_active",
}

// SMS
var SmsService string = "sms"

type SmsStruct struct {
	Send string
	BulkSend string
	GetHistory string
	GetRate string
	GetStatus string
}

var Sms SmsStruct = SmsStruct{
	"send",
	"bulk_send",
	"get_history",
	"get_rate",
	"query_status",
}

// Voice
var VoiceService string = "voice"

type VoiceStruct struct {
	Call string
	Conference string
	HangUp string
	GetHistory string
	GetRate string
	GetStatus string
}

var Voice VoiceStruct = VoiceStruct{
	"call",
	"conference",
	"hangup",
	"get_history",
	"get_rate",
	"query_status",
}

// Fax
var FaxService string = "fax"

type FaxStruct struct {
	Send string
	GetHistory string
	GetRate string
	GetStatus string
}

var Fax FaxStruct = FaxStruct{
	"send",
	"get_history",
	"get_rate",
	"query_status",
}

// IVR
var IvrService string = "ivr"

type IvrStruct struct {
	Dial string
	Play string
	Gather string
	Record string
	Transfer string
	HangUp string
}

var Ivr IvrStruct = IvrStruct{
	"start/dial",
	"middle/play",
	"middle/gather",
	"middle/record",
	"end/transfer",
	"end/hangup",
}