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

# Variables, types and constants

https://www.youtube.com/watch?v=9fYqg6uo-UU

```sh
go mod init varriable
```

이렇게 초기설정하는 것을 항상 잊지말도록 합니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Variables")
}
```

```go
package main

import "fmt"

func main() {
	var username string = "Jake The dog"
	fmt.Println(username)
}
```

변수 선언입니다.

```go
package main

import "fmt"

func main() {
	var username string = "Jake The dog"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
}
```

> Jake The dog
> Variable is of type: string

이렇게 피드백을 줍니다.

```go
package main

import "fmt"

func main() {
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
}
```

uint8은 256을 초과할 수 없습니다. 0부터 시작합니다. 하지만 필요할 때는 int만 지정해주면 됩니다.

대부분의 경우 int를 먼저 사용하지만 운영체제 관련된 프로그래밍할 때는 unit을 사용하게 될 것입니다.

```go
package main

import "fmt"

func main() {
	var smallFloat float32 = 255.134314232132435
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
}
```

`float`는 클수록 정확성이 증가합니다.

타입별칭들입니다.

```go
package main

import "fmt"

func main() {

	// 기본값과 별칭들입니다.
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)
}
```

이렇게 작성하면 기본적으로 0이 할당됩니다. 문자열을 할당하면 비어있는 문자열이 할당됩니다.

타입추론이 가능합니다.

```go
package main

import "fmt"

func main() {
	// 타입 암시
	var website = "google.com"
	fmt.Println(website)
}
```

no var 스타일이 있습니다. 다른 말로 바다사자 연산자라고 부릅니다.

```go
package main

import "fmt"

// 이것은 허용되지 않습니다.
jwtToken := 3000_000

func main() {
	// 바다사자 연산자
	numberUser := 300_000
	fmt.Println(numberUser)
}
```

바다사자 연산자는 글로벌스코프로 정의할 수 없습니다.

```go
const LoginToken string = "qerwerqw"

```

```go
package main

import "fmt"

const LoginToken string = "qerwerqw"

func main() {
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
```

이렇게 작성하면 다른언어의 public 키워드를 붙인 것과 동일합니다. 대문자로 시작하면 이렇게 동작합니다.

# Comma ok syntax and packages in golang

https://www.youtube.com/watch?v=zYIZtbyUIDY

사용자에게 input을 받는법입니다.

```go
package main

import "fmt"

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
}
```

이렇게 시작합니다.

사용자에게 자료입력을 받습니다.

이제 익숙해질 문법입니다.

https://pkg.go.dev/

여기서는 golang의 패키지를 검색할 수 있습니다.

buffer는 자료를 저장하고 읽고 쓸 수 있습니다.

패키지는 중요한 것은 검색 능력입니다. 무엇이 존재하고 어떻게 사용하는지 자주 사용하는 것은 외우고 유용한 것은 메모만 합니다.

바다사자 연산자는 무엇이 들어올지 모를 때 많이 사용합니다.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok 문법입니다. err ok 문법이라고 합니다.
	// golang은 try catch가 없습니다. 에러를 부울처럼 취급하기 바라면서 설계했습니다.

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

}
```

golang에서 주의할 점이 있습니다. 문자열을 표시할 때 따옴표, 쌍따옴표 다릅니다.

`reader.ReadString('\n')`은 따옴표 문자열을 대입했습니다.

err, ok 혹은 comma, ok 문법입니다. golang을 다루면서 자주 보게될 문법입니다.

err를 만약에 사용하지않으면 언더스코어(`_`)로 작성합니다. err는 에러 핸들링할 때 사용합니다.

이상해보이지만 상당히 큰 장점으로 작용합니다.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok 문법입니다. err ok 문법이라고 합니다.
	// golang은 try catch가 없습니다. 에러를 부울처럼 취급하기 바라면서 설계했습니다.

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this is %T", input)
}
```

이렇게 피드백을 받으면 당연히 예상한 것처럼 문자열 자료형입니다.

# Conversions in golang

https://www.youtube.com/watch?v=3j43y-PFJPI

지난 시간 만든 것을 그대로 다시만듭니다.

파라미터는 url, string은 변환하면서 다양한 문제를 만들 수 있습니다.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
}

```

