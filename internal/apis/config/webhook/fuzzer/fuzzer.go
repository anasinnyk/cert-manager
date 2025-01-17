/*
Copyright 2021 The cert-manager Authors.

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

package fuzzer

import (
	fuzz "github.com/google/gofuzz"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/utils/pointer"

	"github.com/jetstack/cert-manager/internal/apis/config/webhook"
)

// Funcs returns the fuzzer functions for the webhook config api group.
var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(s *webhook.WebhookConfiguration, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again

			if s.HealthzPort == nil {
				s.HealthzPort = pointer.Int(12)
			}
			if s.SecurePort == nil {
				s.SecurePort = pointer.Int(123)
			}
			if s.PprofAddress == "" {
				s.PprofAddress = "something:1234"
			}
		},
	}
}
