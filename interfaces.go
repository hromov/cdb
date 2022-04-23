package cdb

type Contacts interface {
	List(limit, offset int, query string) (*ContactsResponse, error)
	ByID(ID uint64) (*Contact, error)
}

type Leads interface {
	List(limit, offset int, query string) (LeadsResponse, error)
	ByContact(ID uint) (*LeadsResponse, error)
	ByID(ID uint64) (*Lead, error)
}

type Misc interface {
	Sources() ([]Source, error)
	Users() ([]User, error)
	Roles() ([]Role, error)
}
