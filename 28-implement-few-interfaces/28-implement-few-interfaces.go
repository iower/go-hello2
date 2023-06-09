package main

import "fmt"

type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

func writeToStream(writer Writer, text string) {
	writer.write(text)
}

func readFromStream(reader Reader) {
	reader.read()
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

func main() {
	myFile := &File{}
	writeToStream(myFile, "Hello")
	readFromStream(myFile)
}
