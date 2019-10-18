package main

import (
	"fmt"
	"io/ioutil"
	"log"
	simplepb "protocol_buffers/protobuf-golang/src/simple"

	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()

	readAndWrite(sm)
}

func readAndWrite(sm proto.Message) {
	fileName := "protobuf-golang/simple.bin"
	writeToFile(fileName, sm)

	sm2 := &simplepb.SimpleMessage{}
	readFromFile(fileName, sm2)

	fmt.Println("Content from file:", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("Cannot serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Cannot write to file!", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err)
		return err2
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "New name"

	fmt.Println(sm)

	fmt.Println("The ID is:", sm.GetId())

	return &sm
}
