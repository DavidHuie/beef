package server

import "github.com/DavidHuie/beef/Godeps/_workspace/src/github.com/boltdb/bolt"

type Server struct {
	db *bolt.DB
}
