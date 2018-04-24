package docutils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"regexp"
)

const XLS 	= 	".xls"
const XLSX 	= 	".xlsx"
const DOC 	= 	".doc"
const DOCX 	= 	".docx"
const PPT 	= 	".ppt"
const PPTX 	= 	".pptx"
const ODT 	= 	".odt"
const RTF 	= 	".rtf"
const TXT 	= 	".txt"
const ODS 	= 	".ods"
const ODP 	= 	".odp"
const AVI 	= 	".avi"
const MPG 	= 	".mpg"
const MP3 	= 	".mp3"
const PDF 	= 	".pdf"
const JPG 	= 	".jpg"
const JPEG 	= 	".jpeg"
const PNG 	= 	".png"

var SupportedFileTypes = []string{
	XLS, XLSX, DOC, DOCX, PPT,
	PPTX, ODT, RTF, TXT, ODS,
	ODP, PDF, JPG, JPEG, PNG,
}

var OfficeFileTypes = []string{
	XLS, XLSX, DOC, DOCX, PPT, PPTX,
	ODT, RTF, TXT, ODS, ODP,
}

var ImageFileTypes = []string{
	JPEG, JPG, PNG,
}


func Jpeg2SwfPageConverter(fileIn string, fileOut string) {
	args := []string{"-o",
		fileOut,
		fileIn,
	}
	path, err := exec.LookPath("jpeg2swf")
	if err != nil {
		log.Fatal("Cannot find jpeg2swf in PATH")
	}
	fmt.Printf("jpeg2swf is available at %s\n", path)
	//out, err := exec.Command("jpeg2swf", args...).Output()
	cmd := exec.Command("jpeg2swf", args...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Success: %s\n", out)
	}	
	
}

func ImageMagickPageConverter(fileIn string, fileOut string) {
	args := []string{"-density", "150",
          "-quality", "90", "-flatten", "+dither",
		fileIn,
		fileOut,
	}

//	path, err := exec.LookPath("convert")
//	if err != nil {
//		log.Fatal("Cannot find convert in PATH")
//	}
//	fmt.Printf("convert is available at %s\n", path)
	out, err := exec.Command("convert", args...).Output()
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Success: %s\n", out)
	}	
}

func GhostscriptPageExtractor(source string, dest string, page int) {
	firstPage := "-dFirstPage=" + strconv.Itoa(page)
	lastPage := "-dLastPage=" + strconv.Itoa(page)
	outputFile := "-sOutputFile=" + dest
	args := []string{
		"-sDEVICE=pdfwrite",
		"-dNOPAUSE",
		"-dQUIET",
		"-dBATCH",
		firstPage,
		lastPage,
		outputFile,
		"/etc/bigbluebutton/nopdfmark.ps",
		source,
		}
		
	path, err := exec.LookPath("gs")
	if err != nil {
		log.Fatal("Cannot find gs in PATH")
	}
	fmt.Printf("gs is available at %s\n", path)
	cmd := exec.Command("gs", args...)
	printCommand(cmd)
	ExecCommandWithTimeout(cmd, 5)
}

func ConvertOfficeDocToPdf(fileIn string, fileOut string, port int) {
	args := []string{"-f", "pdf",
		"-eSelectPdfVersion=1",
		"-eReduceImageResolution=true",
		"-eMaxImageResolution=300",
		"-p",
		strconv.Itoa(port),
		"-o",
		fileOut,
		fileIn,
	}
//	path, err := exec.LookPath("unoconv")
//	if err != nil {
//		log.Fatal("Cannot find unoconv in PATH")
//	}
//	fmt.Printf("unoconv is available at %s\n", path)
	cmd := exec.Command("unoconv", args...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Success: %s\n", out)
	}
}

func Pdf2SwfPageCounter(fileIn string) (int) {
	args := []string{
		fileIn,
	}
//	path, err := exec.LookPath("pdfinfo")
//	if err != nil {
//		log.Fatal("Cannot find pdfinfo in PATH")
//	}
//	fmt.Printf("pdfinfo is available at %s\n", path)
	out, err := exec.Command("pdfinfo", args...).Output()
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
//		fmt.Printf("Success: %s\n", out)
		dst := string(out[:])
		lines := strings.Split(dst, "\n")
//		fmt.Printf("Success: %s\n", lines)
		for _, v := range lines {
			re := regexp.MustCompile(`Pages:[\s]+(\d+)`)
//			fmt.Printf("Line: %s\n", v)
			result_slice := re.FindStringSubmatch(v)
//			for k, rs := range result_slice {
//    			fmt.Printf("%d. %s\n", k, rs)
//			}
			if result_slice == nil {
//				fmt.Printf("Fail: %s\n", result_slice)
			} else {
//				fmt.Printf("Success: %s\n", result_slice[1])
				numpages, err := strconv.Atoi(result_slice[1])
				if (err == nil) {
					return numpages
				}
				
			}
		}

 		//fmt.Println(len(dst), dst)
 		return 0
	}	

	return 0
}

