# TCONF

<p align="center">
  <img alt="golangci-lint logo" src="tconf.png" height="300" />
  <p align="center">A new impressive, easy and pleasant config!</p>
</p>

A new impressive, easy and pleasant config!

# INSTALLATION

```sh
$ go get github.com/ivm97/tconf
```

# EXAMPLES
<h3> Create and save config file:</h3>

```go
    file := New()
	file.AddSection("project")
	file.AddKeyValue("paths", "cmd,libs,pkg")
	file.AddKeyValue("files", "cmd/app/main.cpp")
	file.AddSection("dirs")
	file.AddKeyValue("compiler", "C:/gcc/bin")
	if err := file.Save("settings.tc"); err != nil {
		t.Error(err)
	}
```

<h3> Read config from *os.File or path :</h3>

```go
   tc, err := Open("tests/conf.hconf")
	if err != nil {
		t.Error(err)
	}

  if value, ok := tc.From("section").Get("red"); ok{
    fmt.Println(value)
  }
```