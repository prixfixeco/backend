package testutils

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/prixfixeco/api_server/pkg/types"
)

// ContextMatcher is a matcher for use with testify/mock's MatchBy function. It provides some level of type
// safety reassurance over mock.Anything, in that the resulting function will panic if anything other than
// a context.Context.
var ContextMatcher interface{} = mock.MatchedBy(func(context.Context) bool {
	return true
})

// MapOfStringToInterfaceMatcher is a matcher for use with testify/mock's MatchBy function. It provides some level of type
// safety reassurance over mock.Anything, in that the resulting function will panic if anything other than
// a context.Context.
var MapOfStringToInterfaceMatcher interface{} = mock.MatchedBy(func(map[string]interface{}) bool {
	return true
})

// HTTPRequestMatcher is a matcher for use with testify/mock's MatchBy function. It provides some level of type
// safety reassurance over mock.Anything, in that the resulting function will panic if anything other than
// a *http.Request.
var HTTPRequestMatcher interface{} = mock.MatchedBy(func(*http.Request) bool {
	return true
})

// HTTPResponseWriterMatcher is a matcher for the http.ResponseWriter interface. It provides some level of type
// safety reassurance over mock.Anything, in that the resulting function will panic if anything other than
// a http.ResponseWriter.
var HTTPResponseWriterMatcher interface{} = mock.MatchedBy(func(http.ResponseWriter) bool {
	return true
})

// PreWriteMessageMatcher matches the types.PreWriteMessage type.
func PreWriteMessageMatcher(*types.PreWriteMessage) bool { return true }

// PreUpdateMessageMatcher matches the types.PreUpdateMessage type.
func PreUpdateMessageMatcher(*types.PreUpdateMessage) bool { return true }

// PreArchiveMessageMatcher matches the types.PreArchiveMessage type.
func PreArchiveMessageMatcher(*types.PreArchiveMessage) bool { return true }

// DataChangeMessageMatcher matches the types.PreWriteMessage type.
func DataChangeMessageMatcher(*types.DataChangeMessage) bool { return true }
