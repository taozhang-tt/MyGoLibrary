package main

import "MyGoLibrary/LeetCode"

func main()  {
	listCode := new(LeetCode.ListCode)
	l1 := listCode.GenerateList(1,2,3)
	ret := listCode.ReverseList2(l1)
	listCode.PrintList(ret)
}
