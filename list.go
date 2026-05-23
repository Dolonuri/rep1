package main

import "fmt"

type Node struct {
	Value any
	Next  *Node
	Prev  *Node
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
}

func NewDoubleLinkedList(value any) *DoubleLinkedList {
	node := &Node{Value: value}
	return &DoubleLinkedList{
		Head: node,
		Tail: node,
	}
}

// Append - элемент в конец
func (list *DoubleLinkedList) Append(value any) {
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
func (list *DoubleLinkedList) Prepend(value any) {
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

//AddAfterIndex - добавление нового элемента после определенного
func (list *DoubleLinkedList) AddAfterIndex(value any, index int) {
	newNode := &Node{Value: value}

	if list.Head == nil { //ошибка, если список пуст
		fmt.Print("Error!")
	}

	current := list.Head //стартовая позиция
	position := 0

	for current != nil && position < index { //поиск по индексу, начиная от Head = 0
		current = current.Next
		position++
	}

	if current == list.Tail { //если элемент - хвост, то просто Append
		current.Next = newNode
		newNode.Prev = current
		list.Tail = newNode
	} else {
		nextNode := current.Next //сохраняем адрес следующего после current элемента
		newNode.Next = nextNode  //связываем новый узел со следующим элементом
		newNode.Prev = current   //связываем новый узел с предыдущим элементом
		current.Next = newNode   //связываем следующий элемент с новым узлом
		nextNode.Prev = newNode  //связываем предыдущий элемент с новым узлом
	}
}

//RemoveAfterIndex - удалить элемент идущий после определенного
func (list *DoubleLinkedList) RemoveAfterIndex(index int) {

	if list.Head == nil { //ошибка, если список пуст
		fmt.Print("Error!")
	}

	current := list.Head //стартовая позиция
	position := 0

	for current != nil && position < index { //поиск по индексу, начиная от Head = 0
		current = current.Next
		position++
	}

	if current.Next == nil { //если сказали удалить элемент после хвоста
		fmt.Print("Error! Nothing to delete")
	}

	nodeToDelete := current.Next //сохраняем ссылку на удаляемый элемент

	// Перенастраиваем указатели
	current.Next = nodeToDelete.Next

	if nodeToDelete.Next != nil {
		// Если есть узел после удаляемого, обновляем его Prev
		nodeToDelete.Next.Prev = current
	} else {
		// Если удаляем последний элемент (Tail), обновляем Tail
		list.Tail = current
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

	list.Append("a") //добавляем в конец a, 2 и false
	list.Append(2)
	list.Append(false)

	list.Print() // a 2 false

	list.Prepend(2) //Добавляем 2 в начало списка

	list.Print() // 2 a 2 false

	list.RemoveHead() //Удаляем элемент в начале списка
	list.RemoveTail() //Удаляем элемент в конце списка

	list.AddAfterIndex(15, 0)

	list.Print() // a 15 2

	list.RemoveAfterIndex(1)

	list.Print() // a 15
}
