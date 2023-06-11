package Observer

import "fmt"

type Customer struct {
	ID string
}

func (c *Customer) Update(str string) {
	fmt.Printf("Sending email to customer %s for item %s \n", c.GetID(), str)
}

func (c *Customer) GetID() string {
	return c.ID
}
