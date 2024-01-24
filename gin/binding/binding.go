package binding

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/poplknight/helper"
	"reflect"
	"strings"
	"sync"
)

const (
	json uint8 = iota
	xml
	form
	query
	formPost
	formMultipart
	protoBuf
	msgPack
	yaml
	uri
	header
	toml
)

var anyBindings = new(sync.Map)

func Register(a any) {
	if _, ok := anyBindings.Load(a); !ok {
		anyBindings.Store(a, assertAny(a))
	}
}

func assertAny(a any) []uint8 {
	tf := reflect.TypeOf(a)
	if t, ok := a.(reflect.Type); ok {
		tf = t
	}
	if tf.Kind() == reflect.Ptr {
		return assertAny(tf.Elem())
	}

	var bs []uint8
	if tf.Kind() == reflect.Struct {
		for i := 0; i < tf.NumField(); i++ {
			field := tf.Field(i)
			tag := field.Tag
			fmt.Println("field:", field.Tag.Get("json"))
			if field.Type.Kind() == reflect.Ptr || field.Type.Kind() == reflect.Struct {
				bs = append(bs, assertAny(field.Type)...)
			}
			if _, ok := tag.Lookup(binding.JSON.Name()); ok {
				bs = append(bs, json)
			}
			if _, ok := tag.Lookup(binding.XML.Name()); ok {
				bs = append(bs, xml)
			}
			if _, ok := tag.Lookup(binding.Form.Name()); ok {
				bs = append(bs, form)
			}
			if _, ok := tag.Lookup(binding.Query.Name()); ok {
				bs = append(bs, query)
			}
			if _, ok := tag.Lookup(binding.FormPost.Name()); ok {
				bs = append(bs, formPost)
			}
			if _, ok := tag.Lookup(binding.FormMultipart.Name()); ok {
				bs = append(bs, formMultipart)
			}
			if _, ok := tag.Lookup(binding.ProtoBuf.Name()); ok {
				bs = append(bs, protoBuf)
			}
			if _, ok := tag.Lookup(binding.MsgPack.Name()); ok {
				bs = append(bs, msgPack)
			}
			if _, ok := tag.Lookup(binding.YAML.Name()); ok {
				bs = append(bs, yaml)
			}
			if _, ok := tag.Lookup(binding.Uri.Name()); ok {
				bs = append(bs, uri)
			}
			if _, ok := tag.Lookup(binding.Header.Name()); ok {
				bs = append(bs, header)
			}
			if _, ok := tag.Lookup(binding.TOML.Name()); ok {
				bs = append(bs, toml)
			}

			if t, ok := tag.Lookup("binding"); ok && strings.Index(t, "dive") > -1 {
				qValue := reflect.ValueOf(a)
				bs = append(bs, assertAny(qValue.Field(i))...)
				continue
			}
			if t, ok := tag.Lookup("validate"); ok && strings.Index(t, "dive") > -1 {
				qValue := reflect.ValueOf(a)
				bs = append(bs, assertAny(qValue.Field(i))...)
			}
		}
	}

	mb := make(map[uint8]struct{})
	for _, b := range bs {
		mb[b] = struct{}{}
	}

	bs = make([]uint8, 0, len(mb))
	for b := range mb {
		bs = append(bs, b)
	}
	return helper.Distinct(bs, func(i uint8) uint8 { return i })
}

func bind(c *gin.Context, a any) error {
	bindings, ok := anyBindings.Load(a)
	if !ok {
		bindings = assertAny(a)
		anyBindings.Store(a, bindings)
	}
	if bindings == nil || len(bindings.([]uint8)) == 0 {
		return nil
	}

	for _, tp := range bindings.([]uint8) {
		switch tp {
		case json:
			if err := c.ShouldBindWith(a, binding.JSON); err != nil {
				return err
			}
		case xml:
			if err := c.ShouldBindWith(a, binding.XML); err != nil {
				return err
			}
		case form:
			if err := c.ShouldBindWith(a, binding.Form); err != nil {
				return err
			}
		case query:
			if err := c.ShouldBindWith(a, binding.Query); err != nil {
				return err
			}
		case formPost:
			if err := c.ShouldBindWith(a, binding.FormPost); err != nil {
				return err
			}
		case formMultipart:
			if err := c.ShouldBindWith(a, binding.FormMultipart); err != nil {
				return err
			}
		case protoBuf:
			if err := c.ShouldBindWith(a, binding.ProtoBuf); err != nil {
				return err
			}
		case msgPack:
			if err := c.ShouldBindWith(a, binding.MsgPack); err != nil {
				return err
			}
		case yaml:
			if err := c.ShouldBindWith(a, binding.YAML); err != nil {
				return err
			}
		case uri:
			if err := c.ShouldBindUri(a); err != nil {
				return err
			}
		case header:
			if err := c.ShouldBindWith(a, binding.Header); err != nil {
				return err
			}
		case toml:
			if err := c.ShouldBindWith(a, binding.TOML); err != nil {
				return err
			}
		}
	}

	return nil
}

func GetRequest[T any](c *gin.Context) (T, error) {
	var req T
	var err error
	if reflect.ValueOf(req).Kind() == reflect.Ptr {
		err = bind(c, req)
	} else {
		err = bind(c, &req)
	}

	return req, err
}
