package sort_df

import "testing"

func TestSort(t *testing.T) {
	if !check(bubble()) {
		t.Error("xxxx")
	}
	if !check(selekt()) {
		t.Error("xxxx")
	}
	if !check(insert()) {
		t.Error("xxxx")
	}
	if !check(quick()) {
		t.Error("xxxx")
	}
	if !check(Heap()) {
		t.Error("xxxx")
	}
}

func check(arr []int) bool {
	length := len(arr)
	for index, a := range arr {
		if index == length-1 {
			return true
		}
		if arr[index+1] < a {
			return false
		}
	}
	return true
}
