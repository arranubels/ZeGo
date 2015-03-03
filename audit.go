package zego

type AuditArray struct {
			  Pagination
        Audits       []Audit 
}

type Audit struct {
        Id                    int64
        Ticket_id             int64
        Author_id             int64
        Metadata              interface{}
        Created_at            string
        Events                []map[string]interface{}
}

type Event struct {
	Id	int64
	Type string
	Body string
	Html_body string
	Public bool
	Author_id	int64
	Attachments	[]Attachment
	Via	Via
	Metadata	interface{}
	Created_at	string
}

type ExternalTicketEvent struct {
	Id	int64
	Type string
	Resource string
	Body string
	Success string
}

type TicketShareEvent struct {
	Id	int64
	Type string
	Agreement_id int64
	Action string
}

type TweetEvent struct {
	Id	int64
	Type string
	Direct_message string
	Body string
	Recipients []int64
}

type SatisfactionRatingEvent struct {
	Id	int64
	Type string
	Score string
	Assignee_id int64
	Body string
}

type PushEvent struct {
	Id	int64
	Type string
	Value string
	Value_reference string
}

type LogMeInTranscriptEvent struct {
	Id	int64
	Type string
	Body string
}

type TicketErrorEvent struct {
	Id	int64
	Type string
	Message string
}

type TicketCCEvent struct {
	Id	int64
	Type string
	Recipients []int64
	Via	Via
}

type FacebookCommentEvent struct {
	Id	int64
	Type string
	Data string
	Body string
	Html_body string
	Public bool
	Trusted bool
	Author_id int64
	Graph_object_id int64
}

type FacebookEvent struct {
	Id	int64
	Type string
	Page string
	Communication string
	Ticket_via string
	Body string
}

type NotificationEvent struct {
	Id	int64
	Type string
	Subject string
	Body string
	Recipients []int64
	Via	Via
}

type TicketCreateEvent struct {
	Id	int64
	Type string
	Field_name string
	Value string
}

type TicketChangeEvent struct {
	Id	int64
	Type string
	Field_name string
	Value string
	Previous_value string
}

type TicketCommentPrivacyChangeEvent struct {
	Id	int64
	Type string
	Public bool
	Comment_id	int64
}

type VoiceCommandEvent struct {
	Id	int64
	Type string
	Data string
	Body string
	Formatted_from string
	Formatted_to string
	Html_body string
	Public bool
	Trusted bool
	Author_id	int64
	Attachments	[]Attachment
}

func (a Auth) GetTicketAudit(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + "/audits.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetTicketAuditById(ticket_id string, audit_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + "/audits/"+audit_id+".json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

