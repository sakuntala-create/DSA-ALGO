/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
   dummy := &ListNode{Next: head}
   curr := head
   prev := dummy

   for curr!=nil{
    if curr.Next !=nil && curr.Val == curr.Next.Val{
        dupVal:= curr.Val

    for curr!=nil && curr.Val == dupVal{
        curr= curr.Next
    }
    prev.Next = curr

   }else{
    prev = curr
    curr=curr.Next
   }
   }

  return dummy.Next

}
