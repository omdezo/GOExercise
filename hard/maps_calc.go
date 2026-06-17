package hard

// Maps helpers: pure functions that operate on maps data structures.

// CalcAddOrUpdateInventory sets quantity for a product.
func CalcAddOrUpdateInventory(inv map[string]int, product string, qty int) {
	inv[product] = qty
}

// CalcSearchInventory returns quantity and whether the product exists.
func CalcSearchInventory(inv map[string]int, product string) (int, bool) {
	q, ok := inv[product]
	return q, ok
}

// CalcAddOrUpdatePhonebook sets phone number for a contact.
func CalcAddOrUpdatePhonebook(pb map[string]string, name, number string) {
	pb[name] = number
}

// CalcSearchPhonebook returns number and whether the contact exists.
func CalcSearchPhonebook(pb map[string]string, name string) (string, bool) {
	n, ok := pb[name]
	return n, ok
}

