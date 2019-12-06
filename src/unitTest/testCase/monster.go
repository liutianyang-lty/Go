package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

//待测试的方法
func (this *Monster) Store() bool {

	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err=",err)
		return false
	}

	//保存到文件
	filePath := "e:/monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err=", err)
		return false
	}

	return true
}

//待测试的方法
func (this *Monster) ReStore() bool {

	//先读取
	filePath := "e:/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err=", err)
		return false
	}

	//反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("Unmarshal err=", err)
		return false
	}

	return true
}
