package xmlcompare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnorderedElementsShouldbeTrue(t *testing.T) {
	eq, err := EqualString(
		`<?xml version="1.0"?>
<person id="13">
  <name>
    <first>John</first>
    <last>Doe</last>
  </name>
  <age>42</age>
</person>`,
		`<?xml version="1.0"?>
<person id="13">
  <name>
    <last>Doe</last>
    <first>John</first>
  </name>
  <age>42</age>
</person>`,
	)
	assert.NoError(t, err)
	assert.True(t, eq)
}
func TestNotEqualAttrValue(t *testing.T) {
	eq, err := EqualString(
		`<?xml version="1.0"?>
<person id="1">
  <name>
    <first>John</first>
    <last>Doe</last>
  </name>
  <age>42</age>
</person>`,
		`<?xml version="1.0"?>
<person id="13">
  <name>
    <last>Doe</last>
    <first>John</first>
  </name>
  <age>42</age>
</person>`,
	)
	assert.NoError(t, err)
	assert.False(t, eq)
}

func TestNotEqualValue(t *testing.T) {
	eq, err := EqualString(
		`<?xml version="1.0"?>
<person id="13">
  <name>
    <first>Jo</first>
    <last>Doe</last>
  </name>
  <age>42</age>
</person>`,
		`<?xml version="1.0"?>
<person id="13">
  <name>
    <last>Doe</last>
    <first>John</first>
  </name>
  <age>42</age>
</person>`,
	)
	assert.NoError(t, err)
	assert.False(t, eq)
}
