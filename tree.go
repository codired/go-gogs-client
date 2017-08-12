package gogs

import (
	"github.com/gogits/gogs/models"
	"fmt"
)

func (c *Client) GetTree(user, repo, tree string) (*models.TreeFile, error) {
	files := &models.TreeFile{}
	return files, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/tree?path=%s", user, repo, tree), nil, nil, files)
}
