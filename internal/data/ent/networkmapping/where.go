// Code generated by ent, DO NOT EDIT.

package networkmapping

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldName, v))
}

// FkGatewayID applies equality check predicate on the "fk_gateway_id" field. It's identical to FkGatewayIDEQ.
func FkGatewayID(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldFkGatewayID, v))
}

// GatewayPort applies equality check predicate on the "gateway_port" field. It's identical to GatewayPortEQ.
func GatewayPort(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldGatewayPort, v))
}

// ComputerPort applies equality check predicate on the "computer_port" field. It's identical to ComputerPortEQ.
func ComputerPort(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldComputerPort, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldStatus, v))
}

// FkComputerID applies equality check predicate on the "fk_computer_id" field. It's identical to FkComputerIDEQ.
func FkComputerID(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldFkComputerID, v))
}

// FkUserID applies equality check predicate on the "fk_user_id" field. It's identical to FkUserIDEQ.
func FkUserID(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldFkUserID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldContainsFold(FieldName, v))
}

// FkGatewayIDEQ applies the EQ predicate on the "fk_gateway_id" field.
func FkGatewayIDEQ(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldFkGatewayID, v))
}

// FkGatewayIDNEQ applies the NEQ predicate on the "fk_gateway_id" field.
func FkGatewayIDNEQ(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNEQ(FieldFkGatewayID, v))
}

// FkGatewayIDIn applies the In predicate on the "fk_gateway_id" field.
func FkGatewayIDIn(vs ...uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldIn(FieldFkGatewayID, vs...))
}

// FkGatewayIDNotIn applies the NotIn predicate on the "fk_gateway_id" field.
func FkGatewayIDNotIn(vs ...uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNotIn(FieldFkGatewayID, vs...))
}

// FkGatewayIDGT applies the GT predicate on the "fk_gateway_id" field.
func FkGatewayIDGT(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGT(FieldFkGatewayID, v))
}

// FkGatewayIDGTE applies the GTE predicate on the "fk_gateway_id" field.
func FkGatewayIDGTE(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGTE(FieldFkGatewayID, v))
}

// FkGatewayIDLT applies the LT predicate on the "fk_gateway_id" field.
func FkGatewayIDLT(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLT(FieldFkGatewayID, v))
}

// FkGatewayIDLTE applies the LTE predicate on the "fk_gateway_id" field.
func FkGatewayIDLTE(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLTE(FieldFkGatewayID, v))
}

// GatewayPortEQ applies the EQ predicate on the "gateway_port" field.
func GatewayPortEQ(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldGatewayPort, v))
}

// GatewayPortNEQ applies the NEQ predicate on the "gateway_port" field.
func GatewayPortNEQ(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNEQ(FieldGatewayPort, v))
}

// GatewayPortIn applies the In predicate on the "gateway_port" field.
func GatewayPortIn(vs ...int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldIn(FieldGatewayPort, vs...))
}

// GatewayPortNotIn applies the NotIn predicate on the "gateway_port" field.
func GatewayPortNotIn(vs ...int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNotIn(FieldGatewayPort, vs...))
}

// GatewayPortGT applies the GT predicate on the "gateway_port" field.
func GatewayPortGT(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGT(FieldGatewayPort, v))
}

// GatewayPortGTE applies the GTE predicate on the "gateway_port" field.
func GatewayPortGTE(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGTE(FieldGatewayPort, v))
}

// GatewayPortLT applies the LT predicate on the "gateway_port" field.
func GatewayPortLT(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLT(FieldGatewayPort, v))
}

// GatewayPortLTE applies the LTE predicate on the "gateway_port" field.
func GatewayPortLTE(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLTE(FieldGatewayPort, v))
}

// ComputerPortEQ applies the EQ predicate on the "computer_port" field.
func ComputerPortEQ(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldComputerPort, v))
}

// ComputerPortNEQ applies the NEQ predicate on the "computer_port" field.
func ComputerPortNEQ(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNEQ(FieldComputerPort, v))
}

