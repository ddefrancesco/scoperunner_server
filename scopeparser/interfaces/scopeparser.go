package interfaces

type ScopeParser interface {
	ParseMap() (any, error)
	InitMap() map[any]any
}
