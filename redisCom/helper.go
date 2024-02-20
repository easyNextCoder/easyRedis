package redisCom

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
)

type RedisRet struct {
	V any
}

func (r RedisRet) String() string {
	return fmt.Sprintf("return type (%+v) value (%+v)", reflect.TypeOf(r.V), r.V)
}

//func runFuncName() string {
//	pc := make([]uintptr, 1)
//	runtime.Callers(3, pc)
//	f := runtime.FuncForPC(pc[0])
//	return f.Name()
//}

type Wrapper struct {
	Fname string
}

func (w *Wrapper) Run(ret any, err error) any {

	if err != nil {
		fmt.Printf("%s err %s\n", w.Fname, err)
		panic("")
		return nil
	}
	fmt.Printf("%s %s\n", w.Fname, RedisRet{ret})
	return ret
}

type RedisI interface {
	GetFuncMap() map[string]func()
	GetKey() string
}

func BindFuncMap(ri RedisI, conn redis.Conn) {

	tpv := reflect.ValueOf(ri)
	tp := reflect.TypeOf(ri)

	for i := 0; i < tpv.NumMethod(); i++ {
		name := reflect.New(reflect.TypeOf(""))
		name.Elem().SetString(tp.Method(i).Name)
		ri.GetFuncMap()[tp.Method(i).Name] = func(idx int) func() {
			return func() {
				index := idx
				tpv.Method(index).Call([]reflect.Value{name.Elem()})
				conn.Do(Del, ri.GetKey())
			}
		}(i)
	}
}
