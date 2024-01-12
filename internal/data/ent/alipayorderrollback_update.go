// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/alipayorderrollback"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// AlipayOrderRollbackUpdate is the builder for updating AlipayOrderRollback entities.
type AlipayOrderRollbackUpdate struct {
	config
	hooks    []Hook
	mutation *AlipayOrderRollbackMutation
}

// Where appends a list predicates to the AlipayOrderRollbackUpdate builder.
func (aoru *AlipayOrderRollbackUpdate) Where(ps ...predicate.AlipayOrderRollback) *AlipayOrderRollbackUpdate {
	aoru.mutation.Where(ps...)
	return aoru
}

// SetNotifyID sets the "notify_id" field.
func (aoru *AlipayOrderRollbackUpdate) SetNotifyID(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetNotifyID(s)
	return aoru
}

// SetNotifyType sets the "notify_type" field.
func (aoru *AlipayOrderRollbackUpdate) SetNotifyType(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetNotifyType(s)
	return aoru
}

// SetNotifyTime sets the "notify_time" field.
func (aoru *AlipayOrderRollbackUpdate) SetNotifyTime(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetNotifyTime(s)
	return aoru
}

// SetCharset sets the "charset" field.
func (aoru *AlipayOrderRollbackUpdate) SetCharset(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetCharset(s)
	return aoru
}

// SetVersion sets the "version" field.
func (aoru *AlipayOrderRollbackUpdate) SetVersion(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetVersion(s)
	return aoru
}

// SetSignType sets the "sign_type" field.
func (aoru *AlipayOrderRollbackUpdate) SetSignType(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetSignType(s)
	return aoru
}

// SetSign sets the "sign" field.
func (aoru *AlipayOrderRollbackUpdate) SetSign(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetSign(s)
	return aoru
}

// SetFundBillList sets the "fund_bill_list" field.
func (aoru *AlipayOrderRollbackUpdate) SetFundBillList(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetFundBillList(s)
	return aoru
}

// SetReceiptAmount sets the "receipt_amount" field.
func (aoru *AlipayOrderRollbackUpdate) SetReceiptAmount(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetReceiptAmount(s)
	return aoru
}

// SetInvoiceAmount sets the "invoice_amount" field.
func (aoru *AlipayOrderRollbackUpdate) SetInvoiceAmount(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetInvoiceAmount(s)
	return aoru
}

// SetBuyerPayAmount sets the "buyer_pay_amount" field.
func (aoru *AlipayOrderRollbackUpdate) SetBuyerPayAmount(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetBuyerPayAmount(s)
	return aoru
}

// SetPointAmount sets the "point_amount" field.
func (aoru *AlipayOrderRollbackUpdate) SetPointAmount(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetPointAmount(s)
	return aoru
}

// SetVoucherDetailList sets the "voucher_detail_list" field.
func (aoru *AlipayOrderRollbackUpdate) SetVoucherDetailList(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetVoucherDetailList(s)
	return aoru
}

// SetPassbackParams sets the "passback_params" field.
func (aoru *AlipayOrderRollbackUpdate) SetPassbackParams(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetPassbackParams(s)
	return aoru
}

// SetTradeNo sets the "trade_no" field.
func (aoru *AlipayOrderRollbackUpdate) SetTradeNo(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetTradeNo(s)
	return aoru
}

// SetAppID sets the "app_id" field.
func (aoru *AlipayOrderRollbackUpdate) SetAppID(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetAppID(s)
	return aoru
}

// SetOutTradeNo sets the "out_trade_no" field.
func (aoru *AlipayOrderRollbackUpdate) SetOutTradeNo(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetOutTradeNo(s)
	return aoru
}

// SetOutBizNo sets the "out_biz_no" field.
func (aoru *AlipayOrderRollbackUpdate) SetOutBizNo(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetOutBizNo(s)
	return aoru
}

// SetBuyerID sets the "buyer_id" field.
func (aoru *AlipayOrderRollbackUpdate) SetBuyerID(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetBuyerID(s)
	return aoru
}

// SetSellerID sets the "seller_id" field.
func (aoru *AlipayOrderRollbackUpdate) SetSellerID(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetSellerID(s)
	return aoru
}

