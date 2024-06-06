package filepathx

import "testing"

func Test1(t *testing.T) {
	equal(t, Stem0("asd.x.tar.gz"), "asd")
	equal(t, Stem("asd.x.tar.gz"), "asd.x.tar")
	equal(t, Stem1("asd.x.tar.gz"), "asd.x.tar")
	equal(t, Stem2("asd.x.tar.gz"), "asd.x")
	equal(t, Ext0("asd.x.tar.gz"), ".x.tar.gz")
	equal(t, Ext2("asd.x.tar.gz"), ".tar.gz")
}

func equal(t *testing.T, a, b string) {
	if a != b {
		t.Fatal(a, "!=", b)
	}
}
