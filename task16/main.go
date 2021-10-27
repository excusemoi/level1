package task16

import "fmt"

func Partition(arr []int64, low, high int64) ([]int64, int64) { //разделение массива на две части перестановками элементов
	//в результате слева от индекса i элементы arr[k] < arr[i], справа элементы arr[k] > arr[i]
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func QuickSort(arr []int64, low, high int64) []int64 {
	if low < high { //база рекурсии - длина подмассива > 2
		var p int64
		arr, p = Partition(arr, low, high) //слева от p - элементы arr[i] < arr[p], справа от p - элементы arr[i] > arr[p]
		arr = QuickSort(arr, low, p-1) //рекурсивный вызов алгоритма для левого подмассива
		arr = QuickSort(arr, p+1, high) //рекурсивный вызов для правого подмассива
	}
	return arr
}

func Solve(arr []int64){
	fmt.Println(QuickSort(arr, 0, int64(len(arr)-1)))
}