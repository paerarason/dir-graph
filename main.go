package main
import (
	"log"
	"fmt"
	"os/exec"
	"strings"
)

func check(lines string){
	    
		fold:=strings.Split(string(lines), ".")
        if len(fold)>2{
            fmt.Println(fold,"->",fold)
		}else{	 
		       check(string(fold[0]))
			}	
	}
func main(){

    cmd := exec.Command("/")
 	output, err := cmd.Output()
     if err != nil {
         log.Fatal("Error in Find",err)
     }	 


	lines := strings.Split(string(output), "\n")
    for _, folder:=range lines{
		fold:=strings.Split(string(folder), ".")
		//fmt.Println(fold)
		if len(fold)==1{
           fmt.Println(string(fold[0]))
		   //check(string(fold[0]))
		}else{
		fmt.Println("|__",folder)
          //check(string(fold[0]))
		}
	}
}