package task17

import (
	"fmt"
)

func BinarySearch(arr []int64, elemToSearch int64) (int64, bool) {
	var (
		length = int64(len(arr))
		l int64 = 0
		r = length-1
		m int64
	)
	arr = QuickSort(
		arr,
		0,
		length-1,
	)

	//коллекция должна быть отсортированной
	//соотв. порядок коллекции позволяет с помощью укорачивания границ поиска
	//с левой стороны на m-l, в случае, если средний элемент коллекции arr[m] меньше искомого
	//с правой на r-m, если arr[m] больше искомого
	//если они равны - возвращаем индекс элемента и true
	//если искомого элемента в коллекции не оказалось - возвращаем 0 и false

	for ; l <= r;  {
		m = (l + r) / 2
		if elemToSearch < arr[m] {
			r = m - 1
		} else if elemToSearch > arr[m] {
			l = m + 1
		} else {
			return m, true
		}
	}
	return 0, false

}

func Solve(arr []int64, elemToSearch int64) {
	if i, ok := BinarySearch(arr, elemToSearch); ok {
		fmt.Printf("Index of elem %d in the array %v is %d\n", elemToSearch, arr, i)
	} else {
		fmt.Printf("Elem %d doesn't exist in the array %v\n", elemToSearch, arr)
	}
}

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
