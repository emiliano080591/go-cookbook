package mockgen

import (
	"errors"
	"github.com/emiliano080591/go-cookbook/mockgen/internal"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestExample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetSetter := internal.NewMockGetSetter(ctrl)

	var k string
	mockGetSetter.EXPECT().Get("aqui ponemos lo que sea").Do(func(key string) {
		k = key
	}).Return("", nil)

	customError := errors.New("algo fallo")
	mockGetSetter.EXPECT().Get(gomock.Any()).Return("", customError)

	if _, err := mockGetSetter.Get("aqui ponemos lo que sea"); err != nil {
		t.Errorf("got %#v; want %#v", err, nil)
	}
	if k != "aqui ponemos lo que sea" {
		t.Errorf("bad key")
	}
	if _, err := mockGetSetter.Get("key"); err == nil {
		t.Errorf("got %#v; want %#v", err, customError)
	}
}
