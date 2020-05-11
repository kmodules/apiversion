package apiversion

import (
	"reflect"
	"testing"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *Version
		wantErr bool
	}{
		{
			name: "v1",
			args: args{s: "v1"},
			want: &Version{
				X: 1,
				Y: "",
				Z: 0,
			},
			wantErr: false,
		},
		{
			name: "v1alpha10",
			args: args{s: "v1alpha10"},
			want: &Version{
				X: 1,
				Y: "alpha",
				Z: 10,
			},
			wantErr: false,
		},
		{
			name: "v1rc10",
			args: args{s: "v1rc10"},
			want: &Version{
				X: 1,
				Y: "rc",
				Z: 10,
			},
			wantErr: false,
		},
		{
			name:    "v1dev10",
			args:    args{s: "v1dev10"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVersion(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_Compare(t *testing.T) {
	type args struct {
		one   string
		other string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "v2 <-> v10",
			args: args{
				one:   "v2",
				other: "v10",
			},
			want: -1,
		},
		{
			name: "v10 <-> v2",
			args: args{
				one:   "v10",
				other: "v2",
			},
			want: 1,
		},
		{
			name: "v2 <-> v2",
			args: args{
				one:   "v2",
				other: "v2",
			},
			want: 0,
		},
		{
			name: "v1alpha10 <-> v1alpha2",
			args: args{
				one:   "v1alpha10",
				other: "v1alpha2",
			},
			want: 1,
		},
		{
			name: "v1alpha2 <-> v1alpha10",
			args: args{
				one:   "v1alpha2",
				other: "v1alpha10",
			},
			want: -1,
		},
		{
			name: "v1alpha10 <-> v1alpha10",
			args: args{
				one:   "v1alpha10",
				other: "v1alpha10",
			},
			want: 0,
		},
		{
			name: "v1alpha10 <-> v1beta1",
			args: args{
				one:   "v1alpha10",
				other: "v1beta1",
			},
			want: -1,
		},
		{
			name: "v1beta1 <-> v1alpha10",
			args: args{
				one:   "v1beta1",
				other: "v1alpha10",
			},
			want: 1,
		},
		{
			name: "v1alpha10 <-> v1rc10",
			args: args{
				one:   "v1alpha10",
				other: "v1rc10",
			},
			want: -1,
		},
		{
			name: "v1rc10 <-> v1alpha10",
			args: args{
				one:   "v1rc10",
				other: "v1alpha10",
			},
			want: 1,
		},
		{
			name: "v1beta10 <-> v1rc10",
			args: args{
				one:   "v1beta10",
				other: "v1rc10",
			},
			want: -1,
		},
		{
			name: "v1rc10 <-> v1beta10",
			args: args{
				one:   "v1rc10",
				other: "v1beta10",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			one, err := NewVersion(tt.args.one)
			if err != nil {
				t.Errorf("NewVersion() error = %v", err)
				return
			}
			other, err := NewVersion(tt.args.other)
			if err != nil {
				t.Errorf("NewVersion() error = %v", err)
				return
			}
			if got := one.Compare(*other); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
