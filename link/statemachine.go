package link

import (
	"fmt"
	"log"
)

type Event int
type State int

type Transition struct {
	CurrentState State
	Event        Event
	NextState    State
}

type StateMachine struct {
	Transitions  []Transition
	CurrentState State
}

func NewStateMachine(transitions []Transition, initialState State) *StateMachine {
	return &StateMachine{
		Transitions:  transitions,
		CurrentState: initialState,
	}
}

func (sm *StateMachine) HandleEvent(event Event) error {
	for _, transition := range sm.Transitions {
		if transition.CurrentState == sm.CurrentState && transition.Event == event {
			backup := transition.CurrentState
			sm.CurrentState = transition.NextState

			if backup != sm.CurrentState {
				log.Printf("[FSM] Change from %d to %d ( by evt %d ) \n", backup, sm.CurrentState, event)
			}

			return nil
		}
	}

	return fmt.Errorf("invalid transition for event %d in state %d", event, sm.CurrentState)
}
