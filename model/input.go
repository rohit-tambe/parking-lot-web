package model

// SlotRequest slot assign request
type SlotRequest struct {
	Slot int `json:"slot" validate:"required"`
}

// CarDetails request
type CarDetails struct {
	RagistrationNo string `json:"ragistrationNo" validate:"required"`
	Colour         string `json:"colour" validate:"required"`
}

// ResponseDTO api response
type ResponseDTO struct {
	Status  string   `json:"status,omitempty"`
	Reason  string   `json:"reason,omitempty"`
	Result  string   `json:"result,omitempty"`
	Details []Detail `json:"details,omitempty"`
}

// Detail muster details
type Detail struct {
	Slot           int    `json:"slot"`
	RagistrationNo string `json:"carNumber"`
	Colour         string `json:"carColour"`
}

// GetFailResponse fail response with error
func (v *ResponseDTO) GetFailResponse(err error) {
	v.Status = "Failed"
	if err != nil {
		v.Reason = err.Error()
	}
}
