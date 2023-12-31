// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIPAddress holds the string denoting the ip_address field in the database.
	FieldIPAddress = "ip_address"
	// EdgeSentences holds the string denoting the sentences edge name in mutations.
	EdgeSentences = "sentences"
	// Table holds the table name of the user in the database.
	Table = "users"
	// SentencesTable is the table that holds the sentences relation/edge. The primary key declared below.
	SentencesTable = "user_sentences"
	// SentencesInverseTable is the table name for the Sentence entity.
	// It exists in this package in order to avoid circular dependency with the "sentence" package.
	SentencesInverseTable = "sentences"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldIPAddress,
}

var (
	// SentencesPrimaryKey and SentencesColumn2 are the table columns denoting the
	// primary key for the sentences relation (M2M).
	SentencesPrimaryKey = []string{"user_id", "sentence_id"}
)

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
	// IPAddressValidator is a validator for the "ip_address" field. It is called by the builders before save.
	IPAddressValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIPAddress orders the results by the ip_address field.
func ByIPAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIPAddress, opts...).ToFunc()
}

// BySentencesCount orders the results by sentences count.
func BySentencesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSentencesStep(), opts...)
	}
}

// BySentences orders the results by sentences terms.
func BySentences(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSentencesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSentencesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SentencesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SentencesTable, SentencesPrimaryKey...),
	)
}
