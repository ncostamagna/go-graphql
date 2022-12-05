package internal

import "fmt"

type Service struct{}

func (s *Service) Print() {
	fmt.Println("Entra")
}
