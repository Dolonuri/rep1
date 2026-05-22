package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
}

func NewDoubleLinkedList(value int) *DoubleLinkedList {
	node := &Node{Value: value}
	return &DoubleLinkedList{
		Head: node,
		Tail: node,
	}
}

// Append - элемент в конец
func (list *DoubleLinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil { //если список пуст
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Prev = list.Tail //новый узел ссылается на текущ. хвост
		list.Tail.Next = newNode //текущий хвост ссылается на  новый узел
		list.Tail = newNode      //обновляем указатель Tail на новый узел
	}
}

// Prepend - элемент в начало списка
func (list *DoubleLinkedList) Prepend(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Next = list.Head //новый узел ссылается на текущ. голову
		list.Head.Prev = newNode //текущ голова ссылается на новый узел
		list.Head = newNode      //обновляем указатель Head на новый узел
	}
}

// RemoveHead - удаление элемента из начала списка
func (list *DoubleLinkedList) RemoveHead() {
	if list.Head == nil {
		fmt.Println("Error")
		return
	}

	if list.Head == list.Tail { // Если в списке только один элемент
		list.Head = nil
		list.Tail = nil
	} else {
		list.Head = list.Head.Next // Сдвигаем голову вперед
		list.Head.Prev = nil       // Обнуляем обратный указатель у новой головы
	}
}

//RemoveTail - удаление элемента из конца списка
func (list *DoubleLinkedList) RemoveTail() {
	if list.Tail == nil {
		fmt.Println("Error")
		return
	}

	if list.Head == list.Tail { // Если в списке только один элемент
		list.Head = nil
		list.Tail = nil
	} else {
		list.Tail = list.Tail.Prev // Сдвигаем хвост назад
		list.Tail.Next = nil       // Обнуляем прямой указатель у нового хвоста
	}
}

// Print - для проверки (выводит список)
func (list *DoubleLinkedList) Print() {
	if list.Head == nil {
		fmt.Println("Empty!")
		return
	}

	current := list.Head
	fmt.Print("Head -> ")
	for current != nil {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print(" ")
		}
		current = current.Next
	}
	fmt.Println(" <- Tail")
}

//Тест
func main() {
	list := &DoubleLinkedList{} //Создаем список

	list.Append(3) //добавляем в конец 3, 2 и 1
	list.Append(2)
	list.Append(1)

	list.Print() // 3 2 1

	list.Prepend(2) //Добавляем 2 в начало списка

	list.Print() // 2 3 2 1

	list.RemoveHead() //Удаляем элемент в начале списка
	list.RemoveTail() //Удаляем элемент в конце списка

	list.Print() // 3 2
}
