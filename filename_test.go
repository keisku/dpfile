package main

import (
	"reflect"
	"testing"
)

func Test_newFilename(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    filename
		wantErr bool
	}{
		{
			args: args{
				v: "test.jpg",
			},
			want: filename{
				basename:  "test",
				extention: ".jpg",
			},
		},
		{
			args: args{
				v: "test/test.jpg",
			},
			want: filename{
				basename:  "test",
				extention: ".jpg",
			},
		},
		{
			args: args{
				v: "dir/test.jpg",
			},
			want: filename{
				basename:  "test",
				extention: ".jpg",
			},
		},
		{
			args: args{
				v: "../test.jpg",
			},
			want: filename{
				basename:  "test",
				extention: ".jpg",
			},
		},
		{
			args: args{
				v: "../test/test/test.jpg",
			},
			want: filename{
				basename:  "test",
				extention: ".jpg",
			},
		},
		{
			args: args{
				v: "/path/to/test.jpg",
			},
			want: filename{
				basename:  "test",
				extention: ".jpg",
			},
		},
		{
			args: args{
				v: "test",
			},
			want: filename{
				basename:  "test",
				extention: "",
			},
		},
		{
			args: args{
				v: "test/",
			},
			want: filename{
				basename:  "",
				extention: "",
			},
			wantErr: true,
		},
		{
			args: args{
				v: "path/to/test/",
			},
			want: filename{
				basename:  "",
				extention: "",
			},
			wantErr: true,
		},
		{
			args: args{
				v: "../test/",
			},
			want: filename{
				basename:  "",
				extention: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newFilename(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("newFilename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
