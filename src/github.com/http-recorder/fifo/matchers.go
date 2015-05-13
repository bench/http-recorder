package fifo

import (
	"fmt"
	"github.com/http-recorder/entities"
	"strings"
)

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

// searchRequestInFifo is a generic research function.
// It returns the first request that matches rule defined by requestMatcher and cond
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

/**** Note : Feel free to create your own matchers ****/

func containsPathMatcher(hr *entities.HttpRequest, cond interface{}) bool {
	if strings.Contains(hr.Path, cond.(string)) {
		return true
	}
	return false
}

func containsBodyMatcher(hr *entities.HttpRequest, cond interface{}) bool {
	if strings.Contains(hr.Body, cond.(string)) {
		return true
	}
	return false
}

func MethodMatcher(hr *entities.HttpRequest, cond interface{}) bool {
	if strings.TrimSpace(cond.(string)) == hr.Method {
		return true
	}
	return false
}
