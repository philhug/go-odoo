package odoo

import (
	"fmt"
)

// AccountInvoice represents account.invoice model.
type AccountInvoice struct {
	LastUpdate                     *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken                    *String    `xmlrpc:"access_token,omptempty"`
	AccessUrl                      *String    `xmlrpc:"access_url,omptempty"`
	AccessWarning                  *String    `xmlrpc:"access_warning,omptempty"`
	AccountId                      *Many2One  `xmlrpc:"account_id,omptempty"`
	ActivityDateDeadline           *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityIds                    *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                  *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                 *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                 *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AmountByGroup                  *String    `xmlrpc:"amount_by_group,omptempty"`
	AmountTax                      *Float     `xmlrpc:"amount_tax,omptempty"`
	AmountTaxSigned                *Float     `xmlrpc:"amount_tax_signed,omptempty"`
	AmountTotal                    *Float     `xmlrpc:"amount_total,omptempty"`
	AmountTotalCompanySigned       *Float     `xmlrpc:"amount_total_company_signed,omptempty"`
	AmountTotalSigned              *Float     `xmlrpc:"amount_total_signed,omptempty"`
	AmountUntaxed                  *Float     `xmlrpc:"amount_untaxed,omptempty"`
	AmountUntaxedInvoiceSigned     *Float     `xmlrpc:"amount_untaxed_invoice_signed,omptempty"`
	AmountUntaxedSigned            *Float     `xmlrpc:"amount_untaxed_signed,omptempty"`
	AuthorizedTransactionIds       *Relation  `xmlrpc:"authorized_transaction_ids,omptempty"`
	BankAccountRequired            *Bool      `xmlrpc:"bank_account_required,omptempty"`
	BvrReference                   *String    `xmlrpc:"bvr_reference,omptempty"`
	CampaignId                     *Many2One  `xmlrpc:"campaign_id,omptempty"`
	CashRoundingId                 *Many2One  `xmlrpc:"cash_rounding_id,omptempty"`
	Comment                        *String    `xmlrpc:"comment,omptempty"`
	CommercialPartnerId            *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyCurrencyId              *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                      *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                     *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                           *Time      `xmlrpc:"date,omptempty"`
	DateDue                        *Time      `xmlrpc:"date_due,omptempty"`
	DateInvoice                    *Time      `xmlrpc:"date_invoice,omptempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omptempty"`
	ExtractCanShowResendButton     *Bool      `xmlrpc:"extract_can_show_resend_button,omptempty"`
	ExtractCanShowSendButton       *Bool      `xmlrpc:"extract_can_show_send_button,omptempty"`
	ExtractRemoteid                *Int       `xmlrpc:"extract_remoteid,omptempty"`
	ExtractState                   *Selection `xmlrpc:"extract_state,omptempty"`
	ExtractWordIds                 *Relation  `xmlrpc:"extract_word_ids,omptempty"`
	FiscalPositionId               *Many2One  `xmlrpc:"fiscal_position_id,omptempty"`
	HasOutstanding                 *Bool      `xmlrpc:"has_outstanding,omptempty"`
	Id                             *Int       `xmlrpc:"id,omptempty"`
	IncotermId                     *Many2One  `xmlrpc:"incoterm_id,omptempty"`
	IncotermsId                    *Many2One  `xmlrpc:"incoterms_id,omptempty"`
	InvoiceIcon                    *String    `xmlrpc:"invoice_icon,omptempty"`
	InvoiceLineIds                 *Relation  `xmlrpc:"invoice_line_ids,omptempty"`
	InvoiceZero                    *Bool      `xmlrpc:"invoice_zero,omptempty"`
	IsrReference                   *String    `xmlrpc:"isr_reference,omptempty"`
	JournalId                      *Many2One  `xmlrpc:"journal_id,omptempty"`
	L10NChCurrencyName             *String    `xmlrpc:"l10n_ch_currency_name,omptempty"`
	L10NChIsrNumber                *String    `xmlrpc:"l10n_ch_isr_number,omptempty"`
	L10NChIsrNumberSpaced          *String    `xmlrpc:"l10n_ch_isr_number_spaced,omptempty"`
	L10NChIsrOpticalLine           *String    `xmlrpc:"l10n_ch_isr_optical_line,omptempty"`
	L10NChIsrPostal                *String    `xmlrpc:"l10n_ch_isr_postal,omptempty"`
	L10NChIsrPostalFormatted       *String    `xmlrpc:"l10n_ch_isr_postal_formatted,omptempty"`
	L10NChIsrSent                  *Bool      `xmlrpc:"l10n_ch_isr_sent,omptempty"`
	L10NChIsrSubscription          *String    `xmlrpc:"l10n_ch_isr_subscription,omptempty"`
	L10NChIsrSubscriptionFormatted *String    `xmlrpc:"l10n_ch_isr_subscription_formatted,omptempty"`
	L10NChIsrValid                 *Bool      `xmlrpc:"l10n_ch_isr_valid,omptempty"`
	MediumId                       *Many2One  `xmlrpc:"medium_id,omptempty"`
	MessageAttachmentCount         *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds              *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds             *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter         *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageIds                     *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower              *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId        *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction              *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter       *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds              *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                  *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter           *Int       `xmlrpc:"message_unread_counter,omptempty"`
	MoveId                         *Many2One  `xmlrpc:"move_id,omptempty"`
	MoveName                       *String    `xmlrpc:"move_name,omptempty"`
	Name                           *String    `xmlrpc:"name,omptempty"`
	Number                         *String    `xmlrpc:"number,omptempty"`
	Origin                         *String    `xmlrpc:"origin,omptempty"`
	OutstandingCreditsDebitsWidget *String    `xmlrpc:"outstanding_credits_debits_widget,omptempty"`
	PartnerBankId                  *Many2One  `xmlrpc:"partner_bank_id,omptempty"`
	PartnerId                      *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerShippingId              *Many2One  `xmlrpc:"partner_shipping_id,omptempty"`
	PaymentIds                     *Relation  `xmlrpc:"payment_ids,omptempty"`
	PaymentModeId                  *Many2One  `xmlrpc:"payment_mode_id,omptempty"`
	PaymentMoveLineIds             *Relation  `xmlrpc:"payment_move_line_ids,omptempty"`
	PaymentOrderOk                 *Bool      `xmlrpc:"payment_order_ok,omptempty"`
	PaymentTermId                  *Many2One  `xmlrpc:"payment_term_id,omptempty"`
	PaymentsWidget                 *String    `xmlrpc:"payments_widget,omptempty"`
	PaypalUrl                      *String    `xmlrpc:"paypal_url,omptempty"`
	PlaneId                        *Many2One  `xmlrpc:"plane_id,omptempty"`
	PurchaseId                     *Many2One  `xmlrpc:"purchase_id,omptempty"`
	Reconciled                     *Bool      `xmlrpc:"reconciled,omptempty"`
	Reference                      *String    `xmlrpc:"reference,omptempty"`
	ReferenceType                  *Selection `xmlrpc:"reference_type,omptempty"`
	RefundInvoiceId                *Many2One  `xmlrpc:"refund_invoice_id,omptempty"`
	RefundInvoiceIds               *Relation  `xmlrpc:"refund_invoice_ids,omptempty"`
	Residual                       *Float     `xmlrpc:"residual,omptempty"`
	ResidualCompanySigned          *Float     `xmlrpc:"residual_company_signed,omptempty"`
	ResidualSigned                 *Float     `xmlrpc:"residual_signed,omptempty"`
	SaleId                         *Many2One  `xmlrpc:"sale_id,omptempty"`
	Sent                           *Bool      `xmlrpc:"sent,omptempty"`
	SequenceNumberNext             *String    `xmlrpc:"sequence_number_next,omptempty"`
	SequenceNumberNextPrefix       *String    `xmlrpc:"sequence_number_next_prefix,omptempty"`
	SlipIds                        *Relation  `xmlrpc:"slip_ids,omptempty"`
	SourceEmail                    *String    `xmlrpc:"source_email,omptempty"`
	SourceId                       *Many2One  `xmlrpc:"source_id,omptempty"`
	State                          *Selection `xmlrpc:"state,omptempty"`
	TaxLineIds                     *Relation  `xmlrpc:"tax_line_ids,omptempty"`
	TeamId                         *Many2One  `xmlrpc:"team_id,omptempty"`
	TimesheetCount                 *Int       `xmlrpc:"timesheet_count,omptempty"`
	TimesheetIds                   *Relation  `xmlrpc:"timesheet_ids,omptempty"`
	TransactionId                  *String    `xmlrpc:"transaction_id,omptempty"`
	TransactionIds                 *Relation  `xmlrpc:"transaction_ids,omptempty"`
	Type                           *Selection `xmlrpc:"type,omptempty"`
	UserId                         *Many2One  `xmlrpc:"user_id,omptempty"`
	VendorBillId                   *Many2One  `xmlrpc:"vendor_bill_id,omptempty"`
	VendorBillPurchaseId           *Many2One  `xmlrpc:"vendor_bill_purchase_id,omptempty"`
	VendorDisplayName              *String    `xmlrpc:"vendor_display_name,omptempty"`
	WebsiteMessageIds              *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountInvoices represents array of account.invoice model.
type AccountInvoices []AccountInvoice

// AccountInvoiceModel is the odoo model name.
const AccountInvoiceModel = "account.invoice"

// Many2One convert AccountInvoice to *Many2One.
func (ai *AccountInvoice) Many2One() *Many2One {
	return NewMany2One(ai.Id.Get(), "")
}

// CreateAccountInvoice creates a new account.invoice model and returns its id.
func (c *Client) CreateAccountInvoice(ai *AccountInvoice) (int64, error) {
	return c.Create(AccountInvoiceModel, ai)
}

// UpdateAccountInvoice updates an existing account.invoice record.
func (c *Client) UpdateAccountInvoice(ai *AccountInvoice) error {
	return c.UpdateAccountInvoices([]int64{ai.Id.Get()}, ai)
}

// UpdateAccountInvoices updates existing account.invoice records.
// All records (represented by ids) will be updated by ai values.
func (c *Client) UpdateAccountInvoices(ids []int64, ai *AccountInvoice) error {
	return c.Update(AccountInvoiceModel, ids, ai)
}

// DeleteAccountInvoice deletes an existing account.invoice record.
func (c *Client) DeleteAccountInvoice(id int64) error {
	return c.DeleteAccountInvoices([]int64{id})
}

// DeleteAccountInvoices deletes existing account.invoice records.
func (c *Client) DeleteAccountInvoices(ids []int64) error {
	return c.Delete(AccountInvoiceModel, ids)
}

// GetAccountInvoice gets account.invoice existing record.
func (c *Client) GetAccountInvoice(id int64) (*AccountInvoice, error) {
	ais, err := c.GetAccountInvoices([]int64{id})
	if err != nil {
		return nil, err
	}
	if ais != nil && len(*ais) > 0 {
		return &((*ais)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.invoice not found", id)
}

// GetAccountInvoices gets account.invoice existing records.
func (c *Client) GetAccountInvoices(ids []int64) (*AccountInvoices, error) {
	ais := &AccountInvoices{}
	if err := c.Read(AccountInvoiceModel, ids, nil, ais); err != nil {
		return nil, err
	}
	return ais, nil
}

// FindAccountInvoice finds account.invoice record by querying it with criteria.
func (c *Client) FindAccountInvoice(criteria *Criteria) (*AccountInvoice, error) {
	ais := &AccountInvoices{}
	if err := c.SearchRead(AccountInvoiceModel, criteria, NewOptions().Limit(1), ais); err != nil {
		return nil, err
	}
	if ais != nil && len(*ais) > 0 {
		return &((*ais)[0]), nil
	}
	return nil, fmt.Errorf("account.invoice was not found")
}

// FindAccountInvoices finds account.invoice records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoices(criteria *Criteria, options *Options) (*AccountInvoices, error) {
	ais := &AccountInvoices{}
	if err := c.SearchRead(AccountInvoiceModel, criteria, options, ais); err != nil {
		return nil, err
	}
	return ais, nil
}

// FindAccountInvoiceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountInvoiceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountInvoiceId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.invoice was not found")
}
