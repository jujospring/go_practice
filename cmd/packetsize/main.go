package main

func getPacketSize(message string) int {
	messageLength := len(message)
	largest := 0

	for size := 1; size < messageLength; size++ {
		if messageLength%size == 0 {
			numPackets := messageLength / size
			if isComposite(numPackets) {
				largest = size
			}
		}
	}
	return largest
}

func isComposite(num int) bool {
	return num > 1 && !isPrime(num)
}

func isPrime(num int) bool {
	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
