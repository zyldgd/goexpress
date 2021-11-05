package internal

import "container/list"

type Stack interface {
	Push(elem interface{})
	Pop() interface{}
	Peak() interface{}
	Len() int
	IsEmpty() bool
}

type DefaultStack struct {
	list *list.List
}

func NewDefaultStack() Stack {
	return &DefaultStack{list: list.New()}
}

//func (s *Stack) PushList(list ...interface{}) {
//	if s != nil {
//		for _, e := range list {
//			s.list.PushBack(e)
//		}
//	}
//}

func (s *DefaultStack) Push(elem interface{}) {
	if s != nil {
		s.list.PushBack(elem)
	}
}

func (s *DefaultStack) Pop() interface{} {
	if s != nil {
		e := s.list.Back()
		if e != nil {
			s.list.Remove(e)
			return e.Value
		}
	}
	return nil
}

func (s *DefaultStack) Peak() interface{} {
	if s != nil {
		e := s.list.Back()
		if e != nil {
			return e.Value
		}
	}
	return nil
}

func (s *DefaultStack) Len() int {
	if s != nil {
		return s.list.Len()
	}
	return 0
}

func (s *DefaultStack) IsEmpty() bool {
	if s != nil {
		return s.list.Len() == 0
	}
	return true
}
