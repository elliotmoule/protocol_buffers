package main

import (
	"fmt"
	"io/ioutil"
	"log"
	complexpb "protocol_buffers/protobuf-golang/src/complex"
	enumpb "protocol_buffers/protobuf-golang/src/enums"
	simplepb "protocol_buffers/protobuf-golang/src/simple"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()

	readAndWrite(sm)
	jsonFunctions(sm)

	doEnum()

	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	fmt.Println("Successfully initialised complex proto:", cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY

	fmt.Println("Successfully initialised enumeration proto:", em)
}

func jsonFunctions(sm proto.Message) {
	smAsString := toJSON(sm)

	fmt.Println("Simple Message as JSON:", smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)

	if err != nil {
		log.Fatalln("Cannot convert to JSON", err)
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)

	if err != nil {
		log.Fatalln("Could not unmarshal JSON into the pb struct", err)
	}
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

	fmt.Println("Initialised Simple Message:", sm)

	sm.Name = "New name"

	fmt.Println("'Name' Field changed:", sm)

	fmt.Println("The ID is:", sm.GetId())

	return &sm
}
