package docutils

import (
	"fmt"
	"io"
	"os"
	"log"
	"bytes"
	"os/exec"
	"time"
	"strings"
	"bufio"
	"regexp"
	"path/filepath"
)

func ExecuteCommand(app string, args []string) {
	cmd := exec.Command(app, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
	    fmt.Println(fmt.Sprint(err) + ": " + string(output))
	    return
	} else {
	    fmt.Println(string(output))
	}

}

func printCommand(cmd *exec.Cmd) {
  fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func checkError(err error) {
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
}

func countTags(output string) (numPlacementTags, numTextTags, numImageTags int) {
	numPlacements := regexp.MustCompile(`DEBUG(?:\s+)Using(?:.+)`)
	numTexts := regexp.MustCompile(`VERBOSE(?:\s+)Updating(?:.+)`)
	numImages := regexp.MustCompile(`VERBOSE(?:\s+)Drawing(?:.+)`)

	numPlacementTags = 0
	numTextTags = 0
	numImageTags = 0
	
		//cleaned := strings.Replace(strings.TrimSpace(v), " ", "", -1)
		cleaned := output
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

	return 
}

func execCommandWithTimeout2(cmd *exec.Cmd, timeout time.Duration) (numPlacementTags, numTextTags, numImageTags int) {
    // Create stdout, stderr streams of type io.Reader
    stdout, err := cmd.StdoutPipe()
    checkError(err)
    stderr, err := cmd.StderrPipe()
    checkError(err)

    // Start command
    err = cmd.Start()
    checkError(err)

    
	numPlacementTags = 0
	numTextTags = 0
	numImageTags = 0

	go func() {
		multi := io.MultiReader(stdout, stderr)
		// read command's stdout line by line
		in := bufio.NewScanner(multi)

		for in.Scan() {
//		    log.Printf(in.Text()) // write each line to your log, or anything you need
		    placementTags,  textTags, imageTags := countTags(in.Text())
		    numPlacementTags += placementTags
		    numTextTags += textTags
		    numImageTags += imageTags
		}

		if err := in.Err(); err != nil {
		    log.Printf("error: %s", err)
		}		
	}()


	done := make(chan error, 1)
	go func() {
	    done <- cmd.Wait()
	}()
	select {
	case <-time.After(timeout * time.Second):
	    if err := cmd.Process.Kill(); err != nil {
	        log.Fatal("failed to kill: ", err)
	    }
	    log.Println("process killed as timeout reached")
	case err := <-done:
	    if err != nil {
	        log.Printf("process done with error = %v", err)
	    } else {
	        log.Print("process done gracefully without error")
	    }
	}

	return
 
}

func execCommandWithTimeout3(cmd *exec.Cmd, timeout time.Duration)(outConsole, errConsole string) {
    // Create stdout, stderr streams of type io.Reader
    stdout, err := cmd.StdoutPipe()
    checkError(err)
    stderr, err := cmd.StderrPipe()
    checkError(err)

    // Start command
    err = cmd.Start()
    checkError(err)

    outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, stdout)
        outC <- buf.String()
    }()

    outE := make(chan string)
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, stderr)
        outE <- buf.String()
    }()

	done := make(chan error, 1)
	go func() {
	    done <- cmd.Wait()
	}()
	select {
	case <-time.After(timeout * time.Second):
	    if err := cmd.Process.Kill(); err != nil {
	        log.Fatal("failed to kill: ", err)
	    }
	    log.Println("process killed as timeout reached")
	case err := <-done:
	    if err != nil {
	        log.Printf("process done with error = %v", err)
	    } else {
//	        log.Print("process done gracefully without error")
	    }
	}

	out := <-outC
	outConsole = string(out)
	outErr := <- outE
	errConsole = string(outErr)
	return
    // reading our temp stdout
 //   fmt.Println("previous output:")
 //   fmt.Print(out)

//    numPlacementTags,  numTextTags, numImageTags := parseTags(out)
//    fmt.Printf("Tags: %d %d %d\n", numPlacementTags,  numTextTags, numImageTags )
//   
    // reading our temp stdout
//    fmt.Println("previous error:")
//    fmt.Print(outErr)
}

func ExecCommandWithTimeout(cmd *exec.Cmd, timeout time.Duration) {
    // Create stdout, stderr streams of type io.Reader
 //   stdout, err := cmd.StdoutPipe()
 //   checkError(err)
 //   stderr, err := cmd.StderrPipe()
 //   checkError(err)
 
    // Start command
    err := cmd.Start()
    checkError(err)

    // Non-blockingly echo command output to terminal
 //   go io.Copy(os.Stdout, stdout)
 //   go io.Copy(os.Stderr, stderr)

//    outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
 //   go func() {
 //       var buf bytes.Buffer
 //       io.Copy(&buf, stdout)
 //       outC <- buf.String()
 //   }()

 //   outE := make(chan string)
 //   go func() {
 //       var buf bytes.Buffer
 //       io.Copy(&buf, stderr)
 //       outE <- buf.String()
 //   }()

	done := make(chan error, 1)
	go func() {
	    done <- cmd.Wait()
	}()
	select {
	case <-time.After(timeout * time.Second):
	    if err := cmd.Process.Kill(); err != nil {
	        log.Fatal("failed to kill: ", err)
	    }
	    log.Println("process killed as timeout reached")
	case err := <-done:
	    if err != nil {
	        log.Printf("process done with error = %v", err)
	    } else {
//	        log.Print("process done gracefully without error")
	    }
	}

//	out := <-outC

    // reading our temp stdout
//    fmt.Println("previous output:")
 //   fmt.Print(out)
//   outErr := <- outE
    // reading our temp stdout
 //   fmt.Println("previous error:")
 //   fmt.Print(outErr)
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
    sfi, err := os.Stat(src)
    if err != nil {
        return
    }
    if !sfi.Mode().IsRegular() {
        // cannot copy non-regular files (e.g., directories,
        // symlinks, devices, etc.)
        return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
    }
    dfi, err := os.Stat(dst)
    if err != nil {
        if !os.IsNotExist(err) {
            return
        }
    } else {
        if !(dfi.Mode().IsRegular()) {
            return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
        }
        if os.SameFile(sfi, dfi) {
            return
        }
    }
    if err = os.Link(src, dst); err == nil {
        return
    }
    err = copyFileContents(src, dst)
    return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
}

func GetFileType(filename string) (string) {
    extension := filepath.Ext(filename)
    return extension
}

func GetFilename(path string) (string) {
	_, file := filepath.Split(path)
	return file
}

func GetFileBaseName(basename string) (string) {
	name := strings.TrimSuffix(basename, filepath.Ext(basename))
	return name
}

func isFileSupported(fileExt string) bool {
	for _, v := range SupportedFileTypes {
		if v == strings.ToLower(fileExt) {
			return true
		}
	}
	return false
}

func isOfficeFile(fileExt string) bool {
	for _, v := range OfficeFileTypes {
		if v == strings.ToLower(fileExt) {
			return true
		}
	}
	return false
}

func isPdfFile(fileExt string) bool {
	return PDF == strings.ToLower(fileExt)
}

func isImageFile(fileExt string) bool {
	for _, v := range ImageFileTypes {
		if v == strings.ToLower(fileExt) {
			return true
		}
	}
	return false
}