package zego

type TicketArray struct {
			  Pagination
        Tickets       []Ticket 
}

type Ticket struct {
        Id                    uint32 
        Url                   string 
        External_id           string 
        Created_at            string 
        Updated_at            string 
        Type                  string 
        Subject               string 
        Raw_subject           string 
        Description           string 
        Priority              string 
        Status                string 
        Recipient             string 
        Requester_id          uint32 
        Submitter_id          uint32 
        Assignee_id           uint32 
        Organization_id       uint32 
        Group_id              uint32 
        Collaborator_ids      []uint32 
        Forum_topic_id        uint32 
        Problem_id            uint32 
        Has_incidents         bool 
        Due_at                string 
        Tags                  []string 
        Satisfaction_rating   string 
        Ticket_form_id        uint32 
        Sharing_agreement_ids interface{} 
        Via                   interface{} 
        Custom_Fields         interface{} 
        Fields                interface{} 
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

func (a Auth) DeleteTicket(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + ".json"
	resource, err := api(a, "DELETE", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
