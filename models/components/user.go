// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type User struct {
	ID            string    `json:"id"`
	Email         string    `json:"email"`
	FirstName     *string   `json:"first_name,omitempty"`
	LastName      *string   `json:"last_name,omitempty"`
	Age           *float64  `json:"age,omitempty"`
	PostalCode    *string   `json:"postal_code,omitempty"`
	AssociatedIds []string  `json:"associatedIds,omitempty"`
	Metadata      *Metadata `json:"metadata,omitempty"`
}

func (o *User) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *User) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *User) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *User) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *User) GetAge() *float64 {
	if o == nil {
		return nil
	}
	return o.Age
}

func (o *User) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *User) GetAssociatedIds() []string {
	if o == nil {
		return nil
	}
	return o.AssociatedIds
}

func (o *User) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}
