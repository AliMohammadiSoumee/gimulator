// Code generated by Lorca. DO NOT EDIT.
package main

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"time"
)

var assets = map[string][]byte{}

var FS = &fs{}

type fs struct {}

func (fs *fs) Open(name string) (http.File, error) {
	if name == "/" {
		return fs, nil;
	}
	b, ok := assets[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return &file{name: name, size: len(b), Reader: bytes.NewReader(b)}, nil
}

func (fs *fs) Close() error { return nil }
func (fs *fs) Read(p []byte) (int, error) { return 0, nil }
func (fs *fs) Seek(offset int64, whence int) (int64, error) { return 0, nil }
func (fs *fs) Stat() (os.FileInfo, error) { return fs, nil }
func (fs *fs) Name() string { return "/" }
func (fs *fs) Size() int64 { return 0 }
func (fs *fs) Mode() os.FileMode { return 0755}
func (fs *fs) ModTime() time.Time{ return time.Time{} }
func (fs *fs) IsDir() bool { return true }
func (fs *fs) Sys() interface{} { return nil }
func (fs *fs) Readdir(count int) ([]os.FileInfo, error) {
	files := []os.FileInfo{}
	for name, data := range assets {
		files = append(files, &file{name: name, size: len(data), Reader: bytes.NewReader(data)})
	}
	return files, nil
}

type file struct {
	name string
	size int
	*bytes.Reader 
}

func (f *file) Close() error { return nil }
func (f *file) Readdir(count int) ([]os.FileInfo, error) { return nil, errors.New("not supported") }
func (f *file) Stat() (os.FileInfo, error) { return f, nil }
func (f *file) Name() string { return f.name }
func (f *file) Size() int64 { return int64(f.size) }
func (f *file) Mode() os.FileMode { return 0644 }
func (f *file) ModTime() time.Time{ return time.Time{} }
func (f *file) IsDir() bool { return false }
func (f *file) Sys() interface{} { return nil }

func init() {
	assets["/index.html"] = []byte{0x3c, 0x21, 0x64, 0x6f, 0x63, 0x74, 0x79, 0x70, 0x65, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x0a, 0x0a, 0x0a, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x0a, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x50, 0x61, 0x70, 0x65, 0x72, 0x20, 0x53, 0x6f, 0x63, 0x63, 0x65, 0x72, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x62, 0x6f, 0x64, 0x79, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x41, 0x72, 0x69, 0x61, 0x6c, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x68, 0x6f, 0x72, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x66, 0x69, 0x78, 0x65, 0x64, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7a, 0x2d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x20, 0x31, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x30, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x78, 0x3a, 0x20, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x33, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x66, 0x69, 0x78, 0x65, 0x64, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7a, 0x2d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x20, 0x31, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x30, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x79, 0x3a, 0x20, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x74, 0x6f, 0x70, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x30, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x34, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x34, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x32, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x20, 0x30, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x34, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x6c, 0x65, 0x66, 0x74, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x65, 0x66, 0x74, 0x3a, 0x20, 0x30, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x37, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x72, 0x69, 0x67, 0x68, 0x74, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x30, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x64, 0x65, 0x65, 0x70, 0x73, 0x6b, 0x79, 0x62, 0x6c, 0x75, 0x65, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x33, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x64, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x35, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x65, 0x66, 0x74, 0x3a, 0x20, 0x35, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x28, 0x2d, 0x35, 0x30, 0x25, 0x2c, 0x20, 0x2d, 0x35, 0x30, 0x25, 0x29, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3e, 0x0a, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x0a, 0x0a, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0x0a, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x68, 0x6f, 0x72, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x20, 0x6c, 0x65, 0x66, 0x74, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x76, 0x67, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x64, 0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x73, 0x76, 0x67, 0x43, 0x61, 0x6e, 0x76, 0x61, 0x73, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x73, 0x76, 0x67, 0x3e, 0x0a, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x0a, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x68, 0x6f, 0x72, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x3a, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3b, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x25, 0x3b, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x78, 0x3a, 0x20, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x3b, 0x20, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x2d, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x32, 0x30, 0x25, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x3b, 0x20, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x33, 0x30, 0x70, 0x78, 0x3b, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x61, 0x28, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x30, 0x2e, 0x38, 0x29, 0x3b, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x61, 0x28, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x30, 0x2e, 0x39, 0x35, 0x29, 0x3b, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x62, 0x6f, 0x6c, 0x64, 0x3b, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x30, 0x70, 0x78, 0x20, 0x35, 0x70, 0x78, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x23, 0x31, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x28, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x29, 0x3b, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x36, 0x70, 0x78, 0x20, 0x36, 0x70, 0x78, 0x20, 0x36, 0x70, 0x78, 0x20, 0x30, 0x70, 0x78, 0x3b, 0x20, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x3a, 0x20, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x75, 0x70, 0x70, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x20, 0x31, 0x31, 0x35, 0x25, 0x3b, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x3b, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x32, 0x70, 0x78, 0x20, 0x36, 0x70, 0x78, 0x3b, 0x22, 0x3e, 0x41, 0x6c, 0x69, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x75, 0x70, 0x70, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x3b, 0x20, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x33, 0x30, 0x70, 0x78, 0x3b, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x61, 0x28, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x30, 0x2e, 0x38, 0x29, 0x3b, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x61, 0x28, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x30, 0x2e, 0x39, 0x35, 0x29, 0x3b, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x62, 0x6f, 0x6c, 0x64, 0x3b, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x30, 0x70, 0x78, 0x20, 0x35, 0x70, 0x78, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x34, 0x3a, 0x35, 0x30, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x3a, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3b, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x25, 0x3b, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x78, 0x3a, 0x20, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x3b, 0x20, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x2d, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x32, 0x30, 0x25, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x28, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x29, 0x3b, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x36, 0x70, 0x78, 0x20, 0x36, 0x70, 0x78, 0x20, 0x36, 0x70, 0x78, 0x20, 0x30, 0x70, 0x78, 0x3b, 0x20, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x3a, 0x20, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x20, 0x31, 0x31, 0x35, 0x25, 0x3b, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x3b, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x32, 0x70, 0x78, 0x20, 0x36, 0x70, 0x78, 0x3b, 0x22, 0x3e, 0x54, 0x75, 0x72, 0x6e, 0x3a, 0x20, 0x4d, 0x61, 0x6d, 0x61, 0x64, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x3a, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3b, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x25, 0x3b, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x78, 0x3a, 0x20, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x3b, 0x20, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x2d, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x32, 0x30, 0x25, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x3b, 0x20, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x33, 0x30, 0x70, 0x78, 0x3b, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x61, 0x28, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x30, 0x2e, 0x38, 0x29, 0x3b, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x61, 0x28, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x30, 0x2e, 0x39, 0x35, 0x29, 0x3b, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x62, 0x6f, 0x6c, 0x64, 0x3b, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x30, 0x70, 0x78, 0x20, 0x35, 0x70, 0x78, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x23, 0x32, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x28, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x29, 0x3b, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x36, 0x70, 0x78, 0x20, 0x36, 0x70, 0x78, 0x20, 0x36, 0x70, 0x78, 0x20, 0x30, 0x70, 0x78, 0x3b, 0x20, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x3a, 0x20, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x20, 0x31, 0x31, 0x35, 0x25, 0x3b, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x3b, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x32, 0x70, 0x78, 0x20, 0x36, 0x70, 0x78, 0x3b, 0x22, 0x3e, 0x4d, 0x61, 0x6d, 0x61, 0x64, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x3b, 0x20, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x33, 0x30, 0x70, 0x78, 0x3b, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x61, 0x28, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x30, 0x2e, 0x38, 0x29, 0x3b, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x61, 0x28, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x30, 0x2e, 0x39, 0x35, 0x29, 0x3b, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x62, 0x6f, 0x6c, 0x64, 0x3b, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x30, 0x70, 0x78, 0x20, 0x35, 0x70, 0x78, 0x3b, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x34, 0x3a, 0x31, 0x32, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x0a, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x73, 0x76, 0x67, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x28, 0x27, 0x23, 0x73, 0x76, 0x67, 0x43, 0x61, 0x6e, 0x76, 0x61, 0x73, 0x27, 0x29, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x75, 0x70, 0x70, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x28, 0x27, 0x23, 0x75, 0x70, 0x70, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x27, 0x29, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x75, 0x70, 0x70, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x28, 0x27, 0x23, 0x75, 0x70, 0x70, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x27, 0x29, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x28, 0x27, 0x23, 0x74, 0x75, 0x72, 0x6e, 0x27, 0x29, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x28, 0x27, 0x23, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x27, 0x29, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x28, 0x27, 0x23, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x27, 0x29, 0x3b, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x28, 0x73, 0x74, 0x72, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x76, 0x67, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x73, 0x74, 0x72, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x70, 0x70, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x28, 0x73, 0x74, 0x72, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x75, 0x70, 0x70, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x73, 0x74, 0x72, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x70, 0x70, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x28, 0x73, 0x74, 0x72, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x75, 0x70, 0x70, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x73, 0x74, 0x72, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x75, 0x72, 0x6e, 0x28, 0x73, 0x74, 0x72, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x75, 0x72, 0x6e, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x73, 0x74, 0x72, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x28, 0x73, 0x74, 0x72, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x73, 0x74, 0x72, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x28, 0x73, 0x74, 0x72, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x73, 0x74, 0x72, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x28, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x73, 0x76, 0x67, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x64, 0x74, 0x68, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x28, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x73, 0x76, 0x67, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x4d, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x61, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x28, 0x61, 0x29, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x0a, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x0a, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0x0a, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x0a}
}
