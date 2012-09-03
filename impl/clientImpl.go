package impl

import (
	"fmt"
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
)

type ClientImpl struct {
	appid string
	token string
	baseUri string
}

func NewClient(appid, token string) ClientImpl {
	return ClientImpl{appid, token, "https://secure.hoiio.com/open"}
}

func (client *ClientImpl) GetMethodURI(path ...string) (uri string) {
	path_slice := append([]string{client.baseUri}, path...)
	return strings.Join(path_slice, "/")
}

func (client *ClientImpl) Request(uri string, data url.Values) string {
	app := url.Values{
		"app_id" : {client.appid},
		"access_token" : {client.token},
	}
	params := "?" + app.Encode()
	if data != nil {
		params += "&" + data.Encode()
	}

	response, err := http.Get(uri + params)

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return err.Error()
	} else {
		defer response.Body.Close()
		contents , err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}
		return string(contents)
	}
	return "Dummy string"
}

func (client *ClientImpl) Execute(service string, use string, params url.Values) string {
	methodUri := client.GetMethodURI(service, use)
	return client.Request(methodUri, params)
}