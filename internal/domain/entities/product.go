package entities

type Product struct {
	// omitempty not present when it zero values 
    ID          string  `json:"_id,omitempty"`      // CouchDB uses _id for document ID
    Rev         string  `json:"_rev,omitempty"`     // rev is  for document revision (in couchdb needed it)
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
}
