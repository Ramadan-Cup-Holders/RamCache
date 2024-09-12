package structure

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockHasher struct {
	hashValue int
}

func (m *mockHasher) Write(data []byte) (int, error) {
	return len(data), nil
}

func (m *mockHasher) Sum(b []byte) []byte {
	return []byte{byte(m.hashValue)}
}

func (m *mockHasher) Reset() {}

func (m *mockHasher) Size() int {
	return 4
}

func (m *mockHasher) BlockSize() int {
	return 4
}

func TestHashTable_InsertAndGet(t *testing.T) {
	// Table-driven test cases
	tests := []struct {
		name     string
		actions  []func(*HashTable)
		expected map[string]interface{}
	}{
		{
			name: "Insert single key-value pair",
			actions: []func(*HashTable){
				func(ht *HashTable) {
					err := ht.Insert("key1", "value1")
					require.NoError(t, err)
				},
			},
			expected: map[string]interface{}{
				"key1": "value1",
			},
		},
		{
			name: "Insert multiple key-value pairs",
			actions: []func(*HashTable){
				func(ht *HashTable) {
					err := ht.Insert("key1", "value1")
					require.NoError(t, err)
					err = ht.Insert("key2", "value2")
					require.NoError(t, err)
					err = ht.Insert("key3", "value3")
					require.NoError(t, err)
				},
			},
			expected: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
		},
		{
			name: "Overwrite existing key",
			actions: []func(*HashTable){
				func(ht *HashTable) {
					err := ht.Insert("key1", "value1")
					require.NoError(t, err)
					err = ht.Insert("key1", "newValue1")
					require.NoError(t, err)
				},
			},
			expected: map[string]interface{}{
				"key1": "newValue1",
			},
		},
		{
			name: "Insert with collision and overflow",
			actions: []func(*HashTable){
				func(ht *HashTable) {
					// Same hash for all keys to simulate collision
					err := ht.Insert("key1", "value1")
					require.NoError(t, err)
					err = ht.Insert("key2", "value2")
					require.NoError(t, err)
					err = ht.Insert("key3", "value3")
					require.NoError(t, err)
					err = ht.Insert("key4", "value4")
					require.NoError(t, err)
					err = ht.Insert("key5", "value5")
					require.NoError(t, err)
					err = ht.Insert("key6", "value6")
					require.NoError(t, err)
					err = ht.Insert("key7", "value7")
					require.NoError(t, err)
					err = ht.Insert("key8", "value8")
					require.NoError(t, err)
					err = ht.Insert("key9", "value9")
					require.NoError(t, err)
				},
			},
			expected: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
				"key7": "value7",
				"key8": "value8",
				"key9": "value9",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasher := &mockHasher{hashValue: 1} // Mock hasher that returns fixed hash to simulate collisions
			ht := NewHashTable(hasher, 8)

			// Execute the actions
			for _, action := range tt.actions {
				action(ht)
			}

			// Validate the expected results
			for key, expectedValue := range tt.expected {
				value, ok := ht.Get(key)
				assert.True(t, ok)
				assert.Equal(t, expectedValue, value)
			}
		})
	}
}
