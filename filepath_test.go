package filepath

import "testing"

func Test1(t *testing.T) {
	// equal(t, Stem0("asd.x.tar.gz"), "asd")
	equal(t, Stem("asd.x.tar.gz"), "asd.x.tar")
	equal(t, Stem("ff/asd.x.tar.gz"), "asd.x.tar")
	equal(t, StemN("asd.x.tar.gz", 2), "asd.x")
	equal(t, ExtN("asd.x.tar.gz", 2), ".tar.gz")
	equal(t, ExtN("asd.x.tar.gz", 1), ".gz")
	equal(t, "asd.x.tar.gz", Stem("asd.x.tar.gz")+Ext("asd.x.tar.gz"))
	// equal(t, Stem1("asd.x.tar.gz"), "asd.x.tar")
	// equal(t, Stem2("asd.x.tar.gz"), "asd.x")
	// equal(t, Ext0("asd.x.tar.gz"), ".x.tar.gz")
	// equal(t, Ext2("asd.x.tar.gz"), ".tar.gz")
}

func equal(t *testing.T, a, b string) {
	t.Log(a, b)
	if a != b {
		t.Fatal(a, "!=", b)
	}
}