여기까지 금방왔습니다.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	// 에러를 처리해봅니다.
	numRating, err := strconv.ParseFloat(input, 64)

	// err 속에 무엇인가가 들어있습니다.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating + 1)
	}
}
```

예상과 다르게 형변환을 실패했습니다. 엔터를 하는 순간 `\n`문자열이 뒤에 붙어있는 것입니다.

> strconv.ParseFloat: parsing "1\n": invalid syntax

이런 피드백을 주고 있습니다. 해당하는 `strconv.ParseFloat` 메서드는 처리하는데 문제가 있다고 힌트를 줍니다. Comma ok 표기법에서 에러를 이렇게 바로 접근할 수 있게 해줍니다.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	// 에러를 처리해봅니다.
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// err 속에 무엇인가가 들어있습니다.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
```

이번에는 처리가 성공했습니다. 아까 에러피드백을 단계적으로 보고 프로그래밍의 과정상 효율적으로 코드를 작성하게 되었습니다. 발생할 수 있는 에러를 코드를 절차에 따라 작성하고 에러를 중간 중간 바로 처리한 점이 golang 언어의 설계의도처럼 보입니다.

# Handling time in golang

https://www.youtube.com/watch?v=xsAz24z4Hdg

이번에는 시간을 제어해봅니다.

https://pkg.go.dev/time

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time")

	presentTime := time.Now()
	fmt.Println(presentTime)
}
```

> 2023-02-05 02:59:53.269773 +0900 KST m=+0.000091251

예상한것처럼 동작합니다.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // 02-05-2023 03:02:21 Sunday
}
```

포멧을 이렇게 정리할 수 있습니다.

유용하게 이상한 언어인데 여기는 이상하기만 합니다.

타임 패키지는 확실하게 이상합니다. 왜 1월 2일 2006년 3시 4분 5초를 기억해야 하는지 상당히 이상합니다. 이스터애그가 있을 것 같습니다.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time")
	// 날짜 생성
	createdDate := time.Date(2020, time.May, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
```

이렇게 날짜 생성은 동일하게 포맷팅해서 정리할 수 있습니다.

https://stackoverflow.com/questions/45160822/what-does-20060102150405-mean

이런 의미가 있었습니다. 다시 생각해보면 어느 날로 생각하고 가독성있게 읽을 수 있습니다.

# Build for windows, linux and mac

https://www.youtube.com/watch?v=9vxsO5vMTAw

이번에는 흥미로운 부분을 다룹니다.

golang은 각각의 운영체제마다 실행파일을 컴파일할 수 있게 해줍니다.

```sh
go env
```

> GOOS="darwin"

본인의 운영체제를 확인할 수 있습니다. 저는 맥이라 이런 피드백을 줍니다.

```sh
GOOS="windows" go build
```

윈도우 운영체제로 컴파일합니다.

```sh
GOOS="linux" go build
```

리눅스 운영체제로 컴파일합니다.

각각의 운영체제로 컴파일하면서 하나의 코드베이스로 여러 플랫폼을 대응할 수 있습니다.

# Memory management in golang

https://www.youtube.com/watch?v=G1SP9uDJD0g

이번장에는 자료형을 다루기 시작합니다. 메모리관리를 먼저 다루고 진행합니다.

메모리관리는 golang에는 간단합니다. 개발자의 일이 아닙니다. 메모리 할당은 자동으로 처리됩니다.

메모리관리는 2가지를 메서드로 처리합니다. new(), make()입니다. 차이가 있습니다.

공통점은 모두 데이터를 공간에 할당할 때 활용합니다.

new() 메서드는 메모리는 초기화되지만 할당되지 않습니다. make는 메모리를 초기화하고 할당할 수 있습니다.

zeroed-storage는 처음에 자료를 넣을 수 없습니다.

흥미로운 부분입니다. GC는 자동처리됩니다. 스코프 밖에 nil이 되는 것은 수거대상이 됩니다.

golang에서 CG는 초기에는 비판을 많이 받았습니다. 그리고 나중에 리팩토링하고 성능이 또 향상되었습니다.

https://pkg.go.dev/runtime

runtime 자체가 또 패키지입니다. 여기서 로우레벨 정보를 얻을 수 있습니다. 나중에 고인물되었을 때 다시 방문하기 권장합니다.

GOGC는 CG 변수입니다. 어느정도 용량을 설정하고 초과하면 수거를 시작합니다. 보통은 제어할 일은 없지만 제어가 가능합니다.

https://pkg.go.dev/runtime#NumCPU

프로세서를 확인할 때 사용할 수 있습니다. 언어의 메모리관리의 실용적인 정도까지 다루었습니다.

이주제랑 연결된 포인터를 다음에 다루겠습니다.

# Pointers in golang

https://www.youtube.com/watch?v=-BFJ0dZyxHg

많은 사람들에게 포인터를 필요이상으로 자세히 설명하면서 이해하기 어렵게 만드는 사람들이 있습니다.

여기서는 간단하게만 설명합니다.

포인터가 존재하고 언어에 포인터가 왜 존재하는지 파악해야 합니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("포인터에 온것을 환영합니다.")
}
```

여기서 시작합니다.

포인터가 존재하는 이유입니다. 프로그래밍 세계에서 존재하는 문제입니다. 변수, 상수를 선언할 때마다 메모리에 할당하고 그것을 참조한다는 것을 압니다. 가끔은 이 변수는 같이 할당되지 않고 복사본이 할당될 때가 있습니다. 일관성이 깨질 때가 있습니다. 변수식별자를 할당하는 대신에 포인터를 할당하는 것으로 메모리가 바라보는 대상을 그대로 할당하게 만듭니다. 실제 메모리에 하당하는 그 값을 할당하게 만듭니다. 이 부분을 포인터로 보증하게 할 수 있습니다. 이것이 포인터가 만들어지고 활용하게 된 이유가 됩니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("포인터에 온것을 환영합니다.")
	var one int

}
```

이렇게 할당할 식별자만 정의합니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("포인터에 온것을 환영합니다.")
	// 정수형을 할당할 포인터를 지정하는 방법입니다.
	var ptr *int

	fmt.Println("포인터의 값:", ptr)

}
```

