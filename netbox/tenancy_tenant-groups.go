// Copyright 2017 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by generate_functions.go. DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TenantGroupsService is used in a Client to access NetBox's tenancy/tenant-groups API methods.
type TenantGroupsService struct {
	c *Client
}

// Get retrieves an TenantGroup object from NetBox by its ID.
func (s *TenantGroupsService) Get(id int) (*TenantGroup, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		fmt.Sprintf("api/tenancy/tenant-groups/%d/", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	t := new(TenantGroup)
	err = s.c.Do(req, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// List returns a Page associated with an NetBox API Endpoint.
func (s *TenantGroupsService) List() *Page {
	return NewPage(s.c, "api/tenancy/tenant-groups/", nil)
}

// Extract retrives a list of TenantGroup objects from page.
func (s *TenantGroupsService) Extract(page *Page) ([]*TenantGroup, error) {
	if err := page.Err(); err != nil {
		return nil, err
	}

	var groups []*TenantGroup
	if err := json.Unmarshal(page.data.Results, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}

// Create creates a new TenantGroup object in NetBox and returns the ID of the new object.
func (s *TenantGroupsService) Create(data *TenantGroup) (int, error) {
	req, err := s.c.NewJSONRequest(http.MethodPost, "api/tenancy/tenant-groups/", nil, data)
	if err != nil {
		return 0, err
	}

	g := new(TenantGroup)
	err = s.c.Do(req, g)
	if err != nil {
		return 0, err
	}
	return g.ID, nil
}

// Update changes an existing TenantGroup object in NetBox, and returns the ID of the new object.
func (s *TenantGroupsService) Update(data *TenantGroup) (int, error) {
	req, err := s.c.NewJSONRequest(
		http.MethodPatch,
		fmt.Sprintf("api/tenancy/tenant-groups/%d/", data.ID),
		nil,
		data,
	)
	if err != nil {
		return 0, err
	}

	// g is just used to verify correct api result.
	// data is not changed, because the g is not the full representation that one would
	// get with Get. But if the response was unmarshaled into TenantGroup correctly,
	// everything went fine, and we do not need to update data.
	g := new(TenantGroup)
	err = s.c.Do(req, g)
	if err != nil {
		return 0, err
	}
	return g.ID, nil
}

// Delete deletes an existing TenantGroup object from NetBox.
func (s *TenantGroupsService) Delete(data *TenantGroup) error {
	req, err := s.c.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("api/tenancy/tenant-groups/%d/", data.ID),
		nil,
	)
	if err != nil {
		return err
	}

	return s.c.Do(req, nil)
}
