package gitw_test

import (
	"testing"

	"github.com/gookit/gitw"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
)

var repo = gitw.NewRepo("./").WithFn(func(r *gitw.Repo) {
	r.Git().BeforeExec = gitw.PrintCmdline
})

func TestRepo_Info(t *testing.T) {
	info := repo.Info()
	dump.P(info)

	assert.Nil(t, repo.Err())
	assert.NotNil(t, info)
	assert.Equal(t, "gitw", info.Name)
}

func TestRepo_RemoteInfos(t *testing.T) {
	rs := repo.AllRemoteInfos()
	dump.P(rs)

	assert.NoError(t, repo.Err())
	assert.NotEmpty(t, rs)

	assert.True(t, repo.HasRemote(gitw.DefaultRemoteName))
	assert.NotEmpty(t, repo.RemoteNames())

	rt := repo.DefaultRemoteInfo()
	dump.P(rt)
	assert.NotEmpty(t, rt)
	assert.True(t, rt.Valid())
	assert.False(t, rt.Invalid())
	assert.Equal(t, gitw.DefaultRemoteName, rt.Name)
	assert.Equal(t, "git@github.com:gookit/gitw.git", rt.GitURL())
	assert.Equal(t, "http://github.com/gookit/gitw", rt.URLOfHTTP())
	assert.Equal(t, "https://github.com/gookit/gitw", rt.URLOfHTTPS())

	rt = repo.DefaultRemoteInfo(gitw.RemoteTypePush)
	assert.NotEmpty(t, rt)
}