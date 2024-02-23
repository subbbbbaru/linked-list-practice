package main

import "fmt"

// DO NOT EDIT ---------------------------------------------------------------------------------------------------------
type Node struct {
	value int
	next  *Node
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Next() *Node {
	return n.next
}

func printList(list *Node) {
	curNode := list
	for i := 0; curNode != nil; i++ {
		fmt.Printf("%d: %d\n", i, curNode.Value())
		curNode = curNode.Next()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// pushFront добавляет новый элемент в начало списка
func pushFront(list *Node, value int) {
	// panic("not implemented")
	curr := *list
	*list = Node{next: list, value: value}
    list.next = &curr
}

// pushBack добавляет новый элемент в конец списка
func pushBack(list *Node, value int) {
	// panic("not implemented")
	newNode := &Node{next: nil, value: value}
	curr := list

	if list == nil {
		return
	}

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

// count возвращает кол-во элементов в списке
func count(list *Node) int {
	// panic("not implemented")
	curr := list
	count := 0
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

// popFront возвращает значение первого элемента и удаляет его из списка
func popFront(list *Node) int {
	// panic("not implemented")

	if list != nil {
		curr := *list.next
		value := list.value
		*list = curr
		return value
	}
	return 0
}

// popBack возвращает значение последнего элемента и удаляет его из списка
func popBack(list *Node) int {
	// panic("not implemented")

	if list != nil {
		curr := list
		prevNode := list
		for curr.next != nil {
			prevNode = curr
			curr = curr.next
		}
		value := curr.value
		prevNode.next = nil
		list = prevNode
		return value
	}
	return 0
}

// isValueInList ищет значение в списке и возвращает true, если оно найдено, в ином случае - false
func isValueInList(list *Node, value int) bool {
	// panic("not implemented")
	curr := list
	for curr != nil {
		if curr.value == value {
			return true
		}
		curr = curr.next
	}
	return false
}

// getValueByIndex возвращает значение из списка по индексу, если оно есть, в ином случае - error("index out of range")
func getValueByIndex(list *Node, index int) (int, error) {
	// panic("not implemented")
	count := count(list)
	curr := list
	if index <= 0 || index > count {
		return 0, fmt.Errorf("index out of range")
	}

	for count = 0; count < index; count++ {
		curr = curr.next
	}
	return curr.value, nil
}

// insert добавляет элемент в список в соответствующий индекс
func insert(list *Node, index, value int) {
	// panic("not implemented")
	currIndex := 0
	curr := list
	for curr != nil {
		if currIndex == index-1 {
			newNode := &Node{value: value, next: curr.next}
			curr.next = newNode
			return
		}
		currIndex++
		curr = curr.next
	}
}

// deleteByIndex удаляет элемент из списка по индексу и возвращает его значение. Если индекс неправильный - возвращает error("index out of range")
func deleteByIndex(list *Node, index int) (int, error) {
	// panic("not implemented")

	count := count(list)
	curr := list
	var prevNode *Node

	if index <= 0 || index > count {
		return 0, fmt.Errorf("index out of range")
	}

	for count = 0; count < index; count++ {
		prevNode = curr
		curr = curr.next
	}


	value := curr.value
	prevNode.next = curr.next
	
	return value, nil
}

func bubbleSort(list *Node) {
	if list == nil || list.next == nil {
		return
	}

	swapped := true
	for swapped {
		swapped = false
		cur := list
		for cur.next != nil {
			if cur.value > cur.next.value {
				cur.value, cur.next.value = cur.next.value, cur.value
				swapped = true
			}
			cur = cur.next
		}
	}
}


// sort сортирует список (*)
func sort(list *Node) {
	// panic("not implemented")

	if list == nil || list.next == nil {
		return 
	}
	bubbleSort(list)

}

func main() {
	linkedList := &Node{
		value: 5,
		next: &Node{
			value: 3,
			next: &Node{
				value: 4,
				next: &Node{
					value: 1,
					next: &Node{
						value: 2,
						next:  nil,
					},
				},
			},
		},
	}
	printList(linkedList)


	// Test pushFront function
	fmt.Println("After pushing 10 to the front:")
	pushFront(linkedList, 10)
	printList(linkedList)

	// Test pushBack function
	fmt.Println("After pushing 20 to the back:")
	pushBack(linkedList, 20)
	printList(linkedList)

	// Test count function
	fmt.Println("Count of elements in the list:", count(linkedList))

	// Test popFront function
	val := popFront(linkedList)
	fmt.Println("Popped value from the front:", val)
	printList(linkedList)

	// Test popBack function
	val = popBack(linkedList)
	fmt.Println("Popped value from the back:", val)
	printList(linkedList)

	// Test isValueInList function
	fmt.Println("Is value 4 in the list?", isValueInList(linkedList, 4))

	// Test getValueByIndex function
	if val, err := getValueByIndex(linkedList, 1); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value at index 1:", val)
	}

	// Test insert function
	insert(linkedList, 2, 100)
	fmt.Println("After inserting 100 at index 2:")
	printList(linkedList)


	// Test deleteByIndex function
	if val, err := deleteByIndex(linkedList, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Deleted value at index 2:", val)
		printList(linkedList)
	}

	// Test Bubble Sort
	fmt.Println("Sorting the list using Bubble Sort:")
	bubbleSort(linkedList)
	printList(linkedList)
}
