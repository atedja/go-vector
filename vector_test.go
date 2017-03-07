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
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v2 := v1.Clone()
	v2.Elements[0] = 9.0
	v2.Elements[2] = -1.0
	assert.Equal(0.0, v1.Elements[0])
	assert.Equal(9.0, v2.Elements[0])
	assert.Equal(2.0, v1.Elements[2])
	assert.Equal(-1.0, v2.Elements[2])
}

func TestSet(t *testing.T) {
	assert := assert.New(t)
	v1 := New(4)
	v1.Set([]float64{10.0, 9.9, 9.8, 9.7})
	assert.Equal(10.0, v1.Elements[0])
	assert.Equal(9.9, v1.Elements[1])
	assert.Equal(9.8, v1.Elements[2])
	assert.Equal(9.7, v1.Elements[3])
}

func TestVectorAdd(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v2 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v1.Add(v2)
	assert.Equal(0.0, v1.Elements[0])
	assert.Equal(2.0, v1.Elements[1])
	assert.Equal(4.0, v1.Elements[2])
	assert.Equal(2.0, v1.Elements[3])
}

func TestVectorSubtract(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v2 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v1.Subtract(v2)
	assert.Equal(0.0, v1.Elements[0])
	assert.Equal(0.0, v1.Elements[1])
	assert.Equal(0.0, v1.Elements[2])
	assert.Equal(0.0, v1.Elements[3])
}

func TestVectorScale(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v1.Scale(2.5)
	assert.Equal(0.0, v1.Elements[0])
	assert.Equal(2.5, v1.Elements[1])
	assert.Equal(5.0, v1.Elements[2])
	assert.Equal(2.5, v1.Elements[3])
}

func TestVectorDot(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v2 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	result, err := v1.Dot(v2)
	assert.Nil(err)
	assert.Equal(6.0, result)
}

func TestVectorCross(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0})
	v2 := NewWithValues([]float64{0.0, 3.0, 4.0})
	err := v1.Cross(v2)
	assert.Nil(err)
	assert.Equal(-2.0, v1.Elements[0])
	assert.Equal(0.0, v1.Elements[1])
	assert.Equal(0.0, v1.Elements[2])

	v3 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v4 := NewWithValues([]float64{0.0, 1.0, 2.0})
	err = v3.Cross(v4)
	assert.NotNil(err)
}

func TestVectorResize(t *testing.T) {
	assert := assert.New(t)
	v := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v.Resize(2)
	assert.Equal(len(v.Elements), 2)
	assert.Equal(v.Elements[0], 0.0)
	assert.Equal(v.Elements[1], 1.0)

	v.Resize(8)
	assert.Equal(len(v.Elements), 8)
	assert.Equal(0.0, v.Elements[0])
	assert.Equal(1.0, v.Elements[1])
	assert.Equal(0.0, v.Elements[2])
	assert.Equal(0.0, v.Elements[7])
}

func TestVectorMagnitude(t *testing.T) {
	assert := assert.New(t)
	v := NewWithValues([]float64{3.0, 4.0})
	assert.Equal(5.0, v.Magnitude())
}

func TestVectorUnit(t *testing.T) {
	assert := assert.New(t)
	epsilon := math.Nextafter(1, 2) - 1
	v := NewWithValues([]float64{3.0, 4.0})
	unit := v.Unit()
	assert.InEpsilon(0.6, unit.Elements[0], epsilon)
	assert.InEpsilon(0.8, unit.Elements[1], epsilon)
}

func TestVectorDim(t *testing.T) {
	assert := assert.New(t)
	v := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	assert.Equal(4, v.Dim())
}

func TestVectorZero(t *testing.T) {
	assert := assert.New(t)
	v := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v.Zero()
	assert.Equal(0.0, v.Elements[0])
	assert.Equal(0.0, v.Elements[1])
	assert.Equal(0.0, v.Elements[2])
	assert.Equal(0.0, v.Elements[3])
}
