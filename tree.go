package gogs

import (
	"fmt"
)

type TreeFile struct {
	Name         string `json:"name"`
	LastCommitId string `json:"commitId"`
	IsDir        bool   `json:"isDir"`
	Size         int64  `json:"size"`
}

type Tree struct {
	LastCommitId  string     `json:"lastCommitId"`
	LastCommitMsg string     `json:"lastCommitMsg"`
	Files         []TreeFile `json:"files"`
	Readme        string     `json:"readme"`
	ReadmeName    string     `json:"readmeName"`
	IsReadmeText  bool       `json:"isReadmeText"`
}

func (c *Client) GetTree(user, repo, tree string) (*Tree, error) {
	files := &Tree{}
	return files, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/tree?path=%s", user, repo, tree), nil, nil, files)
}
