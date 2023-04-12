package CUtil

import (
	"github.com/go-resty/resty/v2"
)

// HttpGet  get 请求
func HttpGet(url string, params map[string]string) (*resty.Response, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(params).
		SetHeader("Accept", "application/json").
		//SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		EnableTrace().
		Get(url)
	return resp, err
}

// HttpPost post 请求 url body:键值对
func HttpPost(url string, body map[string]interface{}) (*resty.Response, error) {
	client := resty.New()
	resp, err := client.R().
		SetBody(body).
		//SetBody(map[string]interface{}{"username": "testuser", "password": "testpass"}).
		//SetResult(AuthSuccess).    // or SetResult(AuthSuccess{}).
		//SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
		//SetError(&AuthError{}).       // or SetError(AuthError{}).
		Post(url)

	return resp, err

}