// SetTradeStatus sets the "trade_status" field.
func (aoru *AlipayOrderRollbackUpdate) SetTradeStatus(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetTradeStatus(s)
	return aoru
}

// SetTotalAmount sets the "total_amount" field.
func (aoru *AlipayOrderRollbackUpdate) SetTotalAmount(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetTotalAmount(s)
	return aoru
}

// SetRefundFee sets the "refund_fee" field.
func (aoru *AlipayOrderRollbackUpdate) SetRefundFee(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetRefundFee(s)
	return aoru
}

// SetSubject sets the "subject" field.
func (aoru *AlipayOrderRollbackUpdate) SetSubject(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetSubject(s)
	return aoru
}

// SetBody sets the "body" field.
func (aoru *AlipayOrderRollbackUpdate) SetBody(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetBody(s)
	return aoru
}

// SetGmtCreate sets the "gmt_create" field.
func (aoru *AlipayOrderRollbackUpdate) SetGmtCreate(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetGmtCreate(s)
	return aoru
}

// SetGmtPayment sets the "gmt_payment" field.
func (aoru *AlipayOrderRollbackUpdate) SetGmtPayment(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetGmtPayment(s)
	return aoru
}

// SetGmtClose sets the "gmt_close" field.
func (aoru *AlipayOrderRollbackUpdate) SetGmtClose(s string) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetGmtClose(s)
	return aoru
}

// SetCreateTime sets the "create_time" field.
func (aoru *AlipayOrderRollbackUpdate) SetCreateTime(t time.Time) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetCreateTime(t)
	return aoru
}

// SetUpdateTime sets the "update_time" field.
func (aoru *AlipayOrderRollbackUpdate) SetUpdateTime(t time.Time) *AlipayOrderRollbackUpdate {
	aoru.mutation.SetUpdateTime(t)
	return aoru
}

