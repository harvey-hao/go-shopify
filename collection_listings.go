package goshopify

import (
	"fmt"
	"time"
)

const collectionListingsBasePath = "collection_listings"

// CollectService is an interface for interfacing with the collect endpoints
// of the Shopify API.
// See: https://shopify.dev/docs/admin-api/rest/reference/sales-channels/collectionlisting
type CollectionListingService interface {
	List(interface{}) ([]Collection, error)
	Get(int64, interface{}) (*Collection, error)
}

// CollectServiceOp handles communication with the collect related methods of
// the Shopify API.
type CollectionListingServiceOp struct {
	client *Client
}

// Collect represents a Shopify collect
type Collection struct {
	CollectionID        int64              `json:"collection_id,omitempty"`
	BodyHtml            string             `json:"body_html,omitempty"`
	DefaultProductImage []*CollectionImage `json:"default_product_image,omitempty"`
	Image               *CollectionImage   `json:"image,omitempty"`
	Handle              string             `json:"handle,omitempty"`
	Title               string             `json:"title,omitempty"`
	SortOrder           string             `json:"sort_order,omitempty"`
	PublishedAt         *time.Time         `json:"published_at,omitempty"`
	UpdatedAt           *time.Time         `json:"updated_at,omitempty"`
}

type CollectionImage struct {
	Src string `json:"src,omitempty"`
}

// Represents the result from the collection_listings/X.json endpoint
type CollectionListingResource struct {
	Collection *Collection `json:"collection_listing"`
}

// Represents the result from the collection_listings.json endpoint
type CollectionListingsResource struct {
	Collections []Collection `json:"collection_listings"`
}

// List collects
func (s *CollectionListingServiceOp) List(options interface{}) ([]Collection, error) {
	path := fmt.Sprintf("%s/%s.json", globalApiPathPrefix, collectionListingsBasePath)
	resource := new(CollectionListingsResource)
	err := s.client.Get(path, resource, options)
	return resource.Collections, err
}

// Get individual custom collection
func (s *CollectionListingServiceOp) Get(collectionID int64, options interface{}) (*Collection, error) {
	path := fmt.Sprintf("%s/%s/%d.json", globalApiPathPrefix, collectionListingsBasePath, collectionID)
	resource := new(CollectionListingResource)
	err := s.client.Get(path, resource, options)
	return resource.Collection, err
}
