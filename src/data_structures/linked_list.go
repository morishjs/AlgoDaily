package ds

type List struct {
	Root *ListNode
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
	node.Next = list.Root
	list.Root = &node
	
	list.len++
}

func (list *List) Append(val int) {
	list.lazyInit()

	listIterator := list.Root
	if listIterator == nil {
		list.Root = &ListNode{val, nil}
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

	listIterator := list.Root
	for listIterator != nil {
		arr = append(arr, listIterator.Val)
		listIterator = listIterator.Next
	}

	return arr
}

func (list *List) Contains(val int) bool {
	listIter := list.Root

	for listIter != nil {
		if listIter.Val == val {
			return true
		}

		listIter = listIter.Next
	}

	return false
}

func (list *List) lazyInit() {
	if list.Root == nil {
		list = list.Init()
	}
}

func (list *List) Init() *List {
	list.len = 0

	return list
}
