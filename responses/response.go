package responses

type Response struct {
	Data  interface{} `xml:"data" json:"data"`
	Error *Err        `json:"error"`
}

var FailResponse = Response{Data: "Bad request", Error: &Err{Code: 401, Message: "Bad request"}}

var UnAuthResponse = Response{Data: "UnAuthenticated", Error: &Err{Code: 405, Message: "UnAuthenticated"}}
