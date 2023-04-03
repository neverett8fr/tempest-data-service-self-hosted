package storage

import application "tempest-data-service/pkg/application/entities"

type postNewItem struct {
	Item application.File `json:"item"`
}

type queryParams struct {
	User string `json:"username"`
}

type queryItemParam struct {
	User string `json:"username"`
	Key  string `json:"key"`
}

type queryAllItems struct {
	Query []queryParams `json:"query"`
}

type querySpecificItem struct {
	Query []queryItemParam `json:"query"`
}

func newQueryAllItems(param queryParams) queryAllItems {
	return queryAllItems{
		Query: []queryParams{
			param,
		},
	}
}

func newQuerySpecificItem(param queryItemParam) querySpecificItem {
	return querySpecificItem{
		Query: []queryItemParam{
			param,
		},
	}
}

type queryBodyReponse struct {
	Items []application.File `json:"items"`
}
