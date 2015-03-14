package server

import "github.com/boltdb/bolt"

type Server struct {
	db *bolt.DB
}
