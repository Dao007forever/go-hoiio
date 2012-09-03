go-hoiio
========

Hoiio SDK in Go

## Requirements
Go1

## Compiliation
Install `go-gb` and compile the package.

```terminal
$ gb
```

The package will be in the `_obj` folder.

## Getting started

Create a HoiioRestClient.

### API Credential

The HoiioRestClient needs your Hoiio credentials. You must pass these directly to the constructor.

```go
import . "go-hoiio/impl"

appid := 'YOUR_APPID_HERE'
token := 'YOUR_TOKEN_HERE'
client := NewClient(appid, token)
```

### Make a call

```go
import (
       "net/url"
       . "go-hoiio/impl"
)

appid := "YOUR_APPID_HERE"
token := "YOUR_TOKEN_HERE"
client := NewClient(appid, token)
voiceService := NewVoice(client)
data := url.Values{
     "dest" : {"+6512345678"},
     "dest2": {"+6554326547"},
     "caller_id": {"Trong Dao"},
}

resp = voiceService.call(data)
```

### Send a sms

```go
import (
       "net/url"
       . "go-hoiio/impl"
)

appid := "YOUR_APPID_HERE"
token := "YOUR_TOKEN_HERE"
client := NewClient(appid, token)
smsService := NewSms(client)
data := url.Values{
     "dest" : {"+6512345678"},
     "msg": {"Hello, how are you?"},
}

resp = smsService.send(data)
```
