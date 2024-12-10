# Go Race Condition with Channel and Goroutines

This repository demonstrates a race condition in a Go program that uses channels and goroutines.  The program attempts to send data to a channel after all the goroutines are waiting to receive. This results in unexpected output due to the race condition on the channel read.

## Bug Description

The race condition arises from multiple goroutines simultaneously attempting to receive from the channel (`ch`) after the channel has been closed.  This can lead to unpredictable behavior and even deadlocks. The `wg.Wait()` call will block indefinitely until all goroutines finish.

## Bug Solution

The solution ensures that all goroutines are launched *before* sending the data into the channel. The order of operations is critical in avoiding the race condition.