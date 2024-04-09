package main

func main() {

	// if-else
	ten := 10
	if ten == 20 {
		println("This should'nt be printed because ten!=20")
	} else {
		println("Ten is not equals to 20")
	}

	if "a" == "b" || 10 == 10 || true == false {
		println("10 is equals to 10")
	} else if 11 == 11 && "go" == "go" {
		println("This will not be printed because first condition was satisfied")
	} else {
		println("If no condition will be satisfied then this line will be printed")
	}
}
