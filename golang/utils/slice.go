package util

// UniqueItems takes a slice of strings and remove duplicates
func UniqueItems(s []string) []string {
	var uniqueItems []string
	uniqueItemsMap := make(map[string]bool)
	for _, v := range s {
		if !uniqueItemsMap[v] {
			uniqueItemsMap[v] = true
		}
	}
	for k := range uniqueItemsMap {
		uniqueItems = append(uniqueItems, k)
	}
	return uniqueItems
}
