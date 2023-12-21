// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/s3bucket"
)

// S3Bucket is the model entity for the S3Bucket schema.
type S3Bucket struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
	// bucketName
	BucketName string `json:"bucket_name,omitempty"`
	// CreatedTime holds the value of the "createdTime" field.
	CreatedTime  time.Time `json:"createdTime,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*S3Bucket) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case s3bucket.FieldBucketName:
			values[i] = new(sql.NullString)
		case s3bucket.FieldCreatedTime:
			values[i] = new(sql.NullTime)
		case s3bucket.FieldID, s3bucket.FieldFkUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the S3Bucket fields.
func (s *S3Bucket) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case s3bucket.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case s3bucket.FieldFkUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field fk_user_id", values[i])
			} else if value != nil {
				s.FkUserID = *value
			}
		case s3bucket.FieldBucketName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bucket_name", values[i])
			} else if value.Valid {
				s.BucketName = value.String
			}
		case s3bucket.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdTime", values[i])
			} else if value.Valid {
				s.CreatedTime = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the S3Bucket.
// This includes values selected through modifiers, order, etc.
func (s *S3Bucket) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this S3Bucket.
// Note that you need to call S3Bucket.Unwrap() before calling this method if this S3Bucket
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *S3Bucket) Update() *S3BucketUpdateOne {
	return NewS3BucketClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the S3Bucket entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *S3Bucket) Unwrap() *S3Bucket {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: S3Bucket is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *S3Bucket) String() string {
	var builder strings.Builder
	builder.WriteString("S3Bucket(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("fk_user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.FkUserID))
	builder.WriteString(", ")
	builder.WriteString("bucket_name=")
	builder.WriteString(s.BucketName)
	builder.WriteString(", ")
	builder.WriteString("createdTime=")
	builder.WriteString(s.CreatedTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// S3Buckets is a parsable slice of S3Bucket.
type S3Buckets []*S3Bucket
