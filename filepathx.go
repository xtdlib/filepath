package filepathx

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Stem returns name without extension qwer.tar.gz -> qwer.tar
func Stem(filename string) string {
	return strings.TrimSuffix(filename, Ext(filename))
}

// Stem1 is same with Stem
var Stem1 = Stem

// Stem2 removes 2 comma extension qwer.tar.gz -> qwer
func Stem2(filename string) string {
	return StemN(filename, 2)
}

// StemN removes N comma extension
func StemN(filename string, n int) string {
	stem := filename
	i := 0
	for {
		stem = Stem(stem)
		if stem == Stem(stem) {
			return stem
		}
		i++
		if i >= n {
			return stem
		}
	}
}

// Stem0 removes max comma extension qwer.tar.gz -> qwer
func Stem0(filename string) string {
	return StemN(filename, 99)
}

var Ext1 = Ext

func Ext2(path string) string {
	return strings.TrimPrefix(path, Stem2(path))
}

func ExtN(path string, n int) string {
	return strings.TrimPrefix(path, StemN(path, n))
}

func Ext0(path string) string {
	return strings.TrimPrefix(path, Stem0(path))
}

func IsSymlink(mode fs.FileMode) bool {
	return (mode & os.ModeSymlink) != 0
}

func IsDir(mode fs.FileMode) bool {
	return (mode & os.ModeDir) != 0
}

func IsFile(mode fs.FileMode) bool {
	return (mode&os.ModeDir) == 0 &&
		(mode&os.ModeSymlink) == 0 &&
		(mode&os.ModeDevice) == 0 &&
		(mode&os.ModeNamedPipe) == 0 &&
		(mode&os.ModeSocket) == 0 &&
		(mode&os.ModeCharDevice) == 0
}

var Ext = filepath.Ext
var Abs = filepath.Abs
var Base = filepath.Base
var Clean = filepath.Clean
var Dir = filepath.Dir
var EvalSymlinks = filepath.EvalSymlinks
var FromSlash = filepath.FromSlash
var Glob = filepath.Glob
var IsAbs = filepath.IsAbs
var IsLocal = filepath.IsLocal
var Join = filepath.Join
var Match = filepath.Match
var Rel = filepath.Rel
var Split = filepath.Split
var SplitList = filepath.SplitList
var ToSlash = filepath.ToSlash
var VolumeName = filepath.VolumeName

// var Walk = filepath.Walk # slow, deprecate
var WalkDir = filepath.WalkDir
var SkipDir = filepath.SkipDir
var SkipAll = filepath.SkipAll

type WalkFunc = filepath.WalkFunc
