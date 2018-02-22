// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"encoding/json"
	"fmt"
	"bytes"
)

type MarkdownOption struct {
	Text    string
	Mode    string
	Context string
}

func (c *Client) Markdown(opt MarkdownOption) error {
	body, err := json.Marshal(&opt)
	if err != nil {
		return err
	}
	_, err = c.getResponse("POST", fmt.Sprintf("/markdown"), jsonHeader, bytes.NewReader(body))
	return err
}