func Pdf2SwfBitmapPageConverter(fileIn string, fileOut string, page int) {
	fontsDir := "/usr/share/fonts"
	args := []string{"-T9", 
		"-s", "poly2bitmap",
		"-F",
		fontsDir,
		"-p",
		strconv.Itoa(page),
		fileIn,
		"-o",
		fileOut,
	}
//	path, err := exec.LookPath("pdf2swf")
//	if err != nil {
//		log.Fatal("Cannot find pdf2swf in PATH")
//	}
//	fmt.Printf("pdf2swf is available at %s\n", path)
	out, err := exec.Command("pdf2swf", args...).Output()
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Success: %s\n", out)
	}	
}



func parseTags(output string) (numPlacementTags, numTextTags, numImageTags int) {
	numPlacements := regexp.MustCompile(`DEBUG(?:\s+)Using(?:.+)`)
	numTexts := regexp.MustCompile(`VERBOSE(?:\s+)Updating(?:.+)`)
	numImages := regexp.MustCompile(`VERBOSE(?:\s+)Drawing(?:.+)`)

	numPlacementTags = 0
	numTextTags = 0
	numImageTags = 0
	
	splits := strings.Split(output, "\n")
//	fmt.Printf("%v\n", splits)
	for _, v := range splits {
		//cleaned := strings.Replace(strings.TrimSpace(v), " ", "", -1)
		cleaned := v
//		fmt.Printf("cleaned: %s\n", cleaned)
		if numPlacements.MatchString(cleaned) {
//			fmt.Printf("numPlacements Match: %s\n", cleaned)
//			res := numPlacements.FindAllStringSubmatch(cleaned, -1)
//			fmt.Printf("numPlacements %v\n", res)
			numPlacementTags += 1	
		} else if numTexts.MatchString(cleaned) {
//			fmt.Printf("numTexts Match: %s\n", cleaned)
//			res := numTexts.FindAllStringSubmatch(cleaned, -1)
//			fmt.Printf("numTexts %v\n", res)
			numTextTags += 1
		} else if numImages.MatchString(cleaned) {
//			fmt.Printf("numImages Match: %s\n", cleaned)
//			res := numImages.FindAllStringSubmatch(cleaned, -1)
//			fmt.Printf("numImages %v\n", res)			
			numImageTags += 1
		}
	}

	return 
}


func parseTags1(output string) (numPlacementTags, numTextTags, numImageTags int) {
	numPlacements := regexp.MustCompile(`(\d+)DEBUGUsing`)
	numTexts := regexp.MustCompile(`(\d+)VERBOSEUpdating(?:.+)`)
	numImages := regexp.MustCompile(`(\d+)VERBOSEDrawing(?:.+)`)

	numPlacementTags = 0
	numTextTags = 0
	numImageTags = 0
	
	splits := strings.Split(output, "\n")
//	fmt.Printf("%v\n", splits)
	for _, v := range splits {
		cleaned := strings.Replace(strings.TrimSpace(v), " ", "", -1)
//		fmt.Printf("cleaned: %s\n", cleaned)
		if numPlacements.MatchString(cleaned) {
//			fmt.Printf("numPlacements Match: %s\n", cleaned)
			res := numPlacements.FindAllStringSubmatch(cleaned, -1)[0]
			num, err := strconv.Atoi(res[1])
			if err == nil {
				numPlacementTags = num
			}			
		} else if numTexts.MatchString(cleaned) {
//			fmt.Printf("numTexts Match: %s\n", cleaned)
			res := numTexts.FindAllStringSubmatch(cleaned, -1)[0]
			num, err := strconv.Atoi(res[1])
			if err == nil {
				numTextTags = num
			}
		} else if numImages.MatchString(cleaned) {
			res := numImages.FindAllStringSubmatch(cleaned, -1)[0]
//			fmt.Printf("numImages Match: %s\n", cleaned)
			num, err := strconv.Atoi(res[1])
			if err == nil {
				numImageTags = num
			}
		}
	}

	return 
}

