package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	filePath string
	dbConn   *sql.DB
}

func New(options ...func(*Server)) *Server {
	svr := &Server{}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func (s *Server) Start() error {
	var err error
	s.dbConn, err = sql.Open("sqlite3", s.filePath)
	if err != nil {
		return err
	}
	return err
}

func WithFile(filePath string) func(*Server) {
	return func(s *Server) {
		s.filePath = filePath
	}
}
