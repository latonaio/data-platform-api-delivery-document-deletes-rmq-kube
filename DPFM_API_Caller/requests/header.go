package requests

type Header struct {
	DeliveryDocumentID   int     `json:"DeliveryDocumentID"`
	HeaderDeliveryStatus *string `json:"HeaderDeliveryStatus"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}
