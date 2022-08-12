package fakeassign

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
)

type SomeStructWithTags struct {
	
	Name               string  `faker:"name"`
	Email              string  `faker:"email"`	
	Password           string  `faker:"password"`	
	PhoneNumber        string  `faker:"phone_number"`	
	UserName           string  `faker:"username"`
	
	

}

func Example_withTags() {

	a := SomeStructWithTags{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)