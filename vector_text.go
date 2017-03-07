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
	"encoding/csv"
	"strconv"
)

func (self *Vector) MarshalText() ([]byte, error) {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	strArr := make([]string, len(self.Elements))
	for i, e := range self.Elements {
		strArr[i] = strconv.FormatFloat(e, 'f', -1, 64)
	}

	err := writer.Write(strArr)
	if err != nil {
		return nil, err
	}

	writer.Flush()
	if writer.Error() != nil {
		return nil, writer.Error()
	}

	return buffer.Bytes(), nil
}

func (self *Vector) UnmarshalText(text []byte) error {
	var err error
	var els []string
	reader := csv.NewReader(bytes.NewReader(text))
	els, err = reader.Read()
	if err != nil {
		return err
	}

	self.Elements = make([]float64, len(els))
	for i, e := range els {
		self.Elements[i], err = strconv.ParseFloat(e, 64)
		if err != nil {
			return err
		}
	}

	return nil
}
