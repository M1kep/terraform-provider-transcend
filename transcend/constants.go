package transcend

import "github.com/shurcooL/graphql"

type DataSilo struct {
	id      graphql.String
	title   graphql.String
	link    graphql.String
	catalog struct {
		hasAvcFunctionality graphql.Boolean
	}
}
