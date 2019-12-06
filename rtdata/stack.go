package rtdata

import "fmt"

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/25 11:08
 */

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame // stack is implemented as linked list
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) push(frame *Frame) {
	fmt.Printf("push stack size =%d \n", self.size)
	if self.size > self.maxSize {
		panic("StackOverflowError")
	}

	if self._top != nil {
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	fmt.Printf("pop stack size =%d \n", self.size)
	if self._top == nil {
		panic("jvm stack is empty")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--

	return top
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}

	return self._top
}
