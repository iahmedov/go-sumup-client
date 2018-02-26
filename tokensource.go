package client

import (
	"sync"

	sb "github.com/iahmedov/go-sumup-client/schemas/base"
)

type TokenSource interface {
	AccessToken() (sb.Token, error)
	RefreshToken() (sb.Token, error)
}

type tokenSourceSetter interface {
	setAccessToken(token *sb.Token)
	setRefreshToken(token *sb.Token)
}

type LockableTokenSource struct {
	access, refresh       sb.Token
	mtxAccess, mtxRefresh sync.RWMutex
}

func (ts *LockableTokenSource) AccessToken() (sb.Token, error) {
	ts.mtxAccess.RLock()
	defer ts.mtxAccess.RUnlock()
	return ts.access, nil
}

func (ts *LockableTokenSource) RefreshToken() (sb.Token, error) {
	ts.mtxRefresh.RLock()
	defer ts.mtxRefresh.RUnlock()
	return ts.refresh, nil
}

func (ts *LockableTokenSource) setAccessToken(token *sb.Token) {
	ts.mtxAccess.Lock()
	defer ts.mtxAccess.Unlock()
	ts.access = *token
}

func (ts *LockableTokenSource) setRefreshToken(token *sb.Token) {
	ts.mtxRefresh.Lock()
	defer ts.mtxRefresh.Unlock()
	ts.refresh = *token
}

var _ TokenSource = (*LockableTokenSource)(nil)
var _ tokenSourceSetter = (*LockableTokenSource)(nil)
