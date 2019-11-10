package main

import "fmt"

type Node struct {
  value int
  next *Node
}

type Stack struct {
  size int
  top *Node
}

func (s *Stack) Push(n *Node) {
  n.next = s.top
  s.top = n
  s.size++
}

func (s *Stack) Pop(n *Node){
  if s.size == 0 {
    defer func(){
      fmt.Println("The stack is empty!")
      recover()
    }()
  }
  n = s.top
  s.top = s.top.next
  s.size--
}

func main() {
  s := new(Stack)

  s.Push(&Node{5, nil})
  s.Push(&Node{9, nil})

  fmt.Println(*s)
}
