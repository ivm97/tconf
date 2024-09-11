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
	tc, err := Open("tests/conf.hconf")
	if err != nil {
		t.Error(err)
	}

	res, exist := tc.From("section").Get("red")
	if exist {
		if v, ok := res.(string); ok {
			if v != "color" {
				t.Errorf("Expected red, but got: %s", v)
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
