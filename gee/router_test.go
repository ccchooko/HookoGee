package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRouter("GET", "/", nil)
	r.addRouter("GET", "/hello/:name", nil)
	r.addRouter("GET", "/hello/b/c", nil)
	r.addRouter("GET", "/hi/:name", nil)
	r.addRouter("GET", "/assets/*filepath", nil)
	r.addRouter("GET", "/api/v2/*", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("p/*name/*"), []string{"p", "*name"})

	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRouteWithColon(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRouter("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

}

func TestGetRouteWithStar(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRouter("GET", "/assets/js/jQuery.js")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/assets/*filepath" {
		t.Fatal("should match /assets/*filepath")
	}

	if ps["filepath"] != "js/jQuery.js" {
		t.Fatal("name should be equal to 'js/jQuery.js'")
	}

	fmt.Printf("matched path: %s, params['filepath']: %s\n", n.pattern, ps["filepath"])

}

// 添加一个/api/v2/act的路径测试
func TestGetRouteWithStarNull(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRouter("GET", "/api/v2/act/data")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	fmt.Printf("matched path: %s, params: %s\n", n.pattern, ps)

}