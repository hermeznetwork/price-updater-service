package ports

type Command interface {
	Execute() error
}
