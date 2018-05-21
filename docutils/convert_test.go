package docutils

import (
	"testing"
)

func TestGetFileType(t *testing.T) {
	fileType := GetFileType("foo.ppt")
	if (fileType != ".ppt") {
		t.Error("Expected ppt, got ", fileType)
	}
}

func TestGetFilename(t *testing.T) {
	fileType := GetFilename("/hpme/dev/foo.ppt")
	if (fileType != "foo.ppt") {
		t.Error("Expected foo.ppt, got ", fileType)
	}
}

func TestIsFileSupportedTrue(t *testing.T) {
	supported := isFileSupported(".pdf")
	if supported != true {
		t.Error("Expected true, got ", supported)
	}
}

func TestIsOfficeFileTrue(t *testing.T) {
	supported := isOfficeFile(".pptx")
	if supported != true {
		t.Error("Expected true, got ", supported)
	}
}

func TestIsPdfTrue(t *testing.T) {
	supported := isPdfFile(".pdf")
	if supported != true {
		t.Error("Expected true, got ", supported)
	}
}

func TestIsImageFileTrue(t *testing.T) {
	supported := isImageFile(".png")
	if supported != true {
		t.Error("Expected true, got ", supported)
	}
}

func TestPageCount(t *testing.T) {
	numPages := Pdf2SwfPageCounter("_testdata/ProblemSlides.pdf")
	if numPages != 16 {
		t.Error("Expected 16, got ", numPages)
	}
}

func TestConvertOfficeDocToPdf(t *testing.T) {
	ConvertOfficeDocToPdf("_testdata/lots-pages.ppt", "_testoutput/foo.pdf", 8100)
//	if numPages != 16 {
//		t.Error("Expected 16, got ", numPages)
//	}
}




