package domain

type FesEvents []FesEvent

type FesEvent struct {
	ID      int
	Title   string
	Speaker string
}
