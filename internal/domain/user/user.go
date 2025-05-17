package user

import "fmt"

type UserID string

func (u UserID) String() string {
	return string(u)
}

func (u UserID) validate() error {
	if len(u) == 0 {
		return fmt.Errorf("user id is empty")
	}
	return nil
}
