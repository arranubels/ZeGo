package zego

type Pagination struct {
        Count         int 
        Created       string 
        Next_page     string 
        Previous_page string 
}

type Via struct {
	Channel string
	Source Source
}

type Source struct {
	To SourceObject
	From SourceObject
	Rel string
}

type SourceObject struct {
	Address string
	Name string
	Id string
	Title string
	Topic_id string
	Topic_name string
	Original_recipients string
	Profile_url string
	Username string
}