// ComputerPortIn applies the In predicate on the "computer_port" field.
func ComputerPortIn(vs ...int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldIn(FieldComputerPort, vs...))
}

// ComputerPortNotIn applies the NotIn predicate on the "computer_port" field.
func ComputerPortNotIn(vs ...int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNotIn(FieldComputerPort, vs...))
}

// ComputerPortGT applies the GT predicate on the "computer_port" field.
func ComputerPortGT(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGT(FieldComputerPort, v))
}

// ComputerPortGTE applies the GTE predicate on the "computer_port" field.
func ComputerPortGTE(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGTE(FieldComputerPort, v))
}

// ComputerPortLT applies the LT predicate on the "computer_port" field.
func ComputerPortLT(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLT(FieldComputerPort, v))
}

// ComputerPortLTE applies the LTE predicate on the "computer_port" field.
func ComputerPortLTE(v int32) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLTE(FieldComputerPort, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLTE(FieldStatus, v))
}

// FkComputerIDEQ applies the EQ predicate on the "fk_computer_id" field.
func FkComputerIDEQ(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldFkComputerID, v))
}

// FkComputerIDNEQ applies the NEQ predicate on the "fk_computer_id" field.
func FkComputerIDNEQ(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNEQ(FieldFkComputerID, v))
}

// FkComputerIDIn applies the In predicate on the "fk_computer_id" field.
func FkComputerIDIn(vs ...uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldIn(FieldFkComputerID, vs...))
}

// FkComputerIDNotIn applies the NotIn predicate on the "fk_computer_id" field.
func FkComputerIDNotIn(vs ...uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNotIn(FieldFkComputerID, vs...))
}

// FkComputerIDGT applies the GT predicate on the "fk_computer_id" field.
func FkComputerIDGT(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGT(FieldFkComputerID, v))
}

// FkComputerIDGTE applies the GTE predicate on the "fk_computer_id" field.
func FkComputerIDGTE(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGTE(FieldFkComputerID, v))
}

// FkComputerIDLT applies the LT predicate on the "fk_computer_id" field.
func FkComputerIDLT(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLT(FieldFkComputerID, v))
}

// FkComputerIDLTE applies the LTE predicate on the "fk_computer_id" field.
func FkComputerIDLTE(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLTE(FieldFkComputerID, v))
}

// FkUserIDEQ applies the EQ predicate on the "fk_user_id" field.
func FkUserIDEQ(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldEQ(FieldFkUserID, v))
}

// FkUserIDNEQ applies the NEQ predicate on the "fk_user_id" field.
func FkUserIDNEQ(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNEQ(FieldFkUserID, v))
}

// FkUserIDIn applies the In predicate on the "fk_user_id" field.
func FkUserIDIn(vs ...uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldIn(FieldFkUserID, vs...))
}

// FkUserIDNotIn applies the NotIn predicate on the "fk_user_id" field.
func FkUserIDNotIn(vs ...uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldNotIn(FieldFkUserID, vs...))
}

// FkUserIDGT applies the GT predicate on the "fk_user_id" field.
func FkUserIDGT(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGT(FieldFkUserID, v))
}

// FkUserIDGTE applies the GTE predicate on the "fk_user_id" field.
func FkUserIDGTE(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldGTE(FieldFkUserID, v))
}

// FkUserIDLT applies the LT predicate on the "fk_user_id" field.
func FkUserIDLT(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLT(FieldFkUserID, v))
}

// FkUserIDLTE applies the LTE predicate on the "fk_user_id" field.
func FkUserIDLTE(v uuid.UUID) predicate.NetworkMapping {
	return predicate.NetworkMapping(sql.FieldLTE(FieldFkUserID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NetworkMapping) predicate.NetworkMapping {
	return predicate.NetworkMapping(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NetworkMapping) predicate.NetworkMapping {
	return predicate.NetworkMapping(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NetworkMapping) predicate.NetworkMapping {
	return predicate.NetworkMapping(func(s *sql.Selector) {
		p(s.Not())
	})
}
