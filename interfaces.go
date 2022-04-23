package cdb

import (
	"cdb/contacts"
	"cdb/leads"
	"cdb/misc"
)

type Contacts interface {
	List(limit, offset int, query string) (*contacts.ContactsResponse, error)
	ByID(ID uint64) (*contacts.Contact, error)
}

type Leads interface {
	List(limit, offset int, query string) (*leads.LeadsResponse, error)
	ByContact(ID uint) (*leads.LeadsResponse, error)
	ByID(ID uint64) (*leads.Lead, error)
}

type Misc interface {
	Sources() ([]misc.Source, error)
	Users() ([]misc.User, error)
	Roles() ([]misc.Role, error)
}
