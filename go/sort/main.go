package sort_df

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const count = 100

func init() {
	rand.Seed(time.Now().Unix())
}
func generate() []int {
	list := []int{}
	for i := 0; i < count; i++ {
		list = append(list, rand.Intn(100))
	}
	return list
}

// 冒泡排序是一种简单的排序算法。
// 它重复地走访过要排序的元素列，
// 依次比较相邻的两个元素，
// 如果顺序（如从大到小、首字母从Z到A）错误就把他们交换过来。
// 走访元素的工作是重复地进行直到没有相邻元素需要交换，
// 也就是说该元素列已经排序完成。
// 这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。
// 不管任何时候时间复杂度都是指数相关
func bubble() []int {
	list := generate()
	fmt.Printf("冒泡%v\n", list)
	length := len(list)
	for i := 0; i < length-1; i++ {
		for j := length - 1; j > i; j-- {
			if list[j] < list[j-1] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
	fmt.Printf("冒泡%v\n", list)
	return list
}

// 选择排序同冒泡排序类似
func selekt() []int {
	list := generate()
	fmt.Printf("选择%v\n", list)

	length := len(list)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

// 插入排序
// 假设下标i左侧都是已经排好序的，则第i+1个元素向左插入到合适位置
func insert() []int {
	list := generate()
	fmt.Printf("插入%v\n", list)

	length := len(list)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j > 0; j-- {
			if list[j] >= list[j-1] {
				break
			}
			list[j-1], list[j] = list[j], list[j-1]
		}
	}

	fmt.Printf("插入%v\n", list)
	return list
}

func sort(arr *[]int) *[]int {
	if len(*arr) == 0 {
		return nil
	}
	i := 0
	length := len(*arr)
	j := length - 1
	for i < j {
		for (*arr)[j] >= (*arr)[i] && j > i {
			j--
		}
		(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]

		for (*arr)[i] <= (*arr)[j] && i < j {
			i++
		}
		(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]
	}
	a := (*arr)[:i]
	b := (*arr)[i+1:]
	sort(&a)
	sort(&b)

	return arr
}
func quick() []int {
	list := generate()
	fmt.Printf("快排%v\n", list)
	sort(&list)
	fmt.Printf("快排%v\n", list)
	return list
}

func bigRoot(arr []int) {
	length := len(arr)
	parent := int(math.Ceil(float64(length / 2)))
	for i := parent; i > 0; i-- {
		reBigRoot(arr, i-1, length - 1)
	}
}

func reBigRoot(arr []int, start int, end int) {
	// 第二次开始构造只需要从上往下构造即可
	// 只需要跟两个子节点其中之一来交换
	// 父节点的个数
	// parent := int(math.Ceil(float64(length / 2)))
	list := arr

	right := 0
	left := 0
	for i := start; right <= end; {
		left = 2*i + 1
		right = 2*i + 2
		if left > end {
			return
		}
		if list[i] >= list[left] && (right > end || list[i] >= list[right]) {
			// 比两个子节点都大
			break
		}
		if right <= end && list[right] >= list[left] {
			if list[i] < list[right] {
				list[i], list[right] = list[right], list[i]
				i = right
			}
		} else {
			if list[i] < list[left] {
				list[i], list[left] = list[left], list[i]
				i = left
			}
		}
	}
}

func Heap() []int {
	list := generate()
	fmt.Printf("堆排%v\n", list)
	bigRoot(list)
	length := len(list)
	for i := length - 1; i > 0; i-- {
		list[0], list[i] = list[i], list[0]
		reBigRoot(list,0, i-1)
	}
	fmt.Printf("堆排%v\n", list)
	return list
}
