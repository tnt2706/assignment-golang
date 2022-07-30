.PHONY: start gql-gen

start:
	air /src/server.go

gql-gen:
	# generate tools.go
	printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"\nimport _ "github.com/99designs/gqlgen/graphql/introspection"\n' | gofmt > tools.go

	# generate defines.graphql
	rm -f defines.graphql
	printf 'scalar AWSDateTime\n' > defines.graphql
	printf 'scalar AWSJSON\n' >> defines.graphql
	printf 'scalar AWSEmail\n' >> defines.graphql
	printf 'directive @aws_subscribe(mutations: [String!]!) on FIELD_DEFINITION\n' >> defines.graphql
	printf 'directive @aws_cognito_user_pools on FIELD_DEFINITION | OBJECT\n' >> defines.graphql
	printf 'directive @aws_iam on FIELD_DEFINITION | OBJECT\n' >> defines.graphql
	mv defines.graphql src/graph/schema

	# install gqlgen library
	go mod tidy
	go mod vendor
	go run "github.com/99designs/gqlgen" generate

	# cleanup
	rm -f tools.go
	rm -rf src/graph/generated
	rm -rf src/graph/resolver/generated
	go mod tidy
	go mod vendor
