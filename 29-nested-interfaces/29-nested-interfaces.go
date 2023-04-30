package main

import "fmt"

type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

type ReaderWriter interface {
	Reader
	Writer
}

type File struct {
	text string
}

func (f *File) read() {
	fmt.Println(f.text)
}
func (f *File) write(message string) {
	f.text = message
	fmt.Println("Written", message)
}

func writeToStream(writter ReaderWriter, text string) {
	writter.write(text)
}
func readFromStream(reader ReaderWriter) {
	reader.read()
}

func main() {
	myFile := &File{}
	writeToStream(myFile, "Hello")
	readFromStream(myFile)
	writeToStream(myFile, "Hello2")
	readFromStream(myFile)
}
