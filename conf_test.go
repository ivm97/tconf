package tconf

import "testing"

func TestNew(t *testing.T) {
	file := New()
	file.AddSection("project")
	file.AddKeyValue("paths", "cmd,libs,pkg")
	file.AddKeyValue("files", "cmd/app/main.cpp")
	file.AddSection("dirs")
	file.AddKeyValue("compiler", "C:/gcc/bin")
	if err := file.Save("settings.tc"); err != nil {
		t.Error(err)
	}
}

func TestRead(t *testing.T) {
	tc, err := Open("settings.tc")
	if err != nil {
		t.Error(err)
	}

	res, exist := tc.From("project").Get("paths")
	if exist {
		if v, ok := res.(string); ok {
			if v != "cmd,libs,pkg" {
				t.Errorf("Expected cmd,libs,pkg, but got: %s", v)
			}

		}

	} else {
		t.Error(exist)
	}

}

func TestOpen(t *testing.T) {
	if _, err := Open("a"); err == nil {
		t.Errorf("Expected error, but got %v", err)
	}
}

func TestFail(t *testing.T) {
	tc, err := Open("settings.tc")
	if err != nil {
		t.Error(err)
	}

	for k, v := range *tc.From("project") {
		t.Log("KEY:", k, "VALUE:", v)
	}
}
