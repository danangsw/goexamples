// https://leetcode.com/problems/merge-sorted-array

package problem

func MergeSortedArray(list1 []int, m int, list2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for j >= 0 {
		if i >= 0 && list1[i] > list2[j] {
			list1[k] = list1[i]
			i--
		} else {
			list1[k] = list2[j]
			j--
		}
		k--
	}
}
