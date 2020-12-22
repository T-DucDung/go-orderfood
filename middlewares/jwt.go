package middlewares

import (
	"github.com/astaxie/beego/context"
)

var Log = func(ctx *context.Context) {
	ctx.Request.Header.Add("StoreID", "1")
}