포인터를 만들고 지정을 안하면 nil이 들어있습니다.

> 포인터에 온것을 환영합니다.
> 포인터의 값: <nil>

```go
package main

import "fmt"

func main() {
	fmt.Println("포인터에 온것을 환영합니다.")
	// 정수형을 할당할 포인터를 지정하는 방법입니다.
	// var ptr *int
	// fmt.Println("포인터의 값:", ptr)

	myNumber := 23
	// 포인터로 메모리상 동일한 23을 바라봅니다. &은 참조한다는 의미입니다.
	var ptr = &myNumber
	fmt.Println("포인터의 주소 값: ", ptr)
	fmt.Println("포인터의 실제 값: ", *ptr)
}
```

포인터는 메모리 주소값입니다. 포인터 속에는 무엇이 있는지는 `*`으로 알아냅니다.

포인터의 간단한 예시입니다. 나중에 포인터는 상당히 유용한 기능이라는 것을 배우게 될 것입니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("포인터에 온것을 환영합니다.")
	// 정수형을 할당할 포인터를 지정하는 방법입니다.

	myNumber := 23
	// 포인터로 메모리상 동일한 23을 바라봅니다. &은 참조한다는 의미입니다.
	var ptr = &myNumber
	fmt.Println("포인터의 주소 값: ", ptr)
	fmt.Println("포인터의 실제 값: ", *ptr)
	*ptr = *ptr + 2

	fmt.Println("새로운 값:", myNumber)
}
```

여기서 반환값을 예측할 수 있으면 포인터를 이해한 것입니다.

저는 23을 예상했습니다. 하지만 새로운 값은 25였습니다. 정적언어랑 생각과 다르게 변형이 가해졌습니다. 실제 값에 변형이 가해진다는 것을 배웠습니다.

> 포인터에 온것을 환영합니다.
> 포인터의 주소 값: 0xc00001e0b8
> 포인터의 실제 값: 23
> 새로운 값: 25

```go
package main

import "fmt"

func main() {
	fmt.Println("포인터에 온것을 환영합니다.")

	// myNumber은 0xc00001e0b8에 있는 23을 참조하고 있습니다.
	myNumber := 23
	// ptr은 0xc00001e0b8을 참조합니다.
	var ptr = &myNumber
	fmt.Println("포인터의 주소 값: ", ptr) // 0xc00001e0b8이 됩니다.
	fmt.Println("포인터의 실제 값: ", *ptr) // 23이 됩니다.
	// 0xc00001e0b8의 값인 23에 2를 25로 변형을 가합니다.
	*ptr = *ptr + 2

	fmt.Println("새로운 값:", myNumber)
}
```

자바스크립트처럼 원시형 자료형이 변형이 가해지면 새로운 메모리 주소를 주는 것과 다르게 메모리 주소는 동일하고 새로운 값으로 업데이트 됩니다.

# Array in golang

https://www.youtube.com/watch?v=JoUSa8jtadE

이번에는 배열을 다룹니다. golang은 특이하게 많이 사용하지 않습니다. 공식문서에서 많이 사용하지 않기를 권장합니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Go의 배열에 온 것을 환영합니다.")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("과일 목록입니다: ", fruitList)
}
```

