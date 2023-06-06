package librarianServer

import (
	"database/sql"
	"library/proto/pb"
)

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

type Server struct {
	pb.UnimplementedLibrarianServer
	db *sql.DB
}

func (s *Server) GetBooks(author *pb.Author, stream pb.Librarian_GetBooksServer) error {
	stmt, err := s.db.Prepare("SELECT name, text, encoding FROM Library.Books WHERE AuthorId in (SELECT id FROM Library.Authors where FirstName = ? AND LastName = ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(author.FirstName, author.LastName)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		book := &pb.Book{}
		err = rows.Scan(&book.Name, &book.Text, &book.Encoding)
		if err != nil {
			return err
		}
		err = stream.Send(book)
		if err != nil {
			return err
		}
	}
	return rows.Err()
}

func (s *Server) GetAuthor(book *pb.Book, stream pb.Librarian_GetAuthorServer) error {
	stmt, err := s.db.Prepare("SELECT FirstName, LastName FROM Library.Authors WHERE Id in (SELECT AuthorId FROM Library.Books where Name = ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(book.Name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		author := &pb.Author{}
		err = rows.Scan(&author.FirstName, &author.LastName)
		if err != nil {
			return err
		}
		err = stream.Send(author)
		if err != nil {
			return err
		}
	}
	return rows.Err()
}

func (s *Server) ConnectToDB(config Database) (err error) {
	dbUrl := config.User + ":" + config.Password +
		"@tcp(" + config.Url + ":" + config.Port + ")/" + config.Name
	println(dbUrl)
	s.db, err = sql.Open("mysql", dbUrl)
	if err != nil {
		return err
	}
	err = s.db.Ping()
	return
}

func (s *Server) DisconnectDB() error {
	return s.db.Close()
}
