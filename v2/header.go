package blizzard

import (
	"net/http"
	"time"
)

const (
	headerTimeFormat = "Mon, _2 Jan 2006 15:04:05 MST"
)

// Header Keys
const (
	HeaderKeyBattlenetNamespace   = "Battlenet-Namespace"
	HeaderKeyBattlenetSchema      = "Battlenet-Schema"
	HeaderKeyBattleSchemaRevision = "Battlenet-Schema-Revision"
	HeaderKeyCacheControl         = "Cache-Control"
	HeaderKeyConnection           = "Connection"
	HeaderKeyContentType          = "Content-Type"
	HeaderKeyDate                 = "Date"
	HeaderKeyLastModified         = "Last-Modified"
	HeaderKeyServer               = "Server"
	HeaderKeyVary                 = "Vary"
	HeaderKeyXContentTypeOptions  = "X-Content-Type-Options"
	HeaderKeyXFrameOptions        = "X-Frame-Options"
	HeaderKeyXTraceParentSpanID   = "X-Trace-Parentspanid"
	HeaderKeyXTraceSpanID         = "X-Trace-Spanid"
	HeaderKeyXTraceTraceID        = "X-Trace-Traceid"
)

// Header structure
type Header struct {
	BattlenetNamespace      string    `json:"Battlenet-Namespace"`
	BattlenetSchema         string    `json:"Battlenet-Schema"`
	BattlenetSchemaRevision string    `json:"Battlenet-Schema-Revision"`
	CacheControl            string    `json:"Cache-Control"`
	Connection              string    `json:"Connection"`
	ContentType             string    `json:"Content-Type"`
	Date                    time.Time `json:"Date"`
	LastModified            time.Time `json:"Last-Modified"`
	Server                  string    `json:"Server"`
	Vary                    string    `json:"Vary"`
	XContentTypeOptions     string    `json:"X-Content-Type-Options"`
	XFrameOptions           string    `json:"X-Frame-Options"`
	XTraceSpanID            string    `json:"X-Trace-Spanid"`
	XTraceParentSpanID      string    `json:"X-Trace-Parentspanid"`
	XTraceTraceID           string    `json:"X-Trace-Traceid"`
}

func getHeader(httpHeader http.Header) (*Header, error) {
	var err error
	header := Header{
		BattlenetNamespace:      httpHeader.Get(HeaderKeyBattlenetNamespace),
		BattlenetSchema:         httpHeader.Get(HeaderKeyBattlenetSchema),
		BattlenetSchemaRevision: httpHeader.Get(HeaderKeyBattleSchemaRevision),
		CacheControl:            httpHeader.Get(HeaderKeyCacheControl),
		Connection:              httpHeader.Get(HeaderKeyConnection),
		ContentType:             httpHeader.Get(HeaderKeyContentType),
		Server:                  httpHeader.Get(HeaderKeyServer),
		Vary:                    httpHeader.Get(HeaderKeyVary),
		XContentTypeOptions:     httpHeader.Get(HeaderKeyXContentTypeOptions),
		XFrameOptions:           httpHeader.Get(HeaderKeyXFrameOptions),
		XTraceSpanID:            httpHeader.Get(HeaderKeyXTraceParentSpanID),
		XTraceParentSpanID:      httpHeader.Get(HeaderKeyXTraceSpanID),
		XTraceTraceID:           httpHeader.Get(HeaderKeyXTraceTraceID),
	}
	if httpHeader.Get(HeaderKeyDate) != "" {
		header.Date, err = time.Parse(headerTimeFormat, httpHeader.Get(HeaderKeyDate))
		if err != nil {
			header.Date = time.Time{}
		}
	}
	if httpHeader.Get(HeaderKeyLastModified) != "" {
		header.LastModified, err = time.Parse(headerTimeFormat, httpHeader.Get(HeaderKeyLastModified))
		if err != nil {
			header.LastModified = time.Time{}
		}
	}
	return &header, nil
}
