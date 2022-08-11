package odoo

import (
	"fmt"
)

// ProductProduct represents product.product model.
type ProductProduct struct {
	LastUpdate                               *Time      `xmlrpc:"__last_update,omptempty"`
	Active                                   *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline                     *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityIds                              *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                            *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                          *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                           *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                           *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AlertTime                                *Int       `xmlrpc:"alert_time,omptempty"`
	AssetCategoryId                          *Many2One  `xmlrpc:"asset_category_id,omptempty"`
	AttributeLineIds                         *Relation  `xmlrpc:"attribute_line_ids,omptempty"`
	AttributeValueIds                        *Relation  `xmlrpc:"attribute_value_ids,omptempty"`
	AvailableInPos                           *Bool      `xmlrpc:"available_in_pos,omptempty"`
	Barcode                                  *String    `xmlrpc:"barcode,omptempty"`
	CategId                                  *Many2One  `xmlrpc:"categ_id,omptempty"`
	Code                                     *String    `xmlrpc:"code,omptempty"`
	Color                                    *Int       `xmlrpc:"color,omptempty"`
	CompanyId                                *Many2One  `xmlrpc:"company_id,omptempty"`
	CostCurrencyId                           *Many2One  `xmlrpc:"cost_currency_id,omptempty"`
	CostMethod                               *String    `xmlrpc:"cost_method,omptempty"`
	CreateDate                               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                               *Many2One  `xmlrpc:"currency_id,omptempty"`
	DefaultCode                              *String    `xmlrpc:"default_code,omptempty"`
	DeferredRevenueCategoryId                *Many2One  `xmlrpc:"deferred_revenue_category_id,omptempty"`
	DeliveryCount                            *Int       `xmlrpc:"delivery_count,omptempty"`
	Description                              *String    `xmlrpc:"description,omptempty"`
	DescriptionPicking                       *String    `xmlrpc:"description_picking,omptempty"`
	DescriptionPickingin                     *String    `xmlrpc:"description_pickingin,omptempty"`
	DescriptionPickingout                    *String    `xmlrpc:"description_pickingout,omptempty"`
	DescriptionPurchase                      *String    `xmlrpc:"description_purchase,omptempty"`
	DescriptionSale                          *String    `xmlrpc:"description_sale,omptempty"`
	DisplayName                              *String    `xmlrpc:"display_name,omptempty"`
	EventOk                                  *Bool      `xmlrpc:"event_ok,omptempty"`
	EventTicketIds                           *Relation  `xmlrpc:"event_ticket_ids,omptempty"`
	EventTypeId                              *Many2One  `xmlrpc:"event_type_id,omptempty"`
	ExpensePolicy                            *Selection `xmlrpc:"expense_policy,omptempty"`
	HideExpensePolicy                        *Bool      `xmlrpc:"hide_expense_policy,omptempty"`
	Id                                       *Int       `xmlrpc:"id,omptempty"`
	Image                                    *String    `xmlrpc:"image,omptempty"`
	ImageMedium                              *String    `xmlrpc:"image_medium,omptempty"`
	ImageSmall                               *String    `xmlrpc:"image_small,omptempty"`
	ImageVariant                             *String    `xmlrpc:"image_variant,omptempty"`
	IncomingQty                              *Float     `xmlrpc:"incoming_qty,omptempty"`
	InvoicePolicy                            *Selection `xmlrpc:"invoice_policy,omptempty"`
	IsProductVariant                         *Bool      `xmlrpc:"is_product_variant,omptempty"`
	ItemIds                                  *Relation  `xmlrpc:"item_ids,omptempty"`
	LifeTime                                 *Int       `xmlrpc:"life_time,omptempty"`
	ListPrice                                *Float     `xmlrpc:"list_price,omptempty"`
	LocationId                               *Many2One  `xmlrpc:"location_id,omptempty"`
	LstPrice                                 *Float     `xmlrpc:"lst_price,omptempty"`
	Membership                               *Bool      `xmlrpc:"membership,omptempty"`
	MembershipDateFrom                       *Time      `xmlrpc:"membership_date_from,omptempty"`
	MembershipDateTo                         *Time      `xmlrpc:"membership_date_to,omptempty"`
	MessageAttachmentCount                   *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds                        *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds                       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                          *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter                   *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageIds                               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId                  *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter                 *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                            *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter                     *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                                     *String    `xmlrpc:"name,omptempty"`
	NbrReorderingRules                       *Int       `xmlrpc:"nbr_reordering_rules,omptempty"`
	OptionalProductIds                       *Relation  `xmlrpc:"optional_product_ids,omptempty"`
	OrderpointIds                            *Relation  `xmlrpc:"orderpoint_ids,omptempty"`
	OutgoingQty                              *Float     `xmlrpc:"outgoing_qty,omptempty"`
	PackagingIds                             *Relation  `xmlrpc:"packaging_ids,omptempty"`
	PartnerRef                               *String    `xmlrpc:"partner_ref,omptempty"`
	PosCategId                               *Many2One  `xmlrpc:"pos_categ_id,omptempty"`
	Price                                    *Float     `xmlrpc:"price,omptempty"`
	PriceExtra                               *Float     `xmlrpc:"price_extra,omptempty"`
	PricelistId                              *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	PricelistItemIds                         *Relation  `xmlrpc:"pricelist_item_ids,omptempty"`
	ProductTemplateAttributeValueIds         *Relation  `xmlrpc:"product_template_attribute_value_ids,omptempty"`
	ProductTmplId                            *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	ProductVariantCount                      *Int       `xmlrpc:"product_variant_count,omptempty"`
	ProductVariantId                         *Many2One  `xmlrpc:"product_variant_id,omptempty"`
	ProductVariantIds                        *Relation  `xmlrpc:"product_variant_ids,omptempty"`
	ProjectId                                *Many2One  `xmlrpc:"project_id,omptempty"`
	ProjectTemplateId                        *Many2One  `xmlrpc:"project_template_id,omptempty"`
	PropertyAccountCreditorPriceDifference   *Many2One  `xmlrpc:"property_account_creditor_price_difference,omptempty"`
	PropertyAccountExpenseId                 *Many2One  `xmlrpc:"property_account_expense_id,omptempty"`
	PropertyAccountIncomeId                  *Many2One  `xmlrpc:"property_account_income_id,omptempty"`
	PropertyCostMethod                       *Selection `xmlrpc:"property_cost_method,omptempty"`
	PropertyStockAccountInput                *Many2One  `xmlrpc:"property_stock_account_input,omptempty"`
	PropertyStockAccountOutput               *Many2One  `xmlrpc:"property_stock_account_output,omptempty"`
	PropertyStockInventory                   *Many2One  `xmlrpc:"property_stock_inventory,omptempty"`
	PropertyStockProduction                  *Many2One  `xmlrpc:"property_stock_production,omptempty"`
	PropertyValuation                        *Selection `xmlrpc:"property_valuation,omptempty"`
	PurchaseLineWarn                         *Selection `xmlrpc:"purchase_line_warn,omptempty"`
	PurchaseLineWarnMsg                      *String    `xmlrpc:"purchase_line_warn_msg,omptempty"`
	PurchaseMethod                           *Selection `xmlrpc:"purchase_method,omptempty"`
	PurchaseOk                               *Bool      `xmlrpc:"purchase_ok,omptempty"`
	PurchasedProductQty                      *Float     `xmlrpc:"purchased_product_qty,omptempty"`
	QtyAtDate                                *Float     `xmlrpc:"qty_at_date,omptempty"`
	QtyAvailable                             *Float     `xmlrpc:"qty_available,omptempty"`
	QuotationDescription                     *String    `xmlrpc:"quotation_description,omptempty"`
	QuotationOnlyDescription                 *String    `xmlrpc:"quotation_only_description,omptempty"`
	ReceptionCount                           *Int       `xmlrpc:"reception_count,omptempty"`
	RecurringInvoice                         *Bool      `xmlrpc:"recurring_invoice,omptempty"`
	RemovalTime                              *Int       `xmlrpc:"removal_time,omptempty"`
	Rental                                   *Bool      `xmlrpc:"rental,omptempty"`
	ReorderingMaxQty                         *Float     `xmlrpc:"reordering_max_qty,omptempty"`
	ReorderingMinQty                         *Float     `xmlrpc:"reordering_min_qty,omptempty"`
	ResponsibleId                            *Many2One  `xmlrpc:"responsible_id,omptempty"`
	RouteFromCategIds                        *Relation  `xmlrpc:"route_from_categ_ids,omptempty"`
	RouteIds                                 *Relation  `xmlrpc:"route_ids,omptempty"`
	SaleDelay                                *Float     `xmlrpc:"sale_delay,omptempty"`
	SaleLineWarn                             *Selection `xmlrpc:"sale_line_warn,omptempty"`
	SaleLineWarnMsg                          *String    `xmlrpc:"sale_line_warn_msg,omptempty"`
	SaleOk                                   *Bool      `xmlrpc:"sale_ok,omptempty"`
	SalesCount                               *Float     `xmlrpc:"sales_count,omptempty"`
	SellerIds                                *Relation  `xmlrpc:"seller_ids,omptempty"`
	Sequence                                 *Int       `xmlrpc:"sequence,omptempty"`
	ServicePolicy                            *Selection `xmlrpc:"service_policy,omptempty"`
	ServiceToPurchase                        *Bool      `xmlrpc:"service_to_purchase,omptempty"`
	ServiceTracking                          *Selection `xmlrpc:"service_tracking,omptempty"`
	ServiceType                              *Selection `xmlrpc:"service_type,omptempty"`
	StandardPrice                            *Float     `xmlrpc:"standard_price,omptempty"`
	State                                    *Selection `xmlrpc:"state,omptempty"`
	StockFifoManualMoveIds                   *Relation  `xmlrpc:"stock_fifo_manual_move_ids,omptempty"`
	StockFifoRealTimeAmlIds                  *Relation  `xmlrpc:"stock_fifo_real_time_aml_ids,omptempty"`
	StockMoveIds                             *Relation  `xmlrpc:"stock_move_ids,omptempty"`
	StockQuantIds                            *Relation  `xmlrpc:"stock_quant_ids,omptempty"`
	StockValue                               *Float     `xmlrpc:"stock_value,omptempty"`
	StockValueCurrencyId                     *Many2One  `xmlrpc:"stock_value_currency_id,omptempty"`
	SubscriptionTemplateId                   *Many2One  `xmlrpc:"subscription_template_id,omptempty"`
	SupplierTaxesId                          *Relation  `xmlrpc:"supplier_taxes_id,omptempty"`
	TaxesId                                  *Relation  `xmlrpc:"taxes_id,omptempty"`
	ToWeight                                 *Bool      `xmlrpc:"to_weight,omptempty"`
	Tracking                                 *Selection `xmlrpc:"tracking,omptempty"`
	Type                                     *Selection `xmlrpc:"type,omptempty"`
	UomId                                    *Many2One  `xmlrpc:"uom_id,omptempty"`
	UomName                                  *String    `xmlrpc:"uom_name,omptempty"`
	UomPoId                                  *Many2One  `xmlrpc:"uom_po_id,omptempty"`
	UseTime                                  *Int       `xmlrpc:"use_time,omptempty"`
	ValidArchivedVariantIds                  *Relation  `xmlrpc:"valid_archived_variant_ids,omptempty"`
	ValidExistingVariantIds                  *Relation  `xmlrpc:"valid_existing_variant_ids,omptempty"`
	ValidProductAttributeIds                 *Relation  `xmlrpc:"valid_product_attribute_ids,omptempty"`
	ValidProductAttributeValueIds            *Relation  `xmlrpc:"valid_product_attribute_value_ids,omptempty"`
	ValidProductAttributeValueWnvaIds        *Relation  `xmlrpc:"valid_product_attribute_value_wnva_ids,omptempty"`
	ValidProductAttributeWnvaIds             *Relation  `xmlrpc:"valid_product_attribute_wnva_ids,omptempty"`
	ValidProductTemplateAttributeLineIds     *Relation  `xmlrpc:"valid_product_template_attribute_line_ids,omptempty"`
	ValidProductTemplateAttributeLineWnvaIds *Relation  `xmlrpc:"valid_product_template_attribute_line_wnva_ids,omptempty"`
	Valuation                                *String    `xmlrpc:"valuation,omptempty"`
	VariantSellerIds                         *Relation  `xmlrpc:"variant_seller_ids,omptempty"`
	VirtualAvailable                         *Float     `xmlrpc:"virtual_available,omptempty"`
	Volume                                   *Float     `xmlrpc:"volume,omptempty"`
	WarehouseId                              *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WebsiteDescription                       *String    `xmlrpc:"website_description,omptempty"`
	WebsiteMessageIds                        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	Weight                                   *Float     `xmlrpc:"weight,omptempty"`
	WeightUomId                              *Many2One  `xmlrpc:"weight_uom_id,omptempty"`
	WeightUomName                            *String    `xmlrpc:"weight_uom_name,omptempty"`
	WriteDate                                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProductProducts represents array of product.product model.
type ProductProducts []ProductProduct

// ProductProductModel is the odoo model name.
const ProductProductModel = "product.product"

// Many2One convert ProductProduct to *Many2One.
func (pp *ProductProduct) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProduct(pp *ProductProduct) (int64, error) {
	return c.Create(ProductProductModel, pp)
}

// UpdateProductProduct updates an existing product.product record.
func (c *Client) UpdateProductProduct(pp *ProductProduct) error {
	return c.UpdateProductProducts([]int64{pp.Id.Get()}, pp)
}

// UpdateProductProducts updates existing product.product records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductProducts(ids []int64, pp *ProductProduct) error {
	return c.Update(ProductProductModel, ids, pp)
}

// DeleteProductProduct deletes an existing product.product record.
func (c *Client) DeleteProductProduct(id int64) error {
	return c.DeleteProductProducts([]int64{id})
}

// DeleteProductProducts deletes existing product.product records.
func (c *Client) DeleteProductProducts(ids []int64) error {
	return c.Delete(ProductProductModel, ids)
}

// GetProductProduct gets product.product existing record.
func (c *Client) GetProductProduct(id int64) (*ProductProduct, error) {
	pps, err := c.GetProductProducts([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.product not found", id)
}

// GetProductProducts gets product.product existing records.
func (c *Client) GetProductProducts(ids []int64) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.Read(ProductProductModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProduct finds product.product record by querying it with criteria.
func (c *Client) FindProductProduct(criteria *Criteria) (*ProductProduct, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("product.product was not found")
}

// FindProductProducts finds product.product records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProducts(criteria *Criteria, options *Options) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProductIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductProductId finds record id by querying it with criteria.
func (c *Client) FindProductProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.product was not found")
}
