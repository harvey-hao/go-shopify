package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// ShopService is an interface for interfacing with the shop endpoint of the
// Shopify API.
// See: https://help.shopify.com/api/reference/shop
type ShopService interface {
	Get(options interface{}) (*Shop, error)
	GetShippingZones(options interface{}) ([]*ShippingZone, error)
}

// ShopServiceOp handles communication with the shop related methods of the
// Shopify API.
type ShopServiceOp struct {
	client *Client
}

// Shop represents a Shopify shop
type Shop struct {
	ID                              int64      `json:"id"`
	Name                            string     `json:"name"`
	ShopOwner                       string     `json:"shop_owner"`
	Email                           string     `json:"email"`
	CustomerEmail                   string     `json:"customer_email"`
	CreatedAt                       *time.Time `json:"created_at"`
	UpdatedAt                       *time.Time `json:"updated_at"`
	Address1                        string     `json:"address1"`
	Address2                        string     `json:"address2"`
	City                            string     `json:"city"`
	Country                         string     `json:"country"`
	CountryCode                     string     `json:"country_code"`
	CountryName                     string     `json:"country_name"`
	Currency                        string     `json:"currency"`
	Domain                          string     `json:"domain"`
	Latitude                        float64    `json:"latitude"`
	Longitude                       float64    `json:"longitude"`
	Phone                           string     `json:"phone"`
	Province                        string     `json:"province"`
	ProvinceCode                    string     `json:"province_code"`
	Zip                             string     `json:"zip"`
	MoneyFormat                     string     `json:"money_format"`
	MoneyWithCurrencyFormat         string     `json:"money_with_currency_format"`
	WeightUnit                      string     `json:"weight_unit"`
	MyshopifyDomain                 string     `json:"myshopify_domain"`
	PlanName                        string     `json:"plan_name"`
	PlanDisplayName                 string     `json:"plan_display_name"`
	PasswordEnabled                 bool       `json:"password_enabled"`
	PrimaryLocale                   string     `json:"primary_locale"`
	PrimaryLocationId               int64      `json:"primary_location_id"`
	Timezone                        string     `json:"timezone"`
	IanaTimezone                    string     `json:"iana_timezone"`
	ForceSSL                        bool       `json:"force_ssl"`
	TaxShipping                     bool       `json:"tax_shipping"`
	TaxesIncluded                   bool       `json:"taxes_included"`
	HasStorefront                   bool       `json:"has_storefront"`
	HasDiscounts                    bool       `json:"has_discounts"`
	HasGiftcards                    bool       `json:"has_gift_cards"`
	SetupRequire                    bool       `json:"setup_required"`
	CountyTaxes                     bool       `json:"county_taxes"`
	CheckoutAPISupported            bool       `json:"checkout_api_supported"`
	Source                          string     `json:"source"`
	GoogleAppsDomain                string     `json:"google_apps_domain"`
	GoogleAppsLoginEnabled          bool       `json:"google_apps_login_enabled"`
	MoneyInEmailsFormat             string     `json:"money_in_emails_format"`
	MoneyWithCurrencyInEmailsFormat string     `json:"money_with_currency_in_emails_format"`
	EligibleForPayments             bool       `json:"eligible_for_payments"`
	RequiresExtraPaymentsAgreement  bool       `json:"requires_extra_payments_agreement"`
}

// Represents the result from the admin/shop.json endpoint
type ShopResource struct {
	Shop *Shop `json:"shop"`
}

type ShippingZoneResource struct {
	ShippingZones []*ShippingZone `json:"shipping_zones"`
}

type ShippingZone struct {
	ID                       int64                              `json:"id"`
	Name                     string                             `json:"name"`
	ProfileId                string                             `json:"profile_id"`
	LocationGroupId          string                             `json:"location_group_id"`
	Countries                []*ShippingCountry                 `json:"countries"`
	Provinces                []*ShippingProvince                `json:"provinces"`
	PriceBasedShippingRates  []*ShippingPriceBasedShippingRate  `json:"price_based_shipping_rates"`
	WeightBasedShippingRates []*ShippingWeightBasedShippingRate `json:"weight_based_shipping_rates"`
}

type ShippingCountry struct {
	ID             int64               `json:"id"`
	ShippingZoneId int64               `json:"shipping_zone_id"`
	Name           string              `json:"name"`
	Tax            *decimal.Decimal    `json:"tax,omitempty"`
	Code           string              `json:"code"`
	TaxName        string              `json:"tax_name"`
	Provinces      []*ShippingProvince `json:"provinces"`
}

type ShippingProvince struct {
	ID             int64            `json:"id"`
	Code           string           `json"code"`
	CountryId      int64            `json:"country_id"`
	ShippingZoneId int64            `json:"shipping_zone_id"`
	Name           string           `json:"name"`
	Tax            *decimal.Decimal `json:"tax,omitempty"`
	TaxName        string           `json:"tax_name"`
	TaxType        string           `json:"tax_type"`
	TaxPercentage  *decimal.Decimal `json:"tax_percentage,omitempty"`
}

type ShippingPriceBasedShippingRate struct {
	ID               int64            `json:"id"`
	Name             string           `json:"name"`
	Price            *decimal.Decimal `json:"price,omitempty"`
	ShippingZoneId   int64            `json:"shipping_zone_id"`
	MinOrderSubtotal *decimal.Decimal `json:"min_order_subtotal,omitempty"`
	MaxOrderSubtotal *decimal.Decimal `json:"max_order_subtotal,omitempty"`
}
type ShippingWeightBasedShippingRate struct {
	ID             int64            `json:"id"`
	Name           string           `json:"name"`
	Price          *decimal.Decimal `json:"price"`
	ShippingZoneId int64            `json:"shipping_zone_id"`
	WeightLow      float64          `json:"weight_low"`
	WeightHigh     float64          `json:"weight_high"`
}

// Get shop
func (s *ShopServiceOp) Get(options interface{}) (*Shop, error) {
	resource := new(ShopResource)
	err := s.client.Get(fmt.Sprintf("%s/shop.json", globalApiPathPrefix), resource, options)
	return resource.Shop, err
}

// Get Shipping zones
func (s *ShopServiceOp) GetShippingZones(options interface{}) ([]*ShippingZone, error) {
	resource := new(ShippingZoneResource)
	err := s.client.Get(fmt.Sprintf("%s/shipping_zones.json", globalApiPathPrefix), resource, options)
	return resource.ShippingZones, err
}