func getSwfTags(out1 string) (int, int, int) {
    cmdStr := "echo \"" + out1 + "\" | egrep  'shape id|Updating font|Drawing' | sed 's/  / /g' | cut -d' ' -f 1-3  | sort | uniq -cw 15"
    out, err := exec.Command("bash","-c", cmdStr).Output()
    if err != nil {
    	checkError(err)
        return 0, 0, 0
    }
    return parseTags(string(out))
}



func Pdf2SwfPageConverter1(fileIn string, fileOut string, page int) (bool) {
	fontsDir := "/usr/share/fonts"
	args := []string{"pdf2swf", "-vv", "-T9", "-F",
		fontsDir, 
		"-p",
		strconv.Itoa(page),
		fileIn,
		"-o",
		fileOut, "| egrep  'shape id|Updating font|Drawing' | sed 's/  / /g' | cut -d' ' -f 1-3  | sort | uniq -cw 15",
	}

	cmdStr := strings.Join(args, " ")

//	path, err := exec.LookPath("pdf2swf")
///	if err != nil {
//		log.Fatal("Cannot find pdf2swf in PATH")
//	}
//	fmt.Printf("pdf2swf is available at %s\n", path)
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
//	printCommand(cmd)
	execCommandWithTimeout2(cmd, 5)
	if _, err := os.Stat(fileOut); err == nil {
  		return true
	} 

	return false
}

const MAX_NUM_PLACEMENTS = 8000
const MAX_NUM_IMAGES = 8000
const MAX_NUM_TEXTS = 2500

func Pdf2SwfPageConverter(fileIn string, fileOut string, page int) (bool) {
	fontsDir := "/usr/share/fonts"
	args := []string{"-vv", "-T9", "-F",
		fontsDir, 
		"-p",
		strconv.Itoa(page),
		fileIn,
		"-o",
		fileOut,
	}
//	path, err := exec.LookPath("pdf2swf")
//	if err != nil {
//		log.Fatal("Cannot find pdf2swf in PATH")
//	}
//	fmt.Printf("pdf2swf is available at %s\n", path)
	cmd := exec.Command("pdf2swf", args...)
//	printCommand(cmd)

	outConsole, _ := execCommandWithTimeout3(cmd, 5)

	if _, err := os.Stat(fileOut); err == nil {
		numPlacementTags,  numTextTags, numImageTags := parseTags(outConsole)
//		fmt.Printf("Tags: %d %d %d\n", numPlacementTags,  numTextTags, numImageTags )
		if numPlacementTags > MAX_NUM_PLACEMENTS || numImageTags > MAX_NUM_IMAGES || numTextTags > MAX_NUM_TEXTS {
			return false
		}
  		return true
	} 

	return false
}


func Png2SwfPageConverter(fileIn string, fileOut string) {
	args := []string{"-o",
		fileOut,
		fileIn,
	}
//	path, err := exec.LookPath("png2swf")
//	if err != nil {
//		log.Fatal("Cannot find png2swf in PATH")
//	}
//	fmt.Printf("png2swf is available at %s\n", path)
	cmd := exec.Command("png2swf", args...)
	ExecCommandWithTimeout(cmd, 5)

}

func Pdf2PngPageConverter(fileIn string, fileOut string) {
	args := []string{
		"-density",
		"300x300",
		"-quality",
		"90",
		"+dither",
		"-depth",
		"8",
		"-colors",
		"256",
		fileIn,
		fileOut,
	}
//	path, err := exec.LookPath("convert")
//	if err != nil {
//		log.Fatal("Cannot find convert in PATH")
//	}
//	fmt.Printf("convert is available at %s\n", path)
	out, err := exec.Command("convert", args...).Output()
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Success: %s\n", out)
	}	
}

func Image2PngPageConverter(fileIn string, fileOut string) {
	args := []string{
		fileIn,
		fileOut,
	}
//	path, err := exec.LookPath("convert")
//	if err != nil {
//		log.Fatal("Cannot find convert in PATH")
//	}
//	fmt.Printf("convert is available at %s\n", path)
	out, err := exec.Command("convert", args...).Output()
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Success: %s\n", out)
	}	
}

func Pdf2TextPageConverter(fileIn string, fileOut string, page int) {
	args := []string{
		"-raw",
		"-nopgbrk",
		"-f",
		strconv.Itoa(page),
		"-l",
		strconv.Itoa(page),
		fileIn,
		fileOut,
	}
//	path, err := exec.LookPath("pdftotext")
//	if err != nil {
//		log.Fatal("Cannot find pdftotext in PATH")
//	}
//	fmt.Printf("pdftotext is available at %s\n", path)
	out, err := exec.Command("pdftotext", args...).Output()
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Success: %s\n", out)
	}	
}
