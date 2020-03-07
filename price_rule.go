package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const priceRulesBasePath = "price_rules"

// PriceRuleService is an interface for interfacing with the discount endpoints
// of the Shopify API.
// See: https://shopify.dev/docs/admin-api/rest/reference/discounts/pricerule
type PriceRuleService interface {
	Get(int64) (*PriceRule, error)
}

// PriceRuleServiceOp handles communication with the discount code
// related methods of the Shopify API.
type PriceRuleServiceOp struct {
	client *Client
}

type PriceRule struct {
	ID                                     int64              `json:"id,omitempty"`
	Title                                  string             `json:"title,omitempty"`
	ValueType                              string             `json:"value_type,omitempty"`
	Value                                  *decimal.Decimal   `json:"value,omitempty"`
	CustomerSelection                      string             `json:"customer_selection,omitempty"`
	TargetType                             string             `json:"target_type,omitempty"`
	TargetSelection                        string             `json:"target_selection,omitempty"`
	AllocationMethod                       string             `json:"allocation_method,omitempty"`
	AllocationLimit                        string             `json:"allocation_limit,omitempty"`
	OncePerCustomer                        bool               `json:"once_per_customer,omitempty"`
	UsageLimit                             int32              `json:"usage_limit,omitempty"`
	StartsAt                               *time.Time         `json:"starts_at,omitempty"`
	EndsAt                                 *time.Time         `json:"ends_at,omitempty"`
	CreatedAt                              *time.Time         `json:"created_at,omitempty"`
	UpdatedAt                              *time.Time         `json:"updated_at,omitempty"`
	EntitledProductIds                     []int64            `json:"entitled_product_ids,omitempty"`
	EntitledVariantIds                     []int64            `json:"entitled_variant_ids,omitempty"`
	EntitledCollectionIds                  []int64            `json:"entitled_collection_ids,omitempty"`
	EntitledCountryIds                     []int64            `json:"entitled_country_ids,omitempty"`
	PrerequisiteProductIds                 []int64            `json:"prerequisite_product_ids,omitempty"`
	PrerequisiteVariantIds                 []int64            `json:"prerequisite_variant_ids,omitempty"`
	PrerequisiteCollectionIds              []int64            `json:"prerequisite_collection_ids,omitempty"`
	PrerequisiteSavedSearchIds             []int64            `json:"prerequisite_product_ids,omitempty"`
	PrerequisiteCustomerIds                []int64            `json:"prerequisite_customer_ids,omitempty"`
	PrerequisiteSubtotalRange              *PrerequisiteRange `json:"prerequisite_subtotal_range,omitempty"`
	PrerequisiteQuantityRange              *PrerequisiteRange `json:"prerequisite_quantity_range,omitempty"`
	PrerequisiteShippingPriceRange         *PrerequisiteRange `json:"prerequisite_shipping_price_range,omitempty"`
	PrerequisiteToEntitlementQuantityRatio *PrerequisiteRatio `json:"prerequisite_to_entitlement_quantity_ratio,omitempty"`
}

type PrerequisiteRange struct {
	GreaterThanOrEqualTo string `json:"greater_than_or_equal_to,omitempty"`
	LessThanOrEqualTo    string `json:"less_than_or_equal_to,omitempty"`
}

type PrerequisiteRatio struct {
	PrerequisiteQuantity int32 `json:"prerequisite_quantity,omitempty"`
	EntitledQuantity     int32 `json:"entitled_quantity,omitempty"`
}

type PriceRuleResource struct {
	PriceRule *PriceRule `json:"price_rule"`
}

// Get a single discount code
func (s *PriceRuleServiceOp) Get(priceRuleID int64) (*PriceRule, error) {
	path := fmt.Sprintf("%s/"+priceRulesBasePath+"/%d.json", globalApiPathPrefix, priceRuleID)
	resource := new(PriceRuleResource)
	err := s.client.Get(path, resource, nil)
	return resource.PriceRule, err
}
