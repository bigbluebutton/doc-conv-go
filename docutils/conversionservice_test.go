package docutils_test

import (
	"testing"
	"os"
	"fmt"
	"path/filepath"
	"log"
	"docconv/docutils"
)


func _TestProcessDocument(t *testing.T) {
	curDir, err := os.Getwd()
	if (err != nil) {
		log.Fatal("Cannot get current directory.")
	}

	testDataDir := "_testdata"
	testOutDir := "_testoutput"
	testFile := "long-convert-pdf2swf.pdf"

	inFile := filepath.Join(curDir, testDataDir, testFile)

	os.RemoveAll(filepath.Join(curDir, testOutDir))
	os.MkdirAll(filepath.Join(curDir, testOutDir), os.ModePerm)

	uploadedFile := filepath.Join(curDir, testOutDir, testFile)
	CopyFile(inFile, uploadedFile)

	fmt.Printf("presentation is at %s\n", uploadedFile)

	pres := UploadedPresentation{
		MeetingId: 				"demo_meeting",
		PresentationId: 		"presentation_id",
		PresentationName: 		"Test Presentation",
		Filepath:  				uploadedFile,
		PresentationBaseUrl: 	"http://localhost/pres/",
	}

	fmt.Printf("presentation is at %s\n", pres)
	ProcessDocument(pres)
}
