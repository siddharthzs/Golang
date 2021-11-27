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

	SqlFileRegx, err := regexp.Compile(".*.sql")
	check(err)
	FileContentRegx, err := regexp.Compile("(FUNCTION|PROCEDURE|VIEW) dbo.([a-zA-Z0-9_]+)")
	check(err)

	sqlfile := make(map[string]string)
	duplicateEntry := make(map[string][]string)
	IncorrectFileName := make(map[string][]string)

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


			file, err := os.Open(path)
			check(err)
			b1 := make([]byte, 750)
			n1, err := file.Read(b1)
			if(err == nil){
				indexArr := FileContentRegx.FindStringIndex(string(b1[:n1]))
				functionName := string(b1[indexArr[0]+13:indexArr[1]]) + ".sql"
				if(functionName != info.Name()){
					fmt.Println(functionName,info.Name())
					IncorrectFileName[info.Name()] = append(IncorrectFileName[info.Name()], path)
				}
			}

			file.Close()
			// b1 := make([]byte, 750)
			
			// check(err3)
			// fmt.Println(n1)
			// indexArr := FileContentRegx.FindStringIndex(string(b1[:n1]))
			// functionName := string(b1[indexArr[0]+13:indexArr[1]]) + ".sql"
			// if(functionName != info.Name()){
			// 	IncorrectFileName[info.Name()] = append(IncorrectFileName[info.Name()], path)
			// }
		}
		
		return nil
	})

	if err != nil{
		log.Fatal(err)
	}

	for key, value := range duplicateEntry{
		fmt.Printf("\"%s : %s \";\n", key, strings.Join(value,`, `))
	}

	if(len(IncorrectFileName) > 0){
		println("\n============================================================;")
		for key, value := range IncorrectFileName{
			fmt.Printf("\"%s : %s \";\n", key, strings.Join(value,`, `))
		}
	}

	

	
	// f, err := os.Open("D:/learnGoWithMe/DatabaseCICD/two.sql")
    // check(err)
	// b1 := make([]byte, 750)
    // n1, err := f.Read(b1)
	// check(err)

	// arrval := FileContentRegx.FindStringIndex(string(b1[:n1]))
	// // fmt.Println(arrval[0])
    // fmt.Printf("%s\n",string(b1[arrval[0]+13:arrval[1]]))
}


func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}
