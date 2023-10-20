package main

// import "strings"

type Person struct {
	Name          string
	CriminalScore int
}

type nodeQ struct {
	val  Person
	next *nodeQ
}

type queue struct {
	head *nodeQ
	len  int
}

func (Q *queue) enqueue(input Person) {
	newNode := &nodeQ{val: input, next: nil}
	if Q.head == nil {
		Q.head = newNode
		Q.len++
	} else {
		curr := Q.head
		var temp *nodeQ
		for isPerson1Higher(curr.val, newNode.val) {
			if curr.next == nil {
				curr.next = newNode
				curr = newNode
			} else {
				temp = curr
				curr = curr.next
			}
		}
		if curr == Q.head {
			newNode.next = Q.head
			Q.head = newNode
		} else if temp != nil && curr != newNode {
			newNode.next = curr
			temp.next = newNode
		}
		Q.len++
	}
}

func (Q *queue) dequeue() (output Person) {
	if Q.len == 1 {
		output = Q.head.val
		Q.head = nil
		Q.len--
		return
	} else if Q.len > 1 {
		curr := Q.head
		Q.head = Q.head.next
		curr.next = nil
		output = curr.val
		Q.len--
		return
	}
	return
}

func (Q *queue) getPerson(input string) (output Person) {
	if Q.head != nil {
		curr := Q.head
		if curr.val.Name == input {
			return Q.dequeue()
		} else {
			for curr.next != nil {
				if curr.val.Name == input {
					return curr.val
				}
				curr = curr.next
			}
		}
		return
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func isName1Higher(name1, name2 string) bool {
	minLength := min(len(name1), len(name2))
	i := 0
	for i < minLength {
		if name1[i] != name2[i] {
			break
		}
		i++
	}
	if i == minLength {
		if len(name1) > len(name2) {
			return false
		} else if len(name1) < len(name2) {
			return true
		} else {
			return false // Is equal
		}
	} else {
		if name1[i] > name2[i] {
			return false
		} else {
			return true
		}
	}
}

func isPerson1Higher(person1, person2 Person) bool {
	if person1.CriminalScore < person2.CriminalScore {
		return true
	} else if person1.CriminalScore == person2.CriminalScore {
		if isName1Higher(person1.Name, person2.Name) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// Task 1.a
func LastDayInJail(criminals []Person, chosenPerson string) (onTransport []Person, waiting []Person) {
	// Write your code here
	// --------------------
	Q := queue{}
	for _, person := range criminals {
		Q.enqueue(person)
	}
	i := 5
	personPlus := Q.getPerson(chosenPerson)
	for Q.head != nil && i != 0 {
		if len(onTransport) < 3 {
			onTransport = append(onTransport, Q.dequeue())
		} else {
			waiting = append(waiting, Q.dequeue())
		}
		i--
	}
	if len(onTransport) < 3 && personPlus.Name != "" {
		onTransport = append(onTransport, personPlus)
	} else if personPlus.Name != "" {
		waiting = append(waiting, personPlus)
	}
	return
	// --------------------
}
