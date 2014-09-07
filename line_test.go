package gcss

import (
	"fmt"
	"testing"
)

func Test_line_isEmpty_true(t *testing.T) {
	ln := newLine(1, "")

	if !ln.isEmpty() {
		t.Error("ln.Empty() should return true")
	}
}

func Test_line_isEmpty_false(t *testing.T) {
	ln := newLine(1, "html")

	if ln.isEmpty() {
		t.Error("ln.Empty() should return false")
	}
}

func Test_line_isTopIndent_true(t *testing.T) {
	ln := newLine(1, "html")

	if !ln.isTopIndent() {
		t.Error("ln.isTopIndent() should return true")
	}
}

func Test_line_isTopIndent_false(t *testing.T) {
	ln := newLine(1, "  html")

	if ln.isTopIndent() {
		t.Error("ln.isTopIndent() should return false")
	}
}

func Test_line_childOf_indentInvalidErr(t *testing.T) {
	parent := newElement(newLine(1, "html"), nil)

	ln := newLine(2, "    font-size: 12px")

	_, err := ln.childOf(parent)

	if err == nil {
		t.Error("err should not be nil")
	}

	if expected, actual := fmt.Sprintf("indent is invalid [line: %d]", ln.no), err.Error(); actual != expected {
		t.Errorf("err should be %q [actual: %q]", expected, actual)
	}
}

func Test_line_childOf(t *testing.T) {
	parent := newElement(newLine(1, "html"), nil)

	ln := newLine(2, "  font-size: 12px")

	ok, err := ln.childOf(parent)

	if err != nil {
		t.Errorf("err occurrd [err: %q]", err.Error())
	}

	if !ok {
		t.Error("ok should be true")
	}
}

func Test_newLine(t *testing.T) {
	no := 1
	s := "  html"

	ln := newLine(no, s)

	if ln.no != no {
		t.Errorf("ln.no should be %d [actual: %d]", no, ln.no)
	}

	if ln.s != s {
		t.Errorf("ln.s should be %s [actual: %s]", s, ln.s)
	}

	if ln.indent != indent(s) {
		t.Errorf("ln.indent should be %d [actual: %d]", indent(s), ln.indent)
	}
}

func Test_indent_no_indent(t *testing.T) {
	s := "html"
	expected := 0
	actual := indent(s)

	if actual != expected {
		t.Errorf("%q's indent should be %d [actual: %d]", s, expected, actual)
	}
}

func Test_indent_half_indent(t *testing.T) {
	s := "   html"
	expected := 1
	actual := indent(s)

	if actual != expected {
		t.Errorf("%q's indent should be %d [actual: %d]", s, expected, actual)
	}
}

func Test_indent(t *testing.T) {
	s := "    html"
	expected := 2
	actual := indent(s)

	if actual != expected {
		t.Errorf("%q's indent should be %d [actual: %d]", s, expected, actual)
	}
}