package zego

type TicketArray struct {
			  Pagination
        Tickets       []Ticket 
}

type Ticket struct {
        Id                    int64
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
        Requester_id          int64
        Submitter_id          int64
        Assignee_id           int64
        Organization_id       int64
        Group_id              int64
        Collaborator_ids      []int64
        Forum_topic_id        int64
        Problem_id            int64
        Has_incidents         bool 
        Due_at                string 
        Tags                  []string 
        Satisfaction_rating   string 
        Ticket_form_id        int64
        Sharing_agreement_ids interface{} 
        Via                   interface{} 
        Custom_Fields         []IdValue
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
