// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"sync"

	"go.opentelemetry.io/otel/attribute"
)

// ActiveMetricsObjectMock is a mock implementation of common.ActiveMetricsObject.
//
//	func TestSomethingThatUsesActiveMetricsObject(t *testing.T) {
//
//		// make and configure a mocked common.ActiveMetricsObject
//		mockedActiveMetricsObject := &ActiveMetricsObjectMock{
//			GetActiveMetricsAttributesFunc: func() []attribute.KeyValue {
//				panic("mock out the GetActiveMetricsAttributes method")
//			},
//			IsEndTimeSetFunc: func() bool {
//				panic("mock out the IsEndTimeSet method")
//			},
//		}
//
//		// use mockedActiveMetricsObject in code that requires common.ActiveMetricsObject
//		// and then make assertions.
//
//	}
type ActiveMetricsObjectMock struct {
	// GetActiveMetricsAttributesFunc mocks the GetActiveMetricsAttributes method.
	GetActiveMetricsAttributesFunc func() []attribute.KeyValue

	// IsEndTimeSetFunc mocks the IsEndTimeSet method.
	IsEndTimeSetFunc func() bool

	// calls tracks calls to the methods.
	calls struct {
		// GetActiveMetricsAttributes holds details about calls to the GetActiveMetricsAttributes method.
		GetActiveMetricsAttributes []struct {
		}
		// IsEndTimeSet holds details about calls to the IsEndTimeSet method.
		IsEndTimeSet []struct {
		}
	}
	lockGetActiveMetricsAttributes sync.RWMutex
	lockIsEndTimeSet               sync.RWMutex
}

// GetActiveMetricsAttributes calls GetActiveMetricsAttributesFunc.
func (mock *ActiveMetricsObjectMock) GetActiveMetricsAttributes() []attribute.KeyValue {
	if mock.GetActiveMetricsAttributesFunc == nil {
		panic("ActiveMetricsObjectMock.GetActiveMetricsAttributesFunc: method is nil but ActiveMetricsObject.GetActiveMetricsAttributes was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetActiveMetricsAttributes.Lock()
	mock.calls.GetActiveMetricsAttributes = append(mock.calls.GetActiveMetricsAttributes, callInfo)
	mock.lockGetActiveMetricsAttributes.Unlock()
	return mock.GetActiveMetricsAttributesFunc()
}

// GetActiveMetricsAttributesCalls gets all the calls that were made to GetActiveMetricsAttributes.
// Check the length with:
//
//	len(mockedActiveMetricsObject.GetActiveMetricsAttributesCalls())
func (mock *ActiveMetricsObjectMock) GetActiveMetricsAttributesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetActiveMetricsAttributes.RLock()
	calls = mock.calls.GetActiveMetricsAttributes
	mock.lockGetActiveMetricsAttributes.RUnlock()
	return calls
}

// IsEndTimeSet calls IsEndTimeSetFunc.
func (mock *ActiveMetricsObjectMock) IsEndTimeSet() bool {
	if mock.IsEndTimeSetFunc == nil {
		panic("ActiveMetricsObjectMock.IsEndTimeSetFunc: method is nil but ActiveMetricsObject.IsEndTimeSet was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsEndTimeSet.Lock()
	mock.calls.IsEndTimeSet = append(mock.calls.IsEndTimeSet, callInfo)
	mock.lockIsEndTimeSet.Unlock()
	return mock.IsEndTimeSetFunc()
}

// IsEndTimeSetCalls gets all the calls that were made to IsEndTimeSet.
// Check the length with:
//
//	len(mockedActiveMetricsObject.IsEndTimeSetCalls())
func (mock *ActiveMetricsObjectMock) IsEndTimeSetCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsEndTimeSet.RLock()
	calls = mock.calls.IsEndTimeSet
	mock.lockIsEndTimeSet.RUnlock()
	return calls
}
