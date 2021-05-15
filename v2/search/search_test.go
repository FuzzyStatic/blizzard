package search

import (
	"reflect"
	"testing"
)

func TestField(t *testing.T) {
	tests := []struct {
		name string
		want *FieldSelector
	}{
		{
			name: "test field",
			want: &FieldSelector{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Field(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Field() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldSelector_AND(t *testing.T) {
	type fields struct {
		parts []string
	}
	type args struct {
		field string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FieldSelector
	}{
		{
			name: "AND test",
			args: args{
				field: "a",
				value: "b",
			},
			want: &FieldSelector{
				parts: []string{"a=b"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FieldSelector{
				parts: tt.fields.parts,
			}
			if got := s.AND(tt.args.field, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldSelector.AND() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldSelector_OR(t *testing.T) {
	type fields struct {
		parts []string
	}
	type args struct {
		field  string
		values []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FieldSelector
	}{
		{
			name: "test OR",
			args: args{
				field:  "a",
				values: []string{"foo", "bar"},
			},
			want: &FieldSelector{parts: []string{"a=foo||bar"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FieldSelector{
				parts: tt.fields.parts,
			}
			if got := s.OR(tt.args.field, tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldSelector.OR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldSelector_NOT(t *testing.T) {
	type fields struct {
		parts []string
	}
	type args struct {
		field string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FieldSelector
	}{
		{
			name: "test NOT",
			args: args{
				field: "foo",
				value: "bar",
			},
			want: &FieldSelector{parts: []string{"foo!=bar"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FieldSelector{
				parts: tt.fields.parts,
			}
			if got := s.NOT(tt.args.field, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldSelector.NOT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldSelector_RANGE(t *testing.T) {
	type fields struct {
		parts []string
	}
	type args struct {
		field string
		start int
		stop  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FieldSelector
	}{
		{
			name: "test RANGE",
			args: args{
				field: "bar",
				start: 1,
				stop:  10,
			},
			want: &FieldSelector{parts: []string{"bar=[1,10]"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FieldSelector{
				parts: tt.fields.parts,
			}
			if got := s.RANGE(tt.args.field, tt.args.start, tt.args.stop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldSelector.RANGE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldSelector_MIN(t *testing.T) {
	type fields struct {
		parts []string
	}
	type args struct {
		field string
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FieldSelector
	}{
		{
			name: "test MIN",
			args: args{
				field: "foo",
				value: 10,
			},
			want: &FieldSelector{parts: []string{"foo=[10,]"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FieldSelector{
				parts: tt.fields.parts,
			}
			if got := s.MIN(tt.args.field, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldSelector.MIN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldSelector_MAX(t *testing.T) {
	type fields struct {
		parts []string
	}
	type args struct {
		field string
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FieldSelector
	}{
		{
			name: "test MAX",
			args: args{
				field: "bar",
				value: 420,
			},
			want: &FieldSelector{parts: []string{"bar=[,420]"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FieldSelector{
				parts: tt.fields.parts,
			}
			if got := s.MAX(tt.args.field, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldSelector.MAX() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestFieldSelector_Apply(t *testing.T) {
	type fields struct {
		parts []string
	}
	type args struct {
		v *[]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FieldSelector{
				parts: tt.fields.parts,
			}
			s.Apply(tt.args.v)
		})
	}
}
*/

func TestPage(t *testing.T) {
	type args struct {
		page int
	}
	tests := []struct {
		name string
		args args
		want Opt
	}{
		{
			name: "test PAGE",
			args: args{
				page: 10,
			},
			want: &PageSelector{
				page: 10,
			},
		},
		{
			name: "test PAGE under 1",
			args: args{
				page: -4,
			},
			want: &PageSelector{
				page: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Page(tt.args.page); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Page() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestPageSelector_Apply(t *testing.T) {
	type fields struct {
		page int
	}
	type args struct {
		v *[]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{// TODO: Add test cases.}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PageSelector{
				page: tt.fields.page,
			}
			s.Apply(tt.args.v)
		})
	}
}
*/
func TestPageSize(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *PageSizeSelector
	}{
		{
			name: "test PageSize",
			args: args{
				size: 69,
			},
			want: &PageSizeSelector{size: 69},
		},
		{
			name: "test PageSize over 1000",
			args: args{
				size: 1337,
			},
			want: &PageSizeSelector{size: 1000},
		},
		{
			name: "test PageSize under 1",
			args: args{
				size: -4,
			},
			want: &PageSizeSelector{size: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PageSize(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PageSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestPageSizeSelector_Apply(t *testing.T) {
	type fields struct {
		size int
	}
	type args struct {
		v *[]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PageSizeSelector{
				size: tt.fields.size,
			}
			s.Apply(tt.args.v)
		})
	}
}
*/
func TestOrderBy(t *testing.T) {
	type args struct {
		fields []string
	}
	tests := []struct {
		name string
		args args
		want *OrderBySelector
	}{
		{
			name: "test OrderBy",
			args: args{
				fields: []string{
					"field1:desc",
				},
			},
			want: &OrderBySelector{fields: []string{"field1:desc"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderBy(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestOrderBySelector_Apply(t *testing.T) {
	type fields struct {
		fields []string
	}
	type args struct {
		v *[]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderBySelector{
				fields: tt.fields.fields,
			}
			s.Apply(tt.args.v)
		})
	}
}
*/
func Test_uqe(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test uqe",
			args: args{
				val: "!%#escaped",
			},
			want: "%21%25%23escaped",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uqe(tt.args.val); got != tt.want {
				t.Errorf("uqe() = %v, want %v", got, tt.want)
			}
		})
	}
}
