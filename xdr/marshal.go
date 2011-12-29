package xdr

import (
	"bytes"
	"encoding/binary"
	"log"
	"reflect"
	"errors"
)

func Marshal(val interface{}) ([]byte, error) {
	log.Println("marshal:", val)
	b := new(bytes.Buffer)
        v := reflect.ValueOf(val)
        switch v.Kind() {
        case reflect.Ptr:
                v = v.Elem()
        case reflect.Struct:
		v = v
        default:
                return nil, errors.New("rpc.Marshal: invalid type " + v.Type().String())
        }
        for i := 0; i < v.NumField(); i++ {
                field := v.Field(i)
		log.Println("field:", field)
                switch t := field.Type() ; t.Kind() {
                case reflect.Uint, reflect.Uint32:
			binary.Write(b, binary.BigEndian, uint32(field.Uint()))
		case reflect.Struct, reflect.Interface:
			buf, err := Marshal(field.Interface())
			if err != nil {
				return nil, err
			}
			b.Write(buf)
		case reflect.Slice:
                        switch t.Elem().Kind() {
                        case reflect.Uint8:
				buf := field.Bytes()
				log.Println("slice:", buf)
				binary.Write(b, binary.BigEndian, uint32(len(buf)))
				b.Write(buf)
                        default:
                                panic("slice of unknown type " + t.Elem().Kind().String())
                        }
		default:
			panic("field of unknown type " +t.Elem().Kind().String())
		}
	}
	return b.Bytes(), nil
}
