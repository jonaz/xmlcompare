package xmlcompare

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"sort"
)

const maxDepth = 1000

func EqualString(a, b string) (bool, error) {
	x := &Node{}
	y := &Node{}
	err := xml.Unmarshal([]byte(a), x)
	if err != nil {
		return false, fmt.Errorf("xmlcompare: error unmarshal a: %w", err)
	}

	err = xml.Unmarshal([]byte(b), y)
	if err != nil {
		return false, fmt.Errorf("xmlcompare: error unmarshal b: %w", err)
	}
	return eq(x, y, 0), nil
}
func EqualBytes(a, b []byte) (bool, error) {
	x := &Node{}
	y := &Node{}
	err := xml.Unmarshal(a, x)
	if err != nil {
		return false, fmt.Errorf("xmlcompare: error unmarshal a: %w", err)
	}
	err = xml.Unmarshal(b, y)
	if err != nil {
		return false, fmt.Errorf("xmlcompare: error unmarshal b: %w", err)
	}
	return eq(x, y, 0), nil
}
func Equal(a, b *Node) bool {
	return eq(a, b, 0)
}

func eq(a, b *Node, depth int) bool {
	if depth > maxDepth {
		return false
	}
	if !eqEl(a, b) {
		return false
	}
	if len(a.Nodes) != len(b.Nodes) {
		return false
	}
	if len(a.Nodes) == 0 {
		return bytes.Equal(bytes.TrimSpace(a.Content), bytes.TrimSpace(b.Content))
	}
	sort.Sort(byName(a.Nodes)) // Sort nodes to make sure element order doesnt matter in compare.
	sort.Sort(byName(b.Nodes))
	for i := range a.Nodes {
		if !eq(&a.Nodes[i], &b.Nodes[i], depth+1) {
			return false
		}
	}
	return true
}

const xmlns = "xmlns"

func eqEl(a, b *Node) bool {
	if a.XMLName != b.XMLName {
		return false
	}
	attrs := make(map[xml.Name]string)
	for _, a := range a.Attr {
		if a.Name.Space == xmlns || a.Name.Local == xmlns {
			continue
		}
		attrs[a.Name] = a.Value
	}

	for _, b := range b.Attr {
		if b.Name.Space == xmlns || b.Name.Local == xmlns {
			continue
		}
		if v, ok := attrs[b.Name]; !ok || v != b.Value {
			return false
		}
	}
	return true
}
