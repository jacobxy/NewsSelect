package main

import (
	"io"
	"bufio"
	"os"
	"strings"
	"fmt"
	"regexp"
)

const (
	RegexpMulu =`(^\d[ *\. * \d]+)(.*)`
)

var Mulu map[string]string 
func init(){
	Mulu = make(map[string]string, 40)
	ReadLine("mulu.txt")
}

func ReadLine(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		line = strings.TrimSpace(line)
		fmt.Println(line)
		r := regexp.MustCompile(RegexpMulu)
		matchs := r.FindStringSubmatch(line)
		if len(matchs) < 3{
			fmt.Println(matchs)
			continue
		}
		value := strings.Replace(matchs[1], " ", "", -1)
		content := strings.Split(matchs[2], " ")[0]
		Mulu[value] = content

	}
	return nil
}

type TitleAndContent struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

func ReadContent(fileName string)error{
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
		return err
	}
	buf := bufio.NewReader(f)
	title := ""
	res := ""
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		line = strings.TrimSpace(line)
		r := regexp.MustCompile(RegexpMulu)
		matchs := r.FindStringSubmatch(line)
		if len(res) < 1000 &&  len(res) +len(line) > 1000 && title != ""{ 
			test := AskNewsSummary(title,res)
			fmt.Println(test)
			res = line
		}

		if len(matchs) < 3 {
			res	 += line
			continue
		}


		value := strings.Replace(matchs[1], " ", "", -1)
		content := strings.Split(matchs[2], " ")[0]
		if v,ok := Mulu[value]; ok {
			test := AskNewsSummary(title,res)
			 fmt.Println(test)
			title = v
			res = content 
			return nil
		}
	}
}

func main(){
	ReadContent("content.txt")
}