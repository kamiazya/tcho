package module

type Module interface {
	Build() (err error)
	Run(arguments []string) (err error)
	Close()
}
