package theories

// Switch
func canIDrinkSwitch(age int) bool {
	switch { // switch var { .. case 조건: .. } 로 사용해도 됨, 마찬가지로 variable expression; var도 가능
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func basicSwitch() {
	userType := "admin"
	switch userType {
	case "admin":
		// ...
	case "developer":
		// ...
	}
	// default
	// ...
}