// Mutation returns the AlipayOrderRollbackMutation object of the builder.
func (aoru *AlipayOrderRollbackUpdate) Mutation() *AlipayOrderRollbackMutation {
	return aoru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aoru *AlipayOrderRollbackUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aoru.sqlSave, aoru.mutation, aoru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aoru *AlipayOrderRollbackUpdate) SaveX(ctx context.Context) int {
	affected, err := aoru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aoru *AlipayOrderRollbackUpdate) Exec(ctx context.Context) error {
	_, err := aoru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aoru *AlipayOrderRollbackUpdate) ExecX(ctx context.Context) {
	if err := aoru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aoru *AlipayOrderRollbackUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(alipayorderrollback.Table, alipayorderrollback.Columns, sqlgraph.NewFieldSpec(alipayorderrollback.FieldID, field.TypeInt))
	if ps := aoru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aoru.mutation.NotifyID(); ok {
		_spec.SetField(alipayorderrollback.FieldNotifyID, field.TypeString, value)
	}
	if value, ok := aoru.mutation.NotifyType(); ok {
		_spec.SetField(alipayorderrollback.FieldNotifyType, field.TypeString, value)
	}
	if value, ok := aoru.mutation.NotifyTime(); ok {
		_spec.SetField(alipayorderrollback.FieldNotifyTime, field.TypeString, value)
	}
	if value, ok := aoru.mutation.Charset(); ok {
		_spec.SetField(alipayorderrollback.FieldCharset, field.TypeString, value)
	}
	if value, ok := aoru.mutation.Version(); ok {
		_spec.SetField(alipayorderrollback.FieldVersion, field.TypeString, value)
	}
	if value, ok := aoru.mutation.SignType(); ok {
		_spec.SetField(alipayorderrollback.FieldSignType, field.TypeString, value)
	}
	if value, ok := aoru.mutation.Sign(); ok {
		_spec.SetField(alipayorderrollback.FieldSign, field.TypeString, value)
	}
	if value, ok := aoru.mutation.FundBillList(); ok {
		_spec.SetField(alipayorderrollback.FieldFundBillList, field.TypeString, value)
	}
	if value, ok := aoru.mutation.ReceiptAmount(); ok {
		_spec.SetField(alipayorderrollback.FieldReceiptAmount, field.TypeString, value)
	}
	if value, ok := aoru.mutation.InvoiceAmount(); ok {
		_spec.SetField(alipayorderrollback.FieldInvoiceAmount, field.TypeString, value)
	}
	if value, ok := aoru.mutation.BuyerPayAmount(); ok {
		_spec.SetField(alipayorderrollback.FieldBuyerPayAmount, field.TypeString, value)
	}
	if value, ok := aoru.mutation.PointAmount(); ok {
		_spec.SetField(alipayorderrollback.FieldPointAmount, field.TypeString, value)
	}
	if value, ok := aoru.mutation.VoucherDetailList(); ok {
		_spec.SetField(alipayorderrollback.FieldVoucherDetailList, field.TypeString, value)
	}
	if value, ok := aoru.mutation.PassbackParams(); ok {
		_spec.SetField(alipayorderrollback.FieldPassbackParams, field.TypeString, value)
	}
	if value, ok := aoru.mutation.TradeNo(); ok {
		_spec.SetField(alipayorderrollback.FieldTradeNo, field.TypeString, value)
	}
	if value, ok := aoru.mutation.AppID(); ok {
		_spec.SetField(alipayorderrollback.FieldAppID, field.TypeString, value)
	}
	if value, ok := aoru.mutation.OutTradeNo(); ok {
		_spec.SetField(alipayorderrollback.FieldOutTradeNo, field.TypeString, value)
	}
	if value, ok := aoru.mutation.OutBizNo(); ok {
		_spec.SetField(alipayorderrollback.FieldOutBizNo, field.TypeString, value)
	}
	if value, ok := aoru.mutation.BuyerID(); ok {
		_spec.SetField(alipayorderrollback.FieldBuyerID, field.TypeString, value)
	}
	if value, ok := aoru.mutation.SellerID(); ok {
		_spec.SetField(alipayorderrollback.FieldSellerID, field.TypeString, value)
	}
	if value, ok := aoru.mutation.TradeStatus(); ok {
		_spec.SetField(alipayorderrollback.FieldTradeStatus, field.TypeString, value)
	}
	if value, ok := aoru.mutation.TotalAmount(); ok {
		_spec.SetField(alipayorderrollback.FieldTotalAmount, field.TypeString, value)
	}
	if value, ok := aoru.mutation.RefundFee(); ok {
		_spec.SetField(alipayorderrollback.FieldRefundFee, field.TypeString, value)
	}
	if value, ok := aoru.mutation.Subject(); ok {
		_spec.SetField(alipayorderrollback.FieldSubject, field.TypeString, value)
	}
	if value, ok := aoru.mutation.Body(); ok {
		_spec.SetField(alipayorderrollback.FieldBody, field.TypeString, value)
	}
	if value, ok := aoru.mutation.GmtCreate(); ok {
		_spec.SetField(alipayorderrollback.FieldGmtCreate, field.TypeString, value)
	}
	if value, ok := aoru.mutation.GmtPayment(); ok {
		_spec.SetField(alipayorderrollback.FieldGmtPayment, field.TypeString, value)
	}
	if value, ok := aoru.mutation.GmtClose(); ok {
		_spec.SetField(alipayorderrollback.FieldGmtClose, field.TypeString, value)
	}
	if value, ok := aoru.mutation.CreateTime(); ok {
		_spec.SetField(alipayorderrollback.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := aoru.mutation.UpdateTime(); ok {
		_spec.SetField(alipayorderrollback.FieldUpdateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aoru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{alipayorderrollback.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aoru.mutation.done = true
	return n, nil
}

// AlipayOrderRollbackUpdateOne is the builder for updating a single AlipayOrderRollback entity.
type AlipayOrderRollbackUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AlipayOrderRollbackMutation
}

// SetNotifyID sets the "notify_id" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetNotifyID(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetNotifyID(s)
	return aoruo
}

// SetNotifyType sets the "notify_type" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetNotifyType(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetNotifyType(s)
	return aoruo
}

// SetNotifyTime sets the "notify_time" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetNotifyTime(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetNotifyTime(s)
	return aoruo
}

// SetCharset sets the "charset" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetCharset(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetCharset(s)
	return aoruo
}

// SetVersion sets the "version" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetVersion(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetVersion(s)
	return aoruo
}

