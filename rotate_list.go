/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {

    if head == nil || head.Next==nil || k==0{
        return head
    }
      // Step 1: find length and tail

      length :=1
      tail:= head
     for tail.Next !=nil{
        tail = tail.Next
        length++
     }
     // Step 2: normalize k
    k = k % length
    if k == 0 {
        return head
    }

    // Step 3: make circular
    tail.Next = head
    stepsToNewTail := length - k

    newTail := head
    for i := 1; i < stepsToNewTail; i++ {
        newTail = newTail.Next
    }

    // Step 5: break the circle
    newHead := newTail.Next
    newTail.Next = nil

    return newHead
    
}
