package docutils

import (
	"fmt"
	"os"
	"strconv"
	"path/filepath"
)

type UploadedPresentation struct {
	MeetingId string
	PresentationId string
	PresentationName string
	Filepath string
	PresentationBaseUrl string
}


func Pdf2Swf(fileIn string, fileOut string, page int) {
	success := Pdf2SwfPageConverter(fileIn, fileOut, page)
	if success {
		fmt.Printf("Pdf2Swf success %s\n", fileOut)
	} else {		
		dir, filename := filepath.Split(fileIn)
		filebasename := GetFileBaseName(filename)
		tempPdfFilename := filebasename + "-" + strconv.Itoa(page) + ".pdf"
		tempPdfFile := filepath.Join(dir, tempPdfFilename)
		tempPngFilename := filebasename + "-" + strconv.Itoa(page) + ".png"
		tempPngFile := filepath.Join(dir, tempPngFilename)
		
		fmt.Printf("Creating tempPdffile %s\n", tempPdfFile)

		GhostscriptPageExtractor(fileIn, tempPdfFile, page)	
		fmt.Printf("Creating tempPngfile %s\n", tempPngFile)		
		ImageMagickPageConverter(tempPdfFile, tempPngFile)
		fmt.Printf("Creating swf %s\n", fileOut)
		Png2SwfPageConverter(tempPngFile, fileOut)
		os.Remove(tempPdfFile)
		os.Remove(tempPngFile)

	}
}

func processOfficeDoc(pres UploadedPresentation) {

}

func processPdfDoc(pres UploadedPresentation) {
	numPages := findNumberOfPages(pres.Filepath)
	fmt.Printf("numPages %s\n", strconv.Itoa(numPages))
	convertPdf2Swf(pres, numPages)
	// createTextFiles
	// createThumnails
	// createSVGImages
}

func convertPdf2Swf(pres UploadedPresentation, numPages int) {
	for i := 1; i <= numPages ; i++ {
		dir, filename := filepath.Split(pres.Filepath)
		filebasename := GetFileBaseName(filename)
		swfFilename := filebasename + "-" + strconv.Itoa(i) + ".swf"
		outFile := filepath.Join(dir, swfFilename)
		fmt.Printf("Processing %s\n", outFile)
  		Pdf2Swf(pres.Filepath, outFile, i)
	}
}

func findNumberOfPages(file string) (int) {
	return Pdf2SwfPageCounter(file)
}

func ProcessDocument(pres UploadedPresentation) {
	fileExt := GetFileType(GetFilename(pres.Filepath))
	fmt.Printf("FileExt %s\n", fileExt)
	if isFileSupported(fileExt) {
		fmt.Printf("presentation is supported %s\n", pres.Filepath)
		// send supported notification
		if isOfficeFile(fileExt) {
			processOfficeDoc(pres)
			processPdfDoc(pres)
		} else {
			processPdfDoc(pres)
		}
	} else {
		// send unsupported notification
	}
}