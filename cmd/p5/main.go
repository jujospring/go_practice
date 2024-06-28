package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseBill := calculateBaseBill(costPerMessage, numMessages)
	discount := calculateDiscount(numMessages)
	finalBill := baseBill - (baseBill * discount)
	return finalBill
}

func calculateDiscount(messagesSent int) float64 {
	var discount float64

	if messagesSent > 1000 && messagesSent <= 5000 {
		discount = 0.1
	} else if messagesSent > 5000 {
		discount = 0.2
	}
	return discount
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
