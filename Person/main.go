package main

import "fmt"

type contactInfo struct {
    email string
    zipCode string 
}

type person struct {
    firstName string
    lastName string
    contactInfo
}

func main() {
    brandon := person {
        firstName: "Brandon",
        lastName: "Manning",
        contactInfo: contactInfo {
            email: "manning0218@gmail.com",
            zipCode: "03055",
        },
    }

    brandon.print()

    newContact := &contactInfo {
        email: "manning0218@hotmail.com",
        zipCode: "18201",
    }

    brandon.updateContactInfo(newContact)

    brandon.print()
}

func (p person) print() {
    fmt.Println("Name:", p.firstName, p.lastName)
    fmt.Println("Contact:", p.contactInfo.email)
    fmt.Println("Postal Code:", p.contactInfo.zipCode)
}

func (p *person) updateContactInfo(c *contactInfo) {
    (*p).contactInfo = *c
}
