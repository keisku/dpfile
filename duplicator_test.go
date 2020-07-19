package main

import (
	"reflect"
	"testing"
)

func Test_newDuplicator(t *testing.T) {
	type args struct {
		src      string
		dst      string
		filename string
		o        offset
		l        limit
	}
	tests := []struct {
		name       string
		args       args
		want       duplicator
		wantErrMsg string
	}{
		{
			name: `given src "testdir/helloworld.txt", dst "testdir"`,
			args: args{
				src:      "testdir/helloworld.txt",
				dst:      "testdir",
				filename: "",
				o:        0,
				l:        1,
			},
			want: duplicator{
				src: src{
					path: path{
						dir:       "testdir",
						name:      "helloworld",
						extention: ".txt",
					},
				},
				dst: dst{
					path: path{
						dir:       ".",
						name:      "helloworld_duplicated",
						extention: ".txt",
					},
				},
				offset: 0,
				limit:  1,
			},
		},
		{
			name: `given src "testdir/helloworld.txt", dst "testdir/"`,
			args: args{
				src:      "testdir/helloworld.txt",
				dst:      "testdir/",
				filename: "",
				o:        0,
				l:        1,
			},
			want: duplicator{
				src: src{
					path: path{
						dir:       "testdir",
						name:      "helloworld",
						extention: ".txt",
					},
				},
				dst: dst{
					path: path{
						dir:       "testdir",
						name:      "helloworld_duplicated",
						extention: ".txt",
					},
				},
				offset: 0,
				limit:  1,
			},
		},
		{
			name: `given src "testdir/helloworld.txt", dst "testdir/", filename "hello"`,
			args: args{
				src:      "testdir/helloworld.txt",
				dst:      "testdir/",
				filename: "hello",
				o:        0,
				l:        1,
			},
			want: duplicator{
				src: src{
					path: path{
						dir:       "testdir",
						name:      "helloworld",
						extention: ".txt",
					},
				},
				dst: dst{
					path: path{
						dir:       "testdir",
						name:      "hello",
						extention: ".txt",
					},
				},
				offset: 0,
				limit:  1,
			},
		},
		{
			name: `given src "testdir/helloworld.txt", dst "testdir/", filename "hello.txt"`,
			args: args{
				src:      "testdir/helloworld.txt",
				dst:      "testdir/",
				filename: "hello.txt",
				o:        0,
				l:        1,
			},
			want: duplicator{
				src: src{
					path: path{
						dir:       "testdir",
						name:      "helloworld",
						extention: ".txt",
					},
				},
				dst: dst{
					path: path{
						dir:       "testdir",
						name:      "hello",
						extention: ".txt",
					},
				},
				offset: 0,
				limit:  1,
			},
		},
		{
			name: `given src "testdir/helloworld.txt", dst "testdir2/"`,
			args: args{
				src:      "testdir/helloworld.txt",
				dst:      "testdir2/",
				filename: "",
				o:        0,
				l:        1,
			},
			want:       duplicator{},
			wantErrMsg: "stat testdir2/: no such file or directory",
		},
		{
			name: `given src "dir/helloworld.txt", dst "testdir/"`,
			args: args{
				src:      "dir/helloworld.txt",
				dst:      "testdir/",
				filename: "",
				o:        0,
				l:        1,
			},
			want:       duplicator{},
			wantErrMsg: "stat dir/helloworld.txt: no such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newDuplicator(tt.args.src, tt.args.dst, tt.args.filename, tt.args.o, tt.args.l)
			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("newDuplicator() error = %v, wantErrMsg %v", err, tt.wantErrMsg)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDuplicator() = %v, want %v", got, tt.want)
			}
		})
	}
}
