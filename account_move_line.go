package odoo

import (
	"fmt"
)

// AccountMoveLine represents account.move.line model.
type AccountMoveLine struct {
	LastUpdate             *Time       `xmlrpc:"__last_update,omptempty"`
	AccountId              *Many2One   `xmlrpc:"account_id,omptempty"`
	AccountInternalGroup   *Selection  `xmlrpc:"account_internal_group,omptempty"`
	AccountRootId          *Many2One   `xmlrpc:"account_root_id,omptempty"`
	AccountType            *Selection  `xmlrpc:"account_type,omptempty"`
	AmountCurrency         *Float      `xmlrpc:"amount_currency,omptempty"`
	AmountResidual         *Float      `xmlrpc:"amount_residual,omptempty"`
	AmountResidualCurrency *Float      `xmlrpc:"amount_residual_currency,omptempty"`
	AnalyticDistribution   interface{} `xmlrpc:"analytic_distribution,omptempty"`
	AnalyticLineIds        *Relation   `xmlrpc:"analytic_line_ids,omptempty"`
	AnalyticPrecision      *Int        `xmlrpc:"analytic_precision,omptempty"`
	AssetIds               *Relation   `xmlrpc:"asset_ids,omptempty"`
	Balance                *Float      `xmlrpc:"balance,omptempty"`
	Blocked                *Bool       `xmlrpc:"blocked,omptempty"`
	CompanyCurrencyId      *Many2One   `xmlrpc:"company_currency_id,omptempty"`
	CompanyId              *Many2One   `xmlrpc:"company_id,omptempty"`
	ComputeAllTax          *String     `xmlrpc:"compute_all_tax,omptempty"`
	ComputeAllTaxDirty     *Bool       `xmlrpc:"compute_all_tax_dirty,omptempty"`
	CreateDate             *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One   `xmlrpc:"create_uid,omptempty"`
	Credit                 *Float      `xmlrpc:"credit,omptempty"`
	CumulatedBalance       *Float      `xmlrpc:"cumulated_balance,omptempty"`
	CurrencyId             *Many2One   `xmlrpc:"currency_id,omptempty"`
	CurrencyRate           *Float      `xmlrpc:"currency_rate,omptempty"`
	Date                   *Time       `xmlrpc:"date,omptempty"`
	DateMaturity           *Time       `xmlrpc:"date_maturity,omptempty"`
	Debit                  *Float      `xmlrpc:"debit,omptempty"`
	Discount               *Float      `xmlrpc:"discount,omptempty"`
	DiscountAmountCurrency *Float      `xmlrpc:"discount_amount_currency,omptempty"`
	DiscountBalance        *Float      `xmlrpc:"discount_balance,omptempty"`
	DiscountDate           *Time       `xmlrpc:"discount_date,omptempty"`
	DiscountPercentage     *Float      `xmlrpc:"discount_percentage,omptempty"`
	DisplayName            *String     `xmlrpc:"display_name,omptempty"`
	DisplayType            *Selection  `xmlrpc:"display_type,omptempty"`
	EpdDirty               *Bool       `xmlrpc:"epd_dirty,omptempty"`
	EpdKey                 *String     `xmlrpc:"epd_key,omptempty"`
	EpdNeeded              *String     `xmlrpc:"epd_needed,omptempty"`
	ExpectedPayDate        *Time       `xmlrpc:"expected_pay_date,omptempty"`
	ExpenseId              *Many2One   `xmlrpc:"expense_id,omptempty"`
	FullReconcileId        *Many2One   `xmlrpc:"full_reconcile_id,omptempty"`
	GroupTaxId             *Many2One   `xmlrpc:"group_tax_id,omptempty"`
	Id                     *Int        `xmlrpc:"id,omptempty"`
	IsAccountReconcile     *Bool       `xmlrpc:"is_account_reconcile,omptempty"`
	IsDownpayment          *Bool       `xmlrpc:"is_downpayment,omptempty"`
	IsRefund               *Bool       `xmlrpc:"is_refund,omptempty"`
	IsSameCurrency         *Bool       `xmlrpc:"is_same_currency,omptempty"`
	IsStorno               *Bool       `xmlrpc:"is_storno,omptempty"`
	JournalId              *Many2One   `xmlrpc:"journal_id,omptempty"`
	MatchedCreditIds       *Relation   `xmlrpc:"matched_credit_ids,omptempty"`
	MatchedDebitIds        *Relation   `xmlrpc:"matched_debit_ids,omptempty"`
	MatchingNumber         *String     `xmlrpc:"matching_number,omptempty"`
	MoveAttachmentIds      *Relation   `xmlrpc:"move_attachment_ids,omptempty"`
	MoveId                 *Many2One   `xmlrpc:"move_id,omptempty"`
	MoveName               *String     `xmlrpc:"move_name,omptempty"`
	MoveType               *Selection  `xmlrpc:"move_type,omptempty"`
	Name                   *String     `xmlrpc:"name,omptempty"`
	NonDeductibleTaxValue  *Float      `xmlrpc:"non_deductible_tax_value,omptempty"`
	ParentState            *Selection  `xmlrpc:"parent_state,omptempty"`
	PartnerId              *Many2One   `xmlrpc:"partner_id,omptempty"`
	PaymentId              *Many2One   `xmlrpc:"payment_id,omptempty"`
	PriceSubtotal          *Float      `xmlrpc:"price_subtotal,omptempty"`
	PriceTotal             *Float      `xmlrpc:"price_total,omptempty"`
	PriceUnit              *Float      `xmlrpc:"price_unit,omptempty"`
	ProductId              *Many2One   `xmlrpc:"product_id,omptempty"`
	ProductUomCategoryId   *Many2One   `xmlrpc:"product_uom_category_id,omptempty"`
	ProductUomId           *Many2One   `xmlrpc:"product_uom_id,omptempty"`
	PurchaseLineId         *Many2One   `xmlrpc:"purchase_line_id,omptempty"`
	PurchaseOrderId        *Many2One   `xmlrpc:"purchase_order_id,omptempty"`
	Quantity               *Float      `xmlrpc:"quantity,omptempty"`
	ReconcileModelId       *Many2One   `xmlrpc:"reconcile_model_id,omptempty"`
	Reconciled             *Bool       `xmlrpc:"reconciled,omptempty"`
	Ref                    *String     `xmlrpc:"ref,omptempty"`
	SaleLineIds            *Relation   `xmlrpc:"sale_line_ids,omptempty"`
	Sequence               *Int        `xmlrpc:"sequence,omptempty"`
	StatementId            *Many2One   `xmlrpc:"statement_id,omptempty"`
	StatementLineId        *Many2One   `xmlrpc:"statement_line_id,omptempty"`
	StockValuationLayerIds *Relation   `xmlrpc:"stock_valuation_layer_ids,omptempty"`
	TaxAudit               *String     `xmlrpc:"tax_audit,omptempty"`
	TaxBaseAmount          *Float      `xmlrpc:"tax_base_amount,omptempty"`
	TaxGroupId             *Many2One   `xmlrpc:"tax_group_id,omptempty"`
	TaxIds                 *Relation   `xmlrpc:"tax_ids,omptempty"`
	TaxKey                 *String     `xmlrpc:"tax_key,omptempty"`
	TaxLineId              *Many2One   `xmlrpc:"tax_line_id,omptempty"`
	TaxRepartitionLineId   *Many2One   `xmlrpc:"tax_repartition_line_id,omptempty"`
	TaxTagIds              *Relation   `xmlrpc:"tax_tag_ids,omptempty"`
	TaxTagInvert           *Bool       `xmlrpc:"tax_tag_invert,omptempty"`
	TermKey                *String     `xmlrpc:"term_key,omptempty"`
	WriteDate              *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// AccountMoveLines represents array of account.move.line model.
type AccountMoveLines []AccountMoveLine

// AccountMoveLineModel is the odoo model name.
const AccountMoveLineModel = "account.move.line"

// Many2One convert AccountMoveLine to *Many2One.
func (aml *AccountMoveLine) Many2One() *Many2One {
	return NewMany2One(aml.Id.Get(), "")
}

// CreateAccountMoveLine creates a new account.move.line model and returns its id.
func (c *Client) CreateAccountMoveLine(aml *AccountMoveLine) (int64, error) {
	return c.Create(AccountMoveLineModel, aml)
}

// UpdateAccountMoveLine updates an existing account.move.line record.
func (c *Client) UpdateAccountMoveLine(aml *AccountMoveLine) error {
	return c.UpdateAccountMoveLines([]int64{aml.Id.Get()}, aml)
}

// UpdateAccountMoveLines updates existing account.move.line records.
// All records (represented by ids) will be updated by aml values.
func (c *Client) UpdateAccountMoveLines(ids []int64, aml *AccountMoveLine) error {
	return c.Update(AccountMoveLineModel, ids, aml)
}

// DeleteAccountMoveLine deletes an existing account.move.line record.
func (c *Client) DeleteAccountMoveLine(id int64) error {
	return c.DeleteAccountMoveLines([]int64{id})
}

// DeleteAccountMoveLines deletes existing account.move.line records.
func (c *Client) DeleteAccountMoveLines(ids []int64) error {
	return c.Delete(AccountMoveLineModel, ids)
}

// GetAccountMoveLine gets account.move.line existing record.
func (c *Client) GetAccountMoveLine(id int64) (*AccountMoveLine, error) {
	amls, err := c.GetAccountMoveLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if amls != nil && len(*amls) > 0 {
		return &((*amls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.move.line not found", id)
}

// GetAccountMoveLines gets account.move.line existing records.
func (c *Client) GetAccountMoveLines(ids []int64) (*AccountMoveLines, error) {
	amls := &AccountMoveLines{}
	if err := c.Read(AccountMoveLineModel, ids, nil, amls); err != nil {
		return nil, err
	}
	return amls, nil
}

// FindAccountMoveLine finds account.move.line record by querying it with criteria.
func (c *Client) FindAccountMoveLine(criteria *Criteria) (*AccountMoveLine, error) {
	amls := &AccountMoveLines{}
	if err := c.SearchRead(AccountMoveLineModel, criteria, NewOptions().Limit(1), amls); err != nil {
		return nil, err
	}
	if amls != nil && len(*amls) > 0 {
		return &((*amls)[0]), nil
	}
	return nil, fmt.Errorf("account.move.line was not found")
}

// FindAccountMoveLines finds account.move.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveLines(criteria *Criteria, options *Options) (*AccountMoveLines, error) {
	amls := &AccountMoveLines{}
	if err := c.SearchRead(AccountMoveLineModel, criteria, options, amls); err != nil {
		return nil, err
	}
	return amls, nil
}

// FindAccountMoveLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountMoveLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountMoveLineId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.move.line was not found")
}
