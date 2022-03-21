// Code generated by vfsgen; DO NOT EDIT.

package authz

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// defaults statically implements the virtual filesystem provided to vfsgen.
var defaults = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2021, 11, 29, 20, 16, 39, 227583853, time.UTC),
		},
		"/cluster_role_cluster_read.yaml": &vfsgen۰CompressedFileInfo{
			name:             "cluster_role_cluster_read.yaml",
			modTime:          time.Date(2020, 7, 22, 2, 57, 32, 903625078, time.UTC),
			uncompressedSize: 264,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x74\xcb\xb1\x4a\x04\x31\x10\xc6\xf1\x3e\x4f\xf1\x91\x52\xcc\x8a\x9d\xa4\xb5\xb0\xb1\x5a\xd0\x46\xb6\x98\xcd\x8e\x5e\xb8\x5c\x26\xcc\x24\x27\xe7\xd3\xcb\x9d\x0a\x36\x57\x0d\xf3\xff\xf8\x51\xcb\xaf\xac\x96\xa5\x46\xe8\x4a\x69\xa2\xd1\x77\xa2\xf9\x8b\x7a\x96\x3a\xed\x1f\x6c\xca\x72\x77\xbc\x77\xfb\x5c\xb7\x88\xc7\x32\xac\xb3\xce\x52\xd8\x1d\xb8\xd3\x46\x9d\xa2\x03\x2a\x1d\x38\x22\xfd\xac\xc1\x92\x34\x0e\xca\xb4\x85\xbf\xa4\x67\x01\x14\x5a\xb9\xd8\x59\x00\x4a\xef\x74\x0a\xca\x85\x4e\x11\xbe\xeb\x60\xef\x74\x14\xb6\xe8\x02\xa8\xe5\x27\x95\xd1\x2c\xe2\xcd\xdf\xf8\xc5\x01\xca\x26\x43\x13\xff\x4b\x47\xd6\xf5\xf2\x7e\x70\xf7\xb7\xf0\x25\xdb\xe5\x7e\x52\x4f\x3b\xbf\xb8\x80\x2a\x75\xfe\x75\x2f\xf3\xf3\x35\xba\xb8\xef\x00\x00\x00\xff\xff\xcd\x64\x2e\x25\x08\x01\x00\x00"),
		},
		"/cluster_role_cluster_write.yaml": &vfsgen۰CompressedFileInfo{
			name:             "cluster_role_cluster_write.yaml",
			modTime:          time.Date(2020, 7, 22, 2, 57, 32, 905274906, time.UTC),
			uncompressedSize: 244,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x74\x8b\xbd\x4e\xc5\x30\x0c\x46\xf7\x3c\x85\x95\x11\x29\x45\x6c\x28\x2b\x03\x0b\x53\x25\x58\x10\x83\x9b\x1a\x61\x35\x8d\x23\x3b\x29\x2a\x4f\x8f\xaa\xfb\xa3\xbb\xdc\xf1\x3b\xdf\x39\x58\xf9\x83\xd4\x58\x4a\x04\x9d\x30\x0d\xd8\xdb\x8f\x28\xff\x61\x63\x29\xc3\xf2\x6c\x03\xcb\xe3\xf6\xe4\x16\x2e\x73\x84\x97\xdc\xad\x91\x8e\x92\xc9\xad\xd4\x70\xc6\x86\xd1\x01\x14\x5c\x29\x42\x3a\xbd\xc1\x92\x54\x0a\xbf\xca\x8d\xc2\x85\xe9\x91\x00\x64\x9c\x28\xdb\x91\x00\x28\x7e\xe3\x1e\x94\x32\xee\x11\x7c\xd3\x4e\xde\x69\xcf\x64\xd1\x05\xc0\xca\xaf\x2a\xbd\x5a\x84\x4f\xff\xe0\xbf\x1c\x80\x92\x49\xd7\x44\x37\x68\x23\x9d\xae\x33\x40\x91\x32\x9e\xa5\xf7\xf1\xed\x9e\xf7\x1f\x00\x00\xff\xff\x95\xf6\xc7\xe6\xf4\x00\x00\x00"),
		},
		"/cluster_role_full_access.yaml": &vfsgen۰CompressedFileInfo{
			name:             "cluster_role_full_access.yaml",
			modTime:          time.Date(2020, 7, 22, 2, 57, 32, 905556892, time.UTC),
			uncompressedSize: 236,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x74\xca\xb1\x6a\xc3\x30\x10\x80\xe1\x5d\x4f\x71\x68\x2c\xc8\xa5\x5b\xd1\xda\xa1\x4b\x27\x43\xbb\x94\x0c\x67\xf9\x4c\x84\xcf\x3a\x73\x27\x19\x9c\xa7\x0f\x26\x21\x64\xc9\xf8\xff\x7c\xb8\xe6\x3f\x52\xcb\x52\x22\xe8\x80\xa9\xc3\x56\xcf\xa2\xf9\x82\x35\x4b\xe9\xe6\x4f\xeb\xb2\xbc\x6f\x1f\x6e\xce\x65\x8c\xf0\xc5\xcd\x2a\x69\x2f\x4c\x6e\xa1\x8a\x23\x56\x8c\x0e\xa0\xe0\x42\x11\xa6\xc6\x1c\x30\x25\x32\x0b\xe9\x26\x83\x1e\x14\x80\x71\x20\xb6\x83\x02\x28\x4e\xb8\x07\x25\xc6\x3d\x82\xaf\xda\xc8\x3b\x6d\x4c\x16\x5d\x00\x5c\xf3\xb7\x4a\x5b\x2d\xc2\xbf\x7f\xf3\x27\x07\xa0\x64\xd2\x34\xd1\xd3\xda\x48\x87\x47\x06\x28\x52\xfa\x3b\xfa\xed\x7f\x5e\xb9\x6b\x00\x00\x00\xff\xff\xf5\x6f\xf5\x6c\xec\x00\x00\x00"),
		},
		"/cluster_role_namespace_read.yaml": &vfsgen۰CompressedFileInfo{
			name:             "cluster_role_namespace_read.yaml",
			modTime:          time.Date(2020, 7, 22, 2, 57, 32, 905819499, time.UTC),
			uncompressedSize: 257,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x74\xcb\xb1\x4e\xc4\x30\x0c\xc6\xf1\x3d\x4f\x61\x65\x44\xa4\x88\x0d\x65\x65\x60\x61\xaa\x04\x0b\xea\xf0\x35\x35\x5c\x74\xb9\xa4\xb2\x9d\x43\xe5\xe9\x51\x0b\x03\x0b\x93\xe5\x4f\xbf\x3f\xd6\xfc\xca\xa2\xb9\xd5\x48\x32\x23\x0d\xe8\x76\x6a\x92\xbf\x60\xb9\xd5\xe1\xfc\xa0\x43\x6e\x77\xd7\x7b\x77\xce\x75\x89\xf4\x58\xba\x1a\xcb\xd8\x0a\xbb\x0b\x1b\x16\x18\xa2\x23\xaa\xb8\x70\x24\x61\x2c\x01\x29\xb1\x6a\x48\x3f\x32\xc8\x4e\x89\x0a\x66\x2e\xba\x53\x22\xc1\x3b\xb6\x20\x5c\xb0\x45\xf2\x26\x9d\xbd\x93\x5e\x58\xa3\x0b\x84\x35\x3f\x49\xeb\xab\x46\x7a\xf3\x37\x7e\x72\x44\xc2\xda\xba\x24\xfe\x33\x5d\x59\xe6\xe3\xfd\x60\xf3\xb7\xe4\x4b\xd6\xe3\x7e\xc2\xd2\xc9\x4f\x2e\x50\x6d\x75\xfc\xed\x5e\xc6\xe7\xff\xd2\xc9\x7d\x07\x00\x00\xff\xff\x36\x06\xdc\xb3\x01\x01\x00\x00"),
		},
		"/cluster_role_namespace_write.yaml": &vfsgen۰CompressedFileInfo{
			name:             "cluster_role_namespace_write.yaml",
			modTime:          time.Date(2020, 7, 22, 2, 57, 32, 906055956, time.UTC),
			uncompressedSize: 300,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x8c\x31\x4f\x03\x31\x0c\x46\xf7\xfc\x0a\x2b\x23\x22\x45\x6c\x28\x2b\x03\x0b\xd3\x49\xb0\xa0\x1b\x7c\xb9\x0f\x35\x6a\x9a\x44\xb6\x53\x54\x7e\x3d\xba\xb6\x03\x48\x6c\xcf\xf6\xf3\xe3\x9e\xdf\x21\x9a\x5b\x8d\x24\x0b\xa7\x1d\x0f\xdb\x37\xc9\xdf\x6c\xb9\xd5\xdd\xe1\x49\x77\xb9\x3d\x9c\x1e\xdd\x21\xd7\x35\xd2\x73\x19\x6a\x90\xa9\x15\xb8\x23\x8c\x57\x36\x8e\x8e\xa8\xf2\x11\x91\xbe\x24\x1b\x02\xa7\x04\xd5\x90\xae\x6a\x90\xcd\x25\x2a\xbc\xa0\xe8\xe6\x12\x09\x7f\xf2\x39\x08\x0a\x9f\x23\x79\x93\x01\xef\x64\x14\x68\x74\x81\xb8\xe7\x17\x69\xa3\x6b\xa4\x0f\x7f\xe7\x67\x47\x24\xd0\x36\x24\xe1\xd7\xea\x04\x59\x2e\x63\x12\xb0\xc1\xdf\x93\x5f\x51\x70\xa5\xce\x96\xf6\x1b\x8c\xbe\x6e\xc7\xd9\x05\xaa\xad\x4e\xb7\xcc\xdb\xf4\xfa\x5f\xa9\x37\xb5\xcb\xf7\xb0\x3f\x91\x5b\x77\x76\x3f\x01\x00\x00\xff\xff\xab\xb5\xff\xa9\x2c\x01\x00\x00"),
		},
		"/namespace.yaml": &vfsgen۰FileInfo{
			name:    "namespace.yaml",
			modTime: time.Date(2021, 11, 29, 20, 16, 39, 227289738, time.UTC),
			content: []byte("\x61\x70\x69\x56\x65\x72\x73\x69\x6f\x6e\x3a\x20\x76\x31\x0a\x6b\x69\x6e\x64\x3a\x20\x4e\x61\x6d\x65\x73\x70\x61\x63\x65\x0a\x6d\x65\x74\x61\x64\x61\x74\x61\x3a\x0a\x20\x20\x6e\x61\x6d\x65\x3a\x20\x6e\x61\x6d\x65\x73\x70\x61\x63\x65\x2d\x6e\x61\x6d\x65\x0a\x20\x20\x6c\x61\x62\x65\x6c\x73\x3a\x0a\x20\x20\x20\x20\x72\x61\x66\x61\x79\x2d\x72\x65\x6c\x61\x79\x3a\x20\x22\x74\x72\x75\x65\x22\x0a"),
		},
		"/relay_default_cluster_role.yaml": &vfsgen۰CompressedFileInfo{
			name:             "relay_default_cluster_role.yaml",
			modTime:          time.Date(2020, 7, 22, 2, 57, 32, 912451614, time.UTC),
			uncompressedSize: 214,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x7c\xcb\x31\xcb\xc2\x40\x0c\x80\xe1\xfd\x7e\x45\xb6\xc2\x07\xd7\x0f\x37\xb9\xd5\xc1\xc5\xa9\xa0\x7b\xda\x46\x0c\xbd\x5e\x4a\x2e\x29\xe8\xaf\x97\xa2\x83\x2e\xae\x2f\xcf\x8b\x0b\x5f\x48\x2b\x4b\x49\xa0\x3d\x0e\x2d\xba\xdd\x44\xf9\x81\xc6\x52\xda\x69\x5f\x5b\x96\xff\x75\x17\x26\x2e\x63\x82\x43\xf6\x6a\xa4\x9d\x64\x0a\x33\x19\x8e\x68\x98\x02\x40\xc1\x99\x12\x28\x65\xbc\xc7\x91\xae\xe8\xd9\xe2\xf0\xb2\x51\x37\xac\x9e\xa9\xa6\x10\x01\x17\x3e\xaa\xf8\x52\xb7\x2d\x42\xf3\xd7\x04\x00\xa5\x2a\xae\x03\x7d\xc5\x95\xb4\xff\x08\x11\x8a\x94\xee\x0d\xcf\xdd\xe9\x97\x7d\x06\x00\x00\xff\xff\xfb\x9f\x23\x09\xd6\x00\x00\x00"),
		},
		"/relay_default_role.yaml": &vfsgen۰CompressedFileInfo{
			name:             "relay_default_role.yaml",
			modTime:          time.Date(2020, 7, 22, 2, 57, 32, 912707605, time.UTC),
			uncompressedSize: 154,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x54\xca\xb1\x0a\xc2\x30\x10\x87\xf1\x3d\x4f\x71\x5b\x41\x68\xc4\x4d\xb2\x3a\xb8\x3b\xb8\x5f\xdb\x03\x8f\xa6\xb9\xf0\xbf\xa4\xa0\x4f\x2f\xdd\x74\xfd\x7d\x1f\x57\x7d\x0a\x5c\xad\x24\xc2\xc4\x73\xe4\xde\x5e\x06\xfd\x70\x53\x2b\x71\xbd\x7a\x54\x3b\xef\x97\xb0\x6a\x59\x12\xdd\x72\xf7\x26\x78\x58\x96\xb0\x49\xe3\x85\x1b\xa7\x40\x54\x78\x93\x44\x90\xcc\xef\x11\x47\x44\xcf\xe2\x29\x8c\xc4\x55\xef\xb0\x5e\xfd\xd8\x46\x1a\x4e\x43\x20\x82\xb8\x75\xcc\xf2\x87\xbb\x60\xfa\x81\x6f\x00\x00\x00\xff\xff\x56\x5a\x2c\x8e\x9a\x00\x00\x00"),
		},
		"/role_read_access.yaml": &vfsgen۰CompressedFileInfo{
			name:             "role_read_access.yaml",
			modTime:          time.Date(2020, 7, 22, 2, 57, 32, 912977857, time.UTC),
			uncompressedSize: 200,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\xca\x31\x4e\xc5\x30\x10\x84\xe1\xde\xa7\x58\xb9\x44\x38\x88\x0e\xf9\x02\xf4\x14\x34\x28\xc5\xc4\x59\xc8\x2a\x8e\x1d\xed\xda\x41\xe1\xf4\x28\xa9\x5e\x35\x9a\x5f\x1f\x76\xf9\x64\x35\xa9\x25\x92\x4e\x48\x03\x7a\x5b\xaa\xca\x1f\x9a\xd4\x32\xac\x6f\x36\x48\x7d\x39\x5e\xdd\x2a\x65\x8e\xf4\x51\x33\xbb\x8d\x1b\x66\x34\x44\x47\x54\xb0\x71\x24\x65\xcc\x01\x29\xb1\x59\xd0\x8b\x10\x65\x4c\x9c\xed\x22\x44\x8a\x6f\x9c\x41\x39\xe3\x8c\xe4\x9b\x76\xf6\x4e\x7b\x66\x8b\x2e\x10\x76\x79\xd7\xda\x77\x8b\xf4\xe5\x9f\xfc\xe8\x88\x94\xad\x76\x4d\xfc\x90\x0e\xd6\xe9\xbe\x3f\xdc\xfc\x33\xf9\x2c\x76\xef\x2f\x5a\x5a\xfc\xe8\xfe\x03\x00\x00\xff\xff\x75\x63\x79\x7a\xc8\x00\x00\x00"),
		},
		"/role_write_access.yaml": &vfsgen۰CompressedFileInfo{
			name:             "role_write_access.yaml",
			modTime:          time.Date(2021, 11, 29, 20, 16, 39, 227674765, time.UTC),
			uncompressedSize: 181,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\xca\x31\x8a\xc3\x30\x10\x46\xe1\x5e\xa7\x18\x54\x2e\xc8\xcb\x76\x8b\x2e\x90\x3e\x45\x9a\x90\xe2\xb7\x3c\x21\x83\x65\xc9\xcc\x48\x0e\xce\xe9\x83\x8b\x40\xca\xf7\xf8\xb0\xca\x85\xd5\xa4\x96\x48\x3a\x22\x0d\xe8\xed\x51\x55\x5e\x68\x52\xcb\x30\xff\xdb\x20\xf5\x77\xfb\x73\xb3\x94\x29\xd2\xb9\x66\x76\x0b\x37\x4c\x68\x88\x8e\xa8\x60\xe1\x48\x4f\x95\xc6\x01\x29\xb1\x59\xd0\xc3\x10\x65\x8c\x9c\xed\x30\x44\x8a\x3b\xf6\xa0\x9c\xb1\x47\xf2\x4d\x3b\x7b\xa7\x3d\xb3\x45\x17\x08\xab\x9c\xb4\xf6\xd5\x22\x5d\xfd\x8f\xbf\x39\x22\x65\xab\x5d\x13\x7f\xad\x8d\x75\xfc\xe4\x3b\x00\x00\xff\xff\x5e\x9c\x4e\x4e\xb5\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/cluster_role_cluster_read.yaml"].(os.FileInfo),
		fs["/cluster_role_cluster_write.yaml"].(os.FileInfo),
		fs["/cluster_role_full_access.yaml"].(os.FileInfo),
		fs["/cluster_role_namespace_read.yaml"].(os.FileInfo),
		fs["/cluster_role_namespace_write.yaml"].(os.FileInfo),
		fs["/namespace.yaml"].(os.FileInfo),
		fs["/relay_default_cluster_role.yaml"].(os.FileInfo),
		fs["/relay_default_role.yaml"].(os.FileInfo),
		fs["/role_read_access.yaml"].(os.FileInfo),
		fs["/role_write_access.yaml"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}