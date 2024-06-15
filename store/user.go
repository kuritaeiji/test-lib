package store

import (
	"errors"

	"github.com/kuritaeiji/test-lib/stack"
)

type UserStore struct {
	users map[string]string
}

func (s *UserStore) Insert(id string, name string) error {
	if _, ok := s.users[id]; ok {
		return stack.NewCallStack(errors.New("すでに同一IDのユーザーが存在します"))
	}
	s.users[id] = name
	return nil
}

func (s *UserStore) Get(id string) (string, bool) {
	name, ok := s.users[id]
	return name, ok
}