// SetSignType sets the "sign_type" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetSignType(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetSignType(s)
	return aoruo
}

// SetSign sets the "sign" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetSign(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetSign(s)
	return aoruo
}

// SetFundBillList sets the "fund_bill_list" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetFundBillList(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetFundBillList(s)
	return aoruo
}

// SetReceiptAmount sets the "receipt_amount" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetReceiptAmount(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetReceiptAmount(s)
	return aoruo
}

// SetInvoiceAmount sets the "invoice_amount" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetInvoiceAmount(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetInvoiceAmount(s)
	return aoruo
}

// SetBuyerPayAmount sets the "buyer_pay_amount" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetBuyerPayAmount(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetBuyerPayAmount(s)
	return aoruo
}

// SetPointAmount sets the "point_amount" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetPointAmount(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetPointAmount(s)
	return aoruo
}

// SetVoucherDetailList sets the "voucher_detail_list" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetVoucherDetailList(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetVoucherDetailList(s)
	return aoruo
}

// SetPassbackParams sets the "passback_params" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetPassbackParams(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetPassbackParams(s)
	return aoruo
}

// SetTradeNo sets the "trade_no" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetTradeNo(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetTradeNo(s)
	return aoruo
}

// SetAppID sets the "app_id" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetAppID(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetAppID(s)
	return aoruo
}

// SetOutTradeNo sets the "out_trade_no" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetOutTradeNo(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetOutTradeNo(s)
	return aoruo
}

// SetOutBizNo sets the "out_biz_no" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetOutBizNo(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetOutBizNo(s)
	return aoruo
}

// SetBuyerID sets the "buyer_id" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetBuyerID(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetBuyerID(s)
	return aoruo
}

// SetSellerID sets the "seller_id" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetSellerID(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetSellerID(s)
	return aoruo
}

// SetTradeStatus sets the "trade_status" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetTradeStatus(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetTradeStatus(s)
	return aoruo
}

// SetTotalAmount sets the "total_amount" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetTotalAmount(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetTotalAmount(s)
	return aoruo
}

// SetRefundFee sets the "refund_fee" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetRefundFee(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetRefundFee(s)
	return aoruo
}

// SetSubject sets the "subject" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetSubject(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetSubject(s)
	return aoruo
}

// SetBody sets the "body" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetBody(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetBody(s)
	return aoruo
}

// SetGmtCreate sets the "gmt_create" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetGmtCreate(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetGmtCreate(s)
	return aoruo
}

// SetGmtPayment sets the "gmt_payment" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetGmtPayment(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetGmtPayment(s)
	return aoruo
}

// SetGmtClose sets the "gmt_close" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetGmtClose(s string) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetGmtClose(s)
	return aoruo
}

// SetCreateTime sets the "create_time" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetCreateTime(t time.Time) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetCreateTime(t)
	return aoruo
}

// SetUpdateTime sets the "update_time" field.
func (aoruo *AlipayOrderRollbackUpdateOne) SetUpdateTime(t time.Time) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.SetUpdateTime(t)
	return aoruo
}

// Mutation returns the AlipayOrderRollbackMutation object of the builder.
func (aoruo *AlipayOrderRollbackUpdateOne) Mutation() *AlipayOrderRollbackMutation {
	return aoruo.mutation
}

// Where appends a list predicates to the AlipayOrderRollbackUpdate builder.
func (aoruo *AlipayOrderRollbackUpdateOne) Where(ps ...predicate.AlipayOrderRollback) *AlipayOrderRollbackUpdateOne {
	aoruo.mutation.Where(ps...)
	return aoruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aoruo *AlipayOrderRollbackUpdateOne) Select(field string, fields ...string) *AlipayOrderRollbackUpdateOne {
	aoruo.fields = append([]string{field}, fields...)
	return aoruo
}

