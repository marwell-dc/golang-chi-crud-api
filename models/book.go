package models

type Book struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	PublishedDate    string `json:"published_date"`
	OriginalLanguage string `json:"original_language"`
}

func (b *Book) NewBook() *Book {
	return &Book{}
}

var books = []Book{
	{
		ID:               "1",
		Title:            "7 habits of Highly Effective People",
		Author:           "Stephen Covey",
		PublishedDate:    "15/08/1989",
		OriginalLanguage: "English",
	},
}

func (b Book) ListBooks() []Book {
	return books
}

func (b Book) GetBookById(id string) *Book {
	for _, book := range books {
		if book.ID == id {
			return &book
		}
	}

	return nil
}

func (b Book) StoreBook(book Book) {
	books = append(books, book)
}

func (b *Book) Update(id string, bookUpdate Book) *Book {
	for i, book := range books {
		if book.ID == id {
			books[i] = bookUpdate
			return &book
		}
	}

	return nil
}

func (b Book) Delete(id string) *Book {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], (books)[i+1:]...)
			return &book
		}
	}

	return nil
}
