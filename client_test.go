package tritonmedia

import (
	"fmt"
	"testing"

	"github.com/gofrs/uuid"
	proto "github.com/tritonmedia/tritonmedia.go/pkg/proto"
	triton "github.com/tritonmedia/tritonmedia.go/pkg/tritonmedia"

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
	c, err := NewClient("", "exampleKey")
	if err != nil {
		t.Errorf("client.Get(): failed to create client: %v", err)
	}

	payload := triton.V1MediaList{
		Metadata: triton.V1Metadata{
			Success: true,
			Host:    uuid.Must(uuid.NewV4()).String(),
		},
		Media: []triton.V1MediaObject{
			triton.V1MediaObject{
				Id:         uuid.Must(uuid.NewV4()).String(),
				Name:       "KonoSuba",
				Creator:    0,
				CreatorId:  "",
				Type:       0,
				Source:     0,
				SourceURI:  "",
				Metadata:   0,
				MetadataId: "",
				Status:     0,
			},
		},
	}

	gock.New(fmt.Sprintf("%s/%s", c.BaseURL, "/v1/media")).
		Reply(200).
		JSON(payload)

	gock.InterceptClient(c.Client)

	var resp triton.V1MediaList
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

	postData := triton.V1MediaObject{
		Id:         uuid.Must(uuid.NewV4()).String(),
		Name:       "Test Card",
		Creator:    proto.Media_API,
		CreatorId:  "",
		Type:       proto.Media_MOVIE,
		Source:     proto.Media_FILE,
		SourceURI:  "file:///tmp/Bunny.mkv",
		Metadata:   proto.Media_KITSU,
		MetadataId: "tt5311514",
		Status:     proto.TelemetryStatusEntry_QUEUED,
	}

	payload := triton.V1Media{
		Metadata: triton.V1Metadata{
			Success: true,
			Host:    uuid.Must(uuid.NewV4()).String(),
		},
		Media: postData,
	}

	gock.New(fmt.Sprintf("%s/%s", c.BaseURL, "/v1/media")).
		Reply(200).
		JSON(payload)

	gock.InterceptClient(c.Client)

	var resp triton.V1Media
	err = c.Post("/v1/media", postData, &resp)
	if err != nil {
		t.Errorf("client.Post(): failed to get /v1/media: %v", err)
	}

	if diff, equal := messagediff.PrettyDiff(resp.Media, postData); !equal {
		fmt.Println(diff)
		t.Errorf("client.Post(): resp was not the same as the mocked data")
	}
}
