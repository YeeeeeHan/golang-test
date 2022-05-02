package server

import (
	"io/ioutil"
	"os"
	"testing"
)

func init() {

}

//// TestPrintScannedTokens adapted from https://gist.github.com/KEINOS/76857bc6339515d7144e00f17adb1090
//func TestPrintScannedTokens(t *testing.T) {
//	userInput := "login yeehan 123"
//
//	funcDefer, err := mockStdin(t, userInput)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	defer funcDefer()
//
//	msgActual := capturer.CaptureOutput(func() {
//		// Init Wallet
//		sessionUser := service.Wallet{}
//		ListenAndServe(&sessionUser)
//	})
//
//	msgContain := "foo bar"
//
//	assert.Contains(t, msgActual, msgContain)
//}

// mockStdin adapted from https://gist.github.com/KEINOS/76857bc6339515d7144e00f17adb1090
func mockStdin(t *testing.T, dummyInput string) (funcDefer func(), err error) {
	t.Helper()

	oldOsStdin := os.Stdin
	tmpfile, err := ioutil.TempFile(t.TempDir(), t.Name())

	if err != nil {
		return nil, err
	}

	content := []byte(dummyInput)

	if _, err := tmpfile.Write(content); err != nil {
		return nil, err
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return nil, err
	}

	// Set stdin to the temp file
	os.Stdin = tmpfile

	return func() {
		// clean up
		os.Stdin = oldOsStdin
		os.Remove(tmpfile.Name())
	}, nil
}
