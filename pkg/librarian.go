package librarian

import (
	"context"
	"database/sql"
	"gopkg.in/yaml.v2"
	"library/proto/pb"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type serverConfig struct {
	Port int `yaml:"port"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

type LibrarianConfig struct {
	Server serverConfig `yaml:"server"`
	DB     Database     `yaml:"db"`
}

type Server struct {
	pb.UnimplementedLibrarianServer
	config                      LibrarianConfig
	db                          *sql.DB
	getBooksStmt, getAuthorStmt *sql.Stmt
}

func (s *Server) Construct(config string) error {
	err := s.parseConfig(config)
	if err != nil {
		return err
	}
	return s.connectToDB()
}

func (s *Server) GetPort() int {
	return s.config.Server.Port
}

func (s *Server) GetBooks(author *pb.Author, stream pb.Librarian_GetBooksServer) error {
	rows, err := s.getBooksStmt.Query(author.FirstName, author.LastName)
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

func (s *Server) GetAuthor(ctx context.Context, book *pb.Book) (author *pb.Author, err error) {
	rows, err := s.getAuthorStmt.Query(book.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		author = &pb.Author{}
		err = rows.Scan(&author.FirstName, &author.LastName)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	return
}

func (s *Server) connectToDB() (err error) {
	dbUrl := s.config.DB.User + ":" + s.config.DB.Password +
		"@tcp(" + s.config.DB.Url + ":" + s.config.DB.Port + ")/" + s.config.DB.Name
	println(dbUrl)
	s.db, err = sql.Open("mysql", dbUrl)
	if err != nil {
		return err
	}
	err = s.db.Ping()
	s.getBooksStmt, err =
		s.db.Prepare("SELECT name, text, encoding FROM Library.Books WHERE AuthorId in (SELECT id FROM Library.Authors where FirstName = ? AND LastName = ?)")
	if err != nil {
		return err
	}
	s.getAuthorStmt, err =
		s.db.Prepare("SELECT FirstName, LastName FROM Library.Authors WHERE Id in (SELECT AuthorId FROM Library.Books where Name = ?)")
	if err != nil {
		return err
	}
	return
}

func (s *Server) parseConfig(c string) error {
	data, err := os.ReadFile(c)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &s.config)
	return err
}

func (s *Server) StopServer() error {
	s.getBooksStmt.Close()
	s.getAuthorStmt.Close()
	return s.db.Close()
}
