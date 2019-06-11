package ds

type List struct {
	root *ListNode
	len int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *List) Prepend(val int) {
	list.lazyInit()

	node := ListNode{}
	node.Val = val
	node.Next = list.root
	list.root = &node
	
	list.len++
}

func (list *List) Append(val int) {
	list.lazyInit()

	listIterator := list.root
	if listIterator == nil {
		list.root = &ListNode{val, nil}
		return
	}

	for listIterator.Next != nil {
		listIterator = listIterator.Next
	}

	listIterator.Next = &ListNode{val, nil}
	list.len++
}

func (list *List) ToArray() []int {
	arr := []int{}

	listIterator := list.root
	for listIterator != nil {
		arr = append(arr, listIterator.Val)
		listIterator = listIterator.Next
	}

	return arr
}

func (list *List) lazyInit() {
	if list.root == nil {
		list = list.Init()
	}
}

func (list *List) Init() *List {
	list.len = 0

	return list
}
