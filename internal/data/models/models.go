package models

import "os"

type Image struct {
	Name string
	File *os.File
}
