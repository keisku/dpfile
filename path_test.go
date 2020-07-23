package main

import (
	"reflect"
	"testing"
)

func Test_newSrc(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		args       args
		want       src
		wantErrMsg string
	}{
		{
			name: "success",
			args: args{
				path: "./testdir/helloworld.txt",
			},
			want: src{
				path: path{
					dir: "testdir",
					filename: filename{
						basename:  "helloworld",
						extention: ".txt",
					},
				},
			},
		},
		{
			name: "err: path is not the destination of file",
			args: args{
				path: "./testdir/",
			},
			want:       src{},
			wantErrMsg: "You should set the file path to --src, -s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newSrc(tt.args.path)
			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("newSrc() error = %v, wantErr %v", err.Error(), tt.wantErrMsg)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSrc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newDst(t *testing.T) {
	var (
		testdir_helloworld, _ = newSrc("./testdir/helloworld.txt")
	)
	type args struct {
		src         src
		destination string
		filename    string
	}
	tests := []struct {
		name    string
		args    args
		want    dst
		wantErr bool
	}{
		{
			args: args{
				src:         testdir_helloworld,
				destination: "./",
			},
			want: dst{
				path{
					dir: dir("."),
					filename: filename{
						basename:  "helloworld_duplicated",
						extention: ".txt",
					},
				},
			},
		},
		{
			args: args{
				src:         testdir_helloworld,
				destination: "../",
				filename:    "hoge",
			},
			want: dst{
				path{
					dir: dir(".."),
					filename: filename{
						basename:  "hoge",
						extention: ".txt",
					},
				},
			},
		},
		{
			args: args{
				src:         testdir_helloworld,
				destination: "./testdir/helloworld.txt",
				filename:    "helloworld",
			},
			want: dst{
				path{
					dir: dir("testdir"),
					filename: filename{
						basename:  "helloworld_duplicated",
						extention: ".txt",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newDst(tt.args.src, tt.args.destination, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("newDst() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDst() = %v, want %v", got, tt.want)
			}
		})
	}
}
