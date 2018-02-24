package client

import (
	"sync"
	"time"
)

type Token struct {
	Value string
	ValidUntil time.Time
}

type TokenSource interface {
	AccessToken() (*Token, error)
	RefreshToken() (*Token, error)
}

type tokenSourceSetter interface {
	setAccessToken(token *Token)
	setRefreshToken(token *Token)
}

type LockableTokenSource struct {
	access, refresh Token
	mtxAccess, mtxRefresh sync.RWMutex
}

func (ts *LockableTokenSource) AccessToken() (*Token, error) {
	ts.mtxAccess.RLock()
	defer ts.mtxAccess.RUnlock()
	t := ts.access
	return &t, nil
}

func (ts *LockableTokenSource) RefreshToken() (*Token, error) {
	ts.mtxRefresh.RLock()
	defer ts.mtxRefresh.RUnlock()
	t := ts.refresh
	return &t, nil
}

func (ts *LockableTokenSource) setAccessToken(token *Token) {
	ts.mtxAccess.Lock()
	defer ts.mtxAccess.Unlock()
	ts.access = *token
}

func (ts *LockableTokenSource) setRefreshToken(token *Token) {
	ts.mtxRefresh.Lock()
	defer ts.mtxRefresh.Unlock()
	ts.refresh = *token
}

var _ TokenSource = (*LockableTokenSource)(nil)
var _ tokenSourceSetter = (*LockableTokenSource)(nil)
