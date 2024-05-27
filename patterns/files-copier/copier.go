package files_copier

type copier struct {
	src    string
	status chan int
}

type copierI interface{}

func NewCopier(src string) copierI {

	return &copier{}
}

func (с *copier) CopyIn(dst string) {

}

type src struct {
	path string
}

func New(path string) {

}
