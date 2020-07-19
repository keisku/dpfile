# dpfile

dpfile is a file cloning tool. You can duplicate a file with any name and sequential number.

## Install

```
go get "github.com/kskumgk63/dpfile"
cd $GOPATH/src/github.com/kskumgk63/dpfile
go install
```

## Usage

```
GLOBAL OPTIONS:
   --src value, -s value       path to the file you want to duplicate
   --dst value, -d value       path to the destination which the file will be duplicated (default: "./")
   --filename value, -f value  
   --offset value, -o value    -1 < offset (default: "0")
   --limit value, -l value     0 < limit < 10001 (default: "1")
   --help, -h                  show help (default: false)
```

```
dpfile --src ./testdir/gopher.png
```

output

```
./
├── gopher_duplicated.png
└── testdir
    ├── gopher.png
    └── helloworld.txt
```

---

```
dpfile --src ./testdir/helloworld.txt --dst ./testdir/
```

output

```
./testdir
├── gopher.png
├── helloworld.txt
└── helloworld_duplicated.txt
```

---

```
dpfile --src ./testdir/helloworld.txt --dst ./testdir/ --limit 10
```

output

```
./testdir
├── gopher.png
├── helloworld.txt
├── helloworld_duplicated.txt
├── helloworld_duplicated1.txt
├── helloworld_duplicated2.txt
├── helloworld_duplicated3.txt
├── helloworld_duplicated4.txt
├── helloworld_duplicated5.txt
├── helloworld_duplicated6.txt
├── helloworld_duplicated7.txt
├── helloworld_duplicated8.txt
└── helloworld_duplicated9.txt
```

---

```
dpfile --src ./testdir/helloworld.txt --dst ./testdir/ --offset 100 --limit 10
```

output

```
./testdir
├── gopher.png
├── helloworld.txt
├── helloworld_duplicated100.txt
├── helloworld_duplicated101.txt
├── helloworld_duplicated102.txt
├── helloworld_duplicated103.txt
├── helloworld_duplicated104.txt
├── helloworld_duplicated105.txt
├── helloworld_duplicated106.txt
├── helloworld_duplicated107.txt
├── helloworld_duplicated108.txt
└── helloworld_duplicated109.txt
```

---

```
dpfile --src ./testdir/gopher.png --dst ./testdir/ --filename mouse
```

output

```
./testdir
├── gopher.png
├── helloworld.txt
└── mouse.png
```

---

```
dpfile --src ./testdir/gopher.png --dst ./testdir/ --filename mouse --offset 999
```

output

```
./testdir
├── gopher.png
├── helloworld.txt
└── mouse999.png
```
