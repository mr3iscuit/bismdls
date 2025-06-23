package analysis

import (
	"sync"

	"github.com/mr3iscuit/bismdls/analysis/types"
)

type State struct {
	documents map[string]types.TextDocumentItem
}

var (
	instance *State
	once     sync.Once
)

func NewStateInstance() *State {
	once.Do(func() {
		instance = &State{
			documents: map[string]types.TextDocumentItem{},
		}
	})
	return instance
}

func (s *State) OpenDocument(uri, text string) {
	s.documents[uri] = types.TextDocumentItem{
		URI:  uri,
		Text: text,
	}
}

func (s *State) UpdateDocument(uri, text string) {
	s.documents[uri] = types.TextDocumentItem{
		URI:  uri,
		Text: text,
	}
}
