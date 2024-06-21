package graphql

import "errors"

func main() {
	myFunc()
}

func myFunc() error {
	return errors.New("something bad happened")
}

func call() error {
	return nil
}
