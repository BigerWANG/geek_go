package empty_interface

import (
	"fmt"
	"testing"
)

/*
空接口与断言
1，空接口可以表示任何类型
2，通过断言来将空接口转换为定制类型 v, ok:=p.(int) ok=true时转换成功
 */

func DoSomething(p interface{}) (string, error) {

	switch v:=p.(type) {
	case int:
		fmt.Println("Integer")
		return fmt.Sprintf("%d is Integer type", v), nil

	case string:
		return fmt.Sprintf("%s is String type", v), nil
	case *Mytype:
		return fmt.Sprintf("%s is Mytype", v), nil
	}

	return "", fmt.Errorf("fail")
}

type Mytype struct {

}

func Test(t *testing.T) {
		if res, err := DoSomething(&Mytype{}); err==nil{
			t.Log(res)
		}else {
			t.Error(err)
		}

}


