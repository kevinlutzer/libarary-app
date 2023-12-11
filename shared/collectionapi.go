package shared

type CollectionPutRequest struct {
	Name    string   `json:"name"`
	BookIDs []string `json:"bookIDs"`
}

func (req *CollectionPutRequest) Validate() error {
	if req.Name == "" || len(req.Name) > 512 {
		return NewError(InvalidArguments, "name is required and must be less than 512 characters")
	}

	if len(req.BookIDs) > 0 {
		for _, id := range req.BookIDs {
			if !IsValidID(id) {
				return NewError(InvalidArguments, "an id specified is not a valid id")
			}
		}
	}

	return nil
}

type CollectionData struct {
	Name    string   `json:"name"`
	BookIDs []string `json:"books"`
}

type CollectionPostRequest struct {
	ID        string          `json:"id"`
	Data      *CollectionData `json:"data"`
	FieldMask []string        `json:"fieldMask"`
}

func (req *CollectionPostRequest) Validate() error {
	if req.ID == "" {
		return NewError(InvalidArguments, "id is required")
	}

	if req.Data == nil {
		req.Data = &CollectionData{}
	}

	return nil
}

//
// ApiCollection Definition
//

type ApiCollection struct {
	ID    string    `json:"id"`
	Name  string    `json:"title"`
	Books []ApiBook `json:"books"`
}

//
// Responses
//

type CollectionPutResponse struct {
	ID string `json:"id"`
}

type CollectionLoadResponse struct {
	Collections []ApiCollection `json:"collections"`
}
