package mock

type Retriever struct {
	Contens string
}

func (r Retriever) Get(url string) string {
	return r.Contens
}
