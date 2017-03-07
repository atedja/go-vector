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
	"bytes"
	"encoding/binary"
)

func (self *Vector) MarshalBinary() ([]byte, error) {
	var buffer bytes.Buffer
	var err error

	length := uint64(len(self.Elements))
	err = binary.Write(&buffer, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}

	for _, e := range self.Elements {
		err = binary.Write(&buffer, binary.BigEndian, e)
		if err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
}

func (self *Vector) UnmarshalBinary(data []byte) error {
	var err error
	var length uint64

	reader := bytes.NewReader(data)
	err = binary.Read(reader, binary.BigEndian, &length)
	if err != nil {
		return err
	}

	self.Elements = make([]float64, length)
	var val float64
	var i uint64
	for i = 0; i < length; i++ {
		err = binary.Read(reader, binary.BigEndian, &val)
		if err != nil {
			return err
		}
		self.Elements[i] = val
	}

	return nil
}
