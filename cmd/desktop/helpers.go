package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// Resource represents a single binary resource, such as an image or font.
// A resource has an identifying name and byte array content.
// The serialised path of a resource can be obtained which may result in a
// blocking filesystem write operation.
type Resource interface {
	Name() string
	Content() []byte
}

// StaticResource is a bundled resource compiled into the application.
// These resources are normally generated by the fyne_bundle command included in
// the Fyne toolkit.
type StaticResource struct {
	StaticName    string
	StaticContent []byte
}

// Name returns the unique name of this resource, usually matching the file it
// was generated from.
func (r *StaticResource) Name() string {
	return r.StaticName
}

// Content returns the bytes of the bundled resource, no compression is applied
// but any compression on the resource is retained.
func (r *StaticResource) Content() []byte {
	return r.StaticContent
}

// NewStaticResource returns a new static resource object with the specified
// name and content. Creating a new static resource in memory results in
// sharable binary data that may be serialised to the location returned by
// CachePath().
func NewStaticResource(name string, content []byte) *StaticResource {
	return &StaticResource{
		StaticName:    name,
		StaticContent: content,
	}
}

// LoadResourceFromURLString creates a new StaticResource in memory using the body of the specified URL.
func loadResourceFromURLString(urlStr string) (Resource, error) {
	res, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	name := filepath.Base(urlStr)
	return NewStaticResource(name, bytes), nil
}
