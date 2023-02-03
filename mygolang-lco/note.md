이 재생목록로 배웁니다.

https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa

# Welcome to series on GO programming language

https://www.youtube.com/watch?v=JoJ8Sw5Yb4c

미국과 중국권 스타트업과 대기업 모두 좋아하는 언어입니다.

golang의 공통점과 차이점을 다루고 VCode에서 작성합니다.

타입, 반복문, 함수 등을 다룹니다.

요청과 응답과 다루고 웹 API도 만듭니다.

MongoDB로 백엔드를 만듭니다.

golang의 특수한 기능도 만들어 봅니다.

공식문서를 읽는 방법을 배웁니다. 공식문서자체를 읽기도 합니다.

클라우드와 인프라에 상당히 유용한 언어입니다.

# Before you start with golang

https://www.youtube.com/watch?v=F3klnY_r8FU

golang은 경력과 무관하게 좋아하게 될 언어입니다. 이 언어가 좋아하게 될 이유는 나중에 알게 될 것입니다.

고퍼가 될 것입니다.

오늘날 디자인된 언어입니다. 많은 언어는 나중에 기능을 추가합니다. 하지만 언어는 이미 프로덕션 코드용 언어입니다.

golang에 자주 나오는 질문입니다.

- 언어는 컴파일될 수 있습니다. 인터프린터처럼 동작해도 됩니다. 자바처럼 중간에 있어도 됩니다.
  - 자바는 컴파일 되지않습니다. 중간 바이너리를 만들고 해결하게 만듭니다.
  - golang은 JVM같은 것이 필요없습니다.
- golang은 운영체제에 각각 다르게 실행파일을 만들어줍니다.

golang으로 만들 수 있는 것들입니다.

- 클라우드로 MSA로 쪼갤 때 많이 사용합니다.

이 언어에 baggage를 가져오지 말도록 합니다. 이 언어의 마인드셋은 다릅니다. 디자인의도가 그렇습니다. 평소에 익숙한 기능이 많이 없습니다.

다른 언어를 배운 것이 분명히 도움될 것이지만 선입견을 갖지 말도록 합니다.

- golang은 OOP같지만 아닙니다.
- 클래스가 없습니다. 간결함을 의도로 만들어졌습니다.

언어의 없는것은 무엇인가?

- try, catch가 없습니다. 처음부터 필요없습니다.
- lexer가 대부분 대신 동작해줍니다. 세미콜론이 없습니다.
- 상당히 많은 기능이 언어에 없습니다.

# Golang installation and hello world

https://www.youtube.com/watch?v=62qGe9yhiJI

최신 버전을 설치해도 정상동작할 것입니다. 하지만 과거 버전은 오히려 권장하지 않습니다. 과거에는 동작원리도 다르기도 했습니다.

설치 이후 패키지 오류도 설정하도록 합니다.

https://velog.io/@seondal/Go-package-main-%EC%98%A4%EB%A5%98

```sh
go version
```

이렇게 확인할 수 있습니다.

홈페이지에서 플레이그라운드는 실험 이외에는 권장하지 않습니다. 에디터는 VSCode를 사용할 것입니다.

모듈명을 그대로 작성하도록 합니다.

새로운 폴더를 만들고 go확장자가 있는 go 파일을 만들었으면 터미널에 다음 명령을 하도록 합니다.

```sh
go mod init
```

npm init과 유사합니다.

```sh
go mod init hello
```

이제는 main.go를 작성합니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go world")
}
```

이렇게 작성하고 터미널에 다음 명령을 합니다.

```sh
go run main.go
```

이렇게 작성하면 코드를 컴파일하지 않고 실행합니다.

# GOPATH and reading go docs

https://www.youtube.com/watch?v=QEZlivtFOZk

go run main.go

go는 go 프로그램 명령입니다. run 명령입니다. main.go는 실행할 파일입니다.

대부분의 언어는 help를 지원합니다.

```txt
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	work        workspace maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

	buildconstraint build constraints
	buildmode       build modes
	c               calling between Go and C
	cache           build and test caching
	environment     environment variables
	filetype        file types
	go.mod          the go.mod file
	gopath          GOPATH environment variable
	gopath-get      legacy GOPATH go get
	goproxy         module proxy protocol
	importpath      import path syntax
	modules         modules, module versions, and more
	module-get      module-aware go get
	module-auth     module authentication using go.sum
	packages        package lists and patterns
	private         configuration for downloading non-public code
	testflag        testing flags
	testfunc        testing functions
	vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.
```

이런 피드백을 받을 수 있습니다.

```sh
go help gopath
```

이렇게 문서를 알게 될 수 있습니다.

```sh
go env GOPATH
```

이명령으로 어디에 설치되었는지 찾을 수 있습니다.

여기로 이동하면 볼 수 있는 것들입니다.

```sh
ls
```

`~go/pkg/mod`에서 다음것들을 볼 수 있습니다.

- cache
- github.com
- golang.org
- honnef.co
- mvdan.cc

나중에 다루지만 여기서 캐시가 중요합니다.

# Lexer in golang and Types

https://www.youtube.com/watch?v=elYPAeX9h1E

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go world")
}
```

golang에는 세미콜론을 넣을 수 있는 곳과 없는 곳이 있습니다.

lexer가 세미콜론을 넣을 수 있도록 합니다.

과거에는 세미콜론을 강제하도록 했습니다. 하지만 언어가 lexer가 있습니다. lexer는 문법이 적절한지 판단합니다. 물론 다른 버그를 막아주지는 않습니다. lexer는 자동으로 세미콜론을 필요한 곳에 붙여버립니다.

golang의 자료형입니다. 대료형은 대소문자를 구분합니다. 또 public, static 지정도 가능합니다.

golang은 변수를 사전에 정의해야 합니다. 대부분의 golang에서 보게 되는 대부분의 것들은 타입입니다.

함수자체도 자료형으로 지원해서 함수형 프로그래밍을 지원하게 만들 수 있습니다.