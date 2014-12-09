package zego

type CommentArray struct {
        Comments       []Comment
}

type Comment struct {
        Id                    uint32 
        Type                  string 
        Body                  string 
        Html_body             string 
        Public                bool 
        Author_id             uint32 
				Attachments           []Attachment
        Via                   interface{} 
        Metadata              interface{} 
        Created_at            string 
}

func (a Auth) GetTicketComments(ticket_id string) (*Resource, error) {
	path := "/tickets/" + ticket_id + "/comments.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}
	return resource, nil
}

