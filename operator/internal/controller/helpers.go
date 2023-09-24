package controller

// Function to check if a finalizer is present in the finalizers slice
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// Function to remove a string from a slice
func removeString(slice []string, s string) []string {
	index := -1
	for i, item := range slice {
		if item == s {
			index = i
			break
		}
	}
	if index != -1 {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}
