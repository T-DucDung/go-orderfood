package responses

type Response struct {
	Data interface{} `xml:"data" json:"data"`
	Code int         `xml:"code" json:"code"`
}

var FailResponse = Response{Data: "Bad request", Code: 401}
