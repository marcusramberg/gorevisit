# gorevisit
--
    import "github.com/revisitors/gorevisit"


## Usage

```go
var (
	//ErrUnsupportedType is returned when a Transform does not support the type(s) passed to it
	ErrUnsupportedType = errors.New("unsupported type")
)
```

#### func  BytesToDataURI

```go
func BytesToDataURI(data []byte, contentType string) string
```
BytesToDataURI returns a data URI encoded string given a byte array and a
content type See RFC2397 - http://tools.ietf.org/html/rfc2397

#### type APIMsg

```go
type APIMsg struct {
	Content *Content     `json:"content"`
	Meta    *MetaContent `json:"meta"`
}
```

APIMsg is a message containing Content, and MetaContent. the MetaContent should
be audio.

#### func  NewAPIMsgFromFiles

```go
func NewAPIMsgFromFiles(mediaPath ...string) (*APIMsg, error)
```
NewAPIMsgFromFiles returns an APImsg struct pointer given a path to an image and
audio file

#### func  NewAPIMsgFromJSON

```go
func NewAPIMsgFromJSON(b []byte) (*APIMsg, error)
```
NewAPIMsgFromJSON returns an APIMsg struct pointer from a json byte array.

#### func (*APIMsg) IsValid

```go
func (a *APIMsg) IsValid() bool
```
IsValid verifies that an APIMsg is valid according to the specification

#### func (*APIMsg) JSON

```go
func (a *APIMsg) JSON() ([]byte, error)
```
JSON serializes a gorevisit.APIMsg back to JSON bytes

#### type Content

```go
type Content struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
```

Content contains a type and a string, the string should be a base64 encoded
image

#### type DecodedContent

```go
type DecodedContent struct {
	Type string
	Data []byte
}
```

DecodedContent contains a type and a byte array, the byte array should be image
data

#### func  DataURIToBytes

```go
func DataURIToBytes(dataURI string) (*DecodedContent, error)
```
DataURIToBytes returns a content type string and an array of bytes given a data
URI encoded string. See RFC2397 - http://tools.ietf.org/html/rfc2397

#### type MetaContent

```go
type MetaContent struct {
	Audio *Content `json:"audio"`
}
```

MetaContent contains a Content pointer

#### type RevisitService

```go
type RevisitService struct {
}
```

RevisitService holds context for a POST handler for revisit

#### func  NewRevisitService

```go
func NewRevisitService(t func(*APIMsg) (*APIMsg, error)) *RevisitService
```
NewRevisitService constructs a new Revisit service given a transform function

#### func (*RevisitService) PostHandler

```go
func (rs *RevisitService) PostHandler(w http.ResponseWriter, r *http.Request)
```
PostHandler handles a POST to a revisit service

#### func (*RevisitService) ServeHTTP

```go
func (rs *RevisitService) ServeHTTP(w http.ResponseWriter, r *http.Request)
```
ServeHTTP implements a Revisit service to be passed to a mux

#### type Transformer

```go
type Transformer interface {
	Transform(*APIMsg) (*APIMsg, error)
}
```

Transformer interface transforms an APIMsg into another APIMsg
