# Problem Statement - Bakery Order Processing

## Problem: Bakery Order Processing

Create a program that simulates a bakery processing orders. Each order represents a batch of cookies. Multiple bakers can receive and fulfill these orders concurrently. The system will continue to accept orders until a total of 15 batches have been fulfilled.

### Requirements:

- **Goroutines**: Implement multiple bakers as separate goroutines.
- **Channel for Orders**: Use a channel to represent incoming orders.
- **Shared Count**: Keep track of the number of fulfilled orders using a shared variable.
- **Output**: Print messages indicating each order received and fulfilled, along with the current count of fulfilled orders.

### Example Output:
##### Baker 1: Received order for batch 1! Current Fulfilled Orders: 1
##### Baker 2: Received order for batch 2! Current Fulfilled Orders: 2
##### Baker 3: Received order for batch 3! Current Fulfilled Orders: 3
###### ...
##### Baker 1: Fulfilled order for batch 1! Total Fulfilled Orders: 15

