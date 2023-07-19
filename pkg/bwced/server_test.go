package bwced_test

import (
	"bufio"
	"encoding/base64"
	"github.com/garethjevans/bwced/pkg/bwced"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadImage(t *testing.T) {
	const filename = "bundle.tar.gz"

	root, err := os.MkdirTemp("", "root")
	assert.NoError(t, err)

	defer os.RemoveAll(root)

	//Set up a pipe to avoid buffering
	pr, pw := io.Pipe()
	//This writers is going to transform
	//what we pass to it to multipart form data
	//and write it to our io.Pipe
	writer := multipart.NewWriter(pw)

	go func() {
		defer writer.Close()
		//we create the form data field 'fileupload'
		//which returns another writer to write the actual file

		part, err := writer.CreateFormFile("file", filename)
		assert.NoError(t, err)

		////https://yourbasic.org/golang/create-image/
		//img := createImage()
		//
		////Encode() takes an io.Writer.
		////We pass the multipart field
		////'fileupload' that we defined
		////earlier which, in turn, writes
		////to our io.Pipe
		// Read entire JPG into byte slice.
		p := filepath.Join("testdata", filename)
		f, err := os.Open(p)
		assert.NoError(t, err)
		reader := bufio.NewReader(f)
		content, err := io.ReadAll(reader)
		assert.NoError(t, err)

		// Encode as base64.
		encoded := base64.StdEncoding.EncodeToString(content)

		_, err = part.Write([]byte(encoded))
		assert.NoError(t, err)
	}()

	//We read from the pipe which receives data
	//from the multipart writer, which, in turn,
	//receives data from png.Encode().
	//We have 3 chained writers !
	request := httptest.NewRequest("POST", "/", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	response := httptest.NewRecorder()
	handler := bwced.NewServer(root, 1024, false, nil)

	handler.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

	_, err = os.Stat(filepath.Join(root, filename))
	if os.IsNotExist(err) {
		t.Error("Expected file to exist")
	}
}
