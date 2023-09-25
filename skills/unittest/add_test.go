package unittest_test

import (
	"testing"
	"wuchenyanghaoshuai/vblog/skills/unittest"
)

func TestSum(t *testing.T) {
	t.Log(unittest.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	// if unittest.Sum(1, 2) == 3 {
	// 	t.Fatal("测试成功")
	// } else {
	// 	t.Log("测试失败")
	// }
    

}
