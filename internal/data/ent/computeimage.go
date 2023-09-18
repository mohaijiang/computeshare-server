// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computeimage"
)

// ComputeImage is the model entity for the ComputeImage schema.
type ComputeImage struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// 显示名
	Name string `json:"name,omitempty"`
	// 镜像名
	Image string `json:"image,omitempty"`
	// 版本名
	Tag string `json:"tag,omitempty"`
	// 端口号
	Port int32 `json:"port,omitempty"`
	// 容器命令
	Command      string `json:"command,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ComputeImage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case computeimage.FieldID, computeimage.FieldPort:
			values[i] = new(sql.NullInt64)
		case computeimage.FieldName, computeimage.FieldImage, computeimage.FieldTag, computeimage.FieldCommand:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ComputeImage fields.
func (ci *ComputeImage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case computeimage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ci.ID = int32(value.Int64)
		case computeimage.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ci.Name = value.String
			}
		case computeimage.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				ci.Image = value.String
			}
		case computeimage.FieldTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag", values[i])
			} else if value.Valid {
				ci.Tag = value.String
			}
		case computeimage.FieldPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field port", values[i])
			} else if value.Valid {
				ci.Port = int32(value.Int64)
			}
		case computeimage.FieldCommand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field command", values[i])
			} else if value.Valid {
				ci.Command = value.String
			}
		default:
			ci.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ComputeImage.
// This includes values selected through modifiers, order, etc.
func (ci *ComputeImage) Value(name string) (ent.Value, error) {
	return ci.selectValues.Get(name)
}

// Update returns a builder for updating this ComputeImage.
// Note that you need to call ComputeImage.Unwrap() before calling this method if this ComputeImage
// was returned from a transaction, and the transaction was committed or rolled back.
func (ci *ComputeImage) Update() *ComputeImageUpdateOne {
	return NewComputeImageClient(ci.config).UpdateOne(ci)
}

// Unwrap unwraps the ComputeImage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ci *ComputeImage) Unwrap() *ComputeImage {
	_tx, ok := ci.config.driver.(*txDriver)
	if !ok {
		panic("ent: ComputeImage is not a transactional entity")
	}
	ci.config.driver = _tx.drv
	return ci
}

// String implements the fmt.Stringer.
func (ci *ComputeImage) String() string {
	var builder strings.Builder
	builder.WriteString("ComputeImage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ci.ID))
	builder.WriteString("name=")
	builder.WriteString(ci.Name)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(ci.Image)
	builder.WriteString(", ")
	builder.WriteString("tag=")
	builder.WriteString(ci.Tag)
	builder.WriteString(", ")
	builder.WriteString("port=")
	builder.WriteString(fmt.Sprintf("%v", ci.Port))
	builder.WriteString(", ")
	builder.WriteString("command=")
	builder.WriteString(ci.Command)
	builder.WriteByte(')')
	return builder.String()
}

// ComputeImages is a parsable slice of ComputeImage.
type ComputeImages []*ComputeImage
