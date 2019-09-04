package main

import "MyGoLibrary/LeetCode"

func main()  {
	listCode := new(LeetCode.ListCode)
	l1 := listCode.GenerateList(1,1,1)
	listCode.PrintList(l1)
	l3 := listCode.DeleteDuplicates3(l1)
	listCode.PrintList(l3)
}
