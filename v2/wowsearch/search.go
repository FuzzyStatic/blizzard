package wowsearch

import (
	"fmt"
	"net/url"
	"strings"
)

// Opt is the interface for options passed to the search functions
type Opt interface {
	Apply(*[]string)
}

// Field specifies what fields to filter the search results on
//
// Example:
//
//  wowsearch.Field().AND("timezone","Europe/Paris").AND("locale", "enGB").NOT("type.type", "PVP")
func Field() *FieldSelector {
	return &FieldSelector{}
}

type FieldSelector struct {
	parts []string
}

// An implicit AND operation is performed by chaining multiple AND's
//
// Example:
//
//  wowsearch.Field().AND("id","123").AND("type","PVP")
func (s *FieldSelector) AND(field, value string) *FieldSelector {
	s.parts = append(s.parts, fmt.Sprintf("%s=%s", uqe(field), uqe(value)))
	return s
}

// An OR operation is performed, you can specify multiple alternatives
//
// Example:
//
//  wowsearch.Field().OR("str","5","10")
//  wowsearch.Field().OR("type","man","bear","pig")
func (s *FieldSelector) OR(field string, values ...string) *FieldSelector {
	for i, val := range values {
		values[i] = uqe(val)
	}
	s.parts = append(s.parts, fmt.Sprintf("%s=%s", uqe(field), strings.Join(values, "||")))
	return s
}

// A NOT operation, multiple NOT's can be added by chaining
//
// Example:
//
//  wowsearch.NOT("race","orc")
//  wowsearch.NOT("race","human")
func (s *FieldSelector) NOT(field, value string) *FieldSelector {
	s.parts = append(s.parts, fmt.Sprintf("%s!=%s", uqe(field), uqe(value)))
	return s
}

// A RANGE operation is performed only on numeric field values
//
// Example:
//
//  wowsearch.RANGE("str",2,99)
func (s *FieldSelector) RANGE(field string, start, stop int) *FieldSelector {
	s.parts = append(s.parts, fmt.Sprintf("%s=[%d,%d]", uqe(field), start, stop))
	return s
}

// A MIN operation performs a minimum value check
//
// Example:
//
//  wowsearch.MIN("str",10)
func (s *FieldSelector) MIN(field string, value int) *FieldSelector {
	s.parts = append(s.parts, fmt.Sprintf("%s=[%d,]", uqe(field), value))
	return s
}

// A MAX operation performs a maximum value check
//
// Example:
//
//  wowsearch.MAX("str",100)
func (s *FieldSelector) MAX(field string, value int) *FieldSelector {
	s.parts = append(s.parts, fmt.Sprintf("%s=[,%d]", uqe(field), value))
	return s
}

func (s *FieldSelector) Apply(v *[]string) {
	*v = append(*v, strings.Join(s.parts, "&"))
}

// Page specifies which page of search results to return. The default page is 1 if not specified
//
// Example:
//
//  wowsearch.Page(10)
func Page(page int) Opt {
	if page < 1 {
		page = 1
	}
	return &PageSelector{page: page}
}

type PageSelector struct {
	page int
}

func (s *PageSelector) Apply(v *[]string) {
	*v = append(*v, fmt.Sprintf("_page=%d", s.page))
}

// PageSize specifies the size of each page in the result. The default value is 100 and can support a minimum of 1 or a maximum of 1000.
// If size provided is < 1 it will be set to 1, if it's > 1000 it will be set to 1000
//
// Example:
//
//  wowsearch.PageSize(20) // Will return 20 items
//  wowsearch.PageSize(1337) // Size will be set to 1000
//  wowsearch.PageSize(-2) // Size will be set to 1
func PageSize(size int) *PageSizeSelector {
	switch {
	case size < 1:
		size = 1
	case size > 1000:
		size = 1000
	}
	return &PageSizeSelector{size: size}
}

type PageSizeSelector struct {
	size int
}

func (s *PageSizeSelector) Apply(v *[]string) {
	*v = append(*v, fmt.Sprintf("_pageSize=%d", s.size))
}

// OrderBy accepts a list of fields and sorts the result in order of those fields. :asc and :desc can be appended to each field to instruct the search endpoint to sort in ascending or descending order.
//
// Example:
//
//  wowsearch.OrderBy("field1")
//  wowsearch.OrderBy("field1:asc")
//  wowsearch.OrderBy("field1:desc")
//  wowsearch.OrderBy("field1", "field2")
//  wowsearch.OrderBy("field1:desc", "field2:asc")
func OrderBy(fields ...string) *OrderBySelector {
	return &OrderBySelector{fields: fields}
}

type OrderBySelector struct {
	fields []string
}

func (s *OrderBySelector) Apply(v *[]string) {
	for i, val := range s.fields {
		s.fields[i] = uqe(val)
	}
	*v = append(*v, fmt.Sprintf("orderby=%s", strings.Join(s.fields, ",")))
}

// uqe is just a shorter name for url.QueryEscape
func uqe(val string) string {
	return url.QueryEscape(val)
}
