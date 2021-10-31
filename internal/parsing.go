package internal

import "fmt"

func parseExpression(expression string) error {
	s := newScanner(expression)

	for s.hasNext() {

		token, err := s.searchToken()
		if err != nil {
			return err
		}
		fmt.Printf("%v \n", token)

	}
	return nil
}
