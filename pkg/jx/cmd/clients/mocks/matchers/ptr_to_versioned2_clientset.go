// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	versioned2 "k8s.io/metrics/pkg/client/clientset/versioned"
	"reflect"
)

func AnyPtrToVersioned2Clientset() *versioned2.Clientset {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*versioned2.Clientset))(nil)).Elem()))
	var nullValue *versioned2.Clientset
	return nullValue
}

func EqPtrToVersioned2Clientset(value *versioned2.Clientset) *versioned2.Clientset {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *versioned2.Clientset
	return nullValue
}
