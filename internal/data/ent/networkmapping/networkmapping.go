// Code generated by ent, DO NOT EDIT.

package networkmapping

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the networkmapping type in the database.
	Label = "network_mapping"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldProtocol holds the string denoting the protocol field in the database.
	FieldProtocol = "protocol"
	// FieldFkGatewayID holds the string denoting the fk_gateway_id field in the database.
	FieldFkGatewayID = "fk_gateway_id"
	// FieldGatewayPort holds the string denoting the gateway_port field in the database.
	FieldGatewayPort = "gateway_port"
	// FieldGatewayIP holds the string denoting the gateway_ip field in the database.
	FieldGatewayIP = "gateway_ip"
	// FieldComputerPort holds the string denoting the computer_port field in the database.
	FieldComputerPort = "computer_port"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldFkComputerID holds the string denoting the fk_computer_id field in the database.
	FieldFkComputerID = "fk_computer_id"
	// FieldFkUserID holds the string denoting the fk_user_id field in the database.
	FieldFkUserID = "fk_user_id"
	// FieldDeleteState holds the string denoting the delete_state field in the database.
	FieldDeleteState = "delete_state"
	// Table holds the table name of the networkmapping in the database.
	Table = "network_mappings"
)

// Columns holds all SQL columns for networkmapping fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldProtocol,
	FieldFkGatewayID,
	FieldGatewayPort,
	FieldGatewayIP,
	FieldComputerPort,
	FieldStatus,
	FieldFkComputerID,
	FieldFkUserID,
	FieldDeleteState,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultProtocol holds the default value on creation for the "protocol" field.
	DefaultProtocol string
	// ProtocolValidator is a validator for the "protocol" field. It is called by the builders before save.
	ProtocolValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
	// DefaultDeleteState holds the default value on creation for the "delete_state" field.
	DefaultDeleteState bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the NetworkMapping queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByProtocol orders the results by the protocol field.
func ByProtocol(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProtocol, opts...).ToFunc()
}

// ByFkGatewayID orders the results by the fk_gateway_id field.
func ByFkGatewayID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFkGatewayID, opts...).ToFunc()
}

// ByGatewayPort orders the results by the gateway_port field.
func ByGatewayPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGatewayPort, opts...).ToFunc()
}

// ByGatewayIP orders the results by the gateway_ip field.
func ByGatewayIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGatewayIP, opts...).ToFunc()
}

// ByComputerPort orders the results by the computer_port field.
func ByComputerPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComputerPort, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByFkComputerID orders the results by the fk_computer_id field.
func ByFkComputerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFkComputerID, opts...).ToFunc()
}

// ByFkUserID orders the results by the fk_user_id field.
func ByFkUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFkUserID, opts...).ToFunc()
}

// ByDeleteState orders the results by the delete_state field.
func ByDeleteState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteState, opts...).ToFunc()
}
