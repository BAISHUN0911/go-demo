package common

import (
	"fmt"
	"time"
)

func TimeDemo1() {
	// 上个月
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	oneMonthAgoFmt := oneMonthAgo.Format("2006-01-02")
	fmt.Println(oneMonthAgo)
	fmt.Println(oneMonthAgoFmt)
}
