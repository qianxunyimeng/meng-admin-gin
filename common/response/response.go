// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/26 20:46:00
// @Desc

package response

import "meng-admin-gin/common/model"

type Response struct {
	// 数据集
	RequestId string `json:"requestId,omitempty"`
	Code      int32  `json:"code"`
	Msg       string `json:"msg,omitempty"`
	error     error  `json:"error,omitempty"`
	Status    string `json:"status,omitempty"`
}

type response struct {
	Response
	Data interface{} `json:"data"`
}

type Page struct {
	Count     int `json:"count"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type page struct {
	Page
	List interface{} `json:"list"`
}

func (e *response) SetData(data interface{}) {
	e.Data = data
}

func (e response) Clone() model.Responses {
	return &e
}

func (e *response) SetTraceID(id string) {
	e.RequestId = id
}

func (e *response) SetMsg(s string) {
	e.Msg = s
}

func (e *response) SetCode(code int32) {
	e.Code = code
}

func (e *response) SetError(err error) {
	e.error = err
}

func (e *response) SetSuccess(success bool) {
	if !success {
		e.Status = "error"
	}
}
