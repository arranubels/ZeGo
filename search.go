package zego

import (
		"net/url"
		)

type ResultsArray struct {
			  Pagination
        Results       []Ticket 
}

func (a Auth) Search(query string) (*Resource, error) {

	path := "/search.json?query=" + url.QueryEscape(query)
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) AnonymousSearch(query string) (*Resource, error) {

	path := "/portal/search.json?query=" + url.QueryEscape(query)
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

