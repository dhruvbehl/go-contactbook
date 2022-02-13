package api

type Contact struct {
	Id			string 		`json:"id"`
	FirstName	string		`json:"firstname"`
	LastName	string		`json:"lastname"`
	Phone		string		`json:"phone"`
	Address		string		`json:"address"`
}

func (c *Contact) GetId() string {return c.Id}
func (c *Contact) GetFistName() string {return c.FirstName}
func (c *Contact) GetLastName() string {return c.LastName}
func (c *Contact) GetPhone() string {return c.Phone}
func (c *Contact) GetAddress() string {return c.Address}

func (c *Contact) SetId(id string) {c.Id = id}
func (c *Contact) SetFistName(firstName string) {c.FirstName = firstName}
func (c *Contact) SetLastName(lastName string) {c.LastName = lastName}
func (c *Contact) SetPhone(phone string) {c.Phone = phone}
func (c *Contact) SetAddress(address string) {c.Address = address}