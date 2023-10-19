package main

import (
	"fmt"
	"strings"
)

type node struct {
	val   byte
	count int
	next  *node
}

type stack struct {
	top *node
}

func (S *stack) push(item byte) {
	newNode := &node{val: item, count: 1, next: nil}
	if S.top == nil {
		S.top = newNode
	} else {
		newNode.next = S.top
		S.top = newNode
	}
}

func (S *stack) pop() (byte, int) {
	if S.top != nil {
		popped := S.top.val
		count := S.top.count
		S.top = S.top.next
		return popped, count
	}
	return 0, 0
}

func instructions(count int, move byte) string {
	var mapTime = map[bool]string{
		true:  "times",
		false: "time",
	}

	var mapMove = map[byte]string{
		'R': "right",
		'L': "left",
		'A': "advance",
	}

	timeWord := false
	if count > 1 {
		timeWord = true
	}

	return fmt.Sprintf("Move %s %d % s", mapMove[move], count, mapTime[timeWord])
}

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// Task 3.a
func RobotTranslatorV2(cmd string) string {
	// Write your code here
	// --------------------

	S := stack{}
	for _, char := range cmd {
		if char == 'X' {
			if S.top != nil {
				if S.top.count > 1 {
					S.top.count--
				} else {
					S.pop()
				}
			}
		} else {
			if S.top != nil && (S.top.val == byte(char)) {
				S.top.count++
			} else {
				S.push(byte(char))
			}
		}
	}
	var arrString []string
	for S.top != nil {
		move, count := S.pop()
		if move != 'R' && move != 'L' && move != 'A' {
			return "Invalid command"
		}
		arrString = append(arrString, instructions(count, move))
	}
	reverseArray(arrString)

	return strings.Join(arrString, "\n")
	// --------------------
}
