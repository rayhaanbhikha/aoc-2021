package language

type Language int

func (l Language) String() string {
	switch l {
	case NODE:
		return "node"
	case GOLANG:
		return "go"
	}
	return "unknown"
}

const (
	GOLANG Language = iota
	NODE
)
