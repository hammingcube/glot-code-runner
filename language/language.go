package language

import (
	"github.com/maddyonline/glot-code-runner/language/assembly"
	"github.com/maddyonline/glot-code-runner/language/bash"
	"github.com/maddyonline/glot-code-runner/language/c"
	"github.com/maddyonline/glot-code-runner/language/clojure"
	"github.com/maddyonline/glot-code-runner/language/coffeescript"
	"github.com/maddyonline/glot-code-runner/language/cpp"
	"github.com/maddyonline/glot-code-runner/language/csharp"
	"github.com/maddyonline/glot-code-runner/language/d"
	"github.com/maddyonline/glot-code-runner/language/elixir"
	"github.com/maddyonline/glot-code-runner/language/erlang"
	"github.com/maddyonline/glot-code-runner/language/fsharp"
	"github.com/maddyonline/glot-code-runner/language/golang"
	"github.com/maddyonline/glot-code-runner/language/haskell"
	"github.com/maddyonline/glot-code-runner/language/idris"
	"github.com/maddyonline/glot-code-runner/language/java"
	"github.com/maddyonline/glot-code-runner/language/javascript"
	"github.com/maddyonline/glot-code-runner/language/julia"
	"github.com/maddyonline/glot-code-runner/language/lua"
	"github.com/maddyonline/glot-code-runner/language/nim"
	"github.com/maddyonline/glot-code-runner/language/ocaml"
	"github.com/maddyonline/glot-code-runner/language/perl"
	"github.com/maddyonline/glot-code-runner/language/php"
	"github.com/maddyonline/glot-code-runner/language/python"
	"github.com/maddyonline/glot-code-runner/language/ruby"
	"github.com/maddyonline/glot-code-runner/language/rust"
	"github.com/maddyonline/glot-code-runner/language/scala"
	"github.com/maddyonline/glot-code-runner/language/swift"
)

type runFn func([]string) (string, string, error)

var languages = map[string]runFn{
	"assembly":     assembly.Run,
	"bash":         bash.Run,
	"c":            c.Run,
	"clojure":      clojure.Run,
	"coffeescript": coffeescript.Run,
	"csharp":       csharp.Run,
	"d":            d.Run,
	"elixir":       elixir.Run,
	"cpp":          cpp.Run,
	"erlang":       erlang.Run,
	"fsharp":       fsharp.Run,
	"haskell":      haskell.Run,
	"idris":        idris.Run,
	"go":           golang.Run,
	"java":         java.Run,
	"javascript":   javascript.Run,
	"julia":        julia.Run,
	"lua":          lua.Run,
	"nim":          nim.Run,
	"ocaml":        ocaml.Run,
	"perl":         perl.Run,
	"php":          php.Run,
	"python":       python.Run,
	"ruby":         ruby.Run,
	"rust":         rust.Run,
	"scala":        scala.Run,
	"swift":        swift.Run,
}

func IsSupported(lang string) bool {
	_, supported := languages[lang]
	return supported
}

func Run(lang string, files []string) (string, string, error) {
	return languages[lang](files)
}
