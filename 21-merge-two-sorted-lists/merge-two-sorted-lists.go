/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
     dummy := &ListNode{}      // sentinel node
    tail := dummy

    p, q := list1, list2
    for p != nil && q != nil {
        if p.Val <= q.Val {
            tail.Next = p
            p = p.Next
        } else {
            tail.Next = q
            q = q.Next
        }
        tail = tail.Next
    }

    // attach sisa elemen (jika ada)
    if p != nil {
        tail.Next = p
    } else {
        tail.Next = q
    }

    return dummy.Next
}