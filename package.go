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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Creates a new vector with the specified size. All elements are set to 0.0
func New(size int) *Vector {
	v := &Vector{
		Elements: make([]float64, size),
	}

	return v
}

// Creates a new vector using the specified array as values.
func NewWithValues(values []float64) *Vector {
	v := &Vector{
		Elements: make([]float64, len(values)),
	}

	copy(v.Elements, values)

	return v
}

// Sums of two vectors, returns the resulting vector.
func Add(v1, v2 *Vector) *Vector {
	result := v1.Clone()
	result.Add(v2)
	return result
}

// Difference of two vectors, returns the resulting vector.
func Subtract(v1, v2 *Vector) *Vector {
	result := v1.Clone()
	result.Subtract(v2)
	return result
}

// Dot products of two vectors.
func Dot(v1, v2 *Vector) (float64, error) {
	return v1.Dot(v2)
}

// Cross products of two vectors.
// Vector dimensionality has to be 3.
func Cross(v1, v2 *Vector) (*Vector, error) {
	// Early error check to prevent redundant cloning
	if v1.Dim() != 3 || v2.Dim() != 3 {
		return nil, ErrVectorInvalidDimension
	}

	result := v1.Clone()
	result.Cross(v2)
	return result, nil
}

func Hadamard(v1, v2 *Vector) (*Vector, error) {
	if v1.Dim() != v2.Dim() {
		return nil, ErrVectorInvalidDimension
	}

	length := v1.Dim()
	result := New(length)
	for i := 0; i < length; i++ {
		result.Elements[i] = v1.Elements[i] * v2.Elements[i]
	}
	return result, nil
}