의도적으로 2번 인덱스에 아무것도 할당하지 않았습니다.

> 과일 목록입니다: [Apple Tomato Peach]

이런 피드백을 받습니다. 공백을 받습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Go의 배열에 온 것을 환영합니다.")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("과일 목록입니다: ", fruitList)
	fmt.Println("과일 목록입니다: ", len(fruitList))
}
```

> 과일 목록입니다: 4

특이하게 개수를 4라고 합니다. 메모리에 공간 4개를 할당해서 그런 것 같습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Go의 배열에 온 것을 환영합니다.")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("과일 목록입니다: ", fruitList)
	fmt.Println("과일 목록입니다: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("야채 목록:", len(vegList))
}
```

> 야채 목록: 3

이렇게 피드백을 줍니다.

sort같은 메서드를 지원하지 않습니다.

# Slices in golang

https://www.youtube.com/watch?v=k7hVj8QL9Co

golang에서 slices는 배열 대신에 많이 사용합니다. 자주 사용하게 될 것입니다. slices는 내부적으로 배열입니다. 더욱더 강력하기 때문에 많이 사용합니다. 잘 이해하면 데이터베이스를 더 쉽게 다룰 수 있습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("슬라이스에 왔습니다.")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("fruitList의 자료형: %T\n", fruitList)
}
```

> 슬라이스에 왔습니다.
> fruitList의 자료형: []string

이렇게 피드백을 줍니다. 슬라이스에 자료를 추가하는 방법입니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("슬라이스에 왔습니다.")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("fruitList의 자료형: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
}
```

> [Apple Tomato Peach Mango Banana]

이렇게 피드백을 받습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("슬라이스에 왔습니다.")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("fruitList의 자료형: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
}
```

> [Apple Tomato Peach Mango Banana] > [Tomato Peach Mango Banana]

첫번째 원소를 짤랐습니다.

range는 마지막에 빠지는 것은 다른 언어랑 유사합니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("슬라이스에 왔습니다.")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("fruitList의 자료형: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
}
```

> [Apple Tomato Peach Mango Banana] > [Tomato Peach]

슬라이스는 정말 자주사용하게 될 것입니다. 자주 사용하기 때문에 잘 알아둬야 합니다.

메모리관리 할 때 new, make 2가지를 다루었습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("슬라이스에 왔습니다.")
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 867
	// highScores[4] = 777 // 오버플로우 오류가 발생합니다.

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)
}
```

초기에는 메모리 공간을 설정했지만 새로운 값을 할당할 때 append를 사용하면 추가할 수 있습니다.

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("슬라이스에 왔습니다.")
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 867
	// highScores[4] = 777 // 오버플로우 오류가 발생합니다.

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)


	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
}
```

> [234 321 456 555 666 867 945]
> true

슬라이스에만 사용할 수 있습니다. 또 정렬된 경우를 판별하게 만드는 것도 가능합니다.

# How to remove a value from slice based on index in golang

https://www.youtube.com/watch?v=931nR5TGCAk

이 연산은 잘 알고 있도록 합니다.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("슬라이스에 왔습니다.")
	var courses = []string{"react.js", "javascript", "swift", "python", "php"}

	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
```

> [react.js javascript python php]

이렇게 하나의 값을 중간에 삭제할 수 있습니다. 나중에 `...`을 다루기는 할 것입니다.

# Maps in golang

https://www.youtube.com/watch?v=_0R6H1m9o78

다른 언어에서 해쉬테이블이라고도 부르는 자료형입니다. 키와 값을 짝으로 갖는 자료형입니다.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang의 map")

	languages := make(map[string]string)
}
```

키와 값의 자료형을 미리 정합니다.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang의 map")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("모든 언어 목록:", languages)
}
```

> 모든 언어 목록: map[JS:JavaScript PY:Python RB:Ruby]

컴마로 분리되어 있지 않습니다.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang의 map")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("모든 언어 목록:", languages)
	fmt.Println("JS:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("모든 언어 목록:", languages)

}
```

