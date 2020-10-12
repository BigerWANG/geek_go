package slice_test

import "testing"

func TestSlice(t *testing.T) {
	var s0 []int  //声明一个silce

	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))  //

	// make ( []Type ,length, capacity ) 两个参数 长度和容量
	s1 := make([]int, 3)  // len 是初始化切片的长度，cap是可访问的容量
	s1 = append(s1, 8)  // 已经初始化的前三个值不会被改变，只会在最后append
	t.Log(s1[0])
	t.Log(s1[1])
	t.Log(s1[2])
	t.Log(s1[3])
}

func TestSilceGrowing(t *testing.T) {
	var s []int
	for i:=0; i<10;i++ {
		s = append(s, i)  // cap的增长是倍数级的 
		t.Log(len(s), cap(s))
	}
	
}

func TestSilceShareMem(t *testing.T) {
	// 引用类型
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}


	t.Log(year, len(year), cap(year))

	year = append(year, "dada")  // 当容量不够时会自动扩容，扩容长度是原silce的两倍
	year = append(year, "dada") // 这个时候容量够用，就不需要扩容了


	t.Log(year, len(year), cap(year))



	Q3 := year[3:6]
	t.Log(Q3, len(Q3), cap(Q3))  // len: 3, cap: 9

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))  // len:3, cap: 7, cap表示的是截取位置到剩下的存储空间的容量：len(year) - 截取的开始位置


	t.Log()

}

