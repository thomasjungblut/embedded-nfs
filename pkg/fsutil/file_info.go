package fsutil

import (
	"os"
	"time"
)

type FileInfo struct {
	id       uint64
	name     string
	perm     os.FileMode
	size     int64
	modTime  time.Time
	aTime    time.Time
	cTime    time.Time
	isDir    bool
	numLinks int
}

func (fi *FileInfo) Name() string {
	return fi.name
}

func (fi *FileInfo) Size() int64 {
	return fi.size
}

func (fi *FileInfo) Mode() os.FileMode {
	return fi.perm
}

func (fi *FileInfo) ModTime() time.Time {
	return fi.modTime
}

func (fi *FileInfo) ATime() time.Time {
	return fi.aTime
}

func (fi *FileInfo) CTime() time.Time {
	return fi.cTime
}

func (fi *FileInfo) IsDir() bool {
	return fi.isDir
}

func (fi *FileInfo) Sys() interface{} {
	return nil
}

func (fi *FileInfo) NumLinks() int {
	return fi.numLinks
}

func NewFileInfo(id uint64, name string, perm os.FileMode, size int64, modTime time.Time, aTime time.Time, cTime time.Time, isDir bool, numLinks int) *FileInfo {
	return &FileInfo{id: id, name: name, perm: perm, size: size, modTime: modTime, aTime: aTime, cTime: cTime, isDir: isDir, numLinks: numLinks}
}
