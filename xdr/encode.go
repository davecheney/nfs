package xdr

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"reflect"
)

func Write(w io.Writer, val interface{}) error {
	log.Println("marshal:", val)
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
	case reflect.Struct:
		v = v
	default:
		return fmt.Errorf("rpc.Write: invalid type: %v ", v.Type().String())
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		log.Println("field:", field)
		switch t := field.Type(); t.Kind() {
		case reflect.Uint, reflect.Uint32:
			binary.Write(w, binary.BigEndian, uint32(field.Uint()))
		case reflect.Struct, reflect.Interface:
			if err := Write(w, field.Interface()); err != nil {
				return err
			}
		case reflect.String:
			l := field.Len()
			binary.Write(w, binary.BigEndian, uint32(l))
			b := []byte(field.String())
			// pad to 32 bits
			b = append(b, make([]byte, l % 4)...)
			w.Write(b)
		case reflect.Slice:
			switch t.Elem().Kind() {
			case reflect.Uint8:
				buf := field.Bytes()
				log.Println("slice:", buf)
				binary.Write(w, binary.BigEndian, uint32(len(buf)))
				w.Write(buf)
			default:
				panic("slice of unknown type " + t.Elem().Kind().String())
			}
		default:
			panic("field of unknown type " + t.Kind().String())
		}
	}
	return nil
}
