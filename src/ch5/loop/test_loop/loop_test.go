package test_loop

import (
	"testing"
)

func TestLoop(t *testing.T){
	res := 0
	//for a:=1; a<=100; a++ {
	//	res += a
	//}

	for {
		res += 1
		if res >= 100 {
			break
		}
	}
	t.Log(res)

}



