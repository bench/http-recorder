package fifo

import (
	"fmt"
	"github.com/http-recorder/entities"
	"strings"
)

/**** List of matchers : Feel free to create your owns ****/

func pathContains(hr *entities.HttpRequest, cond interface{}) bool {
	if strings.Contains(hr.Path, cond.(string)) {
		return true
	}
	return false
}

func bodyContains(hr *entities.HttpRequest, cond interface{}) bool {
	if strings.Contains(hr.Body, cond.(string)) {
		return true
	}
	return false
}

func isMethod(hr *entities.HttpRequest, cond interface{}) bool {
	if cond.(string) == hr.Method {
		return true
	}
	return false
}

func isContentType(hr *entities.HttpRequest, cond interface{}) bool {
	if cond.(string) == hr.Headers["Content-Type"][0] {
		return true
	}
	return false
}

/* Internal mechanisms */

func FindBy(key string, value string) (*entities.HttpRequest, error) {

	var matcher MatcherFunc

	key = strings.ToLower(key)
	value = strings.ToLower(value)

	switch key {
	case "pathcontains":
		matcher = pathContains
	case "bodycontains":
		matcher = bodyContains
	case "method":
		matcher = isMethod
	case "contenttype":
		matcher = isContentType
	default:
		return nil, fmt.Errorf("Unsupported query:", key)
	}

	return searchRequestInFifo(MatcherFunc(matcher), value)
}

type RequestMatcher interface {
	MatchesCond(*entities.HttpRequest, interface{}) bool
}

// The MatcherFunc type is an adapter to allow the use of
// ordinary functions as RequestMatcher.
type MatcherFunc func(*entities.HttpRequest, interface{}) bool

// MatchesCond calls f(hr, cond).
func (f MatcherFunc) MatchesCond(hr *entities.HttpRequest, cond interface{}) bool {
	return f(hr, cond)
}

// Generic research function. It returns the first request that
// matches rule defined by requestMatcher and condition
func searchRequestInFifo(requestMatcher RequestMatcher, cond interface{}) (*entities.HttpRequest, error) {
	elements := requestFifo.getElements()
	for _, element := range elements {
		if requestMatcher.MatchesCond(element.Value.(*entities.HttpRequest), cond) {
			requestFifo.removeElement(element)
			return element.Value.(*entities.HttpRequest), nil
		}
	}
	return &entities.HttpRequest{}, fmt.Errorf("No request matching")
}
