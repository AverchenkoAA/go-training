package collection

import (
	"fmt"
)
//List - основная структура коллекции
type List struct{
	count int
	firstEl, lastEl *ListNode
}
//Add – добавляет элемент в коллекцию 
func (n *List) Add(element string){
	e:=ListNode{nameEl: element}
	if n.count == 0{
		n.firstEl=&e
		n.lastEl=&e
	}else{
		e.prevEl=n.lastEl
		n.lastEl.nextEl=&e
		n.lastEl=&e
	}
	n.count++
}
//Print – выводит коллекцию в консоль 
func (n *List) Print(){
	i:=n.firstEl
	for i!=nil {
		fmt.Print(i.Value()," ")
		i=i.Next()
	}
}
//Remove – удаляет элемент с данным индексом из коллекции
func (n *List) Remove(index int){
	if (index>=0)&&(index<n.count){
		el:=n.Get(index)
		el.Prev().nextEl=el.nextEl
		el.Next().prevEl=el.prevEl
		n.count--
	}
}
//Get - возвращает элемент с данным индексом из коллекции
func (n *List) Get(index int) *ListNode{
	e:=n.firstEl
		for i:=0;i<index;i++{
			e=e.Next()
		}
		return e
}
//First – возвращает первый элемент в коллекции
func (n *List) First() *ListNode{
		return n.firstEl
}
//Last – возвращает последний элемент в коллекции 
func (n *List) Last() *ListNode{
	return n.lastEl
}
//Length – возвращает размер коллекции
func (n *List) Length() int{
	return n.count
}
