package document

type Document struct {
	Name string
}

func NewDocument(name string) *Document {
	return &Document{
		Name: name,
	}
}

func (d *Document) Read(p []byte) (n int, err error) {
	// read document
	return 0, nil
}