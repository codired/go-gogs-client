package gogs

import (
	"fmt"
)

type TreeFile struct {
	Name         string `json:"name"`
	LastCommitId string `json:"commit_id"`
	IsDir        bool `json:"is_dir"`
	Size         int64 `json:"size"`
}

type Tree struct {
	LastCommitId  string `json:"last_commit_id"`
	LastCommitMsg string `json:"last_commit_msg"`
	Files         []TreeFile `json:"files"`
	Readme        string `json:"readme"`
	ReadmeName    string `json:"readme_name"`
	IsReadmeText  bool `json:"is_readme_text"`
}

func (c *Client) GetTree(user, repo, tree string) (*TreeFile, error) {
	files := &TreeFile{}
	return files, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/tree?path=%s", user, repo, tree), nil, nil, files)
}
