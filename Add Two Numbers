/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sum := l1.Val + l2.Val
    l1 = l1.Next
    l2 = l2.Next
    extra := sum / 10
    list1 := ListNode{sum%10, nil}
    address := &list1
    var val1, val2, n int
    for {
        if l1 == nil && l2 == nil && extra == 0{
            break
        }
        val1, val2 = 0, 0
        if l1 != nil {
            val1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            val2 = l2.Val
            l2 = l2.Next
        }
        sum = val1 + val2 + extra
        list2 := ListNode{sum%10, address}
        address = &list2
        extra = sum/10
        n++
    }
    addressPrev := address
    addressNext := address.Next
    address.Next = nil
    for i := 0; i < n; i++ {
        address = addressNext.Next
        addressNext.Next = addressPrev
        addressPrev = addressNext
        addressNext = address
    }
    return addressPrev
}
