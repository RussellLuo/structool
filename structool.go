package structool

import (
	"reflect"

	"github.com/RussellLuo/structs"
	"github.com/mitchellh/mapstructure"
)

type DecodeHookFunc mapstructure.DecodeHookFuncValue
type EncodeHookFunc structs.EncodeHookFunc

type Codec struct {
	tagName string

	decodeHooks []func(DecodeHookFunc) DecodeHookFunc
	encodeHooks []func(EncodeHookFunc) EncodeHookFunc

	decodeHookFunc DecodeHookFunc
	encodeHookFunc EncodeHookFunc
}

func New() *Codec {
	return &Codec{
		tagName:        "structool",
		decodeHookFunc: nilDecodeHookFunc,
		encodeHookFunc: nilEncodeHookFunc,
	}
}

func (c *Codec) TagName(name string) *Codec {
	c.tagName = name
	return c
}

func (c *Codec) DecodeHook(hooks ...func(DecodeHookFunc) DecodeHookFunc) *Codec {
	c.decodeHooks = append(c.decodeHooks, hooks...)

	// Build the final hook function by applying the decoding hooks in the
	// order they are passed.
	if len(c.decodeHooks) > 0 {
		f := c.decodeHooks[len(c.decodeHooks)-1](nilDecodeHookFunc)
		for i := len(c.decodeHooks) - 2; i >= 0; i-- {
			f = c.decodeHooks[i](f)
		}
		c.decodeHookFunc = f
	}

	return c
}

func (c *Codec) EncodeHook(hooks ...func(EncodeHookFunc) EncodeHookFunc) *Codec {
	c.encodeHooks = append(c.encodeHooks, hooks...)

	// Build the final hook function by applying the encoding hooks in the
	// order they are passed.
	if len(c.encodeHooks) > 0 {
		f := c.encodeHooks[len(c.encodeHooks)-1](nilEncodeHookFunc)
		for i := len(c.encodeHooks) - 2; i >= 0; i-- {
			f = c.encodeHooks[i](f)
		}
		c.encodeHookFunc = f
	}

	return c
}

func (c *Codec) Decode(in interface{}, out interface{}) (err error) {
	config := &mapstructure.DecoderConfig{
		DecodeHook: c.decodeHookFunc,
		Squash:     true, // Always squash embedded structs.
		TagName:    c.tagName,
		Result:     out,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(in)
}

func (c *Codec) Encode(in interface{}) (out interface{}, err error) {
	isStruct := isStruct(in)

	if !isStruct {
		// Wrap `in` in a struct.
		in = struct{ In interface{} }{In: in}
	}

	s := structs.New(in)
	s.TagName = c.tagName
	s.EncodeHook = structs.EncodeHookFunc(c.encodeHookFunc)

	m := s.Map()

	if !isStruct {
		for _, v := range m {
			// m has one and only one pair of k/v.
			return v, nil
		}
	}

	return m, nil
}

func isStruct(in interface{}) bool {
	v := reflect.ValueOf(in)

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v.Kind() == reflect.Struct
}

func nilDecodeHookFunc(from, to reflect.Value) (interface{}, error) {
	return from.Interface(), nil
}

func nilEncodeHookFunc(value reflect.Value) (interface{}, error) {
	return value.Interface(), nil
}
