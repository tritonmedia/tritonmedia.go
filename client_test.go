/*
	Copyright 2019 Triton Media authors. All rights reserved.

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
package tritonmedia

import (
	"fmt"
	"testing"

	"github.com/gofrs/uuid"
	api "github.com/tritonmedia/tritonmedia.go/pkg/proto"

	"gopkg.in/d4l3k/messagediff.v1"
	"gopkg.in/h2non/gock.v1"
)

func TestNewClient(t *testing.T) {
	c, _ := NewClient("", "s")
	if c.BaseURL != DefaultBaseURL {
		t.Error("NewClient() didn't set to baked in default when DefaultURL wasn't set")
	}

	c, _ = NewClient("http://s.com", "s")
	if c.BaseURL != "http://s.com" {
		t.Error("NewwClient() didn't set BaseURL when provided")
	}

	c, err := NewClient("http://s.com", "")
	if err == nil {
		t.Error("NewClient() didn't fail when no apiToken was provided")
	}
}

func TestGet(t *testing.T) {
	defer gock.Off()

	c, err := NewClient("", "exampleKey")
	if err != nil {
		t.Errorf("client.Get(): failed to create client: %v", err)
	}

	payload := V1MediaList{
		Metadata: V1Metadata{
			Success: true,
			Host:    uuid.Must(uuid.NewV4()).String(),
		},
		Media: []V1MediaObject{
			V1MediaObject{
				Id:         uuid.Must(uuid.NewV4()).String(),
				Name:       "KonoSuba",
				Creator:    api.Media_API,
				CreatorId:  "",
				Type:       api.Media_TV,
				Source:     api.Media_FILE,
				SourceURI:  "",
				Metadata:   api.Media_KITSU,
				MetadataId: "",
				Status:     api.TelemetryStatusEntry_QUEUED,
			},
		},
	}

	gock.New(c.BaseURL).
		Get("/v1/media").
		Reply(200).
		JSON(payload)

	var resp V1MediaList
	err = c.Get("/v1/media", &resp)
	if err != nil {
		t.Errorf("client.Get(): failed to get /v1/media: %v", err)
	}

	if diff, equal := messagediff.PrettyDiff(resp, payload); !equal {
		fmt.Println(diff)
		t.Errorf("client.Get(): resp was not the same as the mocked data")
	}
}

func TestPost(t *testing.T) {
	c, err := NewClient("", "exampleKey")
	if err != nil {
		t.Errorf("client.Post(): failed to create client: %v", err)
	}

	postData := V1MediaObject{
		Name:       "Test Card",
		Creator:    api.Media_API,
		CreatorId:  "",
		Type:       api.Media_MOVIE,
		Source:     api.Media_FILE,
		SourceURI:  "file:///tmp/Bunny.mkv",
		Metadata:   api.Media_KITSU,
		MetadataId: "tt5311514",
		Status:     api.TelemetryStatusEntry_QUEUED,
	}

	payload := V1Media{
		Metadata: V1Metadata{
			Success: true,
			Host:    uuid.Must(uuid.NewV4()).String(),
		},
		Media: postData,
	}

	gock.New(c.BaseURL).
		Post("/v1/media").
		Reply(200).
		JSON(payload)

	var resp V1Media
	err = c.Post("/v1/media", postData, &resp)
	if err != nil {
		t.Errorf("client.Post(): failed to get /v1/media: %v", err)
	}

	// in case of live testing, we don't generate this
	postData.Id = resp.Media.Id

	if diff, equal := messagediff.PrettyDiff(resp.Media, postData); !equal {
		fmt.Println(diff)
		t.Errorf("client.Post(): resp was not the same as the mocked data")
	}
}
