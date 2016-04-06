package runlib_test

import (
	"bytes"
	"github.com/maddyonline/glot-code-runner/runlib"
	"strings"
	"testing"
)

const INPUT_STR = `{
  "language": "cpp",
  "files": [
    {
      "name": "main.cpp",
      "content": "# include <iostream>\nusing namespace std;\nint main() {string s;while(cin >> s) {cout << s.size() << endl;}}"
    }, {
    "name": "_stdin_",
    "content": "abc\nhello"
    }
  ]
}`

const EXPECTED_OUTPUT = `{"stdout":"3\n5\n","stderr":"","error":""}`

func TestRunner(t *testing.T) {
	var b bytes.Buffer
	runlib.Run(strings.NewReader(INPUT_STR), &b)
	got := strings.Trim(b.String(), " \n")
	if got != EXPECTED_OUTPUT {
		t.Errorf("Expected %q, got %q", EXPECTED_OUTPUT, got)
	}
}
