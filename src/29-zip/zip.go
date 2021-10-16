package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	o, _ := os.Create("data.zip")
	defer o.Close()

	w := zip.NewWriter(o)
	defer w.Close()

	cwd, _ := os.Getwd()
	loc := filepath.Join(cwd, "data")

	filepath.WalkDir(loc, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		relpath, _ := filepath.Rel(loc, path)
		fmt.Println(relpath)

		of, _ := os.Open(path)
		f, _ := w.Create(relpath)
		if _, e := io.Copy(f, of); e != nil {
			return e
		}

		return nil
	})
}
