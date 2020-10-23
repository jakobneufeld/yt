package main
import (
	"os"
	"os/exec"
	"fmt"
	"strings"
	"github.com/fatih/color"

)
func main() {
	if os.Args[1] == "help" {
		fmt.Println("Yt a youtube cli. Put arg 1 for channel name and arg 2 for a search keyword. Only works on macos")
		return
	}
	if os.Args[1] == "--netflix" {
		if len(os.Args) == 3 {
			var httpsBuilder strings.Builder
			httpsBuilder.WriteString("https://netflix.com/")
			httpsBuilder.WriteString("/search?q=")
			httpsBuilder.WriteString(os.Args[2])
			fmt.Println("The url is. I will automatically launch it for you",httpsBuilder.String())
			safaricmd := exec.Command("open", "-a", "Firefox", httpsBuilder.String())
			if safaricmd.Run() == nil {
				color.Green("Succes")
				os.Exit(0)
			} else {
				color.Red("Whoops! Bad url is", safaricmd.String())
			}
		
		} else {
			fmt.Println("Not enough args or too many keywords. One only please.")
			os.Exit(1)
	
		}
	} 

	if len(os.Args) == 3 {
		var httpsBuilder strings.Builder
		httpsBuilder.WriteString("https://youtube.com/c/")
		httpsBuilder.WriteString(os.Args[1])
		httpsBuilder.WriteString("/search?query=")
		httpsBuilder.WriteString(os.Args[2])
		fmt.Println("The url is. I will automatically launch it for you",httpsBuilder.String())
		safaricmd := exec.Command("open", "-a", "Firefox", httpsBuilder.String())
		if safaricmd.Run() == nil {
			color.Green("Succes")
			os.Exit(0)
		} else {
			color.Red("Whoops! Bad url is", safaricmd.String())
		}
	
	} else {
		fmt.Println("Not enough args or too many keywords. One only please.")
		os.Exit(1)

	}
	
}