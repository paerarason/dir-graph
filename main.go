package main
import (
	"log"
	"fmt"
	"os/exec"
	"os"
	 "path/filepath"
	"strings"
)

func check(lines string){	
		cmd := exec.Command("pwd")
		dir, err := cmd.Output()
        currentDir := strings.TrimSpace(string(dir))
	    cmd1 := exec.Command("ls")
		value,err:=hasExecutable(currentDir+"/"+lines)
		//fmt.Println(currentDir+"/"+lines)
		if value || err!=nil{
			return 
		}
	    cmd1.Dir = currentDir+"/"+lines
        //fmt.Println(currentDir)
	    outputs, err := cmd1.CombinedOutput()

	if err != nil {
		log.Fatal("ERROR \n ",currentDir+"/"+lines,"\n", err)
	}
	
	//fmt.Println("string=>",string(outputs))
	line := strings.Split(string(outputs), "\n")
	//fmt.Println(line)
    for _, folder:=range line{

		fold:=strings.Split(string(folder), ".")
		//fmt.Println(fold)	
		//fmt.Println(string(folder))
		if len(fold)==1 {
           fmt.Println("|__",string(fold[0]))
		   check(string(fold[0]))
		}else  {	
		   fmt.Println("  |__",folder)
		   break
		}
	 }
    	
	}

func hasExecutable(dir string) (bool, error) {
    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.Mode().IsRegular() && info.Mode().Perm()&0111 != 0 { // Check for executable permission
            return fmt.Errorf("found executable file: %s", path)
        }

        return nil
    })

    if err == nil {
        return false, nil // No executable found
    }

    // Check if error is due to an executable being found
    if pathError, ok := err.(*os.PathError); ok && pathError.Err == os.ErrPermission {
        return true, nil // Executable found
    }

    return false, err // Other error
}


func run(){
	cmd := exec.Command("ls")
 	output, err := cmd.Output()
     if err != nil {
         log.Fatal("Error in Find",err)
     }	
    
	lines := strings.Split(string(output), "\n")
    for _, folder:=range lines{
		fold:=strings.Split(string(folder), ".")

		//fmt.Println(fold)
		if len(fold)==1 && string(fold[0])!="" {
           fmt.Println(string(fold[0]))
		   check(string(fold[0]))
		}else if string(folder)!="" && fold[1]!="exe"{
		   fmt.Println("|__",folder)
          //check(string(fold[0]))
		}
	}
}


func main(){
    run()
}