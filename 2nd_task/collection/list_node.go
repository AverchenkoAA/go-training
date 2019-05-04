package collection

//ListNode - основная структура элементов коллекции
type ListNode struct{
	nameEl string
	nextEl, prevEl *ListNode
}
//Next – возвращает следующий элемент коллекции
func (n *ListNode) Next() *ListNode{
	if (n!=nil)&&(n.nextEl!=nil){
		return n.nextEl
	}
	return nil
}
//Prev – возвращает предыдущий элемент коллекции
func (n *ListNode) Prev() *ListNode{
	if (n!=nil)&&(n.prevEl!=nil){
		return n.prevEl
	}
	return nil
}
//Value – возвращает значение узла
func (n *ListNode) Value() string{
	return n.nameEl
}
