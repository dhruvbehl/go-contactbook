package contact

import (
	"fmt"

	"github.com/dhruvbehl/address-book/api"
	"github.com/rs/zerolog/log"
)

var contactBook = []api.Contact{{
	Id:        "100",
	FirstName: "MockFirstName",
	LastName:  "MockLastName",
	Phone:     "xxxxxxxxxx",
	Address:   "MockAddress",
}}

func GetAllContact() (*[]api.Contact, error) {
	if len(contactBook) == 0 {
		log.Error().Msgf("no contact found")
		return nil, fmt.Errorf("no contact found")
	}
	return &contactBook, nil
}

func GetContactById(id string) (*api.Contact, error) {
	for _, c := range contactBook {
		if c.GetId() == id {
			return &c, nil
		}
	}
	log.Error().Msgf("no contact found with id: [%v]", id)
	return nil, fmt.Errorf("no contact found with id: [%v]", id)
}

func DeleteContactById(id string) error {
	for i, c := range contactBook {
		if c.GetId() == id {
			contactBook = append(contactBook[:i], contactBook[i+1:]...)
			return nil
		}
	}
	log.Error().Msgf("no contact found with id: [%v]", id)
	return fmt.Errorf("no contact found with id: [%v]", id)
}

func UpsertContact(newContact api.Contact, id string) (*api.Contact, error) {
	if id == "" {
		for _, c := range contactBook {
			if c.GetId() == newContact.GetId() {
				log.Error().Msgf("contact already exists with the id: [%v]", newContact.GetId())
				return nil, fmt.Errorf("contact already exists with the id: [%v]", newContact.GetId())
			}
		}

		contactBook = append(contactBook, newContact)
		return &newContact, nil
	} else {
		for i, c := range contactBook {
			if c.GetId() == id && c.GetId() == newContact.GetId() {
				if newContact.GetFistName() != "" {
					c.SetFistName(newContact.GetFistName())
				}
				if newContact.GetLastName() != "" {
					c.SetLastName(newContact.GetLastName())
				}
				if newContact.GetPhone() != "" {
					c.SetPhone(newContact.GetPhone())
				}
				if newContact.GetAddress() != "" {
					c.SetAddress(newContact.GetAddress())
				}
				contactBook = append(append(contactBook[:i], c), contactBook[i+1:]...)
				return &c, nil
			}
		}
		log.Error().Msgf("no contact found with id: [%v]", newContact.GetId())
		return nil, fmt.Errorf("no contact found with id: [%v]", newContact.GetId())
	}
}
