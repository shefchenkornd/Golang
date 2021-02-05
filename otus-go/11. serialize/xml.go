package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Address struct {
	City string
	Postcode string
}

type Employee struct {
	Name string `xml:"name"`
	Age uint8 `xml:"age"`
	Position string `xml:"position"`
	Salary uint `xml:"salary"`
	Address
}

func main() {
	emp := &Employee{
		Name:     "Jack",
		Age:      56,
		Position: "QA",
		Salary:   300,
	}
	emp.Address = Address{
		City:     "Moscow",
		Postcode: "340000",
	}

	encXml, err := xml.Marshal(emp)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(encXml))
	/** output:
	<Employee><name>Jack</name><age>56</age><position>QA</position><salary>300</salary><City>Moscow</City><Postcode>340000</Postcode></Employee>
	*/

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent(" ", "    ")
	enc.Encode(emp)
	/** output:
	<Employee>
		<name>Jack</name>
		<age>56</age>
		<position>QA</position>
		<salary>300</salary>
		<City>Moscow</City>
		<Postcode>340000</Postcode>
	</Employee>
	*/
}
