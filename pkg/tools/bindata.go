// Code generated by go-bindata. DO NOT EDIT.
// sources:
// deploy/data/virtlet-ds.yaml
package tools

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _deployDataVirtletDsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\x49\x73\xe3\xb8\x15\xbe\xfb\x57\xbc\x1a\x57\xa5\xa7\x0f\xb4\xec\xae\x4c\xba\x47\x95\x1c\x6c\x4b\xe3\x51\xb5\x2d\xa9\x24\xdb\x3d\x37\x15\x08\x3e\x51\x88\x40\x80\x01\x40\xda\xca\xaf\x4f\x01\x5c\xc4\x4d\xb2\xbc\x56\xda\x97\x56\x63\xf9\xf0\xf6\x05\x84\xe7\x79\x47\x24\x66\xf7\xa8\x34\x93\xa2\x0f\x24\x8e\x75\x2f\x3d\x3b\x5a\x33\x11\xf4\x61\x40\x30\x92\x62\x8e\xe6\x28\x42\x43\x02\x62\x48\xff\x08\x40\x90\x08\xfb\x90\x32\x65\x38\x9a\xfc\xff\x3a\x26\x14\xfb\xb0\x4e\x7c\xf4\xf4\x46\x1b\x8c\x8e\x74\x8c\xd4\x2e\xd7\xc8\x91\x1a\xa9\xec\x6f\x80\x88\x18\xba\xba\x26\x3e\x72\x9d\x0d\x00\xa8\x44\x18\x56\x87\x34\x18\xc5\x9c\x18\xcc\xf7\x54\x0e\xb7\x7f\x4d\x02\xec\x1f\xaf\x41\x76\x82\x02\x14\x24\xd9\xbf\x95\xd4\x66\x8c\xe6\x41\xaa\x75\x1f\x8c\x4a\x30\x1f\x0f\x84\x9e\x4a\xce\xe8\xa6\x0f\x97\x3c\xd1\x06\xd5\x1f\x4c\x69\xf3\x83\x99\xd5\x9f\xd9\x96\x7c\xe1\xb1\x83\x98\x8e\x06\xc0\xb4\x03\x00\x23\xe1\xd7\xb3\xcf\x80\x82\xf8\x1c\xe1\xfe\x46\xdb\x11\x9d\xa8\x94\xa5\x58\xd0\x01\x54\x0a\x43\x98\x40\x05\x0a\xb5\x21\x6a\x0b\xf7\xab\x91\xe0\x23\xd0\x15\xd2\x35\x06\x9f\x81\x88\x00\x7e\xfd\xf2\xd9\x82\xe4\x90\x66\x85\x90\x68\x04\xb9\x04\xa1\x51\x18\x54\xc0\x04\x30\xc1\x2a\xb0\x15\xf6\xa6\xa3\x41\x8d\xb5\x63\xf0\xa5\x34\xda\x28\x12\x43\xac\x24\xc5\x20\x51\x08\x02\x31\x70\x94\x52\x85\xc4\x20\x10\x8b\xb5\x64\x61\x44\x62\x8b\x5e\x51\xe9\x56\xd3\x39\xa0\x46\x95\x32\x8a\xe7\x94\xca\x44\x98\x71\x4d\x2d\xe5\x99\x52\xf0\x8d\x55\x07\xdc\xe7\x12\x88\x65\xa0\x41\x0a\xc7\x8d\x90\x01\x6a\x78\x60\x66\x05\xf8\x68\x14\x99\x65\x6a\xfb\x57\x21\x2d\xa7\xd6\x1c\x8a\x2c\x97\x96\xd5\xcd\x56\xc9\x76\xf7\x79\x6b\x14\x40\xe1\x7f\x12\xa6\x30\x18\x24\x8a\x89\x70\x4e\x57\x18\x24\x9c\x89\x70\x14\x0a\x59\x0e\x0f\x1f\x91\x26\xc6\x5a\x7d\x65\x67\x86\x39\xcf\x4d\xf6\x16\x55\xa4\xeb\xd3\x5e\x66\xc1\xc3\xc7\x58\xa1\xb6\x3e\xd3\x98\xb7\x2b\xd6\xb8\xe9\xd7\xd8\x69\xac\x00\x90\x31\x2a\x62\x7d\x02\x46\xa2\x35\x99\x12\x9e\x60\x0b\xd6\x02\x37\x64\x6b\xf9\xbe\x2c\xf4\x5e\x6e\x38\x86\xdb\x15\x36\x8c\x02\xa8\x8c\x19\xea\x02\xe0\x93\x86\x25\xc7\xc7\x54\xf2\x24\x42\x08\x14\x4b\x4b\xbb\x39\xb6\x96\x60\x35\x13\xe0\x92\x24\xdc\x38\xfd\x3b\xad\xf1\x24\x64\x02\x02\xa6\x9c\x61\xa2\xd0\x89\x42\x0d\x66\x45\xb6\x16\xec\xf6\x31\xe5\x64\x67\x8f\xb3\xa6\x85\x01\xf8\x1b\xe0\xcc\xb7\x67\xc3\xdf\x4a\x3f\xc0\x47\xa6\x4d\x61\x06\xd6\x5a\x8f\x0a\x2e\x33\xf7\x8e\x15\xc6\x44\xa1\x67\xf5\x51\x8a\x82\x45\x24\xc4\x3e\x44\x4c\x11\x61\x98\xee\xd5\x63\x40\x3e\x3f\x4d\x38\x2f\x5c\x78\xb4\x1c\x4b\x33\x55\x68\xbd\xa5\x5c\x45\x65\x14\x11\x11\x6c\x25\xec\x41\xaf\x7a\xdc\x89\x5e\x95\x53\x99\x8c\x6e\xac\x7d\xeb\xea\x86\x8c\xc8\xf5\x37\xed\x6d\x25\xe9\x65\x32\xd2\x5e\xc0\x54\x45\x7b\x91\xdd\x3c\x25\x66\xd5\x87\x5e\x2e\x4d\xaf\xbe\xa1\x85\xab\x12\xd1\x02\x50\x32\x26\x21\x71\x06\x0b\x17\x2c\x13\x33\x93\x82\xf0\x1d\x47\x55\x31\x0a\xdc\x40\xd2\x35\x2a\x2d\xe9\x7a\xc7\xa6\x94\x28\xbb\xb1\x97\x2d\x3c\xa9\xad\x2c\x40\xb8\x0c\x77\xec\xb6\x6a\xac\xce\x1e\xc3\x52\xaa\xcc\x54\x98\x08\x9d\xad\x64\x47\x70\xe6\xf7\x72\x93\xe8\x39\x9d\xe9\xcc\x1e\x5c\x5c\xa8\x69\xbc\x38\x34\x25\xca\xe3\xcc\xdf\x73\xb0\xd7\x5c\x52\x32\x8d\xe9\x8e\x6d\xd5\x19\xaf\x25\x07\x4b\x64\xd3\xc0\xba\x93\x8f\x8d\x84\x34\x51\xcc\x6c\xac\x3b\xe2\xa3\xa9\x3a\x6f\xac\x58\xca\x38\x86\x18\xd4\x82\x31\x00\x8a\xb4\x6d\x51\xdf\xef\x2e\x86\x8b\xf1\x64\x30\x5c\x8c\xcf\x6f\x86\x15\x18\x17\x15\xfe\x50\x32\xaa\x07\x86\x25\x43\x1e\xcc\x70\xd9\x0c\x17\xd5\xa4\x9e\x9e\x35\x26\xdd\xa6\x8c\x53\x9b\x12\x4f\xac\xc4\x6d\xf4\x6e\x51\x73\x3f\x9a\xdd\x5e\x0f\x6f\x17\x83\xd1\xfc\xfc\xe2\x7a\xb8\xf8\x7e\x7f\xf3\x34\x49\x59\xfa\xb8\x21\xf1\x77\xdc\x74\x50\x56\x13\xa0\x97\x2d\x6e\x2c\x71\x01\x34\x60\xda\x26\xbd\xc5\x3a\x8d\x1a\xd3\x32\xce\x0c\xbf\x21\xcf\x26\xd1\xf3\xd9\x68\x72\xbf\x98\xdf\x4d\xa7\x93\xd9\xed\x87\x91\xad\x15\x93\xe9\x42\x27\x71\x2c\x95\x79\x19\xe1\x83\xc9\x8f\xf1\xf5\xe4\x7c\xb0\x98\xce\x26\xb7\x93\xcb\xc9\xf5\xc7\xc9\x5c\x3e\x08\x2e\x49\xb0\x88\x95\x34\x92\x4a\xfe\x32\x06\xae\x27\x57\xd7\xc3\xfb\xe1\xc7\xd1\xcd\x65\xc8\x31\xc5\x17\x92\x7b\x79\x7e\x3d\xba\x9c\x2c\xe6\x77\x17\xe3\xe1\xc7\x19\x0a\x25\x9c\x51\xe9\xe9\xc4\x17\xf8\x4c\x43\x19\xdd\x9c\x5f\x0d\x17\xb3\xe1\xd5\xf0\xaf\xe9\xe2\x76\x76\x3e\x9e\x5f\x9f\xdf\x8e\x26\xe3\x0f\xa3\xdd\xc5\xec\x85\xc2\x10\x1f\xe3\x85\x51\x44\x68\xee\x32\xd3\xcb\xe4\x3f\x3b\xff\xb1\x18\x0c\xef\x47\x97\xc3\xf9\x87\x71\xa0\xc8\xc3\x22\x40\x5b\xbd\xea\x17\x3a\x69\x1e\x12\xaf\x27\x57\x57\xa3\xf1\xd5\x87\x87\x45\x2e\xc3\x90\x89\xe6\x92\x43\x2d\x7e\x7a\xb7\xb8\x99\x0c\x3e\xd0\x43\x69\x9c\x78\x91\x0c\x5e\xea\xa2\x99\xc5\x57\x4c\x7d\xbe\x18\x8c\x66\x4d\xea\xfb\xd0\x43\x43\x8b\xa4\x9d\x57\x16\x45\xb5\x4c\x5b\x95\x72\x59\xd0\x64\x95\xc8\xc1\x55\xe6\x31\x8c\x04\x50\xa2\x11\x1e\x6c\xa1\xfd\x6f\xa4\x06\xb8\xa4\x84\x97\xc5\xad\x43\xb0\xb3\x0f\x44\x18\x5b\x51\xdb\xae\x8d\x19\x10\xd2\x80\x5c\x2e\x19\x65\x84\xf3\x0d\x90\x94\x30\xee\x3a\x3b\x29\xf0\x0d\x8a\xd8\x9c\x91\x43\xea\xd7\x6a\xb1\xa3\x37\xba\xb7\xd4\x3d\x1a\x2a\x99\xc4\xad\x52\xa7\x31\x5c\xdf\x6a\x6b\xa4\x48\x06\x09\xaf\xb9\x51\xb6\xb1\x3d\xae\x90\x04\x13\xc1\x37\x2d\x65\x57\x21\x6d\x8f\xda\xc2\x6a\x0c\x1e\x04\x54\x2f\xa2\x5f\x53\x58\xbf\xae\x36\xec\xde\xdd\x34\x3a\xd8\x61\x8c\xed\xdd\xb6\x3e\x7f\x62\xb7\x67\x0b\x77\x34\xba\x62\xb2\xb6\xcd\xe2\x32\x74\x0d\x1c\x2b\x5b\xb3\x15\x2a\x04\x1f\x29\x71\xd7\x0a\x66\x85\xea\x81\x69\x2c\xdb\xb5\x07\xc6\x39\xc4\x4a\x06\x09\x45\x40\xa5\xa4\xaa\x42\x72\xb6\x46\x30\x2b\x56\x31\xac\x63\xb8\xcb\xaf\x2a\xa4\xed\xe0\xbc\xfc\x4e\x81\xae\x88\x0a\x30\x85\x25\xe3\x08\x9f\x32\x19\xc8\xb0\x97\x46\xba\x47\x96\xc1\xd7\xdf\x7c\xdf\xf7\xbe\xe1\xef\x5f\xbd\xb3\x33\xfc\xea\xfd\xfe\xdb\x3f\xce\xbc\xd3\x2f\x7f\xff\x72\x4a\xe8\xe9\xe9\xe9\xe9\x97\x1e\x65\x4a\x49\xed\xa5\xd1\xe2\xf4\x84\xcb\xf0\x53\x1f\xc6\x12\x74\x42\x57\x19\xa2\x54\x65\xdb\xb9\x69\x77\x0e\x91\xf6\x76\xb7\x2c\x15\x52\xda\x8d\x4e\x2e\xcc\xa7\x77\xb7\x95\xf6\x9c\xd6\xe3\x25\xcd\x83\xf5\x00\x26\x50\xeb\xa9\x92\x3e\x56\xb7\xe0\xe3\xf6\x92\x2b\xfb\x6b\x85\x8a\x8c\xc4\x9e\xcf\x44\xaf\x12\x2a\xb2\x51\x8f\x36\x06\xb4\xa4\xc4\x80\x07\x77\xe3\xd1\x5f\xfd\xa6\x01\xf6\xaa\x06\xe7\x29\x09\xff\xb4\x9c\xf5\x44\xc2\x79\x23\xc8\x76\xb6\xea\xff\xef\x41\xf6\x90\xe8\xf9\x76\x61\xe6\x38\x0b\x7e\xee\x7e\xa5\x1a\x59\x81\x28\x2c\xef\xb4\xc0\xdf\x80\x4e\x62\x54\x11\x13\x3f\x49\x50\xfe\xb8\x36\xfc\xb5\x21\xfd\x63\x83\x72\x1d\x25\xd1\x8e\x06\xeb\xbe\xee\x9a\x48\x09\x34\xa8\xcb\x1b\xa3\xfc\xaa\xa8\x97\x99\x64\xcf\x2e\x6b\x1d\x74\xc0\x75\x54\x37\xdf\xf9\x21\xbd\x58\x06\x6d\x93\xb1\xa8\x76\xa2\xf3\x5a\xeb\x10\x49\xbf\x3c\x0e\x57\x57\x74\x54\x76\x4d\x4a\xdd\xb0\x67\x7f\x7b\x95\xa6\xa4\x1d\xd8\x1d\x37\x4f\xd3\x52\x93\xc6\x71\x91\x32\x97\x2e\xdb\x90\x50\x48\x6d\x18\x85\x38\x51\xb1\xd4\xf8\x96\xd9\xe3\x67\xca\x07\xd6\xee\x73\xad\xb8\xdb\xc2\x3d\x09\xa0\xa2\xd9\xf7\xbb\xc7\x4d\x23\xfd\xdc\xf2\xf7\x2d\x82\xcc\x9f\x52\x9b\x5b\x79\xd9\xf8\x08\xf3\xfa\x30\xf3\x36\xae\xf3\xb6\x6e\xbe\x9b\xd7\xe7\x05\xfa\x5d\x09\x69\x7f\x2a\xcb\x34\x5a\xf9\xda\x61\x51\x2b\x15\xad\x75\xcf\x95\xd4\xc6\x76\xf8\x90\x75\xf8\x40\x28\x45\xad\x4b\x7b\x74\xdf\xc6\x2c\x7e\xd5\xb1\xda\x14\x36\xb9\xd9\xbb\xb1\xbb\x87\xea\xe8\xa0\xf6\xa2\x74\x65\xee\x2e\x31\xed\x05\xa9\xa5\xe5\x56\xa6\xde\xbb\xb5\x5a\xb7\x34\x2b\x99\x63\xb8\x9d\x0c\x26\x7d\x08\xa4\xf8\x64\xc0\x16\xf4\x54\x06\x98\x7f\x52\x80\x2c\xb7\xb9\x0a\xcd\x5a\x89\x6b\x2c\xb6\x1b\x57\x4c\x67\xbd\x44\x5e\xc5\xc0\xe5\x6c\x64\xfb\x8a\xc7\x0d\x30\xa1\x0d\xe1\x59\xa4\xb6\x45\x5c\xf5\x40\x26\x32\x55\x3a\x8b\xd8\x7e\x7d\x3c\x39\x84\x95\x7d\x5f\x32\x76\x7c\x0c\x79\x12\xaf\x2b\x4a\x74\xc5\x88\x83\x80\x9a\xce\xde\x15\x02\x9e\x06\xaa\x44\x85\xe6\xd7\x99\xbd\x9b\x5f\x51\x6d\x1c\x58\x6b\x1c\x24\x84\xce\x88\xb4\x33\x1e\x1d\x02\xd9\x54\x4c\xed\xa3\xd0\x21\xf2\x2c\x8b\x8c\x6a\x3c\xed\x8a\xc3\x07\x81\xed\xd5\xf2\x73\xc0\xba\x0a\xcc\x7d\xe5\xe5\x41\xd4\x75\x88\xbd\x51\x1b\x79\xdb\x0b\xc0\xfe\xae\xcc\xe8\x65\x75\x57\x67\xc9\xb5\xbf\x30\x6b\x3e\x3f\x51\x3e\xa1\x27\x24\x31\x2b\xa9\xd8\x7f\xdd\x9a\x93\xf5\x37\x7d\xc2\x64\x2f\x3d\xf3\xd1\x90\xe2\x61\x4a\xfe\x32\x63\x26\x39\x5e\x30\x11\x30\x11\xee\x79\xa1\xa2\x24\xc7\xfc\xe6\x92\xc4\xec\xca\xc6\xe2\x3d\x27\x1d\x01\xb4\xce\x68\x41\xea\xc4\xb7\x0d\xa9\xee\x1f\x79\xf9\xea\x79\xed\x29\xc4\xe1\xaf\x64\xac\x04\xda\xe7\x3d\x4f\x26\x2f\x78\x9c\xa3\x6c\x32\xb1\xeb\xbd\x52\x26\x79\x4a\xf5\xe0\x97\x5f\xdc\x0f\x85\x5a\x26\x8a\x62\x39\x5e\x3e\x0b\xd1\xf9\x80\x7b\xbc\xe1\x7e\xa7\xa8\xfc\xed\x3a\x77\xe7\x93\xff\x27\x44\xf3\x16\x5a\xee\xe0\xb1\x24\xc7\xb3\xb5\x2f\xaa\x82\xa7\x06\x47\x39\x3f\x35\x6e\x1a\xbc\x94\xd4\x67\xe4\xda\x7f\x39\xd3\xd9\x8f\x07\x62\xe8\xea\x9d\x38\x28\xdc\x27\xd1\xa8\xec\xcc\xab\x19\xf1\x6c\xeb\xa0\xb2\x18\xd0\x60\xea\x5d\x3d\xad\xc8\x1a\xd6\x20\x3c\x3f\x5f\xf6\x86\x6e\xd7\x52\x75\xd5\xff\x9e\x03\x7e\x95\x17\x62\x19\x6c\xe6\x0b\xfd\xcc\x8c\xdf\x37\x14\x45\x5b\x25\xbf\x83\x7c\x76\x19\xd2\x4f\x12\xa6\x3c\xaa\x82\xdd\x46\x4f\x62\x86\x8f\x06\x85\x7b\x6c\x95\x63\x76\x39\x42\xa2\x8d\x8c\x8a\xc1\x00\xdd\xab\xb0\x3c\x15\x55\x7c\x21\x0f\x4e\xed\x63\x8a\x26\x76\xfd\x4d\x77\xa0\xe7\xb3\x2e\x8f\x45\x24\x8e\x99\x08\x75\x75\xa2\xb4\xd0\x62\xa6\x72\x64\x19\x4b\xde\xdd\x0f\x6b\xf2\x7c\x7b\xf3\xb2\xb0\x6f\x6b\x52\x8d\x57\x2a\x9d\x80\x2f\xc8\x6e\xff\x0b\x00\x00\xff\xff\xec\xc3\x8e\x14\xd6\x2a\x00\x00")

func deployDataVirtletDsYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployDataVirtletDsYaml,
		"deploy/data/virtlet-ds.yaml",
	)
}

func deployDataVirtletDsYaml() (*asset, error) {
	bytes, err := deployDataVirtletDsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/data/virtlet-ds.yaml", size: 10966, mode: os.FileMode(420), modTime: time.Unix(1522279343, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"deploy/data/virtlet-ds.yaml": deployDataVirtletDsYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"deploy": &bintree{nil, map[string]*bintree{
		"data": &bintree{nil, map[string]*bintree{
			"virtlet-ds.yaml": &bintree{deployDataVirtletDsYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

