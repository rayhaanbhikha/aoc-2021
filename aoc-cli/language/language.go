package language

import (
	"errors"
	"fmt"
	"strings"
)

type LanguageType int

const (
	GOLANG LanguageType = iota
	NODE
)

type Language struct {
	languageType LanguageType
}

func NewLanguage(languageType string) (*Language, error) {
	switch strings.ToLower(languageType) {
	case "node":
		return &Language{NODE}, nil
	case "go":
		return &Language{GOLANG}, nil
	}

	return nil, fmt.Errorf("%s language does not have an implementation", languageType)
}

func (l *Language) Create(basePath string, inputData []byte) error {
	switch l.languageType {
	case NODE:
		return NewNodeLanguage(basePath, inputData)
	case GOLANG:
		return NewGoLanguage(basePath, inputData)
	}
	return errors.New("language doesn't exist")
}

func (l *Language) String() string {
	switch l.languageType {
	case NODE:
		return "node"
	case GOLANG:
		return "go"
	}
	return "unknown"
}
