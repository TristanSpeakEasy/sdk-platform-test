// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/tristanspeakeasy/sdk-platform-test/internal/utils"
)

type Metadata struct {
	Allergies            *string           `json:"allergies,omitempty"`
	AdditionalProperties map[string]string `additionalProperties:"true" json:"-"`
}

func (m Metadata) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *Metadata) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Metadata) GetAllergies() *string {
	if o == nil {
		return nil
	}
	return o.Allergies
}

func (o *Metadata) GetAdditionalProperties() map[string]string {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}
