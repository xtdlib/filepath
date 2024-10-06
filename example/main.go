package main

import (
	"io/fs"
	"log"

	"github.com/xtdlib/filepath"
)

func main() {
	filepathx.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		log.Println(path, ", ", d.Type())
		if filepathx.IsSymlink(d.Type()) {
			log.Println("* SYM * ", path)
		} else if filepathx.IsDir(d.Type()) {
			log.Println("* DIR * ", path)
		} else if filepathx.IsFile(d.Type()) {
			log.Println("* FILE * ", path)
		} else {
			panic(path)
		}
		return nil
	})
}
