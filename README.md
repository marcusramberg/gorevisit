gorevisit
=========

A Story
-------

![Alt text](/public/images/happyfrodo.jpg?raw=true "excited frodo")

**"I'm going to make cool glitches for revisit.link with golang!"**

![Alt text](/public/images/killjoyaragorn.jpg?raw=true "buzzkill aragorn")

**"You're going to need to learn how to write web services first"**

![Alt text](/public/images/worriedfrodo.jpg?raw=true "worried frodo")

**"But.. I just want to play with pixels and make cool art..."**

![Alt text](/public/images/killjoyaragorn.jpg?raw=true "buzzkill aragorn")

**"You need to know about encoding and decoding, and serialization and deserialization..."**

![Alt text](/public/images/scaredfrodo.jpg?raw=true "scared frodo")

 **"But..I just want to glitch pictures of cats and stuff this is too much work..."**

![Alt text](/public/images/killjoyaragorn.jpg?raw=true "buzzkill aragorn")

**rambling on and on about http headers and image type detection**

![Alt text](/public/images/sickfrodo.jpg?raw=true "sick frodo")

 **"computers are the worst..."**

![Alt text](/public/images/whataboutgorevisit.jpg?raw=true "what about gorevisit")

**"He doesn't even know about gorevisit"**

![Alt text](/public/images/helpfuleowyn.jpg?raw=true "let's tell him about it")

**"Gorevisit let's you just concentrate on hackin dem pixels!"**

![Alt text](/public/images/happyfrodo.jpg?raw=true "excited frodo")

**"computers are the BEST!"**

Example
-------
```go
package main

import (
	revisit "github.com/revisitors/gorevisit"
	"image"
	"image/color"
	"math/rand"
)

func noise(src image.Image, dst image.RGBA) error {
	// get the boundary box of the original image
	orig := src.Bounds()

	// copy it into the destination image buffer
	for x := orig.Min.X; x < orig.Max.X; x++ {
		for y := orig.Min.Y; y < orig.Max.Y; y++ {
			dst.Set(x, y, src.At(x, y))
		}
	}

	// shift some colors
	numToMod := (orig.Max.X * orig.Max.Y) / 2
	for i := 0; i < numToMod; i++ {
		x := rand.Intn(orig.Max.X)
		y := rand.Intn(orig.Max.Y)
		if prev, ok := src.At(x, y).(color.RGBA); ok {
			prev.R += 30
			prev.B -= 30
			prev.G += 30
			dst.Set(x, y, prev)
		}
		if prev, ok := src.At(x, y).(color.YCbCr); ok {
			prev.Cr += 30
			prev.Cb -= 30
			prev.Y += 30
			dst.Set(x, y, prev)
		}
	}

	return nil

}

func main() {
	// make a RevisitService instance and pass it our glitcher
	s := revisit.NewRevisitService(noise)

	// run it!
	s.Run(":8080")
}
```

Docs
----

# gorevisit
--
    import "github.com/revisitors/gorevisit"


## Usage

#### func  BytesToDataURI

```go
func BytesToDataURI(data []byte, contentType string) string
```
BytesToDataURI given a byte array and a content type, creates a Data URI of the
content

#### type AudioData

```go
type AudioData struct {
	Data string `json:data"`
}
```

AudioData holds a reference to the data URI of sound data in a Revisit.link
message See: https://developer.mozilla.org/en-US/docs/Web/HTTP/data_URIs

#### type ImageData

```go
type ImageData struct {
	Data string `json:"data"`
}
```

ImageData holds a reference the data URI of image data in a Revisit.link message
See: https://developer.mozilla.org/en-US/docs/Web/HTTP/data_URIs

#### func (ImageData) ByteReader

```go
func (i ImageData) ByteReader() io.Reader
```
ByteReader returns an io.Reader for the image data in a Revisit message

#### type MetaData

```go
type MetaData struct {
	Audio AudioData `json:"audio"`
}
```

MetaData wraps the Audio data of a Revisit.link message as per the specification
See: http://revisit.link/spec.html

#### type RevisitMsg

```go
type RevisitMsg struct {
	Content ImageData `json:"content"`
	Meta    MetaData  `json:"meta"`
}
```

RevisitMsg holds a decoded Revisit.link message See:
http://revisit.link/spec.html

#### func  ImageRevisitor

```go
func ImageRevisitor(m *RevisitMsg, t func(src image.Image, dst image.RGBA) error) (*RevisitMsg, error)
```
ImageRevisitor given a RevisitMsg and an image transformation function, runs the
image data through the transformation and returns a new RevisitMsg with the
transformed image

#### func  NewRevisitMsgFromFiles

```go
func NewRevisitMsgFromFiles(mediaPath ...string) (*RevisitMsg, error)
```
NewRevisitMsgFromFiles given the path to an image file and optional path to an
audio file, creates a JSON encoded Revisit.link message

#### func (*RevisitMsg) ImageType

```go
func (r *RevisitMsg) ImageType() string
```
ImageType gets the type of image that is in the message

#### type RevisitService

```go
type RevisitService struct {
}
```

RevisitService holds the necessary context for a Revisit.link service.
Currently, this consists of an imageTransformer

#### func  NewRevisitService

```go
func NewRevisitService(it func(image.Image, image.RGBA) error) *RevisitService
```
NewRevisitService given an image transformation function, returns a new
Revisit.link service

#### func (*RevisitService) Run

```go
func (rs *RevisitService) Run(port string)
```
Run starts the Revisit.link service
