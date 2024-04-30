// Package hostinfo shows host info
// 不用反复采集，一次性读取的信息
// 一般包含主机上不会变化的信息
package hostinfo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/klauspost/cpuid/v2"
)

const Path = "/hostinfo"

type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，表明内容类型是 JSON
	w.Header().Set("Content-Type", "application/json")

	// 序列化数据为 JSON 字符串
	jsonBytes, err := json.Marshal(cpuid.CPU)
	if err != nil {
		// 如果序列化过程中出现错误，通常应记录错误并可能返回一个错误状态
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 将 JSON 字符串写入响应体
	_, err = w.Write(jsonBytes)
	if err != nil {
		// 如果写入响应体时出错，同样应该记录错误
		// 在生产环境中通常会有更完善的错误处理机制
		fmt.Println(err)
	}
}
