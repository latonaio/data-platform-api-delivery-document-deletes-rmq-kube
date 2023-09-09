package requests

type Header struct {
	DeliveryDocument	 int     `json:"DeliveryDocument"`
	HeaderDeliveryStatus *string `json:"HeaderDeliveryStatus"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}
