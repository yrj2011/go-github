// Copyright 2017 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestRepositoriesService_GetCommunityHealthMetrics(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/community/profile", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeRepositoryCommunityHealthMetricsPreview)
		fmt.Fprintf(w, `{
				"health_percentage": 100,
				"files": {
					"code_of_conduct": {
						"name": "Contributor Covenant",
						"key": "contributor_covenant",
						"url": null,
						"html_url": "http://github.com/octocat/Hello-World/blob/master/CODE_OF_CONDUCT.md"
					},
					"contributing": {
						"url": "http://api.github.com/repos/octocat/Hello-World/contents/CONTRIBUTING",
						"html_url": "http://github.com/octocat/Hello-World/blob/master/CONTRIBUTING"
					},
					"license": {
						"name": "MIT License",
						"key": "mit",
						"url": "http://api.github.com/licenses/mit",
						"html_url": "http://github.com/octocat/Hello-World/blob/master/LICENSE"
					},
					"readme": {
						"url": "http://api.github.com/repos/octocat/Hello-World/contents/README.md",
						"html_url": "http://github.com/octocat/Hello-World/blob/master/README.md"
					}
				},
				"updated_at": "2017-02-28T00:00:00Z"
			}`)
	})

	got, _, err := client.Repositories.GetCommunityHealthMetrics(context.Background(), "o", "r")
	if err != nil {
		t.Errorf("Repositories.GetCommunityHealthMetrics returned error: %v", err)
	}

	updatedAt := time.Date(2017, time.February, 28, 0, 0, 0, 0, time.UTC)
	want := &CommunityHealthMetrics{
		HealthPercentage: Int(100),
		UpdatedAt:        &updatedAt,
		Files: &CommunityHealthFiles{
			CodeOfConduct: &Metric{
				Name:    String("Contributor Covenant"),
				Key:     String("contributor_covenant"),
				HTMLURL: String("http://github.com/octocat/Hello-World/blob/master/CODE_OF_CONDUCT.md"),
			},
			Contributing: &Metric{
				URL:     String("http://api.github.com/repos/octocat/Hello-World/contents/CONTRIBUTING"),
				HTMLURL: String("http://github.com/octocat/Hello-World/blob/master/CONTRIBUTING"),
			},
			License: &Metric{
				Name:    String("MIT License"),
				Key:     String("mit"),
				URL:     String("http://api.github.com/licenses/mit"),
				HTMLURL: String("http://github.com/octocat/Hello-World/blob/master/LICENSE"),
			},
			Readme: &Metric{
				URL:     String("http://api.github.com/repos/octocat/Hello-World/contents/README.md"),
				HTMLURL: String("http://github.com/octocat/Hello-World/blob/master/README.md"),
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Repositories.GetCommunityHealthMetrics:\ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
	}
}
