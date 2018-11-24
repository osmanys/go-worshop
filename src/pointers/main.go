package main

import "fmt"

func main(){
  list := []int{12, 4, 5, 6, 20, 10, 8, 15, 15}
  root := new(Tree)
  for _, v := range list {
    Add(root, v)
  }
  fmt.Println(root)
  //printInOrder(root)
}

type Tree struct {
  value int
  left, right *Tree
}

func Add(t *Tree, value int) *Tree {
  if t == nil {
    aux := Tree{value: value}
    return &aux
  }
  if value == t.value {
    return t
  }
  if value < t.value {
    t.left = Add(t.left, value)
  } else {
    t.right = Add(t.right, value)
  }
  return t
}

