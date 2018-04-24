package main

import (
	"fmt"
	"strconv"
	"os/exec"
	docutils "docconv/docutils"
)

func getCPUmodel() string {
        cmd := "cat /proc/cpuinfo | egrep '^model name' | uniq | awk '{print substr($0, index($0,$4))}'"
        out, err := exec.Command("bash","-c",cmd).Output()
        if err != nil {
                return fmt.Sprintf("Failed to execute command: %s", cmd)
        }
        return string(out)
}

func main() {
	getCPUmodel()

	filename := docutils.GetFileType("/home/dev/foo.pptx")
	fmt.Printf("Filename: %s\n", filename)

	fileType := docutils.GetFileType("foo.tex")
	fmt.Printf("FileType: %s\n", fileType)

//	ConvertOfficeDocToPdf("_testdata/foo.pptx", "_testoutput/foo.pdf", 8100)
	docutils.GhostscriptPageExtractor("_testdata/BigSlidesFromCindy.pdf", "_testoutput/foo.pdf", 2)
//	ImageMagickPageConverter("_testdata/big.pdf", "_testoutput/big1.jpeg")
//	Jpeg2SwfPageConverter("_testdata/sampleslide.jpeg", "_testoutput/sampleslide.swf")
	docutils.Pdf2SwfPageCounter("_testdata/ProblemSlides.pdf")
//	Pdf2SwfPageConverter("_testdata/SeveralBigPagesPresentation.pdf", "_testoutput/f1.swf", 1)
//	Pdf2SwfBitmapPageConverter("_testdata/SeveralBigPagesPresentation.pdf", "_testoutput/f2.swf", 2)
//	Png2SwfPageConverter("_testdata/sample_7.png", "_testoutput/sample_2.swf")
//	Pdf2PngPageConverter("_testdata/big.pdf", "_testoutput/big.png")
//	Image2PngPageConverter("_testdata/sampleslide.jpeg", "_testoutput/sampleslide.png")
//	Pdf2TextPageConverter("_testdata/GoodPresentation.pdf", "_testoutput/page1.txt", 1)
//	
//	Pdf2SwfPageConverter("_testdata/BigSlidesFromCindy.pdf", "_testoutput/qcindy-" + strconv.Itoa(1) + ".swf", 1)
//	Pdf2SwfPageConverter("_testdata/BigSlidesFromCindy.pdf", "_testoutput/qcindy-" + strconv.Itoa(2) + ".swf", 2)
//	Pdf2SwfPageConverter("_testdata/BigSlidesFromCindy.pdf", "_testoutput/qcindy-" + strconv.Itoa(3) + ".swf", 3)

//	for i := 1; i < 46; i++ {
//		fmt.Println("Processing cindy-" + strconv.Itoa(i) + ".swf")
//  	Pdf2SwfPageConverter("_testdata/BigSlidesFromCindy.pdf", "_testoutput/cindy-" + strconv.Itoa(i) + ".swf", i)
//	}

//	for i := 1; i < 3; i++ {
//		fmt.Println("Processing plan_reseau-" + strconv.Itoa(i) + ".swf")
//    	Pdf2SwfPageConverter("_testdata/plan_reseau_2008-2009.pdf", "_testoutput/plan_reseau-" + strconv.Itoa(i) + ".swf", i)
//	}


//	for i := 1; i < 7; i++ {
//		fmt.Println("Processing FinalProject-" + strconv.Itoa(i) + ".swf")
//    	Pdf2SwfPageConverter("_testdata/FinalProject.pdf", "_testoutput/final-project-" + strconv.Itoa(i) + ".swf", i)
//	}


	inFile := "_testdata/BigSlidesFromCindy.pdf"
	outFile := "_testoutput/BigSlidesFromCindy.pdf"
	docutils.CopyFile(inFile, outFile)
/*	
	for i := 1; i <= 45; i++ {
		fmt.Println("Processing BigSlidesFromCindy-" + strconv.Itoa(i) + ".swf")
  		Pdf2Swf(outFile, "_testoutput/BigSlidesFromCindy-" + strconv.Itoa(i) + ".swf", i)
	}
*/
	inFile = "_testdata/long-convert-pdf2swf.pdf"
	outFile = "_testoutput/long-convert-pdf2swf.pdf"
	docutils.CopyFile(inFile, outFile)

	for i := 1; i <= 8; i++ {
		fmt.Println("Processing long-convert-pdf2swf-" + strconv.Itoa(i) + ".swf")
    	docutils.Pdf2Swf(outFile, "_testoutput/long-convert-pdf2swf-" + strconv.Itoa(i) + ".swf", i)
	}

	inFile = "_testdata/Problem2.pdf"
	outFile = "_testoutput/Problem2.pdf"
	docutils.CopyFile(inFile, outFile)

	for i := 1; i <= 1; i++ {
		fmt.Println("Processing Problem2" + strconv.Itoa(i) + ".swf")
    	docutils.Pdf2Swf(outFile, "_testoutput/Problem2-" + strconv.Itoa(i) + ".swf", i)
	}
}
