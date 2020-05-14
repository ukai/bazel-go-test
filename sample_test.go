package sample

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

func fileCopy(t *testing.T, src, dst string) {
	t.Helper()
	t.Logf("%s -> %s", src, dst)
	err := os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		t.Fatal(err)
	}
	r, err := os.Open(src)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	fi, err := r.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if fi.IsDir() {
		err = os.MkdirAll(dst, 0755)
		if err != nil {
			t.Fatal(err)
		}
		return
	}
	w, err := os.Create(dst)
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(w, r)
	if err != nil {
		w.Close()
		t.Fatal(err)
	}
	err = w.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoad(t *testing.T) {
	dir, err := bazel.NewTmpDir("load")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(cwd)
	err = os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}
	files, err := bazel.ListRunfiles()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", files)
	for _, f := range files {
		fileCopy(t, f.Path, f.ShortPath)
	}
	_, err = Load("testdata/foo")
	if err != nil {
		t.Errorf("Load %v; want nil err", err)
	}
}
