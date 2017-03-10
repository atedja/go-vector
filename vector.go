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

type Vector []float64

// Clone this vector, returning a new Vector.
func (self Vector) Clone() Vector {
	return NewWithValues(self)
}

// Sets the values of this vector.
func (self Vector) Set(values []float64) {
	copy(self, values)
}

// Scale this vector (performs scalar multiplication) by the specified value.
func (self Vector) Scale(value float64) {
	length := len(self)
	for i := 0; i < length; i++ {
		self[i] *= value
	}
}

// Returns the magnitude of this vector.
func (self Vector) Magnitude() float64 {
	result := 0.0
	for _, e := range self {
		result += e * e
	}
	return math.Sqrt(result)
}

// Zeroes this vector
func (self Vector) Zero() {
	for i, _ := range self {
		self[i] = 0.0
	}
}

// Iterates through the elements of this vector and for each element invokes
// the function.
func (self Vector) ApplyFn(applyFn func(float64) float64) {
	for i, e := range self {
		self[i] = applyFn(e)
	}
}

// Iterates through the elements of this vector and for each element invokes
// the function with index.
func (self Vector) ApplyFnWithIndex(applyFn func(int, float64) float64) {
	for i, e := range self {
		self[i] = applyFn(i, e)
	}
}
