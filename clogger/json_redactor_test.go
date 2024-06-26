package clogger

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FooDecimal struct {
	decimal.Decimal
}

func (d *FooDecimal) MarshalJSON() ([]byte, error) {
	return json.Marshal("0x" + d.BigInt().Text(16))
}

func TestRedactJSONObject(t *testing.T) {
	d := FooDecimal{decimal.NewFromInt(100)}
	o, err := json.Marshal(&d)
	assert.NoError(t, err)

	fmt.Println("====> ", string(o))

	var t1 = map[string]interface{}{
		"a": &d,
	}

	_, err = redactJSONObject(t1, []string{"b"})
	assert.NoError(t, err)
}

func TestRedactJSON(t *testing.T) {
	var t1 = map[string]interface{}{
		"a": 1,
		"b": "foo",
		"c": map[string]interface{}{"d": 2},
		"e": []interface{}{1, 2, map[string]interface{}{"f": 3}},
	}

	in, err := json.Marshal(t1)
	assert.NoError(t, err)

	out, err := redactJSON(in, []string{"f"})
	assert.NoError(t, err)

	assert.Equal(t, `{"a":1,"b":"foo","c":{"d":2},"e":[1,2,{"f":"redacted"}]}`, string(out))
}
