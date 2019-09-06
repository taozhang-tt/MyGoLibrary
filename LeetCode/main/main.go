package main

import "MyGoLibrary/LeetCode"

func main()  {
	listCode := new(LeetCode.ListCode)
	l1 := listCode.GenerateList()
	ret := listCode.SwapPairs(l1)
	listCode.PrintList(ret)
}
