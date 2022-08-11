package odoo

import (
	"fmt"
)

// ResPartner represents res.partner model.
type ResPartner struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	Active                        *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline          *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityIds                   *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                 *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AdditionalInfo                *String    `xmlrpc:"additional_info,omptempty"`
	AssociateMember               *Many2One  `xmlrpc:"associate_member,omptempty"`
	BankAccountCount              *Int       `xmlrpc:"bank_account_count,omptempty"`
	BankIds                       *Relation  `xmlrpc:"bank_ids,omptempty"`
	Barcode                       *String    `xmlrpc:"barcode,omptempty"`
	Birthdate                     *String    `xmlrpc:"birthdate,omptempty"`
	BirthdateDate                 *Time      `xmlrpc:"birthdate_date,omptempty"`
	CalendarLastNotifAck          *Time      `xmlrpc:"calendar_last_notif_ack,omptempty"`
	CategoryId                    *Relation  `xmlrpc:"category_id,omptempty"`
	ChannelIds                    *Relation  `xmlrpc:"channel_ids,omptempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omptempty"`
	City                          *String    `xmlrpc:"city,omptempty"`
	CityId                        *Many2One  `xmlrpc:"city_id,omptempty"`
	Color                         *Int       `xmlrpc:"color,omptempty"`
	Comment                       *String    `xmlrpc:"comment,omptempty"`
	CommercialCompanyName         *String    `xmlrpc:"commercial_company_name,omptempty"`
	CommercialPartnerId           *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyName                   *String    `xmlrpc:"company_name,omptempty"`
	CompanyType                   *Selection `xmlrpc:"company_type,omptempty"`
	ContactAddress                *String    `xmlrpc:"contact_address,omptempty"`
	ContractIds                   *Relation  `xmlrpc:"contract_ids,omptempty"`
	ContractsCount                *Int       `xmlrpc:"contracts_count,omptempty"`
	CountryEnforceCities          *Bool      `xmlrpc:"country_enforce_cities,omptempty"`
	CountryId                     *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	Credit                        *Float     `xmlrpc:"credit,omptempty"`
	CreditLimit                   *Float     `xmlrpc:"credit_limit,omptempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omptempty"`
	Customer                      *Bool      `xmlrpc:"customer,omptempty"`
	CustomerPaymentModeId         *Many2One  `xmlrpc:"customer_payment_mode_id,omptempty"`
	Date                          *Time      `xmlrpc:"date,omptempty"`
	Debit                         *Float     `xmlrpc:"debit,omptempty"`
	DebitLimit                    *Float     `xmlrpc:"debit_limit,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	Email                         *String    `xmlrpc:"email,omptempty"`
	Email2                        *String    `xmlrpc:"email2,omptempty"`
	EmailFormatted                *String    `xmlrpc:"email_formatted,omptempty"`
	Employee                      *Bool      `xmlrpc:"employee,omptempty"`
	EventCount                    *Int       `xmlrpc:"event_count,omptempty"`
	Firstname                     *String    `xmlrpc:"firstname,omptempty"`
	FollowupStatus                *Selection `xmlrpc:"followup_status,omptempty"`
	FreeMember                    *Bool      `xmlrpc:"free_member,omptempty"`
	Function                      *String    `xmlrpc:"function,omptempty"`
	HasUnreconciledEntries        *Bool      `xmlrpc:"has_unreconciled_entries,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	IdNumbers                     *Relation  `xmlrpc:"id_numbers,omptempty"`
	ImStatus                      *String    `xmlrpc:"im_status,omptempty"`
	Image                         *String    `xmlrpc:"image,omptempty"`
	ImageMedium                   *String    `xmlrpc:"image_medium,omptempty"`
	ImageSmall                    *String    `xmlrpc:"image_small,omptempty"`
	IndustryId                    *Many2One  `xmlrpc:"industry_id,omptempty"`
	InvoiceIds                    *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoiceWarn                   *Selection `xmlrpc:"invoice_warn,omptempty"`
	InvoiceWarnMsg                *String    `xmlrpc:"invoice_warn_msg,omptempty"`
	IsBlacklisted                 *Bool      `xmlrpc:"is_blacklisted,omptempty"`
	IsCompany                     *Bool      `xmlrpc:"is_company,omptempty"`
	IsPublished                   *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized                *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	JournalItemCount              *Int       `xmlrpc:"journal_item_count,omptempty"`
	Lang                          *Selection `xmlrpc:"lang,omptempty"`
	LastTimeEntriesChecked        *Time      `xmlrpc:"last_time_entries_checked,omptempty"`
	Lastname                      *String    `xmlrpc:"lastname,omptempty"`
	MeetingCount                  *Int       `xmlrpc:"meeting_count,omptempty"`
	MeetingIds                    *Relation  `xmlrpc:"meeting_ids,omptempty"`
	MemberLines                   *Relation  `xmlrpc:"member_lines,omptempty"`
	MembershipAmount              *Float     `xmlrpc:"membership_amount,omptempty"`
	MembershipCancel              *Time      `xmlrpc:"membership_cancel,omptempty"`
	MembershipStart               *Time      `xmlrpc:"membership_start,omptempty"`
	MembershipState               *Selection `xmlrpc:"membership_state,omptempty"`
	MembershipStop                *Time      `xmlrpc:"membership_stop,omptempty"`
	MessageAttachmentCount        *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageBounce                 *Int       `xmlrpc:"message_bounce,omptempty"`
	MessageChannelIds             *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds            *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError               *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter        *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageIds                    *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower             *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId       *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction             *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter      *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds             *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                 *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter          *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Mobile                        *String    `xmlrpc:"mobile,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	OcnToken                      *String    `xmlrpc:"ocn_token,omptempty"`
	OnlinePartnerBankAccount      *String    `xmlrpc:"online_partner_bank_account,omptempty"`
	OnlinePartnerVendorName       *String    `xmlrpc:"online_partner_vendor_name,omptempty"`
	OpportunityCount              *Int       `xmlrpc:"opportunity_count,omptempty"`
	OpportunityIds                *Relation  `xmlrpc:"opportunity_ids,omptempty"`
	ParentId                      *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentName                    *String    `xmlrpc:"parent_name,omptempty"`
	PartnerGid                    *Int       `xmlrpc:"partner_gid,omptempty"`
	PartnerLedgerLabel            *String    `xmlrpc:"partner_ledger_label,omptempty"`
	PartnerShare                  *Bool      `xmlrpc:"partner_share,omptempty"`
	PaymentNextActionDate         *Time      `xmlrpc:"payment_next_action_date,omptempty"`
	PaymentTokenCount             *Int       `xmlrpc:"payment_token_count,omptempty"`
	PaymentTokenIds               *Relation  `xmlrpc:"payment_token_ids,omptempty"`
	Phone                         *String    `xmlrpc:"phone,omptempty"`
	PhoneCompany                  *String    `xmlrpc:"phone_company,omptempty"`
	PickingWarn                   *Selection `xmlrpc:"picking_warn,omptempty"`
	PickingWarnMsg                *String    `xmlrpc:"picking_warn_msg,omptempty"`
	PlaneIds                      *Relation  `xmlrpc:"plane_ids,omptempty"`
	PosOrderCount                 *Int       `xmlrpc:"pos_order_count,omptempty"`
	PropertyAccountPayableId      *Many2One  `xmlrpc:"property_account_payable_id,omptempty"`
	PropertyAccountPositionId     *Many2One  `xmlrpc:"property_account_position_id,omptempty"`
	PropertyAccountReceivableId   *Many2One  `xmlrpc:"property_account_receivable_id,omptempty"`
	PropertyPaymentTermId         *Many2One  `xmlrpc:"property_payment_term_id,omptempty"`
	PropertyProductPricelist      *Many2One  `xmlrpc:"property_product_pricelist,omptempty"`
	PropertyPurchaseCurrencyId    *Many2One  `xmlrpc:"property_purchase_currency_id,omptempty"`
	PropertyStockCustomer         *Many2One  `xmlrpc:"property_stock_customer,omptempty"`
	PropertyStockSupplier         *Many2One  `xmlrpc:"property_stock_supplier,omptempty"`
	PropertySupplierPaymentTermId *Many2One  `xmlrpc:"property_supplier_payment_term_id,omptempty"`
	PurchaseOrderCount            *Int       `xmlrpc:"purchase_order_count,omptempty"`
	PurchaseWarn                  *Selection `xmlrpc:"purchase_warn,omptempty"`
	PurchaseWarnMsg               *String    `xmlrpc:"purchase_warn_msg,omptempty"`
	Ref                           *String    `xmlrpc:"ref,omptempty"`
	RefCompanyIds                 *Relation  `xmlrpc:"ref_company_ids,omptempty"`
	SaleOrderCount                *Int       `xmlrpc:"sale_order_count,omptempty"`
	SaleOrderIds                  *Relation  `xmlrpc:"sale_order_ids,omptempty"`
	SaleWarn                      *Selection `xmlrpc:"sale_warn,omptempty"`
	SaleWarnMsg                   *String    `xmlrpc:"sale_warn_msg,omptempty"`
	Self                          *Many2One  `xmlrpc:"self,omptempty"`
	SignatureCount                *Int       `xmlrpc:"signature_count,omptempty"`
	SignupExpiration              *Time      `xmlrpc:"signup_expiration,omptempty"`
	SignupToken                   *String    `xmlrpc:"signup_token,omptempty"`
	SignupType                    *String    `xmlrpc:"signup_type,omptempty"`
	SignupUrl                     *String    `xmlrpc:"signup_url,omptempty"`
	SignupValid                   *Bool      `xmlrpc:"signup_valid,omptempty"`
	StateId                       *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                        *String    `xmlrpc:"street,omptempty"`
	Street2                       *String    `xmlrpc:"street2,omptempty"`
	SubscriptionCount             *Int       `xmlrpc:"subscription_count,omptempty"`
	Supplier                      *Bool      `xmlrpc:"supplier,omptempty"`
	SupplierInvoiceCount          *Int       `xmlrpc:"supplier_invoice_count,omptempty"`
	SupplierPaymentModeId         *Many2One  `xmlrpc:"supplier_payment_mode_id,omptempty"`
	TaskCount                     *Int       `xmlrpc:"task_count,omptempty"`
	TaskIds                       *Relation  `xmlrpc:"task_ids,omptempty"`
	TeamId                        *Many2One  `xmlrpc:"team_id,omptempty"`
	Title                         *Many2One  `xmlrpc:"title,omptempty"`
	TotalDue                      *Float     `xmlrpc:"total_due,omptempty"`
	TotalInvoiced                 *Float     `xmlrpc:"total_invoiced,omptempty"`
	TotalOverdue                  *Float     `xmlrpc:"total_overdue,omptempty"`
	Trust                         *Selection `xmlrpc:"trust,omptempty"`
	Type                          *Selection `xmlrpc:"type,omptempty"`
	Tz                            *Selection `xmlrpc:"tz,omptempty"`
	TzOffset                      *String    `xmlrpc:"tz_offset,omptempty"`
	UnreconciledAmlIds            *Relation  `xmlrpc:"unreconciled_aml_ids,omptempty"`
	UseParentAddress              *Bool      `xmlrpc:"use_parent_address,omptempty"`
	UserId                        *Many2One  `xmlrpc:"user_id,omptempty"`
	UserIds                       *Relation  `xmlrpc:"user_ids,omptempty"`
	Vat                           *String    `xmlrpc:"vat,omptempty"`
	Website                       *String    `xmlrpc:"website,omptempty"`
	WebsiteDescription            *String    `xmlrpc:"website_description,omptempty"`
	WebsiteId                     *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds             *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription        *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords           *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg              *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle              *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePrivate                *Bool      `xmlrpc:"website_private,omptempty"`
	WebsitePublished              *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteShortDescription       *String    `xmlrpc:"website_short_description,omptempty"`
	WebsiteUrl                    *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
	XFax                          *String    `xmlrpc:"x_fax,omptempty"`
	XMfgtMemberNumber             *String    `xmlrpc:"x_mfgt_member_number,omptempty"`
	XMfgtMemberSince              *Time      `xmlrpc:"x_mfgt_member_since,omptempty"`
	XMfgtMemberUntil              *Time      `xmlrpc:"x_mfgt_member_until,omptempty"`
	Zip                           *String    `xmlrpc:"zip,omptempty"`
	ZipId                         *Many2One  `xmlrpc:"zip_id,omptempty"`
}

// ResPartners represents array of res.partner model.
type ResPartners []ResPartner

// ResPartnerModel is the odoo model name.
const ResPartnerModel = "res.partner"

// Many2One convert ResPartner to *Many2One.
func (rp *ResPartner) Many2One() *Many2One {
	return NewMany2One(rp.Id.Get(), "")
}

// CreateResPartner creates a new res.partner model and returns its id.
func (c *Client) CreateResPartner(rp *ResPartner) (int64, error) {
	return c.Create(ResPartnerModel, rp)
}

// UpdateResPartner updates an existing res.partner record.
func (c *Client) UpdateResPartner(rp *ResPartner) error {
	return c.UpdateResPartners([]int64{rp.Id.Get()}, rp)
}

// UpdateResPartners updates existing res.partner records.
// All records (represented by ids) will be updated by rp values.
func (c *Client) UpdateResPartners(ids []int64, rp *ResPartner) error {
	return c.Update(ResPartnerModel, ids, rp)
}

// DeleteResPartner deletes an existing res.partner record.
func (c *Client) DeleteResPartner(id int64) error {
	return c.DeleteResPartners([]int64{id})
}

// DeleteResPartners deletes existing res.partner records.
func (c *Client) DeleteResPartners(ids []int64) error {
	return c.Delete(ResPartnerModel, ids)
}

// GetResPartner gets res.partner existing record.
func (c *Client) GetResPartner(id int64) (*ResPartner, error) {
	rps, err := c.GetResPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.partner not found", id)
}

// GetResPartners gets res.partner existing records.
func (c *Client) GetResPartners(ids []int64) (*ResPartners, error) {
	rps := &ResPartners{}
	if err := c.Read(ResPartnerModel, ids, nil, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindResPartner finds res.partner record by querying it with criteria.
func (c *Client) FindResPartner(criteria *Criteria) (*ResPartner, error) {
	rps := &ResPartners{}
	if err := c.SearchRead(ResPartnerModel, criteria, NewOptions().Limit(1), rps); err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("res.partner was not found")
}

// FindResPartners finds res.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartners(criteria *Criteria, options *Options) (*ResPartners, error) {
	rps := &ResPartners{}
	if err := c.SearchRead(ResPartnerModel, criteria, options, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindResPartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerId finds record id by querying it with criteria.
func (c *Client) FindResPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.partner was not found")
}
