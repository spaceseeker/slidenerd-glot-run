package language

import (
	"slidenerd-glot-run/language/assembly"
	"slidenerd-glot-run/language/ats"
	"slidenerd-glot-run/language/bash"
	"slidenerd-glot-run/language/c"
	"slidenerd-glot-run/language/clojure"
	"slidenerd-glot-run/language/cobol"
	"slidenerd-glot-run/language/coffeescript"
	"slidenerd-glot-run/language/cpp"
	"slidenerd-glot-run/language/crystal"
	"slidenerd-glot-run/language/csharp"
	"slidenerd-glot-run/language/d"
	"slidenerd-glot-run/language/elixir"
	"slidenerd-glot-run/language/elm"
	"slidenerd-glot-run/language/erlang"
	"slidenerd-glot-run/language/fsharp"
	"slidenerd-glot-run/language/golang"
	"slidenerd-glot-run/language/groovy"
	"slidenerd-glot-run/language/haskell"
	"slidenerd-glot-run/language/idris"
	"slidenerd-glot-run/language/java"
	"slidenerd-glot-run/language/javascript"
	"slidenerd-glot-run/language/julia"
	"slidenerd-glot-run/language/kotlin"
	"slidenerd-glot-run/language/lua"
	"slidenerd-glot-run/language/mercury"
	"slidenerd-glot-run/language/nim"
	"slidenerd-glot-run/language/ocaml"
	"slidenerd-glot-run/language/perl"
	"slidenerd-glot-run/language/perl6"
	"slidenerd-glot-run/language/php"
	"slidenerd-glot-run/language/python"
	"slidenerd-glot-run/language/ruby"
	"slidenerd-glot-run/language/rust"
	"slidenerd-glot-run/language/scala"
	"slidenerd-glot-run/language/swift"
	"slidenerd-glot-run/language/typescript"
	"slidenerd-glot-run/language/solidity"
)

type runFn func([]string, string) (string, string, error)

var languages = map[string]runFn{
	"assembly":     assembly.Run,
	"ats":          ats.Run,
	"bash":         bash.Run,
	"c":            c.Run,
	"clojure":      clojure.Run,
	"cobol":        cobol.Run,
	"coffeescript": coffeescript.Run,
	"crystal":      crystal.Run,
	"csharp":       csharp.Run,
	"d":            d.Run,
	"elixir":       elixir.Run,
	"elm":          elm.Run,
	"cpp":          cpp.Run,
	"erlang":       erlang.Run,
	"fsharp":       fsharp.Run,
	"haskell":      haskell.Run,
	"idris":        idris.Run,
	"go":           golang.Run,
	"groovy":       groovy.Run,
	"java":         java.Run,
	"javascript":   javascript.Run,
	"julia":        julia.Run,
	"kotlin":       kotlin.Run,
	"lua":          lua.Run,
	"mercury":      mercury.Run,
	"nim":          nim.Run,
	"ocaml":        ocaml.Run,
	"perl":         perl.Run,
	"perl6":        perl6.Run,
	"php":          php.Run,
	"python":       python.Run,
	"ruby":         ruby.Run,
	"rust":         rust.Run,
	"scala":        scala.Run,
	"swift":        swift.Run,
	"typescript":   typescript.Run,
	"solidity":		solidity.Run,
}

func IsSupported(lang string) bool {
	_, supported := languages[lang]
	return supported
}

func Run(lang string, files []string, stdin string) (string, string, error) {
	return languages[lang](files, stdin)
}
