package app

type Response struct {
	Code    int         `json:"code" example:"200"`
	Data    interface{} `json:"data" `
	Message string      `json:"message"`
}

type Page struct {
	Data      interface{} `json:"data"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

type PageResponse struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data Page `json:"data"`
	// 消息
	Msg string `json:"msg"`
}

func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

func (res *PageResponse) ReturnOK() *PageResponse {
	res.Code = 200
	return res
}
