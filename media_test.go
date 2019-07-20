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

func TestListMedia(t *testing.T) {
	c, err := NewClient("", "exampleKey")
	if err != nil {
		t.Errorf("client.ListMedia(): failed to create a client: %v", err)
	}

	expectedMediaObject := V1MediaObject{
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

	resp := V1MediaList{
		Metadata: V1Metadata{
			Success: true,
			Host:    "host",
		},
		Media: []V1MediaObject{expectedMediaObject},
	}

	gock.New(c.BaseURL).
		Get("/v1/media").
		Reply(200).
		JSON(resp)

	m, err := c.ListMedia()
	if err != nil {
		t.Errorf("client.ListMedia(): failed to get media %v", err)
	}

	if diff, equal := messagediff.PrettyDiff(m[0], expectedMediaObject); !equal {
		fmt.Println(diff)
		t.Errorf("client.ListMedia(): resp was not the same as the mocked data")
	}
}

func TestMedia(t *testing.T) {
	c, err := NewClient("", "exampleKey")
	if err != nil {
		t.Errorf("client.Media(): failed to create a client: %v", err)
	}

	expectedMediaObject := V1MediaObject{
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

	resp := V1Media{
		Metadata: V1Metadata{
			Success: true,
			Host:    "host",
		},
		Media: expectedMediaObject,
	}

	gock.New(c.BaseURL).
		Get("/v1/media/s").
		Reply(200).
		JSON(resp)

	m, err := c.Media("s")
	if err != nil {
		t.Errorf("client.Media(): failed to get media %v", err)
	}

	if diff, equal := messagediff.PrettyDiff(m, expectedMediaObject); !equal {
		fmt.Println(diff)
		t.Errorf("client.Media(): resp was not the same as the mocked data")
	}
}

func TestNewMedia(t *testing.T) {
	c, err := NewClient("", "exampleKey")
	if err != nil {
		t.Errorf("client.NewMedia(): failed to create client: %v", err)
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

	m, err := c.NewMedia(postData)
	if err != nil {
		t.Errorf("client.NewMedia(): failed to create new media %v", err)
	}

	// in case of live testing, we don't generate this
	postData.Id = m.Id

	if diff, equal := messagediff.PrettyDiff(m, postData); !equal {
		fmt.Println(diff)
		t.Errorf("client.Post(): resp was not the same as the mocked data")
	}
}
