package main

import (
	"fmt"
	"os"
	"path/filepath"
	"log"
	"regexp"
	"strings"
)

func main(){
	mydir, err := os.Getwd()
	if err != nil{
		fmt.Println(err)
	}
	idx := strings.Index(mydir,"DatabaseCICD")
	mydir = mydir[:idx+12]
	// fmt.Println(mydir)

	SqlFileRegx, err := regexp.Compile(".*.sql")
	if err != nil{
		log.Fatal(err)
	}
	sqlfile := make(map[string]string)
	duplicateEntry := make(map[string][]string)


	err = filepath.Walk(mydir, func(path string, info os.FileInfo, err error) error{
		if (err == nil && SqlFileRegx.MatchString(info.Name())){
			if val, ok := sqlfile[info.Name()]; ok {  // if present

				// send to duplicateEntry
				if _, ok := duplicateEntry[info.Name()]; !ok {
					duplicateEntry[info.Name()] = append(duplicateEntry[info.Name()], val)
				}
				duplicateEntry[info.Name()] = append(duplicateEntry[info.Name()], path)
			}else{
				sqlfile[info.Name()] = path
			}
		}
		return nil
	})

	if err != nil{
		log.Fatal(err)
	}

	for key, value := range duplicateEntry{
		fmt.Printf("\"%s : %s \";\n", key, strings.Join(value,`, `))
		// fmt.Println("\"",key,":",strings.Join(value,`, `),"\";")
	}
}