package items

type Item struct {
	ID                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Pictures          []Picture   `json:"pictures"`
	Description       Description `json:"description"`
	Video             string      `json:"video"`
	Price             float64     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Picture struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

type Description struct {
	PlaintText string `json:"plain_text"`
	HTML       string `json:"html"`
}
