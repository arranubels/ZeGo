package zego

type AuditArray struct {
			  Pagination
        Audits       []Audit 
}

type Audit struct {
        Id                    uint64 
        Ticket_id             uint64 
        Author_id             uint64 
				Metadata              []map[string]interface{}
				Created_at            string
				Events                []map[string]interface{}
}

type Event struct {
	Id	uint64
	Type string
	Body string
	Html_body string
	Public bool
	Author_id	uint64
	Attachments	[]Attachment
	Via	Via
	Metadata	interface{}
	Created_at	string
}

type ExternalTicketEvent struct {
	Id	uint64
	Type string
	Resource string
	Body string
	Success string
}

type TicketShareEvent struct {
	Id	uint64
	Type string
	Agreement_id uint64
	Action string
}

type TweetEvent struct {
	Id	uint64
	Type string
	Direct_message string
	Body string
	Recipients []uint64
}

type SatisfactionRatingEvent struct {
	Id	uint64
	Type string
	Score string
	Assignee_id uint64
	Body string
}

type PushEvent struct {
	Id	uint64
	Type string
	Value string
	Value_reference string
}

type LogMeInTranscriptEvent struct {
	Id	uint64
	Type string
	Body string
}

type TicketErrorEvent struct {
	Id	uint64
	Type string
	Message string
}

type TicketCCEvent struct {
	Id	uint64
	Type string
	Recipients []uint64
	Via	Via
}

type FacebookCommentEvent struct {
	Id	uint64
	Type string
	Data string
	Body string
	Html_body string
	Public bool
	Trusted bool
	Author_id uint64
	Graph_object_id uint64
}

type FacebookEvent struct {
	Id	uint64
	Type string
	Page string
	Communication string
	Ticket_via string
	Body string
}

type NotificationEvent struct {
	Id	uint64
	Type string
	Subject string
	Body string
	Recipients []uint64
	Via	Via
}

type TicketCreateEvent struct {
	Id	uint64
	Type string
	Field_name string
	Value string
}

type TicketChangeEvent struct {
	Id	uint64
	Type string
	Field_name string
	Value string
	Previous_value string
}

type TicketCommentPrivacyChangeEvent struct {
	Id	uint64
	Type string
	Public bool
	Comment_id	uint64
}

type VoiceCommandEvent struct {
	Id	uint64
	Type string
	Data string
	Body string
	Formatted_from string
	Formatted_to string
	Html_body string
	Public bool
	Trusted bool
	Author_id	uint64
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

