package main

import (
	ds "go-algo-daily/src/data_structures"
)

func getIntersection(listOne *ds.List, listTwo *ds.List) *ds.List {
	listOneArr := listOne.ToArray()
	listOneArrLen := len(listOneArr)

	findInsert := func(l []int) <-chan int {
		ch := make(chan int)

		go func() {
			for _, v := range l {
				if listTwo.Contains(v) == true {
					ch <- v
				}
			}

			close(ch)
		}()
		
		return ch
	}

	chan1 := findInsert(listOneArr[:(listOneArrLen / 2)])
	chan2 := findInsert(listOneArr[(listOneArrLen / 2):])

	newList := new(ds.List)

	for v := range chan1 {
		newList.Append(v)
	}
	for v := range chan2 {
		newList.Append(v)
	}

	return newList
}
