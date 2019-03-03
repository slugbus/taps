//   Copyright 2019 The SlugBus++ Authors.

//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at

//        http://www.apache.org/licenses/LICENSE-2.0

//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Package slugger is a go client library
// for the UCSC TAPS API
package slugger

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	// The deefault query should be okay.
	t.Run("OK Query", func(t *testing.T) {
		got, err := Query()
		if err != nil {
			t.Errorf("got err: %v, wanted: nil", err)
		}
		if got == nil {
			t.Errorf("got no response")
		}
	})

	// These urls should fail.
	tURLs := []string{"https://www.random.org/bad_url", "bad_url", "https://jsonplaceholder.typicode.com/todos"}
	for _, url := range tURLs {
		t.Run(fmt.Sprintf("Bad Query: %s", url), func(t *testing.T) {
			OverrideURL(url)
			got, err := Query()
			if err == nil {
				t.Errorf("got no error, wanted error")
			}
			if got != nil {
				t.Errorf("got a response")
			}
		})
	}

	// Reseting the URL should make the
	// first test pass, despite the changes above.
	RestoreURL()
	t.Run("OK Query", func(t *testing.T) {
		got, err := Query()
		if err != nil {
			t.Errorf("got err: %v, wanted: nil", err)
		}
		if got == nil {
			t.Errorf("got no response")
		}
	})

	// If we get a successful query, test to make
	// sure none of the results are their 0 values.
	t.Run("Default Check", func(t *testing.T) {
		got, err := Query()
		if err != nil {
			t.Errorf("got err: %v, wanted: nil", err)
		}

		for _, bus := range got {

			if bus.ID == "" {
				t.Errorf("got default value for ID, got %q , did not want %q", bus.ID, bus.ID)
			}

			if bus.Type == "" {
				t.Errorf("got default value for Type, got %q , did not want %q", bus.Type, bus.Type)
			}

			if bus.Lat == 0 {
				t.Errorf("got default value for Lat, got %v , did not want %v", bus.Lat, bus.Lat)
			}

			if bus.Lon == 0 {
				t.Errorf("got default value for Lon, got %v , did not want %v", bus.Lon, bus.Lon)
			}

		}

	})
}