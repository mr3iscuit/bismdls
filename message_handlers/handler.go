package message_handlers

import (
	"sync"

	"github.com/mr3iscuit/bismdls/logger"
)

type HandlerStrategy interface {
	Handle(msg []byte) (any, error)
}

type handlerRegistry struct {
	handlers map[string]HandlerStrategy
}

func newHandlerRegistry() *handlerRegistry {
	return &handlerRegistry{
		make(map[string]HandlerStrategy),
	}
}

var (
	handlerRegistryInstance *handlerRegistry
	once                    sync.Once
)

func (registry *handlerRegistry) Register(name string, strategy HandlerStrategy) error {
	if registry.handlers[name] != nil {
		panic("Handler already registered")
	}

	registry.handlers[name] = strategy

	return nil
}

func (registry *handlerRegistry) GetHandler(name string) HandlerStrategy {
	value, exists := registry.handlers[name]

	if !exists {
		logger.GetLogger().Printf("Handler not found for method [%s]. Available handlers: %v", name, registry.getHandlerNames())
		return nil
	}

	return value
}

func (registry *handlerRegistry) getHandlerNames() []string {
	names := make([]string, 0, len(registry.handlers))
	for name := range registry.handlers {
		names = append(names, name)
	}
	return names
}

func HandlerRegistryInstannce() *handlerRegistry {
	once.Do(func() {
		handlerRegistryInstance = newHandlerRegistry()
	})
	return handlerRegistryInstance
}

func init() {
	handlerRegistry := HandlerRegistryInstannce()
	handlerRegistry.Register("initialize", &Initialize{})
	handlerRegistry.Register("initialized", &Initialized{})
	handlerRegistry.Register("textDocument/didOpen", &TextDocumentDidOpen{})
	handlerRegistry.Register("textDocument/didChange", &TextDocumentDidChange{})
	handlerRegistry.Register("textDocument/didClose", &TextDocumentDidClose{})
	handlerRegistry.Register("textDocument/hover", &TextDocumentOnHover{})
}
