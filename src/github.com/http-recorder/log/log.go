package log

import (
	"fmt"
	"time"
)

const timeLayout = "Jan 2, 2006 at 3:04:05"

func RecorderInfo(str ...interface{}) {
	fmt.Println(time.Now().Format(timeLayout), "[HTTP-RECORDER]", str)
}

func RetrieverInfo(str ...interface{}) {
	fmt.Println(time.Now().Format(timeLayout), "[HTTP-RETRIEVER]", str)
}
