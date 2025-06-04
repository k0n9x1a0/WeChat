package middleware

import (
	"github.com/astaxie/beego/context"
	"sync"
)

var (
	requestCountMu sync.Mutex
	requestCount   int
)

var BaseAuthLog = func(ctx *context.Context) {
	// 在请求处理前递增请求计数器
	incrementRequestCount()

	// 检查请求计数是否超出阈值
	if getRequestCount() > 2000000 {
		resp := map[string]interface{}{
			"status":  "error",
			"message": "请求过多",
		}
		ctx.Output.JSON(resp, false, false)
		ctx.Abort(403, "")
		return
	}

	// 在这里执行其他身份验证逻辑或记录请求的任何其他操作
}

func incrementRequestCount() {
	requestCountMu.Lock()
	defer requestCountMu.Unlock()
	requestCount++
}

func getRequestCount() int {
	requestCountMu.Lock()
	defer requestCountMu.Unlock()
	return requestCount
}
