// Copyright 2019 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package tool

import (
	"reflect"
	"testing"

	"github.com/grailbio/reflow/ec2cluster/instances"
)

func TestFilterInstanceTypes(t *testing.T) {
	instanceTypes := []string{"a", "b", "c", "d"}
	existing := map[string]instances.VerifiedStatus{
		"a": {true, true, 10},
		"b": {true, false, 70},
		"c": {false, false, -1},
	}
	for _, tt := range []struct {
		instanceTypes      []string
		existing           map[string]instances.VerifiedStatus
		retry              bool
		verified, toverify []string
	}{
		{instanceTypes, map[string]instances.VerifiedStatus{}, false, []string{}, instanceTypes},
		{instanceTypes, existing, false, []string{"a"}, []string{"c", "d"}},
		{instanceTypes, existing, true, []string{"a"}, []string{"b", "c", "d"}},
		{[]string{"a"}, existing, false, []string{"a"}, []string{"c"}},
		{[]string{"a"}, existing, true, []string{"a"}, []string{"b", "c"}},
	} {
		verified, toverify := instancesToVerify(tt.instanceTypes, tt.existing, tt.retry)
		if len(tt.verified) == 0 {
			if len(verified) != 0 {
				t.Errorf("got %v want %v", verified, tt.verified)
			}
		} else if got, want := verified, tt.verified; !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		if got, want := toverify, tt.toverify; !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	}
}
