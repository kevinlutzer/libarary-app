package shared

type CollectionCreateRequest struct {
	Name    string   `json:"name" binding:"required"`
	BookIDs []string `json:"bookIDs"`
}

func (req *CollectionCreateRequest) Validate() error {
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
	Name    string   `json:"name" `
	BookIDs []string `json:"bookIDs"`
}

type CollectionUpdateRequest struct {
	ID        string          `json:"id" binding:"required"`
	Data      *CollectionData `json:"data"`
	FieldMask []string        `json:"fieldMask"`
}

func (req *CollectionUpdateRequest) Validate() error {
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
	ID    string    `json:"id" binding:"required"`
	Name  string    `json:"title" binding:"required"`
	Books []ApiBook `json:"books"`
}

//
// Responses
//

type CollectionCreateResponseData struct {
	ID string `json:"id" binding:"required"`
}

// Swag definition
type CollectionCreateResponse = ApiResponse[CollectionCreateResponseData]

type CollectionGetResponseData struct {
	Collections []ApiCollection `json:"collections" binding:"required"`
}

// Swag definition
type CollectionGetResponse = ApiResponse[CollectionGetResponseData]
