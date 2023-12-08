// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/s3user"
)

// S3User is the model entity for the S3User schema.
type S3User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
	// accessKey
	AccessKey string `json:"access_key,omitempty"`
	// secretKey
	SecretKey string `json:"secret_key,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the S3UserQuery when eager-loading is set.
	Edges        S3UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// S3UserEdges holds the relations/edges for other nodes in the graph.
type S3UserEdges struct {
	// Buckets holds the value of the buckets edge.
	Buckets []*S3Bucket `json:"buckets,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BucketsOrErr returns the Buckets value or an error if the edge
// was not loaded in eager-loading.
func (e S3UserEdges) BucketsOrErr() ([]*S3Bucket, error) {
	if e.loadedTypes[0] {
		return e.Buckets, nil
	}
	return nil, &NotLoadedError{edge: "buckets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*S3User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case s3user.FieldAccessKey, s3user.FieldSecretKey:
			values[i] = new(sql.NullString)
		case s3user.FieldID, s3user.FieldFkUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the S3User fields.
func (s *S3User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case s3user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case s3user.FieldFkUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field fk_user_id", values[i])
			} else if value != nil {
				s.FkUserID = *value
			}
		case s3user.FieldAccessKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_key", values[i])
			} else if value.Valid {
				s.AccessKey = value.String
			}
		case s3user.FieldSecretKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_key", values[i])
			} else if value.Valid {
				s.SecretKey = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the S3User.
// This includes values selected through modifiers, order, etc.
func (s *S3User) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryBuckets queries the "buckets" edge of the S3User entity.
func (s *S3User) QueryBuckets() *S3BucketQuery {
	return NewS3UserClient(s.config).QueryBuckets(s)
}

// Update returns a builder for updating this S3User.
// Note that you need to call S3User.Unwrap() before calling this method if this S3User
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *S3User) Update() *S3UserUpdateOne {
	return NewS3UserClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the S3User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *S3User) Unwrap() *S3User {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: S3User is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *S3User) String() string {
	var builder strings.Builder
	builder.WriteString("S3User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("fk_user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.FkUserID))
	builder.WriteString(", ")
	builder.WriteString("access_key=")
	builder.WriteString(s.AccessKey)
	builder.WriteString(", ")
	builder.WriteString("secret_key=")
	builder.WriteString(s.SecretKey)
	builder.WriteByte(')')
	return builder.String()
}

// S3Users is a parsable slice of S3User.
type S3Users []*S3User