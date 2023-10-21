package Impl

import (
	"github.com/graphql-go/graphql"
)

// Schema
var todoSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"completed": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

// Queries
var todoQueries = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TodoQuries",
		Fields: graphql.Fields{
			"todos": &graphql.Field{
				Type: graphql.NewList(todoSchema),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetTodos(), nil
				},
			},
			"todo": &graphql.Field{
				Type: todoSchema,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						todo := GetTodoByID(id)
						return todo, nil
					}
					return nil, nil
				},
			},
		},
	},
)

// Mutations
var todoMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TodoMutation",
		Fields: graphql.Fields{
			"addTodo": &graphql.Field{
				Type: todoSchema,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					title, ok := p.Args["title"].(string)
					if ok {
						return AddTodoItem(title), nil
					}
					return nil, nil
				},
			},
			"updateTodo": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"completed": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					title, ok2 := p.Args["title"].(string)
					completed, ok3 := p.Args["completed"].(bool)
					if ok && ok2 && ok3 {
						updatedTodo := UpdateTodoItem(id, title, completed)
						return updatedTodo, nil
					}
					return nil, nil
				},
			},
			"deleteTodo": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						return DeleteTodoItem(id), nil
					}
					return nil, nil
				},
			},
		},
	},
)

var TodoSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    todoQueries,
		Mutation: todoMutation,
	},
)
