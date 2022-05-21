package validation

func IsValidClient(username string, email string, age uint) bool {
	if (len(username) >= 3 && IsValidEmail(email) && age >= 18) {
		return true
	}
	return false
}
