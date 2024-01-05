package gphoto2_test

import (
	"testing"

	. "github.com/weebney/gphoto2-go"
)

func TestPhoto(t *testing.T) {
	context := NewContext()
	defer context.Free()

	camera, err := NewCamera()
	if err != nil {
		t.Fatal(err)
	}

	err = camera.Init(context)
	if err != nil {
		t.Fatal(err)
	}

	path, err := camera.Capture(CAPTURE_IMAGE, context)
	if err != nil {
		t.Fatal(err)
	}

	file, err := camera.File(path.Folder, path.Name, FILE_TYPE_NORMAL, context)
	if err != nil {
		t.Fatal(err)
	}

	err = file.Save("image.jpg")
	if err != nil {
		t.Fatal(err)
	}

	err = camera.Free()
	if err != nil {
		t.Fatal(err)
	}
}
