## Discoveries in GO

I chose to not use an array in my Stack and Queue structure but rather followed from the [godash](https://godoc.org/github.com/Kairi/godash). I am trying to avoid using as many non Standard Library packages as possible, so I opted in writing my own implementations of Pop, Push, etc but couldn't help but be inspired from godash. I have later found that my approach follows the [Linked List](https://en.wikipedia.org/wiki/Linked_list) pattern which I now look forward to going through.

I did find this pattern hard to reason about when approaching the QueueStack and StackQueue implementations, and nearly reverted to using Arrays, but I persisted and with a bit of help from this [blog post](http://www.geeksforgeeks.org/queue-using-stacks/), I was able to figure it out.
