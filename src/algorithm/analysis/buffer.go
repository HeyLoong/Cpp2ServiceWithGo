package analysis

import "C"
import (
	"micoService/DLL"
	"strconv"
	"unsafe"

	"github.com/gin-gonic/gin"
)

// 调用dll中点缓冲区算法:
func PointBuffer(c *gin.Context) {
	//获取函数反射
	PointBuffer := DLL.Lib.NewProc("c_PointBuffer")
	//url参数获取 http://localhost:8080/pointbuffer?req=POINT(1%202)&distance=0.5&n=16
	req := c.Query("req")
	distanceStr := c.Query("distance")
	nStr := c.Query("n")

	// 检查必需的参数是否存在
	if req == "" || distanceStr == "" || nStr == "" {
		c.JSON(400, gin.H{"error": "Missing required parameters"})
		return
	}

	// 解析 distance 和 n 参数为数值类型
	distance, err := strconv.ParseFloat(distanceStr, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid distance parameter"})
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid n parameter"})
		return
	}

	//调用算法PointBuffer
	retPointer, _, err := PointBuffer.Call(DLL.StrPtr(req), DLL.FloatPtr(distance), DLL.IntPtr(n))
	if err != nil {
		// 处理错误
	}

	//将字符串结果解码输出
	//way：读取 C 字符串直到遇到结束符 '\0'  UTF-8
	var retBytes []byte
	for i := uintptr(0); ; i++ {
		b := *(*byte)(unsafe.Pointer(retPointer + i))
		if b == 0 {
			break
		}
		retBytes = append(retBytes, b)
	}

	// 返回的数据使用了 UTF-8 编码
	retStr := string(retBytes)
	c.JSON(200, gin.H{"result": retStr})

}
