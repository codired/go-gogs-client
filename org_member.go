// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AddOrgMembershipOption struct {
	Role string `json:"role" binding:"Required"`
}

func (c *Client) AddOrgMembership(org, user string, opt AddOrgMembershipOption) error {
	body, err := json.Marshal(&opt)
	if err != nil {
		return err
	}
	_, err = c.getResponse("PUT", fmt.Sprintf("/orgs/%s/membership/%s", org, user), jsonHeader, bytes.NewReader(body))
	return err
}

func (c *Client) AddMemberToOrganizationTeam(org, team, user string) error {

	_, err := c.getResponse("GET", fmt.Sprintf("/orgs/%s/teams/%s/add/%s", org, team, user), jsonHeader, nil)

	return err
}

func (c *Client) RemoveMemberFromOrganizationTeam(org, team, user string) error {
	_, err := c.getResponse("GET", fmt.Sprintf("/orgs/%s/teams/%s/remove/%s", org, team, user), jsonHeader, nil)

	return err
}
