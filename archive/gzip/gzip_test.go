package gzip_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/heaven-chp/common-library-go/archive/gzip"
	"github.com/heaven-chp/common-library-go/file"
)

func TestCompress(t *testing.T) {
	name := uuid.New().String() + string(filepath.Separator) + uuid.New().String() + ".gz"
	defer os.RemoveAll(filepath.Dir(name))

	input := uuid.New().String() + string(filepath.Separator)
	if err := os.Mkdir(input, os.ModePerm); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(input)

	output := uuid.New().String() + string(filepath.Separator)
	defer os.RemoveAll(output)

	path := input + uuid.New().String() + ".txt"
	data := []string{"aaa"}
	flag := int(os.O_WRONLY | os.O_APPEND | os.O_CREATE)
	if err := file.Write(path, data, flag, 0600); err != nil {
		t.Fatal(err)
	}

	if err := gzip.Compress(name, path); err != nil {
		t.Fatal(err)
	}

	fileName := uuid.New().String() + ".txt"
	if err := gzip.Decompress(name, fileName, output); err != nil {
		t.Fatal(err)
	} else if result, err := file.Read(output + fileName); err != nil {
		t.Fatal(err)
	} else if result[0] != data[0] {
		t.Fatal("invalid data - ", result[0], ", ", data[0])
	}
}

func TestDecompress(t *testing.T) {
	TestCompress(t)
}