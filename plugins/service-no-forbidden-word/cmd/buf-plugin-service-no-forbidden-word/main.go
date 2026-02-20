// Copyright 2025 buf-build-plugins contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"strings"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checkutil"
	"buf.build/go/bufplugin/option"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const forbiddenWordOptionKey = "forbidden_words"

var serviceNoForbiddenWordRuleSpec = &check.RuleSpec{
	ID:      "SERVICE_NO_FORBIDDEN_WORD",
	Default: true,
	Purpose: "Checks that service names do not contain any of the configured forbidden words.",
	Type:    check.RuleTypeLint,
	Handler: checkutil.NewServiceRuleHandler(checkServiceNoForbiddenWord, checkutil.WithoutImports()),
}

func main() {
	check.Main(&check.Spec{
		Rules: []*check.RuleSpec{serviceNoForbiddenWordRuleSpec},
	})
}

func checkServiceNoForbiddenWord(
	_ context.Context,
	responseWriter check.ResponseWriter,
	request check.Request,
	serviceDescriptor protoreflect.ServiceDescriptor,
) error {
	serviceName := string(serviceDescriptor.Name())
	serviceNameLower := strings.ToLower(serviceName)

	forbiddenWords, err := option.GetStringSliceValue(request.Options(), forbiddenWordOptionKey)
	if err != nil {
		return err
	}
	if len(forbiddenWords) == 0 {
		return nil
	}

	for _, word := range forbiddenWords {
		if strings.Contains(serviceNameLower, strings.ToLower(word)) {
			responseWriter.AddAnnotation(
				check.WithDescriptor(serviceDescriptor),
				check.WithMessagef("service name %q should not contain the word %q", serviceName, word),
			)
			break
		}
	}
	return nil
}
