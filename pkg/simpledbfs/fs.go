package simpledbfs

import (
	"github.com/pkg/errors"
	"github.com/smallfz/libnfs-go/fs"
	"github.com/thomasjungblut/embeded-nfs/pkg/fsutil"
	"github.com/thomasjungblut/go-sstables/simpledb"
	"os"
)

type SimpleDBFS struct {
	db *simpledb.DB
}

func (s *SimpleDBFS) Open(name string) (fs.File, error) {
	return s.OpenFile(name, os.O_RDONLY, os.FileMode(0o644))
}

// TODO(thomas): how do we store? JSON for metadata? bytes along-side? simpledb is strings, but it could also be protobuf

func (s *SimpleDBFS) OpenFile(name string, flag int, perm os.FileMode) (fs.File, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) Stat(name string) (fs.FileInfo, error) {
	_, err := s.db.Get(name)
	if err != nil {
		if errors.Is(err, simpledb.ErrNotFound) {
			return nil, os.ErrNotExist
		}
		return nil, err
	}

	return &fsutil.FileInfo{}, nil
}

func (s *SimpleDBFS) Chmod(name string, perm os.FileMode) error {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) Chown(name string, uid, gid int) error {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) Symlink(oldName, newName string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) Readlink(name string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) Link(oldName, newName string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) Rename(oldName, newName string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) Remove(name string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) MkdirAll(name string, perm os.FileMode) error {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) GetFileId(fi fs.FileInfo) uint64 {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) GetRootHandle() []byte {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) GetHandle(info fs.FileInfo) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) ResolveHandle(fh []byte) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleDBFS) Attributes() *fs.Attributes {
	//TODO implement me
	panic("implement me")
}

func NewSimpleDBFS(basePath string) (fs.FS, error) {
	db, err := simpledb.NewSimpleDB(basePath, simpledb.MemstoreSizeBytes(128*1024*1024))
	if err != nil {
		return nil, err
	}

	if err := db.Open(); err != nil {
		return nil, err
	}

	return &SimpleDBFS{db}, nil
}
