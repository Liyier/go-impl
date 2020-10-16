package sort

// 冒泡排序 O(n2) 稳定
func bubbleSort(arr []int)  {
	var (
		length = len(arr)
		flag bool
	)
	for j := length-1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if arr[i] > arr[i+1] {
				// swap
				arr[i], arr[i+1] = arr[i+1], arr[i]
				// 存在交换
				flag = true
			}
		}
		if !flag {
			// 不存在交换， 数据已经有序
			break
		}
	}
}
