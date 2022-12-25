package main

import (
	"fmt"
	"sort"
)

type metricsProvider struct {
	metrics []string
}

type MetricsProvider interface {
	Add(metrics ...string)
	View()
}

func InitMetricsProvider() MetricsProvider {
	return &metricsProvider{} //make([]commonpb.ComposerMetric, 1)
}

func (m *metricsProvider) Add(metrics ...string) {
	m.metrics = append(m.metrics, metrics...)
}

func (m *metricsProvider) View() {
	fmt.Println(m.metrics, "     len = ", len(m.metrics), "     cap = ", cap(m.metrics))
}

func main() {
	test1()
}

func test3() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{11, 22, 33, 44, 55}

	arr3 := append(arr1, arr2...)
	fmt.Println(arr3)
}

func test2() {
	arr := []int{1, 1, 1, 2, 2, 3, 3, 4, 1, 2, 3, 4, 1, 1, 2, 2, 3, 4, 4, 3}
	researchSort(arr)
	fmt.Println(arr)
}

func test1() {
	mp := InitMetricsProvider()

	sl := make([]string, 0, 0)

	mp.Add(sl...)
	mp.View()
}

func researchSlice() {
	a := make([]int, 3)
	fmt.Println(a, len(a), cap(a))

	b := make([]int, 0, 3)
	b = append(b, 1)
	fmt.Println(b, len(b), cap(b))

	fmt.Println("\nResearch sort.Slice()")
	//sl := []int{1, 3, 5, 7, 9, 8, 6, 4, 2, 1, 2, 3, 4, 9, 8, 7, 6, 5}
	sl := make([]int, 0)
	fmt.Println("sl = ", sl)
	sort.Slice(sl, func(i, j int) bool {
		if sl[i] > sl[j] {
			return true
		}
		return false
	})
	fmt.Println("after sort sl = ", sl)

}

func researchSort(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println("arr in func after sort", arr)
}
