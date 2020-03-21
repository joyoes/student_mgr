package main

import "fmt"

//学生管理系统
//有一个物件：
// 1.它保存了一些数据	--> 结构体的字段
// 2.他有4个功能	    --> 结构体的方法
type student struct {
	id   int64
	name string
}

//造一个学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}

//查看学生
func (s studentMgr) showAllStudent() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", stu.id, stu.name)
	}
}

//添加学生
func (s studentMgr) addStudent() {
	//向allStudent中添加一个新学生
	//1.创建一个新学生
	//1.1 获取用户输入
	var (
		stuID   int64
		stuName string
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&stuID)
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&stuName)
	//1.2 创建一个学生
	newStu := student{
		id:   stuID,
		name: stuName,
	} //newStu是student的指针类型
	//2.添加到allStudent这个map中
	s.allStudent[newStu.id] = newStu
}

//删除学生
func (s studentMgr) deleteStudent() {
	// 1.请用户输入要删除的学生序号
	var deleteID int64
	fmt.Print("请输入要删除的学生序号:")
	fmt.Scanln(&deleteID)
	// 2.先去map中查看有没有这个学生,如果没有打印查无此人的提示
	_, ok := s.allStudent[deleteID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	//删除map中的元素
	delete(s.allStudent, deleteID)
}

// 编辑学生
func (s studentMgr) editStudent() {
	//1.获取用户输入的学号
	var editID int64
	fmt.Print("请输入要编辑的学生序号:")
	fmt.Scanln(&editID)
	//2.展示该学号对应的学生信息,如果没有提示查无此人
	stuObj, ok := s.allStudent[editID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息如下:学号:%d 姓名:%s\n", stuObj.id, stuObj.name)
	//3.请输入需要修改的学生姓名
	var newName string
	fmt.Print("请输入学生的新姓名:")
	fmt.Scanln(&newName)
	//4.修改学生姓名
	// s.allStudent[editID].name = newName
	stuObj.name = newName
	s.allStudent[editID] = stuObj
}
