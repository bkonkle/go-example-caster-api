package tests

import (
	"testing"

	"caster/graph/generated"
	"caster/graph/resolvers"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
)

func TestUsers(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(resolvers.NewResolver())))

	var resp struct {
		CreateTodo struct{ ID string }
	}
	c.MustPost(`mutation { createTodo(input:{text:"Fery important", userId: "test-user"}) { id } }`, &resp)

	require.Equal(t, "1", resp.CreateTodo.ID)

	t.Run("update the todo text", func(t *testing.T) {
		var resp struct {
			UpdateTodo struct{ Text string }
		}
		c.MustPost(
			`mutation($id: ID!, $text: String!) { updateTodo(id: $id, changes:{text:$text}) { text } }`,
			&resp,
			client.Var("id", 1),
			client.Var("text", "Very important"),
		)

		require.Equal(t, "Very important", resp.UpdateTodo.Text)
	})
}
