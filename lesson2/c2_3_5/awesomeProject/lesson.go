package main

func getOsName() string {
	return "mac"
}

func main() {
	os := getOsName()
	switch os {
	case "mac":
		println("Mac!!")
	case "windows":
		println("Windows!!")
	default:
		println("Default!!")
	}
}
