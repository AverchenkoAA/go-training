package collection

import (
	"fmt"
)
//Count int - количество элементов в коллекции
var Count int
//Last - указатель на последний элемент коллекции
var last *ListNode
//ListNode - основная структура элементов коллекции
type ListNode struct{
	NameL string
	NextL, PrevL *ListNode
}
//Add – добавляет элемент в коллекцию 
func (n *ListNode) Add(element string){
	e := ListNode{NameL: element}
	if Count == 0{
		*n=e
		last=n
	}else{
		e.PrevL=last
		(*last).NextL=&e
		last=&e
	}
	Count++
}
//Next – возвращает следующий элемент коллекции
func (n *ListNode) Next() *ListNode{
	if (n!=nil)&&(n.NextL!=nil){
		return n.NextL
	}
	return nil
}
//Prev – возвращает предыдущий элемент коллекции
func (n *ListNode) Prev() *ListNode{
	if (n!=nil)&&(n.PrevL!=nil){
		return n.PrevL
	}
	return nil
}
//Value – возвращает значение узла
func (n *ListNode) Value() string{
	return n.NameL
}
//Print – выводит коллекцию в консоль 
func (n *ListNode) Print(){
	for n!=nil {
		fmt.Print(" ",n.Value())
		n=n.Next()
	}
}
//Remove – удаляет элемент с данным индексом из коллекции
func (n *ListNode) Remove(index int){
	if (index>=0)&&(index<Count){
		el:=n.Get(index)
		el.Prev().NextL=el.NextL
		el.Next().PrevL=el.PrevL
		Count--
	}
}
//Get - возвращает элемент с данным индексом из коллекции
func (n *ListNode) Get(index int) *ListNode{
	if (index>=0)&&(index<Count){
		for i:=0;i<index;i++{
			n=n.Next()
		}
		return n
	}
	return nil
}
//First – возвращает первый элемент в коллекции
func (n *ListNode) First() *ListNode{
		return n
}
//Last – возвращает последний элемент в коллекции 
func (n *ListNode) Last() *ListNode{
	return last
}
//Length – возвращает размер коллекции
func (n *ListNode) Length() int{
	return Count
}
