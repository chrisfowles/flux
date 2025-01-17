// Code generated by vfsgen; DO NOT EDIT.

package install

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

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 9, 5, 10, 37, 50, 538727977, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(2019, 8, 6, 11, 23, 6, 137011753, time.UTC),
			uncompressedSize: 836,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4b\xaf\xd3\x30\x10\x85\xf7\xfe\x15\x47\xba\x8b\x0b\xe8\x26\xa8\x3b\x94\x5d\xdb\x05\x0b\x10\x8b\xf0\xd8\x20\x16\x63\x7b\x42\x4d\x5d\x3b\xf2\x23\x3c\xa2\xfc\x77\x94\xa4\x95\x9a\xb6\x20\x55\xba\x3b\x7b\x7c\xc6\x73\xe6\xe8\x2b\x8a\x42\x3c\xe0\xd3\x8e\x11\x39\x74\x46\x31\x48\x29\x9f\x5d\x7a\x82\xb2\x39\x26\x0e\x08\xde\x72\x7c\x02\x39\xbd\x28\x41\x1a\xa7\x8d\xfb\x0e\x0a\x2c\x1e\xe0\x9d\xfd\x0d\xc7\xac\x59\xa3\xf1\x01\xef\xb2\xe4\xe0\x38\x71\xc4\x4f\x93\x76\x53\x4b\x21\x29\xb2\x1e\x27\x70\x8c\x50\xde\xa5\xe0\x2d\x5e\xd4\x9b\xf5\xf6\x65\x29\xa8\x35\x5f\x38\x44\xe3\x5d\x85\x6e\x25\xf6\xc6\xe9\x0a\x1f\x67\x57\xeb\xd9\x94\x38\x70\x22\x4d\x89\x2a\x01\x58\x92\x6c\xe3\x78\x02\x1c\x1d\xb8\x42\x63\xf3\x2f\x71\x7e\xe9\x7b\x98\x06\xe5\x07\x3a\x70\x6c\x49\x31\x86\xe1\xf8\x3e\x5d\x2b\xf4\xfd\xf2\xb5\xef\xc1\x4e\x0f\x83\x18\x73\x39\x37\x14\x24\xa9\x92\x72\xda\xf9\x60\xfe\x50\x32\xde\x95\xfb\x37\xb1\x34\xfe\x75\xb7\x92\x9c\xe8\xe4\x77\x3b\x27\x54\x7b\xcb\xf7\x9a\x15\x21\x5b\x9e\x24\x05\xa8\x35\x6f\x83\xcf\x6d\xac\xf0\xf5\xf1\xd5\xe3\xb7\xa9\x2f\x70\xf4\x39\x28\x5e\x14\x3b\x0e\xf2\xac\x50\xc0\x79\x57\x1f\x85\x9f\xeb\xf7\xff\xd6\x3e\xc3\x86\x9b\x99\x80\xfb\x17\xf5\x96\x6b\x6e\x46\xd1\x69\xd1\xff\xcc\x17\xc0\x75\xb6\x8b\xff\x62\x96\x3f\x58\xa5\x63\x76\x37\xc1\xb9\xb2\x73\x89\xc1\x25\x27\xb7\xc8\xb0\x71\x3c\x69\x6e\x28\xdb\x34\xa3\x32\x12\xf5\x37\x00\x00\xff\xff\xfd\x7f\x67\x6a\x44\x03\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(2019, 9, 5, 10, 37, 50, 530727937, time.UTC),
			uncompressedSize: 6452,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x5f\x6f\x1b\xb9\x11\x7f\xf7\xa7\x18\x28\x0f\x49\x00\x69\x65\xc7\xb9\x43\xb1\x57\x1f\x90\x4b\x2e\x6e\x9a\x8b\x63\xc4\x4d\x8b\x3e\x35\x14\x77\xa4\x25\xc4\x25\xb7\x1c\x52\x3a\xc1\xb8\xef\x5e\x0c\xb9\x7f\xb8\x96\x9c\x1c\xf2\xd6\x3c\xc4\x36\x77\xfe\x0f\xf9\x9b\x1f\xb9\x58\x2c\xce\x44\xab\xfe\x89\x8e\x94\x35\x25\x88\xb6\xa5\xe5\xee\xe2\x6c\xab\x4c\x55\xc2\x1b\x6c\xb5\x3d\x34\x68\xfc\x59\x83\x5e\x54\xc2\x8b\xf2\x0c\xc0\x88\x06\x4b\x58\xeb\xf0\xfb\xfd\x3d\xa8\x35\x14\x37\xa2\x41\x6a\x85\x44\xf8\xe3\x8f\xee\x7b\xfc\xb3\x84\xfb\xfb\xe9\xd7\xfb\x7b\x40\x53\xb1\x18\xb5\x28\xd9\x98\xc3\x56\x2b\x29\xa8\x84\x8b\x33\x00\x42\x8d\xd2\x5b\xc7\x5f\x00\x1a\xe1\x65\xfd\x9b\x58\xa1\xa6\xb4\x90\xfb\x66\x69\xef\x84\xc7\xcd\x21\x7d\xf4\x87\x16\x4b\xf8\x84\xd2\xa1\xf0\x78\x06\xe0\xb1\x69\xb5\xf0\xd8\x19\xcb\x32\xe0\x7f\xc2\x18\xeb\x85\x57\xd6\x0c\xc6\x01\x5a\x67\x1b\xf4\x35\x06\x2a\x94\x5d\xb6\xd6\xf9\x12\x66\x97\xe7\x97\x17\x33\x78\x02\x1e\xb5\xce\x24\xc0\x5b\x20\xe9\x44\x8b\xb0\x6c\xd0\x3b\x25\x89\x93\x6b\xad\x32\xfe\x29\x01\x2b\x17\x9d\x61\x3d\xc9\xe1\x41\x16\x00\x7d\x2d\xe2\xef\xe8\x76\x4a\xe2\x2b\x29\x6d\x30\xfe\x66\x2a\x08\xb0\xb3\x3a\x34\x38\x98\x5a\x74\xa6\x36\xca\x2f\xb6\x78\x18\x1c\x10\x57\xc1\x8f\x0e\xfb\x95\xd1\xde\x82\x55\xaa\xd8\xe0\x4c\xaa\xc2\xb5\x08\xda\x7f\xb0\x15\x96\x70\xfe\xf2\xfc\x1c\x9e\xc0\xbe\x46\x03\x0d\x47\x83\x15\x38\x14\xd5\xc2\x1a\x7d\x98\xc3\x1e\x61\x6f\xcd\x53\x0f\x2b\x04\xb1\xd2\xc8\xf5\x90\x75\x63\xab\xb3\xce\xe0\x13\xf8\x47\xad\x08\x14\x81\x00\xdf\xb4\x6b\x82\x40\x58\xc1\xda\x3a\xd8\xa0\x41\x27\xbc\x32\x1b\xb8\xbb\xfb\x1b\x6c\xf1\x40\x05\xbc\x33\xf0\xfe\x2f\x04\x3f\x5f\xc1\x45\x71\x71\x3e\x1f\xac\xf4\xbe\x53\x0a\x04\xc2\x61\x1e\x07\x59\x0e\xc5\x20\x56\x20\x80\xb0\x15\xbc\x29\xba\x42\xc1\x1e\x07\x33\x52\x18\xd8\x3b\xe5\x39\xd0\xe2\x74\xfd\x36\x68\x86\x62\x60\xd3\xfa\xc3\x1b\xe5\xf2\x22\x36\x58\xa9\xd0\x94\xf0\x01\x1b\xeb\x0e\x79\x9e\x08\x6b\xab\xb5\xdd\x73\x46\x9d\x6b\x45\x31\xd5\x40\xbc\x26\x40\x06\xf2\xb6\x51\x5c\x81\xad\xb1\x7b\xf3\x9f\xda\x92\xa7\xc1\xc4\x5a\x69\x9c\xc3\xbe\x56\xb2\x86\x83\x0d\xb0\x57\x5a\xa7\xa4\xbc\x85\xca\xf2\x39\xe3\x65\x56\xe2\x5f\x1c\xd8\xbd\xe1\xb0\x07\x03\x0e\x5b\x0b\x4e\xf8\x1a\x1d\xf8\x5a\x98\xce\xf1\x46\xf9\x3a\xac\xc0\xf2\x22\x82\x56\x5b\x2c\xe0\xdf\x36\x3c\xd5\x1a\x84\x26\xdb\xbb\x98\x16\x1b\x94\x07\x65\xbc\x8d\x3a\xd2\x1a\x2f\x94\x41\x37\x87\x15\x6a\xbb\x2f\xe0\x0e\xc7\xaa\xd6\xde\xb7\x54\x2e\x97\x95\x95\x54\xf0\xc6\x92\x15\x1f\x1d\x34\x4b\x3e\x7a\xe4\x97\x9b\xa0\x2a\xa4\x65\x20\x5c\xb4\x4e\xed\x84\xc7\xb8\xf5\x38\x91\xa2\xf6\x8d\x1e\x2c\xf5\xbd\x20\xaa\x17\xd2\x9a\xb5\xda\x0c\x9f\x00\xd2\xc2\x07\xd1\x96\xd9\x62\x7e\x90\x16\x99\xda\xf7\xf6\xa5\xd8\x86\x15\x2e\x93\x91\x71\xfb\x7d\xb3\x27\x7b\x45\x35\xaf\xd4\x62\x87\x20\xa0\x52\xeb\x35\x3a\x06\xcd\xde\x42\x77\xaa\x46\x60\x8c\x2d\x48\xe6\xf2\x26\x30\xb8\xec\x54\x85\x7d\xd9\xd7\x6a\xd3\x88\x76\x0c\x44\xf9\x1a\x84\x01\x34\xde\x1d\x62\x0e\x5f\x92\xd0\x97\x39\x08\x53\x41\x30\xd2\x36\x8c\xd6\x51\x3f\x65\xfb\x21\xb6\x53\x98\x6a\xb0\x82\x66\x17\x2d\x28\xa4\xae\x9f\x47\x1d\xe0\x32\x7c\x47\x07\x32\xb5\x6f\x76\x20\x22\x81\xb7\xa0\x1a\xc6\x49\xb8\xbe\xbd\x8e\x20\x00\xcf\x38\x2d\x52\x1b\xa3\xcc\xe8\x9c\x93\xdb\xa1\x53\x6b\x25\x23\x60\x43\x1b\x5c\x6b\x09\xe9\xf9\x9f\x28\xe4\x60\x25\xc1\x47\xaa\x22\x17\x88\xfd\xfd\x89\xc2\x81\x70\x9b\xf1\x98\x3e\x52\xb1\x4d\xbb\x61\xfc\xa0\xac\x34\x53\x08\x7e\xf2\x08\x08\x1f\xeb\x9d\x00\xe1\xbe\x9c\xc3\x49\x3c\xc2\xff\x6c\x42\x74\x55\x77\x18\x71\xd2\x58\x98\x95\xe9\x24\xce\x40\x35\x62\x83\x69\xf7\xb3\x42\x01\x6f\x95\xa9\x62\xce\x0d\xc3\x8a\x43\x39\xee\xda\x04\x29\x1a\x05\x21\x83\x47\x54\xe5\x26\x30\x4f\x00\xe1\x87\x73\x5f\x87\x55\x51\x59\xb9\x45\x57\x48\xdb\x2c\xdd\x32\x61\x40\xfc\xb1\xf4\x62\x28\x5d\xdf\x47\x9e\xf7\xcc\x05\xd8\xab\x17\x1b\xe0\x48\x8b\x41\x26\xba\x29\xa1\x33\xa8\x6c\x6e\xad\xbc\x28\x2e\x5e\x16\x2f\xa6\xb2\xb7\x41\xeb\x5b\xab\x95\x3c\x94\xf0\x6e\x7d\x63\xfd\xad\x43\xca\xb3\x70\x48\x36\x38\x89\x94\xe3\xb8\xc3\xff\x06\x24\x3f\x59\x03\x90\x6d\x28\xe1\x87\xf3\x66\xb2\xd8\x44\xa8\x2f\xe1\xc7\x97\x1f\xd4\x48\x13\xac\xcb\x95\x17\x63\x67\x6e\x23\x65\xb8\x3c\xbf\xe4\xc9\xa9\xcc\xda\xba\x26\x6e\x59\xa1\x07\x69\xad\x76\x68\x90\xe8\xd6\xd9\x15\xe6\x11\x70\x49\xaf\xa7\x53\x3b\xb9\x4a\x06\xa7\xcb\xc2\xd7\x25\x2c\x45\xab\x52\xa5\x77\x3f\x2e\x55\x85\xc6\x2b\x7f\x28\xda\xb0\xca\x64\x95\x51\x5e\x09\xfd\x06\xb5\x38\xdc\xf1\xf9\xac\xa8\x84\x1f\x32\x01\xaf\x1a\xb4\xc1\x9f\xf8\xc6\x43\x56\xfd\x7f\x84\x9a\x1d\xda\x49\x63\x4e\xd3\x23\x48\x63\xee\x36\x45\x86\x5e\xc6\xc8\xaa\x25\x51\xcd\x3c\xcf\x26\xe6\x09\xda\x76\x78\xb3\xe1\x96\x81\x32\x69\xcf\x3d\xa5\xa4\x43\x54\x2f\x27\x30\xd9\xd7\xec\xa3\xd1\x87\x12\xbc\x0b\xc8\xd6\x98\x03\x45\x84\x5a\x75\xc0\xce\x47\xaa\x45\xb7\xb6\x4e\x22\x1b\x4d\xa4\x87\x39\xcf\x63\x81\xe7\xbc\x64\x1a\xfb\x4e\xb8\x2e\xf6\x24\xf6\x7d\xe1\x67\x67\xf4\x9d\x91\x3a\x44\xe4\x64\xea\x96\x06\x5c\x8f\xaa\x89\x1b\x7c\x83\xca\xf4\x64\xe6\x27\x56\x7d\x40\x33\x06\x74\x85\x0a\xa5\x16\x8e\x29\xdb\xca\xee\x32\x00\xf8\x0a\x0d\x48\xf0\x98\x27\xef\xac\xf5\xcb\x82\xa8\x7e\x34\x01\x61\x26\x5e\x67\xe3\x88\x9a\x25\xcf\xf3\x5e\x24\xb3\x80\x66\xa7\x9c\x35\x71\x20\xa4\x59\x3b\x7b\xff\xf9\x97\x5f\x5f\x7f\xbc\x79\xfb\xee\x7a\x96\x46\xc0\x9c\xeb\x61\x77\xe8\xdc\x74\x5e\x67\x66\xe2\x88\x5b\x1d\xd2\x34\xf5\xfa\x54\x8e\x47\x83\xf6\x38\xc7\x71\x73\xb2\xf0\xa3\x89\xf2\xcc\xe3\x8b\x47\xef\x8d\x21\x3a\xa3\x22\x5d\x74\xb1\x27\x99\x89\x87\x84\x26\x6f\x7a\x64\x33\x3d\xf5\x16\x06\x84\xf6\xe8\x0c\x53\xeb\xa3\x88\xd7\xce\x36\xbc\x2d\x7a\xc6\x32\x07\x41\xbc\xdd\xba\xa9\xca\x65\xd0\x56\x6e\xe9\xb8\xd9\x68\x76\xe5\x89\xba\x8c\xe5\x9e\xd4\x65\x27\x74\xc0\xa3\x9a\x7c\x6b\x13\x3f\xdc\x03\xfd\xcc\xfd\xca\x0e\xe0\x91\x3f\x1d\xf5\x5f\x19\xf6\x8f\xec\x4b\x96\x4a\xec\x66\x22\x37\xc5\x87\x31\x68\x76\x59\x4e\x72\x48\x6d\x48\x57\x34\xac\x78\x10\x49\x21\x6b\xac\xb8\xb2\x79\x6b\x07\x56\xc9\x4d\xe4\xb2\xcc\x33\x2b\xd6\x75\xb4\x31\x53\xe8\xae\x98\x51\x71\x1e\x9d\xf0\xd5\x88\x42\xdb\xea\x03\x17\x82\xf2\x52\x8c\xe4\xcd\xef\x2d\x47\x19\xb8\xa5\x71\xc3\xc5\xfb\x70\xec\x03\xd4\x76\x1f\xaf\x7f\xd6\x18\x94\x3e\x12\x3b\x3f\x2d\xdd\x62\x31\x24\x10\xb9\x3f\x3b\xbf\x1a\x96\x8a\x8e\xf3\x14\xb4\x93\x85\xd4\x81\x3c\xba\x82\xf1\x4b\xe7\x25\xf9\x4c\xe9\xa8\x8d\xa5\x78\x9d\x44\xdf\xdd\x4e\x92\xe2\x53\x47\xe8\xe3\xf5\x72\xda\xd8\x31\x86\x5e\x9e\x2f\xf1\xde\xb1\x64\xbc\xf0\x65\x08\x9c\x47\xdc\x49\x5f\x9d\x4d\x48\x96\x22\x68\x02\xc5\x0b\x70\xac\x9e\xc2\x2a\xed\xa6\x55\xc4\xf5\x48\x71\xe2\xbd\xf7\x59\x7f\x99\x7c\x9e\xc7\xd2\x9f\xad\xb4\x0b\x99\x95\x65\xd7\xdf\x49\x20\x8c\x85\x09\xdf\x17\x95\x72\x57\x47\xa8\x9f\x87\xf5\x29\x23\x58\x63\xf3\x3e\x7f\xfa\x2d\xdd\xcf\x85\xd9\xa4\x6f\xd7\xca\xc7\x3b\x23\x29\x6f\xdd\x61\x40\xab\xb7\x4c\x0c\x27\xce\x79\x06\x05\xa7\xaf\xee\xef\xa1\xb8\x56\x9e\x2d\xc5\x67\x9e\xa9\xc4\xca\x09\x23\xeb\x5e\xe8\x97\xf8\x57\x7a\xf0\x51\xeb\xb8\xc4\x67\x83\x4e\x69\x32\x3f\x60\xbd\xbb\xd8\x06\xfa\xbb\x55\x26\x53\x98\xcd\x67\xdd\xbb\x91\x26\xcc\xd5\x99\x5e\x1d\x8f\xaa\xbd\x30\x71\xfb\x39\xe4\xae\xca\xc4\xe8\x1b\x61\xd4\x9a\xf9\x1e\x6f\x50\x52\x15\xba\x94\xeb\x03\xd6\x1c\xef\xbb\x96\x10\x82\xa9\xd0\x3d\x28\xa0\x43\x2d\xbc\xda\x61\xa4\x33\xd4\xb7\x77\x33\x29\xe2\x83\x0d\x3f\x24\x47\x61\x55\x29\x77\x31\x4f\x3f\x5f\x0c\x8f\x60\x63\x71\xe2\x23\xd7\xa9\xe2\xc4\x97\xa3\xbe\xaa\xbd\xd4\x09\x03\x9f\x09\xdd\x29\xfd\x40\xe8\x86\xce\xb1\x0c\x9c\xd6\xff\xb5\x11\xea\x64\x00\xc8\x1f\x7a\x0b\xbd\xd4\xf8\x8c\x77\x12\x74\x91\xcf\xe9\xde\x72\x41\xd1\xc4\xa7\x21\xae\x13\x4f\x03\xe5\x1f\x5c\xee\xf2\x5a\x75\xb8\xda\xa1\xe6\xd5\x57\x60\xb4\xd7\xe8\x6c\xb1\xd6\xd5\x5f\xb7\x78\x00\x55\xfd\x3c\x88\x7d\x65\x54\x66\x51\xb1\x09\xe1\x83\xc3\xc9\x0d\xf3\x84\xaf\xf8\xf9\xb0\x18\xe4\x69\x82\x05\x3d\x14\x82\xf2\x50\x0b\x8a\x30\x6f\x8d\x3e\x80\x90\x12\x29\xc1\x65\x8d\xe9\x91\xe6\x59\xff\x1e\xf0\x65\x2d\x34\xe1\x97\xe7\x27\xbc\xf5\xfa\xd3\x02\x93\x77\x41\xfa\xe4\x68\x1f\xef\x78\x3c\xf7\x83\x07\x3a\x18\x09\x2b\x6b\xb7\x5b\xc4\x96\xb7\xeb\xe0\x63\xb6\x51\x7e\x36\x87\x06\x05\x57\x8a\x8f\x39\x88\x78\xf1\xea\x76\x70\x68\xc9\x3b\x14\xcd\xb0\x95\x1f\x46\xc3\xa6\x17\xe4\x85\xc7\xab\x8d\xf2\x8f\x37\xdc\xe0\xef\xbe\xef\x7a\x36\x07\x84\x81\x59\xef\x63\xd6\xa3\x74\x66\xe4\x19\x16\x9b\x62\x0e\xff\x42\xa6\x1b\xaf\xb5\x0d\xd5\xf3\x22\xbe\x1a\x78\xbb\x65\xd2\x4a\xd0\x0a\xe7\x95\x0c\x5a\xb8\xbe\x8a\x9d\x95\x87\x03\xa6\xf3\x7a\xb5\x27\xbe\x98\x4a\xb6\x55\xec\xd9\x6e\xb1\xb7\x6e\x4b\xc3\x0d\xe4\x81\x5a\x74\x74\x25\x56\xf2\xe2\xc5\xe5\xf1\xff\x79\xc2\x77\xe8\x76\x27\x1e\x7b\x99\x6b\x8d\xd3\x95\xb7\xea\x4f\x39\xcc\x8b\x2d\x8f\x87\xd4\x2b\x42\x9f\xbd\x20\x3f\xcd\x1e\xa1\xb3\xd7\x64\x4e\x31\xbe\x8a\x44\xbe\x33\x05\x63\xad\xc8\xa3\x59\x74\x21\x5c\x95\x97\xe7\x97\x17\x67\xdd\x31\x7e\x55\x55\x2a\xdd\x35\x19\xc4\x5f\x31\x87\x99\xe0\xe5\xf8\x7d\x9c\xe3\xf7\xf7\xe0\xe2\x48\xf8\x86\xf6\x22\x3e\xe5\x4f\x8e\xfe\xf8\x5b\xef\xe0\x63\xdb\x99\x7f\x73\x73\xd7\x0f\x60\x9a\x77\xbc\x30\xb8\x6e\x1c\x83\xa9\xac\x27\xb0\x51\x18\x1a\x71\x88\x77\x74\xbd\x1b\x5f\x6a\x0c\x69\x6b\xb7\xa1\x05\x45\x14\x90\xc0\x1a\x20\xdb\x20\xbc\x0f\x2b\x74\x06\x3d\x12\x5b\x0f\x2d\x8d\x0f\x31\x95\xa1\xfe\x19\x60\x76\x63\x0d\xce\xf2\x2f\xaf\x63\x00\xf9\x53\x4c\x72\x4e\xd3\xd7\x99\x9e\xdf\xc5\xf8\x26\x5f\x06\xea\x39\xbb\x98\x9d\xfd\x2f\x00\x00\xff\xff\x52\x89\xd6\x57\x34\x19\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(2019, 8, 6, 11, 23, 6, 137011753, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(2019, 8, 6, 11, 23, 6, 137011753, time.UTC),
			uncompressedSize: 874,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\xcd\x6e\x9c\x40\x10\x84\xef\x3c\x45\x49\x7b\x0d\x1b\x61\x69\x2f\xdc\xa2\x38\x89\x2c\x25\xd6\x5e\x9c\x7b\x7b\x68\xf0\x28\xf3\x97\xe9\x66\xb3\x04\xf9\xdd\xa3\xd9\x5f\x36\xf6\x9c\x80\xaa\xaf\xa7\xa6\x80\xba\xae\xab\x15\x3c\x7b\x43\xe6\x85\x3b\x74\x9c\x5c\x9c\x3c\x07\xc5\x28\xdc\xe1\x79\xc2\x57\x37\xee\xa1\x11\x07\x47\xb5\x82\x89\x41\xc9\x06\xce\xb0\x9e\x06\x86\x67\xa5\x8e\x94\xd6\x15\x25\xfb\x93\xb3\xd8\x18\x5a\x50\x4a\xf2\x71\xd7\x54\xbf\x6c\xe8\x5a\xdc\x5f\xc6\x56\x67\x7b\x5b\x01\x81\x3c\xb7\xd7\xdd\xe7\x19\xb6\xc7\xfa\x91\x3c\x4b\x22\xc3\x78\x7d\x3d\x99\x0e\xb7\x2d\xe6\xf9\x56\x9d\x67\x70\xe8\x8a\x4d\x12\x9b\x32\x31\x73\x72\xd6\x90\xb4\x68\x2a\x40\xd8\xb1\xd1\x98\x8b\x02\x78\x52\xf3\xf2\x9d\x9e\xd9\xc9\xf1\xc1\x9b\x00\x15\xa0\xec\x93\x23\xe5\x13\xb2\x08\x5b\x96\xbb\xa1\xdf\xe3\x81\x73\x94\xb2\x2e\x5d\x5d\x98\xfa\x5d\xa6\xac\x43\x9b\x0b\xa1\x6d\xd6\x9b\x75\xb3\xb9\xd5\xb7\xa3\x73\xdb\xe8\xac\x99\x5a\x3c\xf4\x8f\x51\xb7\x99\xa5\xd4\x7a\x76\x51\x1e\x16\xf9\x6a\xd4\x1e\x9b\xe6\x0e\xc0\x0a\x3f\x68\x6f\xfd\xe8\xcb\x0e\x31\x4f\xe5\x95\x8e\xc2\x1f\x60\x03\x3c\x0f\xf4\x3c\x29\xcb\x12\x7c\xc0\xc6\xe3\x06\x14\xfb\x97\xd1\xc7\x8c\x18\x18\x56\xd9\x2f\xed\x09\x4d\x73\xd7\x34\x58\xe1\x9e\x7b\x1a\x9d\x22\xc5\x7c\xcd\xb5\x2a\x9e\xdd\xee\x78\xf9\x14\x4c\xf4\x87\x8f\x4c\x23\x06\x56\xb8\x38\x08\x62\x0f\x26\xf3\x82\xcc\xbf\x47\x16\x05\x85\x0e\x99\x25\xc5\x20\xbc\xbe\x0c\x2a\x53\x6f\x4e\x78\xec\xd3\x38\xcb\x41\xaf\x07\x58\x74\xbf\x8d\x59\xdb\x63\xba\x8b\x2c\x6c\xc6\x6c\x75\xfa\x1c\x83\xf2\x5e\xdb\x05\x97\xc7\xf0\x49\x9e\x84\xf3\xff\xcc\x49\xfa\x96\xe3\x98\xde\x6a\xe4\x5c\xfc\xb3\xcd\x76\x67\x1d\x0f\xfc\x45\x0c\x39\xd2\xc3\xaf\xd0\x93\x13\xae\xfe\x05\x00\x00\xff\xff\x5d\x9a\x63\xab\x6a\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(2019, 8, 6, 11, 23, 6, 137011753, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
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
