// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "html/template"

var tmpl *template.Template

func init() {
	tmpl = template.Must(
		template.New("list.tmpl").
			ParseFiles("list.tmpl"))
}
