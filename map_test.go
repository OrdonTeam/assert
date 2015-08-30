package assert

import "testing"

func TestThatMapIsEmptyHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatMap(make(map[interface{}]interface{})).IsEmpty()
	mockT.HasNoErrors()
}

func TestThatMapIsEmptyHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	m := make(map[interface{}]interface{})
	m[true] = true
	assert.ThatMap(m).IsEmpty()
	mockT.HasErrorMessages(
		"Expected len <0>, but was <1>.",
	)
}