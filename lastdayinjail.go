package main

type Person struct {
	Name          string
	CriminalScore int
}

type priorityQueue struct {
	items []Person
}

func (pq *priorityQueue) push(person Person) {
	pq.items = append(pq.items, person)
	pq.upHeapify(len(pq.items) - 1)
}

func (pq *priorityQueue) pop() (output Person) {
	if len(pq.items) != 0{
        output = pq.items[0]
        last := pq.items[len(pq.items)-1]
        pq.items[0] = last
        pq.items = pq.items[:len(pq.items)-1]
        pq.downHeapify(0)
        return
    }
	return
}

func (pq *priorityQueue) upHeapify(idx int) {
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if isPerson1Higher(pq.items[idx], pq.items[parentIdx]) {
			pq.items[idx], pq.items[parentIdx] = pq.items[parentIdx], pq.items[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (pq *priorityQueue) downHeapify(idx int) {
	leftChild := 2*idx + 1
	rightChild := 2*idx + 2
	smallest := idx

	if leftChild < len(pq.items) && isPerson1Higher(pq.items[leftChild], pq.items[smallest]) {
		smallest = leftChild
	}
	if rightChild < len(pq.items) && isPerson1Higher(pq.items[rightChild], pq.items[smallest]) {
		smallest = rightChild
	}

	if smallest != idx {
		pq.items[idx], pq.items[smallest] = pq.items[smallest], pq.items[idx]
		pq.downHeapify(smallest)
	}
}

func (pq *priorityQueue) getPersonByName(name string) (output Person) {
	for i, person := range pq.items {
		if person.Name == name {
			return pq.items[i]
		}
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

func LastDayInJail(criminals []Person, chosenPerson string) (onTransport []Person, waiting []Person) {
	pq := priorityQueue{}

	for _, person := range criminals {
		pq.push(person)
	}

	i := 5
	chosen := pq.getPersonByName(chosenPerson)

	for i > 0 && len(pq.items) > 0 {
		person := pq.pop()
		if len(onTransport) < 3 {
			onTransport = append(onTransport, person)
		} else {
			waiting = append(waiting, person)
		}
		i--
	}
	if len(onTransport) < 3 && chosen.Name != "" {
		onTransport = append(onTransport, chosen)
	} else if chosen.Name != "" {
		waiting = append(waiting, chosen)
	}
	return
}
