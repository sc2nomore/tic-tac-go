// Code generated by mockery v1.0.0
package mocks

import core "github.com/sc2nomore/tic-tac-go/core"
import mock "github.com/stretchr/testify/mock"

// PlayerMapper is an autogenerated mock type for the PlayerMapper type
type PlayerMapper struct {
	mock.Mock
}

// Player provides a mock function with given fields: player
func (_m *PlayerMapper) Player(player int) core.Player {
	ret := _m.Called(player)

	var r0 core.Player
	if rf, ok := ret.Get(0).(func(int) core.Player); ok {
		r0 = rf(player)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Player)
		}
	}

	return r0
}

// PlayerSymbol provides a mock function with given fields: player
func (_m *PlayerMapper) PlayerSymbol(player int) (string, error) {
	ret := _m.Called(player)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(player)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(player)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}