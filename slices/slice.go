package main

// slice

// /most used  concept
// usefull methods are included
// slice is a dynamic data set
func main() {

	////uninitialized slice is nil
	// var numbers []int
	// fmt.Println(numbers == n il)
	// fmt.Println(len(numbers))

	////make([]int,length,capasity)

	// var nums = make([]int, 2, 5)
	// var nums = make([]int, 0, 5)
	// fmt.Println(nums == nil)
	// fmt.Println(cap(nums))
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	//copy Function
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 1)
	// var num2 = make([]int, len(nums))
	// copy(num2, nums)
	// fmt.Println(nums, num2)

	//Slice Operator
	// variable[Start:End]

	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[1:3])

	//Compare two array
	// use full to compare two arrays
	// var num = []int{1, 2, 3}
	// var num2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(num, num2))

	// var nums = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}
	// fmt.Println(nums)

}
