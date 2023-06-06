package test

import (
	"context"
	"library/pkg"
	"library/proto/pb"
	"log"
	"testing"
)

type streamBooksMock struct {
	pb.Librarian_GetBooksServer
	books *[]*pb.Book
}

func (s streamBooksMock) Send(book *pb.Book) error {
	*s.books = append(*s.books, book)
	return nil
}

func TestGetBooks(t *testing.T) {
	server := &librarian.Server{}
	err := server.Construct("./config/testConfig.yaml")
	if err != nil {
		log.Fatal(err)
	}
	author := &pb.Author{FirstName: "Lev", LastName: "Tolstoy"}
	streamBooks := make([]*pb.Book, 0)
	stream := streamBooksMock{books: &streamBooks}
	err = server.GetBooks(author, &stream)
	if err != nil {
		log.Fatal(err)
	}
	expected := map[string]bool{}
	expected["War and peace"] = true
	expected["Sunday"] = true
	expected["Childhood"] = true
	if len(*stream.books) != 3 {
		log.Fatal("Wrong number of books")
	}
	for _, value := range *stream.books {
		if _, ok := expected[value.Name]; !ok {
			log.Fatal("Wrong book")
		}
		delete(expected, value.Name)
	}
	streamBooks = make([]*pb.Book, 0)

	author = &pb.Author{FirstName: "Alexey", LastName: "Tolstoy"}
	server.GetBooks(author, stream)
	if len(*stream.books) > 0 {
		log.Fatal("Wrong number of books")
	}

	author = &pb.Author{FirstName: "Robert", LastName: "Dauni"}
	server.GetBooks(author, stream)
	if len(*stream.books) > 0 {
		log.Fatal("Wrong number of books")
	}
}

func TestGetAuthor(t *testing.T) {
	librarian := &librarian.Server{}
	err := librarian.Construct("./config/testConfig.yaml")
	if err != nil {
		log.Fatal(err)
	}
	book := &pb.Book{Name: "War and Peace"}
	author, err := librarian.GetAuthor(context.Background(), book)
	if err != nil {
		log.Fatal(err)
	}
	expected := pb.Author{FirstName: "Lev", LastName: "Tolstoy"}
	if author.GetFirstName() != expected.GetFirstName() || author.GetLastName() != author.GetLastName() {
		log.Fatal("Wrong author")
	}
	book = &pb.Book{Name: "Zolushka"}
	author, err = librarian.GetAuthor(context.Background(), book)
	if err != nil {
		log.Fatal(err)
	}
	if author != nil {
		log.Fatal("No such book")
	}
}
