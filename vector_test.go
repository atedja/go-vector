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

func TestClone(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v2 := v1.Clone()
	assert.Equal(v1, v2)
	v2[0] = 9.0
	v2[2] = -1.0
	assert.Equal(0.0, v1[0])
	assert.Equal(9.0, v2[0])
	assert.Equal(2.0, v1[2])
	assert.Equal(-1.0, v2[2])
}

func TestSet(t *testing.T) {
	assert := assert.New(t)
	v1 := New(4)
	v1.Set([]float64{10.0, 9.9, 9.8, 9.7})
	assert.Equal(10.0, v1[0])
	assert.Equal(9.9, v1[1])
	assert.Equal(9.8, v1[2])
	assert.Equal(9.7, v1[3])

	v1.Set([]float64{1.0, 2.0, 3.0, 4.0})
	assert.Equal(1.0, v1[0])
	assert.Equal(2.0, v1[1])
	assert.Equal(3.0, v1[2])
	assert.Equal(4.0, v1[3])
}

func TestVectorScale(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v1.Scale(2.5)
	assert.Equal(0.0, v1[0])
	assert.Equal(2.5, v1[1])
	assert.Equal(5.0, v1[2])
	assert.Equal(2.5, v1[3])
}

func TestVectorMagnitude(t *testing.T) {
	assert := assert.New(t)
	v := NewWithValues([]float64{3.0, 4.0})
	assert.Equal(5.0, v.Magnitude())
}

func TestVectorZero(t *testing.T) {
	assert := assert.New(t)
	v := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v.Zero()
	assert.Equal(0.0, v[0])
	assert.Equal(0.0, v[1])
	assert.Equal(0.0, v[2])
	assert.Equal(0.0, v[3])
}

func TestDo(t *testing.T) {
	assert := assert.New(t)
	v := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v.DoWithIndex(func(i int, e float64) float64 {
		return e * 3.0
	})
	assert.Equal(0.0, v[0])
	assert.Equal(3.0, v[1])
	assert.Equal(6.0, v[2])
	assert.Equal(3.0, v[3])

	v.Do(func(e float64) float64 {
		return e * 2.0
	})
	assert.Equal(0.0, v[0])
	assert.Equal(6.0, v[1])
	assert.Equal(12.0, v[2])
	assert.Equal(6.0, v[3])
}
