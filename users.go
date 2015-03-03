package zego


type UserArray struct {
        Users       []User
}


type User struct {
        Id                    int64
        Url                   string
        Name                  string
        External_id           string
        Alias                 string
        Created_at            string
        Updated_at            string
        Active                bool
        Verified              bool
        Shared                bool
        Shared_agent          bool
        Locale                string
        Locale_id             int64
        Time_zone             string
        Last_login_at         string
        Email                 string
        Phone                 string
        Signature             string
        Details               string
        Notes                 string
        Organization_id       int64
        Role                  string
        Customer_role_id      int64
        Moderator             bool
        Ticket_restriction    string
        Only_private_comments bool
        Tags                  []string
        Restricted_agent      bool
        Suspended             bool
        Photo                 interface{}
        User_fields           interface{}
}



func (a Auth) ListUsers() (*Resource, error) {

	path := "/users.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) FilterUsers(filters string) (*Resource, error) {
	path := "/users.json?" + filters
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) ShowUser(user_id string) (*Resource, error) {

	path := "/users/" + user_id + ".json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) ShowUserRelated(user_id string) (*Resource, error) {

	path := "/users/" + user_id + "/related.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
