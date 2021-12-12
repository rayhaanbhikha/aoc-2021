package session

import (
	"errors"
	"fmt"
	"os"
	"path"
)

const DEFAULT_FILENAME = ".aoc-session"

type Session struct {
	value    string
	filePath string
}

func (s *Session) serialise() error {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("need to set session token by running: aoc session -t <token>")
		}
		return err
	}

	token := string(data)
	if token == "" {
		return fmt.Errorf("no session token. %s is empty", s.filePath)
	}

	s.value = token

	return nil
}

func (s *Session) deserialise() error {
	return os.WriteFile(s.filePath, []byte(s.value), os.FileMode(0600))
}

func (s *Session) Init() error {
	return s.serialise()
}

func (s *Session) SetToken(newSessionToken string) error {
	s.value = newSessionToken
	return s.deserialise()
}

func (s *Session) GetToken() string {
	return s.value
}

func New() (*Session, error) {
	currentDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	session := &Session{
		filePath: path.Join(currentDir, DEFAULT_FILENAME),
	}

	return session, nil
}