언어의 제어흐름을 아직 안 다루었습니다. map을 순회하는 법을 다룹니다. golang에 순회는 상당히 흥미롭습니다.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang의 map")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("모든 언어 목록:", languages)
	fmt.Println("JS:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("모든 언어 목록:", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
```

> For key JS, value is JavaScript
> For key PY, value is Python

바다사자 연산자는 로 응용하는 것은 순회하 때도 동일합니다.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang의 map")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("모든 언어 목록:", languages)
	fmt.Println("JS:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("모든 언어 목록:", languages)

	for _, value := range languages {
		fmt.Printf("value is %v\n", value)

	}

}

```

comma ok 표기법으로 key를 무시하게 만들었습니다.

https://www.youtube.com/watch?v=NMTN543WVQY

https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa

# Structs in golang

https://www.youtube.com/watch?v=NMTN543WVQY

golang에는 클래스가 없습니다. 오직 Structs만 존재합니다. Struct은 더 깊이 있게 이해가 필요한 지식입니다. 이번에는 간단하게 다룹니다.

golang에는 상속이 없습니다. 부모자식 클래스 개념이 없습니다. golang의 설계 의도는 상속을 많이 하면 복잡성이 증가한다고 봤습니다. 그래서 없습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Structs에 온것을 환영합니다.")
	// 상속 부모자식 클래스 개념이 없습니다.
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
```

저장에 가독성있게 자동정렬 됩니다. 클래스랑 비슷하기 때문에 파스칼 케이스로 작성합니다. 모든 사람들이 접근가능하기 때문에 각각의 속성도 파스칼 케이스로 적성해야 합니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Structs에 온것을 환영합니다.")
	// 상속 부모자식 클래스 개념이 없습니다.
	jake := User{"Jake", "jaketheGod@go.dev", true, 30}
	fmt.Println(jake)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
```

이렇게 데이터를 만듭니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Structs에 온것을 환영합니다.")
	// 상속 부모자식 클래스 개념이 없습니다.
	jake := User{"Jake", "jaketheGod@go.dev", true, 30}
	fmt.Println(jake)
	fmt.Printf("Jake details are: %+v\n", jake)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
```

이런 문법은 생각보다 많이 사용합니다.

> Jake details are: {Name:Jake Email:jaketheGod@go.dev Status:true Age:30}

```go
package main

import "fmt"

func main() {
	fmt.Println("Structs에 온것을 환영합니다.")
	// 상속 부모자식 클래스 개념이 없습니다.
	jake := User{"Jake", "jaketheGod@go.dev", true, 30}
	fmt.Println(jake)
	fmt.Printf("Jake details are: %+v\n", jake)
	fmt.Printf("Name is %v and email is %v.", jake.Name, jake.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
```

> Name is Jake and email is jaketheGod@go.dev.%

상속이 없다는 것과 대소문자 구분을 항상 주의하도록 합니다.

# If else in golang

https://www.youtube.com/watch?v=f_xNeRurjZY

이제부터는 제어흐름을 다룹니다. 다른 언어랑 유사하지만 조심할 부분들이 있습니다.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang if else")

	loginCount := 23

	var result string
	if loginCount < 10 {
		result = "일반 유저"
	}

	fmt.Println(result)

}
```

참고로 golang은 중괄호를 줄바꿈할 수 없습니다. 이부분은 golang의 독선입니다.

이렇게 조건문을 작성할 수 있습니다.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang if else")

	loginCount := 23

	var result string
	if loginCount < 10 {
		result = "일반 유저"
	} else if loginCount > 10 {
		result = "해비유저"
	} else {
		result = "이벤트 달성"
	}

	fmt.Println(result)

}
```

예상한대로 동작합니다.

golang은 조건 문속에서 변수 선언할 수 있습니다.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang if else")

	loginCount := 23

	var result string
	if loginCount < 10 {
		result = "일반 유저"
	} else if loginCount > 10 {
		result = "해비유저"
	} else {
		result = "이벤트 달성"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("짝수")
	} else {
		fmt.Println("홀수")
	}

	// 요청을 받으면서 할당하게 만드는 방법입니다.
	if num := 3; num < 10 {
		fmt.Println("num은 10보다 큽니다.")
	} else {
		fmt.Println("num은 10보다 작습니다.")
	}
}
```

다른 언어를 공부해봤으면 상당히 간편합니다.
