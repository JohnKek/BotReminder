package files

import (
	"TransactionalOutbox/internal/storage"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{
		basePath: basePath,
	}
}

const defaultPerm = 0774

func (s Storage) Save(page *storage.Page) error {
	filePath := filepath.Join(s.basePath, page.UserName)
	if err := os.Mkdir(filePath, defaultPerm); err != nil {
		return err
	}
	fName, err := fileName(page)
	if err != nil {
		return err
	}

	filePath = filepath.Join(filePath, fName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}
	return nil
}

func (s Storage) PickRandom(userName string) (*storage.Page, error) {
	path := filepath.Join(s.basePath, userName)
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no pages")
	}
	n := rand.Intn(len(files))
	file := files[n]
	return s.decodePage(filepath.Join(path, file.Name()))

}

func (s Storage) Remove(p *storage.Page) error {
	fileName, err := fileName(p)
	if err != nil {
		return err
	}
	err = os.Remove(filepath.Join(s.basePath, p.UserName, fileName))
	if err != nil {
		return err
	}
	return nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}

func (s Storage) decodePage(fileStorage string) (*storage.Page, error) {
	f, err := os.Open(fileStorage)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	var p storage.Page
	if err = gob.NewEncoder(f).Encode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (s Storage) IsExists(p *storage.Page) (bool, error) {
	fileName, err := fileName(p)
	if err != nil {
		return false, err
	}
	path := filepath.Join(s.basePath, p.UserName, fileName)
	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		return false, err
	}
	return true, nil
}
