// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/DanielTitkov/weavle/ent/schema"
	"github.com/DanielTitkov/weavle/ent/sentence"
	"github.com/DanielTitkov/weavle/ent/story"
	"github.com/DanielTitkov/weavle/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sentenceFields := schema.Sentence{}.Fields()
	_ = sentenceFields
	// sentenceDescText is the schema descriptor for text field.
	sentenceDescText := sentenceFields[1].Descriptor()
	// sentence.TextValidator is a validator for the "text" field. It is called by the builders before save.
	sentence.TextValidator = sentenceDescText.Validators[0].(func(string) error)
	// sentenceDescCreatedAt is the schema descriptor for created_at field.
	sentenceDescCreatedAt := sentenceFields[2].Descriptor()
	// sentence.DefaultCreatedAt holds the default value on creation for the created_at field.
	sentence.DefaultCreatedAt = sentenceDescCreatedAt.Default.(func() time.Time)
	// sentenceDescOrder is the schema descriptor for order field.
	sentenceDescOrder := sentenceFields[3].Descriptor()
	// sentence.OrderValidator is a validator for the "order" field. It is called by the builders before save.
	sentence.OrderValidator = sentenceDescOrder.Validators[0].(func(int) error)
	// sentenceDescID is the schema descriptor for id field.
	sentenceDescID := sentenceFields[0].Descriptor()
	// sentence.DefaultID holds the default value on creation for the id field.
	sentence.DefaultID = sentenceDescID.Default.(func() uuid.UUID)
	storyFields := schema.Story{}.Fields()
	_ = storyFields
	// storyDescCreatedAt is the schema descriptor for created_at field.
	storyDescCreatedAt := storyFields[2].Descriptor()
	// story.DefaultCreatedAt holds the default value on creation for the created_at field.
	story.DefaultCreatedAt = storyDescCreatedAt.Default.(func() time.Time)
	// storyDescUpdatedAt is the schema descriptor for updated_at field.
	storyDescUpdatedAt := storyFields[3].Descriptor()
	// story.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	story.UpdateDefaultUpdatedAt = storyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// storyDescID is the schema descriptor for id field.
	storyDescID := storyFields[0].Descriptor()
	// story.DefaultID holds the default value on creation for the id field.
	story.DefaultID = storyDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescIPAddress is the schema descriptor for ip_address field.
	userDescIPAddress := userFields[1].Descriptor()
	// user.IPAddressValidator is a validator for the "ip_address" field. It is called by the builders before save.
	user.IPAddressValidator = userDescIPAddress.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
