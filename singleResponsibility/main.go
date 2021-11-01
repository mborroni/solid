package document

func main() {
	doc := NewDocument("test")

	fileStorage := NewFileStorage("./desktop/")
	if err := fileStorage.Save(doc); err != nil {
		// handle error gracefully
	}
}