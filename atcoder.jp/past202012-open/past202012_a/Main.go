package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	ans := "draw"
	if len(s) >= 3 {
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+1] && s[i+1] == s[i+2] {
				ans = string(s[i])
				break
			}
		}
	}

	fmt.Println(ans)
}
