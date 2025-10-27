# Author: Temur Rustamov
### Lab 2: Mutex and Channel-Based Barrier
#### Notes:
> Here is a demonstration of reusable barrier using  atomic counters, mutex locks and an unbuffered channel. Each goroutine completes Part A then synchronizes at the barrier by incrementing a shared atomic counter. The last arriving goroutine signals others via the channel so all routines can proceed together to Part B. The barrier becomes reusable if we reset channel, counter so the group of goroutines to wait for each other multiple times before continuing execution.