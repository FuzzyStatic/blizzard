package blizzard

import (
	"fmt"
	"net/url"
	"strings"
)

// SearchOpts is the interface for options passed to the search functions
type SearchOpt interface {
	Apply(*[]string)
}

// SearchField specifies what fields to filter the results on
//
// Additional search operations can be performed as follows:
//
// AND
// "str=5&dex=10"
// An implicit AND operation is performed by multiple query parameters.
//
// OR
// "str=5||10"
// "type=man||bear||pig"
// An OR operation is performed by placing a pair of bars between two values.
//
// NOT
// "race!=orc"
// A NOT operation is performed by using a combination of an exclamation mark and equal sign between two values. You can combine the NOT and OR operations in a single statement (for example, race!=orc||human).
//
// RANGE
// "str=[2,99]"
// "str=(2,99)"
// A RANGE operation is performed only on numeric field values by using either brackets (inclusive) or parentheses (exclusive) around comma-separated minimum and maximum values.
//
// MIN
// "str=[41,]"
// "str=(41,)"
// A MIN operation performs a minimum value check using the Range syntax.
//
// MAX
// "str=[,77]"
// "str=(,77)"
// A MAX operation performs a maximum value check using the Range syntax.
//
// Example:
// SearchField().AND("timezone","Europe/Paris").AND("locale", "enGB").NOT("type", "PVP")

func SearchField() *SearchFieldSelector {
	return &SearchFieldSelector{}
}

type SearchFieldSelector struct {
	parts []string
}

func (s *SearchFieldSelector) AND(field, value string) *SearchFieldSelector {
	s.parts = append(
		s.parts,
		fmt.Sprintf("%s=%s", uqe(field), uqe(value)),
	)
	return s
}
func (s *SearchFieldSelector) OR(field, firstValue, secondValue string) *SearchFieldSelector {
	s.parts = append(
		s.parts,
		fmt.Sprintf(
			"%s=%s||%s",
			uqe(field),
			uqe(firstValue),
			uqe(secondValue),
		),
	)
	return s
}
func (s *SearchFieldSelector) NOT(field, value string) *SearchFieldSelector {
	s.parts = append(
		s.parts,
		fmt.Sprintf("%s!=%s", uqe(field), uqe(value)),
	)
	return s
}

func (s *SearchFieldSelector) RANGE(field string, start, stop int) *SearchFieldSelector {
	s.parts = append(
		s.parts,
		fmt.Sprintf("%s=[%d,%d]", uqe(field), start, stop),
	)
	return s
}

func (s *SearchFieldSelector) MIN(field string, value int) *SearchFieldSelector {
	s.parts = append(
		s.parts,
		fmt.Sprintf("%s=[%d,]", uqe(field), value),
	)
	return s
}

func (s *SearchFieldSelector) MAX(field string, value int) *SearchFieldSelector {
	s.parts = append(s.parts, fmt.Sprintf("%s=[,%d]", uqe(field), value))
	return s
}

func (s *SearchFieldSelector) Apply(v *[]string) {
	*v = append(*v, strings.Join(s.parts, "&"))
}

// SearchPage specifies which page of search results to return. The default page is 1.
func SearchPage(page int) SearchOpt {
	return &SearchPageSelector{page: page}
}

type SearchPageSelector struct {
	page int
}

func (s *SearchPageSelector) Apply(v *[]string) {
	*v = append(*v, fmt.Sprintf("_page=%d", s.page))
}

// SearchPageSize specifies the size of each page in the result. The default value is 100 and can support a minimum of 1 or a maximum of 1000.
func SearchPageSize(size int) *SearchPageSizeSelector {
	return &SearchPageSizeSelector{size: size}
}

type SearchPageSizeSelector struct {
	size int
}

func (s *SearchPageSizeSelector) Apply(v *[]string) {
	*v = append(*v, fmt.Sprintf("_pageSize=%d", s.size))
}

// SearchOrderBy accepts a list of fields and sorts the result in order of those fields. :asc and :desc can be appended to each field to instruct the search endpoint to sort in ascending or descending order.
// "field1"
// "field1:asc"
// "field1:desc"
// "field1", "field2"
// "field1:desc", "field2:asc"
func SearchOrderBy(fields ...string) *SearchOrderBySelector {
	return &SearchOrderBySelector{fields: fields}
}

type SearchOrderBySelector struct {
	fields []string
}

func (s *SearchOrderBySelector) Apply(v *[]string) {
	for i, val := range s.fields {
		s.fields[i] = uqe(val)
	}

	*v = append(*v, fmt.Sprintf("orderby=%s", strings.Join(s.fields, ",")))
}

// uqe is just a short name for uqe
func uqe(val string) string {
	return url.QueryEscape(val)
}
