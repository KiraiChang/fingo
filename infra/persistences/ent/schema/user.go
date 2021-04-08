package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/KiraiChang/fingo/domain/entities"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive(),
		field.String("name").
			Annotations(
				entgql.OrderField(entities.UserNameField),
			),
		field.String("pwd"),
		field.String("email").
			Annotations(
				entgql.OrderField(entities.UserEmailField),
			),
		field.Int("age").
			Annotations(
				entgql.OrderField(entities.UserAgeField),
			),
	}
}
