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
)

type Vector struct {
	Elements []float64
}

// Clone this vector, returning a new Vector.
func (self *Vector) Clone() *Vector {
	return NewWithValues(self.Elements)
}

// Sets the values of this vector.
func (self *Vector) Set(values []float64) {
	copy(self.Elements, values)
}

// Adds this vector with another.
func (self *Vector) Add(other *Vector) {
	length := min(len(self.Elements), len(other.Elements))
	for i := 0; i < length; i++ {
		self.Elements[i] += other.Elements[i]
	}
}

// Subtracts this vector with another.
func (self *Vector) Subtract(other *Vector) {
	length := min(len(self.Elements), len(other.Elements))
	for i := 0; i < length; i++ {
		self.Elements[i] -= other.Elements[i]
	}
}

// Scale this vector (performs scalar multiplication) by the specified value.
func (self *Vector) Scale(value float64) {
	length := len(self.Elements)
	for i := 0; i < length; i++ {
		self.Elements[i] *= value
	}
}

// Performs dot product with another vector.
func (self *Vector) Dot(other *Vector) (float64, error) {
	if len(self.Elements) != len(other.Elements) {
		return 0.0, ErrVectorNotSameSize
	}

	result := 0.0
	for i, u := range self.Elements {
		result += u * other.Elements[i]
	}

	return result, nil
}

// Performs cross product with another vector.
// Vector dimensionality has to be 3.
func (self *Vector) Cross(other *Vector) error {
	if self.Dim() != 3 || other.Dim() != 3 {
		return ErrVectorInvalidDimension
	}

	x := self.Elements[1]*other.Elements[2] - self.Elements[2]*other.Elements[1]
	y := self.Elements[2]*other.Elements[0] - self.Elements[0]*other.Elements[2]
	z := self.Elements[0]*other.Elements[1] - self.Elements[1]*other.Elements[0]

	self.Elements[0] = x
	self.Elements[1] = y
	self.Elements[2] = z

	return nil
}

// Resizes this vector to a new size, filling extra elements with 0.0,
// or truncating them if new size is smaller.
func (self *Vector) Resize(n int) {
	newElements := make([]float64, n)
	copy(newElements, self.Elements)
	self.Elements = newElements
}

// Returns the magnitude of this vector.
func (self *Vector) Magnitude() float64 {
	result := 0.0
	for _, e := range self.Elements {
		result += e * e
	}
	return math.Sqrt(result)
}

// Returns the unit vector.
func (self *Vector) Unit() *Vector {
	magRec := 1.0 / self.Magnitude()
	unit := self.Clone()
	for i, _ := range unit.Elements {
		unit.Elements[i] *= magRec
	}
	return unit
}

// Returns the dimensionality of this vector, i.e length
func (self *Vector) Dim() int {
	return len(self.Elements)
}

// Zeroes this vector
func (self *Vector) Zero() {
	for i, _ := range self.Elements {
		self.Elements[i] = 0.0
	}
}

// Iterates through the elements of this vector and applies a custom function
func (self *Vector) ApplyFn(applyFn func(int, float64) float64) {
	for i, e := range self.Elements {
		self.Elements[i] = applyFn(i, e)
	}
}
