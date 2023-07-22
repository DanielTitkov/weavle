// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SentencesColumns holds the columns for the "sentences" table.
	SentencesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "order", Type: field.TypeInt},
		{Name: "story_sentences", Type: field.TypeUUID},
	}
	// SentencesTable holds the schema information for the "sentences" table.
	SentencesTable = &schema.Table{
		Name:       "sentences",
		Columns:    SentencesColumns,
		PrimaryKey: []*schema.Column{SentencesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sentences_stories_sentences",
				Columns:    []*schema.Column{SentencesColumns[4]},
				RefColumns: []*schema.Column{StoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// StoriesColumns holds the columns for the "stories" table.
	StoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"open", "closed", "waiting"}, Default: "open"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// StoriesTable holds the schema information for the "stories" table.
	StoriesTable = &schema.Table{
		Name:       "stories",
		Columns:    StoriesColumns,
		PrimaryKey: []*schema.Column{StoriesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "ip_address", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserSentencesColumns holds the columns for the "user_sentences" table.
	UserSentencesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "sentence_id", Type: field.TypeUUID},
	}
	// UserSentencesTable holds the schema information for the "user_sentences" table.
	UserSentencesTable = &schema.Table{
		Name:       "user_sentences",
		Columns:    UserSentencesColumns,
		PrimaryKey: []*schema.Column{UserSentencesColumns[0], UserSentencesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_sentences_user_id",
				Columns:    []*schema.Column{UserSentencesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_sentences_sentence_id",
				Columns:    []*schema.Column{UserSentencesColumns[1]},
				RefColumns: []*schema.Column{SentencesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SentencesTable,
		StoriesTable,
		UsersTable,
		UserSentencesTable,
	}
)

func init() {
	SentencesTable.ForeignKeys[0].RefTable = StoriesTable
	UserSentencesTable.ForeignKeys[0].RefTable = UsersTable
	UserSentencesTable.ForeignKeys[1].RefTable = SentencesTable
}
