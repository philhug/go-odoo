package odoo

import (
	"fmt"
)

// AccountMove represents account.move model.
type AccountMove struct {
	LastUpdate                            *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken                           *String    `xmlrpc:"access_token,omptempty"`
	AccessUrl                             *String    `xmlrpc:"access_url,omptempty"`
	AccessWarning                         *String    `xmlrpc:"access_warning,omptempty"`
	ActivityCalendarEventId               *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline                  *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration           *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon                 *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                           *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                         *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                       *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                      *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                        *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                        *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AlwaysTaxExigible                     *Bool      `xmlrpc:"always_tax_exigible,omptempty"`
	AmountPaid                            *Float     `xmlrpc:"amount_paid,omptempty"`
	AmountResidual                        *Float     `xmlrpc:"amount_residual,omptempty"`
	AmountResidualSigned                  *Float     `xmlrpc:"amount_residual_signed,omptempty"`
	AmountTax                             *Float     `xmlrpc:"amount_tax,omptempty"`
	AmountTaxSigned                       *Float     `xmlrpc:"amount_tax_signed,omptempty"`
	AmountTotal                           *Float     `xmlrpc:"amount_total,omptempty"`
	AmountTotalInCurrencySigned           *Float     `xmlrpc:"amount_total_in_currency_signed,omptempty"`
	AmountTotalSigned                     *Float     `xmlrpc:"amount_total_signed,omptempty"`
	AmountUntaxed                         *Float     `xmlrpc:"amount_untaxed,omptempty"`
	AmountUntaxedSigned                   *Float     `xmlrpc:"amount_untaxed_signed,omptempty"`
	AssetAssetType                        *Selection `xmlrpc:"asset_asset_type,omptempty"`
	AssetDepreciatedValue                 *Float     `xmlrpc:"asset_depreciated_value,omptempty"`
	AssetDepreciationBeginningDate        *Time      `xmlrpc:"asset_depreciation_beginning_date,omptempty"`
	AssetId                               *Many2One  `xmlrpc:"asset_id,omptempty"`
	AssetIdDisplayName                    *String    `xmlrpc:"asset_id_display_name,omptempty"`
	AssetIds                              *Relation  `xmlrpc:"asset_ids,omptempty"`
	AssetNumberDays                       *Int       `xmlrpc:"asset_number_days,omptempty"`
	AssetRemainingValue                   *Float     `xmlrpc:"asset_remaining_value,omptempty"`
	AssetValueChange                      *Bool      `xmlrpc:"asset_value_change,omptempty"`
	AttachmentIds                         *Relation  `xmlrpc:"attachment_ids,omptempty"`
	AuthorizedTransactionIds              *Relation  `xmlrpc:"authorized_transaction_ids,omptempty"`
	AutoPost                              *Selection `xmlrpc:"auto_post,omptempty"`
	AutoPostOriginId                      *Many2One  `xmlrpc:"auto_post_origin_id,omptempty"`
	AutoPostUntil                         *Time      `xmlrpc:"auto_post_until,omptempty"`
	BankPartnerId                         *Many2One  `xmlrpc:"bank_partner_id,omptempty"`
	CampaignId                            *Many2One  `xmlrpc:"campaign_id,omptempty"`
	CommercialPartnerId                   *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyCurrencyId                     *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryCode                           *String    `xmlrpc:"country_code,omptempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                                  *Time      `xmlrpc:"date,omptempty"`
	DepreciationValue                     *Float     `xmlrpc:"depreciation_value,omptempty"`
	DirectionSign                         *Int       `xmlrpc:"direction_sign,omptempty"`
	DisplayInactiveCurrencyWarning        *Bool      `xmlrpc:"display_inactive_currency_warning,omptempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omptempty"`
	DisplayQrCode                         *Bool      `xmlrpc:"display_qr_code,omptempty"`
	DraftAssetIds                         *Bool      `xmlrpc:"draft_asset_ids,omptempty"`
	DuplicatedRefIds                      *Relation  `xmlrpc:"duplicated_ref_ids,omptempty"`
	EdiBlockingLevel                      *Selection `xmlrpc:"edi_blocking_level,omptempty"`
	EdiDocumentIds                        *Relation  `xmlrpc:"edi_document_ids,omptempty"`
	EdiErrorCount                         *Int       `xmlrpc:"edi_error_count,omptempty"`
	EdiErrorMessage                       *String    `xmlrpc:"edi_error_message,omptempty"`
	EdiShowAbandonCancelButton            *Bool      `xmlrpc:"edi_show_abandon_cancel_button,omptempty"`
	EdiShowCancelButton                   *Bool      `xmlrpc:"edi_show_cancel_button,omptempty"`
	EdiState                              *Selection `xmlrpc:"edi_state,omptempty"`
	EdiWebServicesToProcess               *String    `xmlrpc:"edi_web_services_to_process,omptempty"`
	ExpenseSheetId                        *Relation  `xmlrpc:"expense_sheet_id,omptempty"`
	ExtractAttachmentId                   *Many2One  `xmlrpc:"extract_attachment_id,omptempty"`
	ExtractCanShowBanners                 *Bool      `xmlrpc:"extract_can_show_banners,omptempty"`
	ExtractCanShowResendButton            *Bool      `xmlrpc:"extract_can_show_resend_button,omptempty"`
	ExtractCanShowSendButton              *Bool      `xmlrpc:"extract_can_show_send_button,omptempty"`
	ExtractErrorMessage                   *String    `xmlrpc:"extract_error_message,omptempty"`
	ExtractRemoteId                       *Int       `xmlrpc:"extract_remote_id,omptempty"`
	ExtractState                          *Selection `xmlrpc:"extract_state,omptempty"`
	ExtractStatusCode                     *Int       `xmlrpc:"extract_status_code,omptempty"`
	ExtractWordIds                        *Relation  `xmlrpc:"extract_word_ids,omptempty"`
	FiscalPositionId                      *Many2One  `xmlrpc:"fiscal_position_id,omptempty"`
	HasMessage                            *Bool      `xmlrpc:"has_message,omptempty"`
	HasReconciledEntries                  *Bool      `xmlrpc:"has_reconciled_entries,omptempty"`
	HidePostButton                        *Bool      `xmlrpc:"hide_post_button,omptempty"`
	HighestName                           *String    `xmlrpc:"highest_name,omptempty"`
	Id                                    *Int       `xmlrpc:"id,omptempty"`
	InalterableHash                       *String    `xmlrpc:"inalterable_hash,omptempty"`
	InvoiceCashRoundingId                 *Many2One  `xmlrpc:"invoice_cash_rounding_id,omptempty"`
	InvoiceDate                           *Time      `xmlrpc:"invoice_date,omptempty"`
	InvoiceDateDue                        *Time      `xmlrpc:"invoice_date_due,omptempty"`
	InvoiceFilterTypeDomain               *String    `xmlrpc:"invoice_filter_type_domain,omptempty"`
	InvoiceHasOutstanding                 *Bool      `xmlrpc:"invoice_has_outstanding,omptempty"`
	InvoiceIncotermId                     *Many2One  `xmlrpc:"invoice_incoterm_id,omptempty"`
	InvoiceLineIds                        *Relation  `xmlrpc:"invoice_line_ids,omptempty"`
	InvoiceOrigin                         *String    `xmlrpc:"invoice_origin,omptempty"`
	InvoiceOutstandingCreditsDebitsWidget *String    `xmlrpc:"invoice_outstanding_credits_debits_widget,omptempty"`
	InvoicePartnerDisplayName             *String    `xmlrpc:"invoice_partner_display_name,omptempty"`
	InvoicePaymentTermId                  *Many2One  `xmlrpc:"invoice_payment_term_id,omptempty"`
	InvoicePaymentsWidget                 *String    `xmlrpc:"invoice_payments_widget,omptempty"`
	InvoiceSourceEmail                    *String    `xmlrpc:"invoice_source_email,omptempty"`
	InvoiceUserId                         *Many2One  `xmlrpc:"invoice_user_id,omptempty"`
	InvoiceVendorBillId                   *Many2One  `xmlrpc:"invoice_vendor_bill_id,omptempty"`
	IsMoveSent                            *Bool      `xmlrpc:"is_move_sent,omptempty"`
	IsStorno                              *Bool      `xmlrpc:"is_storno,omptempty"`
	JournalId                             *Many2One  `xmlrpc:"journal_id,omptempty"`
	L10NChCurrencyName                    *String    `xmlrpc:"l10n_ch_currency_name,omptempty"`
	L10NChIsQrValid                       *Bool      `xmlrpc:"l10n_ch_is_qr_valid,omptempty"`
	L10NChIsrNeedsFixing                  *Bool      `xmlrpc:"l10n_ch_isr_needs_fixing,omptempty"`
	L10NChIsrNumber                       *String    `xmlrpc:"l10n_ch_isr_number,omptempty"`
	L10NChIsrNumberSpaced                 *String    `xmlrpc:"l10n_ch_isr_number_spaced,omptempty"`
	L10NChIsrOpticalLine                  *String    `xmlrpc:"l10n_ch_isr_optical_line,omptempty"`
	L10NChIsrSent                         *Bool      `xmlrpc:"l10n_ch_isr_sent,omptempty"`
	L10NChIsrSubscription                 *String    `xmlrpc:"l10n_ch_isr_subscription,omptempty"`
	L10NChIsrSubscriptionFormatted        *String    `xmlrpc:"l10n_ch_isr_subscription_formatted,omptempty"`
	L10NChIsrValid                        *Bool      `xmlrpc:"l10n_ch_isr_valid,omptempty"`
	L10NDin5008Addresses                  *String    `xmlrpc:"l10n_din5008_addresses,omptempty"`
	L10NDin5008DocumentTitle              *String    `xmlrpc:"l10n_din5008_document_title,omptempty"`
	L10NDin5008TemplateData               *String    `xmlrpc:"l10n_din5008_template_data,omptempty"`
	LineIds                               *Relation  `xmlrpc:"line_ids,omptempty"`
	LinkedAssetType                       *String    `xmlrpc:"linked_asset_type,omptempty"`
	MadeSequenceHole                      *Bool      `xmlrpc:"made_sequence_hole,omptempty"`
	MediumId                              *Many2One  `xmlrpc:"medium_id,omptempty"`
	MessageAttachmentCount                *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds                    *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                       *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter                *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                    *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                            *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                     *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId               *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                     *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter              *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                     *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MoveType                              *Selection `xmlrpc:"move_type,omptempty"`
	MyActivityDateDeadline                *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                                  *String    `xmlrpc:"name,omptempty"`
	Narration                             *String    `xmlrpc:"narration,omptempty"`
	NeededTerms                           *String    `xmlrpc:"needed_terms,omptempty"`
	NeededTermsDirty                      *Bool      `xmlrpc:"needed_terms_dirty,omptempty"`
	NumberAssetIds                        *Int       `xmlrpc:"number_asset_ids,omptempty"`
	PartnerBankId                         *Many2One  `xmlrpc:"partner_bank_id,omptempty"`
	PartnerCreditWarning                  *String    `xmlrpc:"partner_credit_warning,omptempty"`
	PartnerId                             *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerShippingId                     *Many2One  `xmlrpc:"partner_shipping_id,omptempty"`
	PaymentId                             *Many2One  `xmlrpc:"payment_id,omptempty"`
	PaymentIds                            *Relation  `xmlrpc:"payment_ids,omptempty"`
	PaymentReference                      *String    `xmlrpc:"payment_reference,omptempty"`
	PaymentState                          *Selection `xmlrpc:"payment_state,omptempty"`
	PaymentStateBeforeSwitch              *String    `xmlrpc:"payment_state_before_switch,omptempty"`
	PaymentTermDetails                    *String    `xmlrpc:"payment_term_details,omptempty"`
	PlaneId                               *Many2One  `xmlrpc:"plane_id,omptempty"`
	PosOrderIds                           *Relation  `xmlrpc:"pos_order_ids,omptempty"`
	PosPaymentIds                         *Relation  `xmlrpc:"pos_payment_ids,omptempty"`
	PostedBefore                          *Bool      `xmlrpc:"posted_before,omptempty"`
	PurchaseId                            *Many2One  `xmlrpc:"purchase_id,omptempty"`
	PurchaseOrderCount                    *Int       `xmlrpc:"purchase_order_count,omptempty"`
	PurchaseVendorBillId                  *Many2One  `xmlrpc:"purchase_vendor_bill_id,omptempty"`
	QrCodeMethod                          *Selection `xmlrpc:"qr_code_method,omptempty"`
	QuickEditMode                         *Bool      `xmlrpc:"quick_edit_mode,omptempty"`
	QuickEditTotalAmount                  *Float     `xmlrpc:"quick_edit_total_amount,omptempty"`
	QuickEncodingVals                     *String    `xmlrpc:"quick_encoding_vals,omptempty"`
	RateDiffPartialRecId                  *Many2One  `xmlrpc:"rate_diff_partial_rec_id,omptempty"`
	Ref                                   *String    `xmlrpc:"ref,omptempty"`
	RestrictModeHashTable                 *Bool      `xmlrpc:"restrict_mode_hash_table,omptempty"`
	ReversalMoveId                        *Relation  `xmlrpc:"reversal_move_id,omptempty"`
	ReversedEntryId                       *Many2One  `xmlrpc:"reversed_entry_id,omptempty"`
	SaleId                                *Many2One  `xmlrpc:"sale_id,omptempty"`
	SaleOrderCount                        *Int       `xmlrpc:"sale_order_count,omptempty"`
	SecureSequenceNumber                  *Int       `xmlrpc:"secure_sequence_number,omptempty"`
	SequenceNumber                        *Int       `xmlrpc:"sequence_number,omptempty"`
	SequencePrefix                        *String    `xmlrpc:"sequence_prefix,omptempty"`
	ShowDiscountDetails                   *Bool      `xmlrpc:"show_discount_details,omptempty"`
	ShowNameWarning                       *Bool      `xmlrpc:"show_name_warning,omptempty"`
	ShowPaymentTermDetails                *Bool      `xmlrpc:"show_payment_term_details,omptempty"`
	ShowResetToDraftButton                *Bool      `xmlrpc:"show_reset_to_draft_button,omptempty"`
	SourceId                              *Many2One  `xmlrpc:"source_id,omptempty"`
	State                                 *Selection `xmlrpc:"state,omptempty"`
	StatementLineId                       *Many2One  `xmlrpc:"statement_line_id,omptempty"`
	StatementLineIds                      *Relation  `xmlrpc:"statement_line_ids,omptempty"`
	StockMoveId                           *Many2One  `xmlrpc:"stock_move_id,omptempty"`
	StockValuationLayerIds                *Relation  `xmlrpc:"stock_valuation_layer_ids,omptempty"`
	StringToHash                          *String    `xmlrpc:"string_to_hash,omptempty"`
	SuitableJournalIds                    *Relation  `xmlrpc:"suitable_journal_ids,omptempty"`
	TaxCashBasisCreatedMoveIds            *Relation  `xmlrpc:"tax_cash_basis_created_move_ids,omptempty"`
	TaxCashBasisOriginMoveId              *Many2One  `xmlrpc:"tax_cash_basis_origin_move_id,omptempty"`
	TaxCashBasisRecId                     *Many2One  `xmlrpc:"tax_cash_basis_rec_id,omptempty"`
	TaxClosingEndDate                     *Time      `xmlrpc:"tax_closing_end_date,omptempty"`
	TaxCountryCode                        *String    `xmlrpc:"tax_country_code,omptempty"`
	TaxCountryId                          *Many2One  `xmlrpc:"tax_country_id,omptempty"`
	TaxLockDateMessage                    *String    `xmlrpc:"tax_lock_date_message,omptempty"`
	TaxReportControlError                 *Bool      `xmlrpc:"tax_report_control_error,omptempty"`
	TaxTotals                             *String    `xmlrpc:"tax_totals,omptempty"`
	TeamId                                *Many2One  `xmlrpc:"team_id,omptempty"`
	TimesheetCount                        *Int       `xmlrpc:"timesheet_count,omptempty"`
	TimesheetEncodeUomId                  *Many2One  `xmlrpc:"timesheet_encode_uom_id,omptempty"`
	TimesheetIds                          *Relation  `xmlrpc:"timesheet_ids,omptempty"`
	TimesheetTotalDuration                *Int       `xmlrpc:"timesheet_total_duration,omptempty"`
	ToCheck                               *Bool      `xmlrpc:"to_check,omptempty"`
	TransactionIds                        *Relation  `xmlrpc:"transaction_ids,omptempty"`
	TransferModelId                       *Many2One  `xmlrpc:"transfer_model_id,omptempty"`
	TypeName                              *String    `xmlrpc:"type_name,omptempty"`
	UserId                                *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds                     *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountMoves represents array of account.move model.
type AccountMoves []AccountMove

// AccountMoveModel is the odoo model name.
const AccountMoveModel = "account.move"

// Many2One convert AccountMove to *Many2One.
func (am *AccountMove) Many2One() *Many2One {
	return NewMany2One(am.Id.Get(), "")
}

// CreateAccountMove creates a new account.move model and returns its id.
func (c *Client) CreateAccountMove(am *AccountMove) (int64, error) {
	return c.Create(AccountMoveModel, am)
}

// UpdateAccountMove updates an existing account.move record.
func (c *Client) UpdateAccountMove(am *AccountMove) error {
	return c.UpdateAccountMoves([]int64{am.Id.Get()}, am)
}

// UpdateAccountMoves updates existing account.move records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAccountMoves(ids []int64, am *AccountMove) error {
	return c.Update(AccountMoveModel, ids, am)
}

// DeleteAccountMove deletes an existing account.move record.
func (c *Client) DeleteAccountMove(id int64) error {
	return c.DeleteAccountMoves([]int64{id})
}

// DeleteAccountMoves deletes existing account.move records.
func (c *Client) DeleteAccountMoves(ids []int64) error {
	return c.Delete(AccountMoveModel, ids)
}

// GetAccountMove gets account.move existing record.
func (c *Client) GetAccountMove(id int64) (*AccountMove, error) {
	ams, err := c.GetAccountMoves([]int64{id})
	if err != nil {
		return nil, err
	}
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.move not found", id)
}

// GetAccountMoves gets account.move existing records.
func (c *Client) GetAccountMoves(ids []int64) (*AccountMoves, error) {
	ams := &AccountMoves{}
	if err := c.Read(AccountMoveModel, ids, nil, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountMove finds account.move record by querying it with criteria.
func (c *Client) FindAccountMove(criteria *Criteria) (*AccountMove, error) {
	ams := &AccountMoves{}
	if err := c.SearchRead(AccountMoveModel, criteria, NewOptions().Limit(1), ams); err != nil {
		return nil, err
	}
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("account.move was not found")
}

// FindAccountMoves finds account.move records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoves(criteria *Criteria, options *Options) (*AccountMoves, error) {
	ams := &AccountMoves{}
	if err := c.SearchRead(AccountMoveModel, criteria, options, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountMoveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountMoveModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountMoveId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.move was not found")
}
