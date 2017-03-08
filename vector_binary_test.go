/*
Copyright 2017 Albert Tedja

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUnmarshalBinary(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{-1.0, 0.0, 2.0, 3.14})
	result, err := v1.MarshalBinary()
	assert.Nil(err)

	v2 := New(1)
	v2.UnmarshalBinary(result)
	assert.Equal(v1, v2)
}

func TestMarshalUnmarshalBinaryError(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{-1.0, 0.0, 2.0, 3.14})
	result, err := v1.MarshalBinary()
	assert.Nil(err)

	v2 := New(1)
	missing := result[:(len(result) - 4)]
	err = v2.UnmarshalBinary(missing)
	assert.NotNil(err)

	// test v2 unchanged
	assert.Equal(1, v2.Dim())
}
