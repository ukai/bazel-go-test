package sample

import "io/ioutil"

func Load(fname string) ([]byte, error) {
	return ioutil.ReadFile(fname)
}
