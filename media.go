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
)

// ListMedia returns a list of all media on the server
func (c *Client) ListMedia() ([]V1MediaObject, error) {
	var resp V1MediaList
	err := c.Get("/v1/media", &resp)
	if err != nil {
		return nil, err
	}
	return resp.Media, nil
}

// Media looks up a media by ID and gets it
func (c *Client) Media(id string) (V1MediaObject, error) {
	var resp V1Media
	err := c.Get(fmt.Sprintf("/v1/media/%s", id), &resp)
	if err != nil {
		return V1MediaObject{}, err
	}
	return resp.Media, nil
}

// NewMedia creates a new media object on the server and begins processing
func (c *Client) NewMedia(m V1MediaObject) (V1MediaObject, error) {
	var resp V1Media
	err := c.Post("/v1/media", m, &resp)
	if err != nil {
		return V1MediaObject{}, err
	}
	return resp.Media, nil
}
