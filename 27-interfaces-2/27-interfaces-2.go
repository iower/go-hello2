package main

import "fmt"

type File struct {
	text string
}
type Folder struct{}

type Stream interface {
	read() string
	write(string)
	close()
}

func writeToStream(stream Stream, text string) {
	stream.write(text)
}

func closeStream(stream Stream) {
	stream.close()
}

// implement methods for *File
func (f *File) read() string {
	return f.text
}
func (f *File) write(message string) {
	f.text = message
	fmt.Println("Written:", message)
}
func (f *File) close() {
	fmt.Println("Closed")
}

// implement methods for *Folder
func (f *Folder) close() {
	fmt.Println("Closed")
}

func main() {
	myFile := &File{}
	myFolder := &Folder{}

	writeToStream(myFile, "Hello")
	closeStream(myFile)

	// closeStream(myFolder) // *Folder does not implement Stream
	myFolder.close()
}
