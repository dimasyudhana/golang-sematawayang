package main

import "fmt"

// ListNode adalah struktur data untuk merepresentasikan node pada linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists menggabungkan dua linked list menjadi satu linked list yang terurut
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// dummy node untuk merepresentasikan awal linked list yang dihasilkan
	dummy := &ListNode{}
	// current node sebagai pointer ke node terakhir pada linked list yang dihasilkan
	current := dummy

	// looping sampai salah satu dari linked list habis
	for list1 != nil && list2 != nil {
		// pilih node dengan nilai yang lebih kecil untuk dimasukkan ke linked list yang dihasilkan
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		// pindahkan pointer ke node terakhir pada linked list yang dihasilkan
		current = current.Next
	}

	// sisa node pada linked list yang belum habis dihubungkan ke linked list yang dihasilkan
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// kembalikan linked list yang dihasilkan, dimulai dari node setelah dummy node
	return dummy.Next
}

func main() {
	// contoh penggunaan mergeTwoLists
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	mergedList := mergeTwoLists(list1, list2)

	// print nilai node pada linked list yang dihasilkan
	for mergedList != nil {
		fmt.Printf("%d ", mergedList.Val)
		mergedList = mergedList.Next
	}
}
