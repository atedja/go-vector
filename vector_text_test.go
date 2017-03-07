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

func TestMarshalText(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{-1.0, 0.0, 2.0, 3.14})
	result, err := v1.MarshalText()
	assert.Nil(err)
	assert.Equal("-1,0,2,3.14\n", string(result))
}

func TestUnmarshalText(t *testing.T) {
	assert := assert.New(t)
	v1 := New(1)
	data := []byte(`1.0,-2,3.4,19.0,2.124,5.0,-29.0`)
	err := v1.UnmarshalText(data)
	assert.Nil(err)
	assert.Equal(1.0, v1.Elements[0])
	assert.Equal(-2.0, v1.Elements[1])
	assert.Equal(3.4, v1.Elements[2])
	assert.Equal(19.0, v1.Elements[3])
	assert.Equal(2.124, v1.Elements[4])
	assert.Equal(5.0, v1.Elements[5])
	assert.Equal(-29.0, v1.Elements[6])
}
