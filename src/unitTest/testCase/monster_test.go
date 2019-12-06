package monster

import (
	"testing"
)

//测试用例
func TestStore(t *testing.T) {

	//先创建一个Monster实例
	monster := Monster{
		Name : "红孩儿",
		Age : 10,
		Skill : "吐火",
	}
	res := monster.Store()

	if !res {
		t.Fatalf("monster.Store() 错误，希望为%v 实际为%v", true, res)
	}
	t.Logf("测试成功")
}

func TestReStore(t *testing.T) {

	//先创建一个Monster实例
	var monster Monster
	res := monster.ReStore()

	if !res {
		t.Fatalf("monster.ReStore() 错误，希望为%v 实际为%v", true, res)
	}

	//进一步测试
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() 错误，希望为%v 实际为%v", true, res)
	}
	t.Logf("测试成功")
}
