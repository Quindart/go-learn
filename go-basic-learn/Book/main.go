package Book

type Book struct {
	name   string
	author string
	pages  int
}

// *Book = pointer to Book => &Book  = 0xancas
// New creates a new instance of Book with the provided name, author, and number of pages.
func New(name string, author string, pages int) *Book {
	return &Book{
		name:   name,
		author: author,
		pages:  pages,
	}
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetPages() int {
	return b.pages
}