package validator

import (
	"reflect"
	"testing"
)

// func TestNewValidationTranslator(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want *ValidationTranslator
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewValidationTranslator(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewValidationTranslator() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_バリデーション(t *testing.T) {
	type request struct {
		Name     string `json:"name" validate:"required,min=1,max=50" trans:"名前"`
		Email    string `json:"email" validate:"required,min=1,max=50" trans:"メアド"`
		Password string `json:"password" validate:"required,min=6,max=50" trans:"パスワード"`
	}
	goodRequest := &request{Name: "テスト", Email: "test@test.com", Password: "testtesttest"}
	tests := []struct {
		name   string
		fields *ValidationTranslator
		args   *request
		want   map[string]string
	}{
		{
			name:   "バリデーション成功",
			fields: NewValidationTranslator(),
			args:   goodRequest,
			want:   map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &ValidationTranslator{
				validator:  tt.fields.validator,
				translator: tt.fields.translator,
			}
			if got := this.Validate(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("バリデーション失敗件数: %d", len(got))
				t.Errorf("バリデーション失敗内容: %v", got)
			}
		})
	}

	badRequest := &request{Name: "テスト", Email: "test@test.com", Password: "test"}
	tests = []struct {
		name   string
		fields *ValidationTranslator
		args   *request
		want   map[string]string
	}{
		{
			name:   "バリデーションエラーの場合",
			fields: NewValidationTranslator(),
			args:   badRequest,
			want:   map[string]string{"request.パスワード": "バリデーションエラー（この場合、パスワードエラー）"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &ValidationTranslator{
				validator:  tt.fields.validator,
				translator: tt.fields.translator,
			}

			if got := len((this.Validate(tt.args))); got <= 0 {
				t.Errorf("バリデーション失敗件数: %v", got)
			}
		})
	}
}
