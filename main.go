package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"os"
	"io"
	"bufio"
)

func main() {
	inits()
}

func inits()  {
	readFile()
}

func ajax(name string) string {
	fmt.Printf(name)
	url := "http://www.**.com/student/list.json"
	str1 := strings.Replace(name, "\n", "", -1)
	str:="{\"searchKey\":\""+str1+"\"}"
	payload := strings.NewReader(str)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "0b7932ac-e9c3-4a18-c27a-68f02b85530d")
	req.Header.Add("Cookie", "_const_dession_id_=885ffc7170634e4bb3; path=/; domain=.**.com; Expires=Tue, 19 Jan 2018 03:14:07 GMT;")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//转换为json
	/*if json, err := gojson.NewJson([]byte(body)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(json.Get("data"))
	}*/
	//fmt.Println(res)
	//fmt.Println(string(body))
	return string(body)
}

func readFile(){
	//ch := make(chan string, 100000)
	inputFile, inputError := os.Open("/Users/zzx/test/a2.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		//fmt.Printf("The inputString input was: %s", inputString)
		//Ch <- inputString
		//fmt.Fprintln(inputString)
		strStudentInfo := ajax(inputString)
		bools :=  writeFile(strStudentInfo)
		if bools ==true{

		}
		//time.Sleep(time.Second * 20)
		if readerError == io.EOF {
			return
		}

	}
}

func writeFile(strStudentInfo string) bool{
	outputFile, outputError := os.OpenFile("/Users/zzx/test/a2.json", os.O_WRONLY|os.O_APPEND, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return false
	}

	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(strStudentInfo+",\n")
	outputWriter.Flush()
	outputFile.Close()

	return true
}

