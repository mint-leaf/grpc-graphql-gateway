// Code generated by proroc-gen-graphql, DO NOT EDIT.
package author

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var _ = json.Marshal
var _ = json.Unmarshal

var gql__type_ListAuthorsRequest = graphql.NewObject(graphql.ObjectConfig{
	Name:   "ListAuthorsRequest",
	Fields: graphql.Fields{},
}) // message ListAuthorsRequest in author/author.proto

var gql__type_Author = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
}) // message Author in author/author.proto

var gql__type_ListAuthorsResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "ListAuthorsResponse",
	Fields: graphql.Fields{
		"authors": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(gql__type_Author)),
		},
	},
}) // message ListAuthorsResponse in author/author.proto

var gql__type_GetAuthorRequest = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetAuthorRequest",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
}) // message GetAuthorRequest in author/author.proto

// gql__resolver_AuthorService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.Resolver interface.
type gql__resolver_AuthorService struct {
	conn *grpc.ClientConn
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *gql__resolver_AuthorService) GetQueries() graphql.Fields {
	return graphql.Fields{
		"authors": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(gql__type_Author)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		},
		"author": &graphql.Field{
			Type: gql__type_Author,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *gql__resolver_AuthorService) GetMutations() graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you worried about open/close performance for each handling graphql request,
// then you can call RegisterBookHandler with *grpc.ClientConn manually.
func RegisterBookGraphql(mux *runtime.ServeMux) error {
	return RegisterBookGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when appication will terminate.
// Otherwise, the resolver opens connection automatically, but then you need to define host with ServiceOption like:
//
// service XXXService {
//    option (graphql.service) = {
//        host: "localhost:50051"
//    };
//
//    ...with RPC definitions
// }
func RegisterBookGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) (err error) {
	if conn == nil {
		conn, err = grpc.Dial("localhost:8080", grpc.WithInsecure())
		if err != nil {
			return
		}
	}
	mux.AddHandler(&gql__resolver_AuthorService{conn})
	return
}