// Save executes the query and returns the updated AlipayOrderRollback entity.
func (aoruo *AlipayOrderRollbackUpdateOne) Save(ctx context.Context) (*AlipayOrderRollback, error) {
	return withHooks(ctx, aoruo.sqlSave, aoruo.mutation, aoruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aoruo *AlipayOrderRollbackUpdateOne) SaveX(ctx context.Context) *AlipayOrderRollback {
	node, err := aoruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aoruo *AlipayOrderRollbackUpdateOne) Exec(ctx context.Context) error {
	_, err := aoruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aoruo *AlipayOrderRollbackUpdateOne) ExecX(ctx context.Context) {
	if err := aoruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aoruo *AlipayOrderRollbackUpdateOne) sqlSave(ctx context.Context) (_node *AlipayOrderRollback, err error) {
	_spec := sqlgraph.NewUpdateSpec(alipayorderrollback.Table, alipayorderrollback.Columns, sqlgraph.NewFieldSpec(alipayorderrollback.FieldID, field.TypeInt))
	id, ok := aoruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AlipayOrderRollback.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aoruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, alipayorderrollback.FieldID)
		for _, f := range fields {
			if !alipayorderrollback.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != alipayorderrollback.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aoruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aoruo.mutation.NotifyID(); ok {
		_spec.SetField(alipayorderrollback.FieldNotifyID, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.NotifyType(); ok {
		_spec.SetField(alipayorderrollback.FieldNotifyType, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.NotifyTime(); ok {
		_spec.SetField(alipayorderrollback.FieldNotifyTime, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.Charset(); ok {
		_spec.SetField(alipayorderrollback.FieldCharset, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.Version(); ok {
		_spec.SetField(alipayorderrollback.FieldVersion, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.SignType(); ok {
		_spec.SetField(alipayorderrollback.FieldSignType, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.Sign(); ok {
		_spec.SetField(alipayorderrollback.FieldSign, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.FundBillList(); ok {
		_spec.SetField(alipayorderrollback.FieldFundBillList, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.ReceiptAmount(); ok {
		_spec.SetField(alipayorderrollback.FieldReceiptAmount, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.InvoiceAmount(); ok {
		_spec.SetField(alipayorderrollback.FieldInvoiceAmount, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.BuyerPayAmount(); ok {
		_spec.SetField(alipayorderrollback.FieldBuyerPayAmount, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.PointAmount(); ok {
		_spec.SetField(alipayorderrollback.FieldPointAmount, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.VoucherDetailList(); ok {
		_spec.SetField(alipayorderrollback.FieldVoucherDetailList, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.PassbackParams(); ok {
		_spec.SetField(alipayorderrollback.FieldPassbackParams, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.TradeNo(); ok {
		_spec.SetField(alipayorderrollback.FieldTradeNo, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.AppID(); ok {
		_spec.SetField(alipayorderrollback.FieldAppID, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.OutTradeNo(); ok {
		_spec.SetField(alipayorderrollback.FieldOutTradeNo, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.OutBizNo(); ok {
		_spec.SetField(alipayorderrollback.FieldOutBizNo, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.BuyerID(); ok {
		_spec.SetField(alipayorderrollback.FieldBuyerID, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.SellerID(); ok {
		_spec.SetField(alipayorderrollback.FieldSellerID, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.TradeStatus(); ok {
		_spec.SetField(alipayorderrollback.FieldTradeStatus, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.TotalAmount(); ok {
		_spec.SetField(alipayorderrollback.FieldTotalAmount, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.RefundFee(); ok {
		_spec.SetField(alipayorderrollback.FieldRefundFee, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.Subject(); ok {
		_spec.SetField(alipayorderrollback.FieldSubject, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.Body(); ok {
		_spec.SetField(alipayorderrollback.FieldBody, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.GmtCreate(); ok {
		_spec.SetField(alipayorderrollback.FieldGmtCreate, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.GmtPayment(); ok {
		_spec.SetField(alipayorderrollback.FieldGmtPayment, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.GmtClose(); ok {
		_spec.SetField(alipayorderrollback.FieldGmtClose, field.TypeString, value)
	}
	if value, ok := aoruo.mutation.CreateTime(); ok {
		_spec.SetField(alipayorderrollback.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := aoruo.mutation.UpdateTime(); ok {
		_spec.SetField(alipayorderrollback.FieldUpdateTime, field.TypeTime, value)
	}
	_node = &AlipayOrderRollback{config: aoruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aoruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{alipayorderrollback.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aoruo.mutation.done = true
	return _node, nil
}
