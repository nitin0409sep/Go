package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	// log the err
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat() // Returns File Info or Error in Path

	// if err != nil {
	// 	panic(err);
	// }

	// // fmt.Println(fileInfo);
	// fmt.Println("File Name", fileInfo.Name());
	// fmt.Println("File or Folder", fileInfo.IsDir());
	// fmt.Println("File Size", fileInfo.Size());
	// fmt.Println("File Permission", fileInfo.Mode());
	// fmt.Println("File Modified At", fileInfo.ModTime());

	//! How to read a File

	// Method 1
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }
	
	// defer f.Close()
	
	// buffer := make([]byte, 13)
	
	// d, err := f.Read(buffer)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := range buffer{
	// 	println("data", d, string(buffer[i]))
	// }

	data, err := os.ReadFile("example.txt")

	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(data));
	
	// Read Folders
	dir, err := os.Open(".");
	
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	fileInfo, err := dir.ReadDir(1); // Number of dir u waana list down, if value is less or equal to 0 then it will list all dir's

	for _, fi := range fileInfo {
		fmt.Println(fi.Name())
	}

	// Create a File
	f, err := os.Create("example2.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	f.WriteString("Hey Go!");
	f.WriteString("Nice Language");

	// Write File
	bytes := []byte("Hello Nitin") // File is nothing but array of bytes
	f.Write(bytes);


	//! Very Important -> Read a file from one place and write it to another file -> Streaming -> Dont load whole data at once
	sourceFile, err := os.Open("example.txt");
	
	if(err != nil){
		panic(err)
	}
	
	defer sourceFile.Close()
	
	destFile, err := os.Create("example1.txt")
	
	if(err != nil){
		panic(err)
	}

	defer destFile.Close()


	// For Streaming -> buferio -> Inbuild package -> Provides Prdefined Buffer
	reader := bufio.NewReader(sourceFile);
	writer := bufio.NewWriter(destFile);

	// Pipe
	for {
		b, err := reader.ReadByte()

		if(err != nil) {
			if(err.Error() != "EOF") {
				panic(err)
			}

			// EOF -> End of File
			break;
		}

		e := writer.WriteByte(b);

		if e != nil {
			panic(nil)
		}
	}	
	

	writer.Flush()

	fmt.Println("Written to new File Successfully")

}
