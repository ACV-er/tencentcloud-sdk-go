package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	//"log"
	"net/http"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type BaseResponse struct {
}

type ErrorResponse struct {
	Response struct {
		Error struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
		RequestId string `json:"RequestId"`
	} `json:"Response"`
}

type DeprecatedAPIErrorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

type CMQAPIErrorResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if resp.Response.Error.Code != "" {
		return errors.NewTencentCloudSDKError(resp.Response.Error.Code, resp.Response.Error.Message, resp.Response.RequestId)
	}

	cmqResp := &CMQAPIErrorResponse{}
	err = json.Unmarshal(body, cmqResp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if cmqResp.RequestId != "" {
		return errors.NewTencentCloudSDKError(strconv.Itoa(cmqResp.Code), cmqResp.Message, cmqResp.RequestId)
	}

	deprecated := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, deprecated)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if deprecated.Code != 0 {
		return errors.NewTencentCloudSDKError(deprecated.CodeDesc, deprecated.Message, "")
	}
	return nil
}

func ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if resp.Response.Error.Code != "" {
		return errors.NewTencentCloudSDKError(resp.Response.Error.Code, resp.Response.Error.Message, resp.Response.RequestId)
	}

	cmqResp := &CMQAPIErrorResponse{}
	err = json.Unmarshal(body, cmqResp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if cmqResp.RequestId != "" {
		return errors.NewTencentCloudSDKError(strconv.Itoa(cmqResp.Code), cmqResp.Message, cmqResp.RequestId)
	}

	deprecated := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, deprecated)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if deprecated.Code != 0 {
		return errors.NewTencentCloudSDKError(deprecated.CodeDesc, deprecated.Message, "")
	}
	return nil
}

func ParseFromHttpResponse(hr *http.Response, response Response) (err error) {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body because %s", err)
		return errors.NewTencentCloudSDKError("ClientError.IOError", msg, "")
	}
	if hr.StatusCode != 200 {
		msg := fmt.Sprintf("Request fail with http status code: %s, with body: %s", hr.Status, body)
		return errors.NewTencentCloudSDKError("ClientError.HttpStatusCodeError", msg, "")
	}
	//log.Printf("[DEBUG] Response Body=%s", body)
	err = response.ParseErrorFromHTTPResponse(body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	return
}
