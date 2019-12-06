package model

//定义一个结构体
type student struct{
	Name string
	score float64
}

//因为student结构体首字母小写，因此只能在model包使用
//通过工厂模式来解决这个问题
func NewStudent(n string, s float64) *student {
	return &student{
		Name : n,
		score: s,
	}
}

//如果score字段首字母小写，则在其它包不可以直接调用，我们可以提供一个方法
func (s *student) GetScore() float64 {
	return s.score
}