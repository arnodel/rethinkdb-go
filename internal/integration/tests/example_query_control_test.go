package tests

import (
	"fmt"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

// Return heroes and superheroes.
func ExampleBranch() {
	cur, err := r.DB("examples").Table("marvel").OrderBy("name").Map(r.Branch(
		r.Row.Field("victories").Gt(100),
		r.Row.Field("name").Add(" is a superhero"),
		r.Row.Field("name").Add(" is a hero"),
	)).Run(session)
	if err != nil {
		fmt.Print(err)
		return
	}

	var strs []string
	err = cur.All(&strs)
	if err != nil {
		fmt.Print(err)
		return
	}

	for _, str := range strs {
		fmt.Println(str)
	}

	// Output:
	// Iron Man is a superhero
	// Jubilee is a hero
}

// Return an error
func ExampleError() {
	err := r.Error("this is a runtime error").Exec(session)
	fmt.Println(err)
}

// Suppose we want to retrieve the titles and authors of the table posts. In the
// case where the author field is missing or null, we want to retrieve the
// string "Anonymous".
func ExampleTerm_Default() {
	cur, err := r.DB("examples").Table("posts").Map(map[string]interface{}{
		"title":  r.Row.Field("title"),
		"author": r.Row.Field("author").Default("Anonymous"),
	}).Run(session)
	if err != nil {
		fmt.Print(err)
		return
	}

	var res map[string]interface{}
	err = cur.One(&res)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(res)
}

// Convert a Go integer to a ReQL object
func ExampleExpr_int() {
	cur, err := r.Expr(1).Run(session)
	if err != nil {
		fmt.Print(err)
		return
	}

	var res interface{}
	err = cur.One(&res)
	if err != nil {
		fmt.Print(err)
		return
	}

	jsonPrint(res)

	// Output: 1
}

// Convert a Go slice to a ReQL object
func ExampleExpr_slice() {
	cur, err := r.Expr([]int{1, 2, 3}).Run(session)
	if err != nil {
		fmt.Print(err)
		return
	}

	var res []interface{}
	err = cur.All(&res)
	if err != nil {
		fmt.Print(err)
		return
	}

	jsonPrint(res)

	// Output:
	// [
	//     1,
	//     2,
	//     3
	// ]
}

// Convert a Go slice to a ReQL object
func ExampleExpr_map() {
	cur, err := r.Expr(map[string]interface{}{
		"a": 1,
		"b": "b",
	}).Run(session)
	if err != nil {
		fmt.Print(err)
		return
	}

	var res interface{}
	err = cur.One(&res)
	if err != nil {
		fmt.Print(err)
		return
	}

	jsonPrint(res)

	// Output:
	// {
	//     "a": 1,
	//     "b": "b"
	// }
}

// Convert a Go slice to a ReQL object
func ExampleExpr_struct() {
	type ExampleTypeNested struct {
		N int
	}

	type ExampleTypeEmbed struct {
		C string
	}

	type ExampleTypeA struct {
		ExampleTypeEmbed

		A      int
		B      string
		Nested ExampleTypeNested
	}

	cur, err := r.Expr(ExampleTypeA{
		A: 1,
		B: "b",
		ExampleTypeEmbed: ExampleTypeEmbed{
			C: "c",
		},
		Nested: ExampleTypeNested{
			N: 2,
		},
	}).Run(session)
	if err != nil {
		fmt.Print(err)
		return
	}

	var res interface{}
	err = cur.One(&res)
	if err != nil {
		fmt.Print(err)
		return
	}

	jsonPrint(res)

	// Output:
	// {
	//     "A": 1,
	//     "B": "b",
	//     "C": "c",
	//     "Nested": {
	//         "N": 2
	//     }
	// }
}

// Convert a Go struct (with rethinkdb tags) to a ReQL object. The tags allow
// the field names to be changed.
func ExampleExpr_structTags() {
	type ExampleType struct {
		A int    `rethinkdb:"field_a"`
		B string `rethinkdb:"field_b"`
	}

	cur, err := r.Expr(ExampleType{
		A: 1,
		B: "b",
	}).Run(session)
	if err != nil {
		fmt.Print(err)
		return
	}

	var res interface{}
	err = cur.One(&res)
	if err != nil {
		fmt.Print(err)
		return
	}

	jsonPrint(res)

	// Output:
	// {
	//     "field_a": 1,
	//     "field_b": "b"
	// }
}

// Execute a raw JSON query
func ExampleRawQuery() {
	cur, err := r.RawQuery([]byte(`"hello world"`)).Run(session)
	if err != nil {
		fmt.Print(err)
		return
	}

	var res interface{}
	err = cur.One(&res)
	if err != nil {
		fmt.Print(err)
		return
	}

	jsonPrint(res)

	// Output: "hello world"
}
