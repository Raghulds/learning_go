package gofundamentals

import "fmt"

func Panic() {
	fmt.Println(division(7, 3))
	fmt.Println(division(7, 0))
}

func division(a, b int) (n int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("Err: %v", e)
			n = 0
		}
	}()

	return a / b, nil
}

/*
panic -
Manual panic trigger by using panic() function
recover in defer funcfor catching panics

named return values -
returning values from recover catch
documentation
*/
