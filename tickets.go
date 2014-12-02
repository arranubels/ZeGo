package zego

type TicketArray struct {
        Count         int `json:"count"`
        Created       string `json:"created"`
        Next_page     string `json:"next_page"`
        Previous_page string `json:"previous_page"`
        Tickets       []Ticket `json:"tickets"`
}

type Ticket struct {
        Id                    uint32 `json:"id"`
        Url                   string `json:"url"`
        External_id           string `json:"external_id"`
        Created_at            string `json:"created_at"`
        Updated_at            string `json:"updated_at"`
        Type                  string `json:"type"`
        Subject               string `json:"subject"`
        Raw_subject           string `json:"raw_subject"`
        Description           string `json:"description"`
        Priority              string `json:"priority"`
        Status                string `json:"status"`
        Recipient             string `json:"recipient"`
        Requester_id          uint32 `json:"requester_id"`
        Submitter_id          uint32 `json:"submitter_id"`
        Assignee_id           uint32 `json:"assignee_id"`
        Organization_id       uint32 `json:"organization_id"`
        Group_id              uint32 `json:"group_id"`
        Collaborator_ids      []uint32 `json:"collaborator_ids"`
        Forum_topic_id        uint32 `json:"forum_topic_id"`
        Problem_id            uint32 `json:"problem_id"`
        Has_incidents         bool `json:"has_incidents"`
        Due_at                string `json:"due_at"`
        Tags                  []string `json:"tags"`
        Satisfaction_rating   string `json:"satisfaction_rating"`
        Ticket_form_id        uint32 `json:"ticket_form_id"`
        Sharing_agreement_ids interface{} `json:"sharing_agreement_ids"`
        Via                   interface{} `json:"via"`
        Custom_Fields         interface{} `json:"custom_fields"`
        Fields                interface{} `json:"fields"`
}

func (a Auth) ListTickets() (*Resource, error) {

	path := "/tickets.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetTicket(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + ".json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetMultipleTickets(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + ".json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetTicketComments(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + "/comments.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) DeleteTicket(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + ".json"
	resource, err := api(a, "DELETE", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
