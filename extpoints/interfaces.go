package extpoints

type Hoge interface {
	Before(args []string) error
	Run(args []string) error
	After(args []string) error
	String() string
}
