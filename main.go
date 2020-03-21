package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎光临学生管理系统")
	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.编辑学生
		5.退出系统
	`)
}

var smr studentMgr //全局变量studentMgr

func main() {
	var smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {

		//1.打印菜单
		showMenu()
		fmt.Print("请输入要执行的操作:")
		//2.等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		//3.执行相应的操作
		switch choice {
		case 1:
			smr.showAllStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.deleteStudent()
		case 4:
			smr.editStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("请输入1~5之间的数字")
		}
	}
}
