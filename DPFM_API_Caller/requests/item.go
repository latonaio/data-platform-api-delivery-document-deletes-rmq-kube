package requests

type Item struct {
	DeliveryDocument    	int     `json:"DeliveryDocument"`
	DeliveryDocumentItem	int     `json:"DeliveryDocumentItem"`
	IsMarkedForDeletion 	*bool   `json:"IsMarkedForDeletion"`
}
