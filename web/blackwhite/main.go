package main

import "fmt"

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

type EmpLink struct {
	Head *Emp
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *EmpLink) insert(emp *Emp) {
	cur := this.Head
	fmt.Println("start input emp")
	if cur == nil {
		this.Head = emp
		return
	}
	var pre *Emp = nil // 辅助指针 pre 在cur前面
	// 先去找位置, 找到则直接退出
	for {
		if cur.Next != nil {
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	pre.Next = emp
	emp.Next = cur
}

func (this *EmpLink) showLink(no int) {
	cur := this.Head
	if cur == nil {
		fmt.Println("sorry, your link is nil")
		return
	}
	for {
		if cur != nil {
			fmt.Printf("链表:%d 雇员:%d 名字:%s \n", no, cur.Id, cur.Name)
			cur = cur.Next
		}
		break
	}
}

func (this *HashTable) Insert(emp *Emp) {
	linkNo := this.HashFun(emp.Id)

	this.LinkArr[linkNo].insert(emp)
}

func (this *HashTable) Show() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].showLink(i)
	}
}
func (this *HashTable) HashFun(id int) int {
	return id % 7
}

func main() {
	var key, name string
	var id int
	var hashTable HashTable
	for {
		fmt.Println("======雇员系统菜单======")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示添加雇员")
		fmt.Println("find 表示添加雇员")
		fmt.Println("exit 表示添加雇员")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.Show()
		case "find":
			hashTable.Show()
		case "exit":
			hashTable.Show()
		default:
			fmt.Println("输入错误")
		}
	}
}
