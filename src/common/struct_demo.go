package common

import (
	"fmt"
	"time"
)

type DemoUser struct {
	Username  string
	Role      string
	Status    int // 1: active, 0: inactive
	CreatedAt time.Time
}

func StructDemo() {
	u1 := DemoUser{Username: "admin", Role: "admin", CreatedAt: time.Now()}
	_ = u1

	u2 := new(DemoUser)
	u2.Username = "Alice"
	fmt.Println("u2的名字: ", u2.Username)

	u3 := DemoUser{"Alice", "user", 1, time.Now()}
	_ = u3

	u4 := DemoUser{Username: "abc"}
	u5 := DemoUser{Username: "abc"}
	fmt.Println(u4 == u5)

}
