이 재생목록로 배웁니다.

https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa

https://github.com/hiteshchoudhary/golang

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

# Switch case in golang and online playground

https://www.youtube.com/watch?v=Up4lTPhJBvs

golang의 Switch case 문입니다.

lugo game을 만들어봅니다.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	// 시드 생성
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice: ", diceNumber)
}

```

주사위 기능을 구현했습니다.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	// 시드 생성
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 you can open")
	case 2:
		fmt.Println("You can move 2")
	case 3:
		fmt.Println("You can move 3")
	case 4:
		fmt.Println("You can move 4")
	case 5:
		fmt.Println("You can move 5")
	case 6:
		fmt.Println("You can move 6")
	default:
		fmt.Println("What was that?")
	}
}
```

fallthrough도 만들 수 있습니다. 특정 케이스를 달성하면 그이후도 모두 실행하도록 할 수 있습니다.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	// 시드 생성
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 you can open")
	case 2:
		fmt.Println("You can move 2")
	case 3:
		fmt.Println("You can move 3")
		fallthrough
	case 4:
		fmt.Println("You can move 4")
		fallthrough
	case 5:
		fmt.Println("You can move 5")
	case 6:
		fmt.Println("You can move 6")
	default:
		fmt.Println("What was that?")
	}
}
```

> Switch and case in golang
> Value of dice: 3
> You can move 3
> You can move 4
> You can move 5

# Loop break continue and goto in golang

https://www.youtube.com/watch?v=ZWBA3l818y0

이번에는 반복문을 만듭니다. golang의 반복문은 흥미롭습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("golang의 반복문")

	days := []string{"일요일", "화요일", "수요일", "금요일", "토요일"}

	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
}
```

이렇게 반복문을 만들 수 있습니다. 가장 기본적인 형태의 순회방법입니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("golang의 반복문")

	days := []string{"일요일", "화요일", "수요일", "금요일", "토요일"}

	fmt.Println(days)

	for i := range days {
		fmt.Println(days[i])
	}
}
```

여기서 i는 인덱스를 접근합니다. 조금더 자주 활용합니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("golang의 반복문")

	days := []string{"일요일", "화요일", "수요일", "금요일", "토요일"}

	fmt.Println(days)

	for index, value := range days {
		fmt.Printf("index is %v and value is %v\n", index, value)
	}
}
```

comma ok 표기법처럼 응용도 가능합니다. 인덱스와 값 모두 얻을 수 있습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("golang의 반복문")

	days := []string{"일요일", "화요일", "수요일", "금요일", "토요일"}

	fmt.Println(days)

	rougeValue := 1

	for rougeValue < 10 {
		fmt.Println("값: ", rougeValue)
		rougeValue++
	}
}
```

```go
package main

import "fmt"

func main() {
	fmt.Println("golang의 반복문")

	days := []string{"일요일", "화요일", "수요일", "금요일", "토요일"}

	fmt.Println(days)

	rougeValue := 1

	for rougeValue < 10 {

		// 종료
		if rougeValue == 5 {
			break
		}

		// 하나 건너 뛰기
		if rougeValue == 3 {
			rougeValue++
			continue
		}

		fmt.Println("값: ", rougeValue)
		rougeValue++
	}
}
```

```go
package main

import "fmt"

func main() {
	fmt.Println("golang의 반복문")

	days := []string{"일요일", "화요일", "수요일", "금요일", "토요일"}

	fmt.Println(days)

	rougeValue := 1

	for rougeValue < 10 {

		// 종료
		if rougeValue == 5 {
			break
		}

		// 하나 건너 뛰기
		if rougeValue == 3 {
			rougeValue++
			continue
		}

		if rougeValue == 2 {
			goto loc
		}

		fmt.Println("값: ", rougeValue)
		rougeValue++
	}

loc:
	fmt.Println("go to 기능을 구현합니다.")
}
```

또 반복문에서도 보는 comma ok 표기법을 주의하도록 합니다.

# Functions in golang

https://www.youtube.com/watch?v=rcUST3QvVOQ

함수는 당연히 존재합니다. 언어별로 대부분의 경우 함수는 존재할 수 밖에 없습니다. golang은 처음 사용할 때부터 사용합니다.

func까지 작성하자마자 함수를 작성할 수 있게 해줍니다. 함수 이름은 다른 것도 당연히 작성할 수 있습니다.

```go
package main

func main() {

}
```

이것이 함수입니다. 매개변수도 다룰 것입니다. 함수를 정의만 해주면 됩니다. 프로그래밍을 실행할 때 정의한 함수를 실행해줍니다. 별도의 호출이 필요없습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("여기서는 함수를 다룹니다.")
	greeter()
}

func greeter() {
	fmt.Println("안녕하세요")
}
```

이렇게 함수를 main에 호출해야 실행이 가능합니다. 커맨드라인은 main을 호출하고 main은 greeter를 호출합니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("여기서는 함수를 다룹니다.")
	greeter()

	// 에러를 돌려줍니다.
	func greeterTwo()  {
		fmt.Println("함수속 에 함수 정의는 허용되지 않습니다.")
	}
}

func greeter() {
	fmt.Println("안녕하세요")
}
```

즉시 실행함수, 익명함수도 존재하지만 안다루도록 합니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("여기서는 함수를 다룹니다.")

	result := adder(3, 5)

	fmt.Println("간단한 결과", result)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

```

함수 시그니처가 필요합니다. 인자의 타입과 반환값의 타입을 모두 지정해줘야 합니다.

받을 인자의 개수를 모를 때 활용할 수 있는 문법이 있습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("여기서는 함수를 다룹니다.")

	proResult := proAdder(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("간단한 결과", proResult)
}


func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
```

이렇게 작성하면 sum 함수 비슷하게 구현할 수 있습니다.

몇가지 경우는 반환값이 여러가지인 경우도 존재할 수 있습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("여기서는 함수를 다룹니다.")

	proResult, myMessage := proAdder(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("간단한 결과", proResult, myMessage)
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "adfsdsaf"
}
```

이렇게 작성하면 여러 개의 반환값을 갖도록 만들 수 있습니다. 그리고 반환값에 대한 할당도 comma ok 표기법으로 대응해줘야 합니다.

람다(익명)함수도 분명 존재하지만 나중에 다룰 것입니다.

# Methods in golang

https://www.youtube.com/watch?v=GhYIKwMxz_Y

golang에 대한 메서드입니다.

golang은 클래스가 존재하지 않아 struct를 활용해서 메서드를 만듭니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("메서드를 다룹니다.")
	// 상속 부모자식 클래스 개념이 없습니다.
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() {

}
```

메서드는 이렇게 정의합니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("메서드를 다룹니다.")
	jake := User{"Jake", "jaketheGod@go.dev", true, 30, 23}
	jake.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int // 비공개 속성값은 케멀케이스로 정의합니다.
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
```

> 메서드를 다룹니다.
> Is user active: true

이런 피드백을 받습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("메서드를 다룹니다.")
	// 상속 부모자식 클래스 개념이 없습니다.
	jake := User{"Jake", "jaketheGod@go.dev", true, 30, 23}
	jake.GetStatus()
	jake.NewMail()
	fmt.Println(jake)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int // 비공개 속성값은 케멀케이스로 정의합니다.
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "jakethedog@go.dev"
	fmt.Println("Email: ", u.Email)
}
```

메서드의 피드백과 다르게 원본 인스턴스에는 변형이 없습니다.

> Email: jakethedog@go.dev
> {Jake jaketheGod@go.dev true 30 23}

메서드는 복사되어 할당됩니다. 대부분의 언어는 복사 됩니다. 이런 이유로 포인터를 활용하게 됩니다. 원본 값을 변형하는 예시는 나중에 다룹니다.

# Defer in golang

https://www.youtube.com/watch?v=jiy584-vv38

이번에는 Defer문을 다룹니다. 인터넷에서 혼선이 많이 발생하는 지식입니다. 문서도 혼란스럽습니다.

함수를 실행할 때 줄 순서대로 실행합니다. 하지만 defer를 선언하면 나중에 실행합니다.

LIFO 구조로 실행합니다.

https://go.dev/ref/spec#Defer_statements

공식 문서는 이렇게 설명합니다.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("hello")
}
```

> hello
> World

World를 나중에 실행합니다.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("hello")
}
```

> hello
> Two
> One
> World

LIFO 순서에 맞게 실행됩니다. 이점만 숙지하면 상당히 간단하게 사용할 수 있습니다. golang이 취미인 입장에서는 이정도로 파악하면 됩니다.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("hello")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
```

> hello
> 43210Two
> One
> World

이렇게 실행합니다. defer는 함수 속 반복문에 작성되어 있습니다. hello를 출력하고 defer가 없어서 제일 먼저 실행합니다. 다음으로 myDefer를 실행하고 반복문의 역순으로 출력합니다. 그리고 defer를 역순으로 실행합니다.

문서와 자료를 통해 학습하기 보단 직접 코드를 작성해서 이렇게 이해하는 것이 더 직관적입니다.

# Working with files in golang

https://www.youtube.com/watch?v=Mdg3tlGUXrE

다른 언어처럼 golang도 text파일만 처리할 수 있습니다. 다른 언어처럼 다른 형식도 패키지에 의존해야 합니다. IO랑 defer문법만 알면 간단하게 할 수 있습니다.

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("golang으로 파일처리")
	content := "파일을 읽어볼 것입니다."

	// 파일 생성
	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		// panic은 프로그램을 강제 중단합니다.
		panic(err)
	}

	// 생성한 파일에 글 작성하기
	length, err := io.WriteString(file, content)
	if err != nil {
		// 에러처리는 대부분의 경우 유사합니다.
		panic(err)
	}

	fmt.Println("length is: ", length)
	// 생성한 파일을 닫기합니다.
	file.Close()
}
```

> golang으로 파일처리
> length is: 33

터미널을 이렇게 피드백을 돌려줍니다.

```txt
파일을 읽어볼 것입니다.
```

이렇게 파일이 생성되었습니다.

```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("golang으로 파일처리")
	content := "파일을 읽어볼 것입니다."

	// 파일 생성
	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		// panic은 프로그램을 강제 중단합니다.
		panic(err)
	}

	// 생성한 파일에 글 작성하기
	length, err := io.WriteString(file, content)
	if err != nil {
		// 에러처리는 대부분의 경우 유사합니다.
		panic(err)
	}

	fmt.Println("length is: ", length)
	// 생성한 파일을 닫기합니다.
	defer file.Close()
	readFile("./mylcogofile.txt")
}

// 파일 읽기는 생성과 유사합니다.

func readFile(fileName string) {
	// 파일을 읽을 때는 운영체제 제어보다 더 많은 기능이 필요하고 ioutil이 제공합니다.
	dataByte, err := ioutil.ReadFile(fileName)
	// 파일을 읽을 때는 기본적으로 string으로 되어 있지 않습니다. byte로 되어 있습니다.
	if err != nil {
		panic(err)
	}

	fmt.Println("dataByte의 생김새는...", dataByte)
}
```

> golang으로 파일처리
> length is: 33
> dataByte의 생김새는... [237 140 140 236 157 188 236 157 132 32 236 157 189 236 150 180 235 179 188 32 234 178 131 236 158 133 235 139 136 235 139 164 46]

바이트는 이렇게 생겼습니다.

에러핸들링은 자주하면서 `checkNilErr`같은 함수를 만들 때가 많습니다.

```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("golang으로 파일처리")
	content := "파일을 읽어볼 것입니다."

	// 파일 생성
	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)

	// 생성한 파일에 글 작성하기
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", length)
	// 생성한 파일을 닫기합니다.
	defer file.Close()
	readFile("./mylcogofile.txt")
}

// 파일 읽기는 생성과 유사합니다.

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("dataByte의 생김새는...", dataByte)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

```

이렇게 작성하면 에러를 간략하게 핸들링할 수 있습니다.

# Handling web request in golang

https://www.youtube.com/watch?v=ru53LpdVHn4

이번장에는 web을 다루기 시작합니다. 재미있는 부분입니다. 지금까지 배운 것을 응용하는 성격이 강합니다.

http 패키지로 웹 요청을 주고 받을 수 있습니다. net/http를 가장 자주 사용하고 가장 빠르게 사용할 수 있는 패키지입니다.

https://pkg.go.dev/net/http

header도 제어하는 것은 당연히 가능합니다.

요청을 보낼 때 response의 타입이 아주 중요합니다. 깊게 공부하려면 필요한 지식입니다.

https://pkg.go.dev/net/http#Request

요청은 종료되면 반드시 닫도록 합니다.

```go
	// For client requests, setting this field prevents re-use of
	// TCP connections between requests to the same hosts, as if
	// Transport.DisableKeepAlives were set.
	Close bool
```

이부분을 주의하도록 합니다.

```go
package main

import (
	"fmt"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web에 요청합니다.")
	// get 요청을 날립니다.
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response 타입: %T\n", response)
	// 요청이 종료되면 닫도록 합니다.
	response.Body.Close()
}
```

> LCO web에 요청합니다.
> response 타입: \*http.Response

response의 포인터를 활용합니다.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web에 요청합니다.")
	// get 요청을 날립니다.
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response 타입: %T\n", response)

	// 요청이 종료되면 닫도록 합니다.
	defer response.Body.Close()

	dateBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dateBytes)
	fmt.Println(content)
}
```

html을 응답으로 받게 됩니다.

이것은 가장 기본적인 웹 요청입니다. 대부분 golang을 활용할 때는 API 핸들링입니다.

# Handling URL in golang

https://www.youtube.com/watch?v=cl7_ouTMFh0

웹을 다룰 때는 URL이 안 다룰 수 없습니다. 지난 시간에는 URL로 요청을 보내는 법을 배웠습니다.

net은 라이브러리를 넘어 모듈에 가깝습니다.

```go
package main

import "fmt"

// 가짜 url을 만들어봅니다.

const MY_URL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid:afsdqwerzxv3645"

func main() {
	fmt.Println("URL 다루기에 환영합니다.")
	fmt.Println(MY_URL)
}
```

여기서 시작합니다. 이제는 파싱합니다. 정보를 처리합니다.

```go
package main

import (
	"fmt"
	"net/url"
)

// 가짜 url을 만들어봅니다.

const MY_URL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid:afsdqwerzxv3645"

func main() {
	fmt.Println("URL 다루기에 환영합니다.")
	fmt.Println(MY_URL)

	// 파싱 처리
	result, _ := url.Parse(MY_URL)
	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // lco.dev:3000
	fmt.Println(result.Path)     // /learn
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid:afsdqwerzxv3645
	fmt.Println(result.Port())   // 3000
}
```

각각의 요소를 선택하고 계산에 나중에 응용해 볼 수 있을 것 같습니다.

여기서 변수명으로 저장하려는 본능이 나오겠지만 더 우아한 처리방법이 있습니다.

```go
package main

import (
	"fmt"
	"net/url"
)

// 가짜 url을 만들어봅니다.

const MY_URL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid:afsdqwerzxv3645"

func main() {
	fmt.Println("URL 다루기에 환영합니다.")
	fmt.Println(MY_URL)

	// 파싱 처리
	result, _ := url.Parse(MY_URL)
	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // lco.dev:3000
	fmt.Println(result.Path)     // /learn
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid:afsdqwerzxv3645
	fmt.Println(result.Port())   // 3000

	qparams := result.Query()
	fmt.Printf("Query Param: %T\n", qparams) // Query Param: url.Values

	fmt.Println(qparams["coursename"]) // [reactjs]

	for _, val := range qparams {
		fmt.Println("Param is", val)
	}

	// 포인터를 전달합니다.
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=jake",
	}

	// 문자열로 변환하고 할당합니다.
	auotherURL := partsOfUrl.String()
	fmt.Println(auotherURL)
}
```

> https://lco.dev/tutcss

이렇게 URL을 수동으로 생성하는 것도 가능합니다.

# Creating server for golang frontend

https://www.youtube.com/watch?v=xh79JXJy0yY

golang으로 요청을 작은 체험만 해봤습니다. golang으로 다양한 라우팅 처리하고 핸들링하는 것을 배워야 합니다.

https://github.com/hiteshchoudhary/golang/tree/main/lcowebserver

여기서 서버를 만들고 임시로 만들 수 있습니다.

```js
/*
Part of exercise file for go lang course at
https://web.learncodeonline.in
*/

const express = require("express");
const app = express();
const port = 8000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.status(200).send("Welcome to LearnCodeonline server");
});

app.get("/get", (req, res) => {
  res.status(200).json({ message: "Hello from learnCodeonline.in" });
});

app.post("/post", (req, res) => {
  let myJson = req.body; // your JSON

  res.status(200).send(myJson);
});

app.post("/postform", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
```

이렇게 임시 서버를 만들 수 있습니다.

요청 응답을 처리할 수 있게 golang으로 만들어야 합니다. json데이터를 주고 받게 만드는 것입니다.

요청과 응답을 주고 받게 만드는 것은 직관적으로 할 수 있습니다.

# How to make GET request in golang

https://www.youtube.com/watch?v=V-sxFQ0fWlw

목표는 단순합니다. 자바스크립트 익스레스 서버처럼 만드는 것입니다.

서버의 get요청 처리입니다. post를 json으로 처리하고 또 form을 처리할 것입니다.

문자열 처리를 어렵게 처리하는 방법과 쉬운 방법입니다.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("요청처리 1장입니다.")

	PerformGetRequest()
}

// public 함수로 만듭니다.
func PerformGetRequest() {
	const MY_URL = "http://localhost:8000/get"

	response, err := http.Get(MY_URL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("상태 코드: ", response.StatusCode)
	fmt.Println("콘텐츠 길이: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
```

> 요청처리 1장입니다.
> 상태 코드: 200
> 콘텐츠 길이: 43
> {"message":"Hello from learnCodeonline.in"}

이렇게 응답을 받을 수 있습니다. 데이터의 JSON을 아직 안 다루었습니다. 나중에 다룰 것입니다.

이것을 다른 방식으로 해결하는 것도 가능합니다. 복잡하지만 가치가 있습니다.

문자열을 처리하는 방식입니다.

https://pkg.go.dev/strings#Builder

strings 패키지에서 처리를 도와줄 것입니다.

메서드를 활용해서 제어하게 되는 것입니다. 지금의 방식에는 문제는 없지만 각자 다른 방식으로 작성하면서 문제가 생길 수 있습니다.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("요청처리 1장입니다.")

	PerformGetRequest()
}

// public 함수로 만듭니다.
func PerformGetRequest() {
	const MY_URL = "http://localhost:8000/get"

	response, err := http.Get(MY_URL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("상태 코드: ", response.StatusCode)
	fmt.Println("콘텐츠 길이: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("byteCount: ", byteCount)
	fmt.Println("byteCount: ", responseString.String()) // 장점은 원본 데이터를 달고 있습니다.

}
```

라이브러리 하나만 사용하는 것으로 이렇게 강력하게 사용할 수 있습니다. 초심자들에게는 전자를 권장합니다. 하지만 어느정도 사용해보면 당연히 후자가 더 좋습니다.

# How to make POST request with JSON data in golang

https://www.youtube.com/watch?v=h5NeKZuzUoc

웹서버는 그 라우트의 데이터만 받습니다.

데이터를 서버로 보내야 하는 경우를 처리합니다. 이번에는 JSON을 생성하고 만드는 법을 다룹니다.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("요청처리 1장입니다.")

	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const MY_URL = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"courseName": "Let's go with golang",
			"price": "0",
			"platform": "learnCodeOnline.in"
		}
	`)

	// 두번째 인자 header의 content-type을 주의하도록 합니다.
	response, err := http.Post(MY_URL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
```

주의할 점은 header로 넣게 되는 content-type입니다.

> 요청처리 1장입니다.
> {"courseName":"Let's go with golang","price":"0","platform":"learnCodeOnline.in"}

post 요청이 처리되었습니다.

# How to send form data in golang

https://www.youtube.com/watch?v=U_LjyX4iDbU

form 데이터처럼 이미지 업로처럼 필요한 경우들이 존재합니다. 기타 변수추가하고 다양한 응용은 나중에 다룹니다.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("요청처리 1장입니다.")

	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const MY_URL = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstName", "jake")
	data.Add("lastName", "the dog")
	data.Add("email", "jakethegod@go.dev")

	response, err := http.PostForm(MY_URL, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
```

> {"email":"jakethegod@go.dev","firstName":"jake","lastName":"the dog"}

form 응답은 그렇게 다르게 생기지 않았습니다.

# How to create JSON data in golang

https://www.youtube.com/watch?v=SZ5xZ9OTeEI

이번에는 JSON을 다룹니다. JSON을 응답으로 받았지만 문자열로만 다루었습니다. JSON을 생성하고 처리하는 것을 아는 것은 상당히 중요합니다.

struct도 다시 다루게 될 것입니다.

```go
package main

import "fmt"

type course struct {
	Name     string
	Price    int
	Platform string
	password string
	Tags     []string
}

func main() {
	fmt.Println("JSON입니다.")
}
```

간편하게 struct를 생성합니다.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	password string
	Tags     []string
}

func main() {
	fmt.Println("JSON입니다.")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS BootCamp", 299, "LearnCodeOnline.in", "asdf", []string{"web-dev", "js"}},
		{"MERN BootCamp", 199, "LearnCodeOnline.in", "qwer", []string{"full-stack", "js"}},
		{"Vue BootCamp", 299, "LearnCodeOnline.in", "zcvx", nil},
	}

	// json데이터로 패키징합니다.
	// 첫번째 인자는 인터페이스입니다. struct의 파생된 단어입니다.
	finalJson, err := json.Marshal(lcoCourses)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}
```

> [{"Name":"ReactJS BootCamp","Price":299,"Platform":"LearnCodeOnline.in","Tags":["web-dev","js"]},{"Name":"MERN BootCamp","Price":199,"Platform":"LearnCodeOnline.in","Tags":["full-stack","js"]},{"Name":"Vue BootCamp","Price":299,"Platform":"LearnCodeOnline.in","Tags":null}]

데이터를 읽을 수 있습니다. `nil`은 `null`로 변환이 발생합니다.

`finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")`이렇게 설정하면 들여쓰기가 되서 가독성을 높일 수 있습니다. 하지만 일부 어퍼케이스로 변환이 발생합니다. 대부부의 경우 소문자로 시작하는 것을 권장합니다.

```txt
[
        {
                "Name": "ReactJS BootCamp",
                "Price": 299,
                "Platform": "LearnCodeOnline.in",
                "Tags": [
                        "web-dev",
                        "js"
                ]
        },
        {
                "Name": "MERN BootCamp",
                "Price": 199,
                "Platform": "LearnCodeOnline.in",
                "Tags": [
                        "full-stack",
                        "js"
                ]
        },
        {
                "Name": "Vue BootCamp",
                "Price": 299,
                "Platform": "LearnCodeOnline.in",
                "Tags": null
        }
]
```

struct 엘리어스를 설정할 수 있지만 업데이트가 되었습니다. 이부분은 생략합니다.

# How to consume JSON data in golang

https://www.youtube.com/watch?v=a96veXdifys

이번에는 JSON 데이터의 소비자가 됩니다. JSON을 디코딩합니다.

문자열 형태가 아닌 일반적인 JSON으로 다루게 됩니다.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json: "coursename"`
	Price    int
	Platform string   `json: "website"`
	password string   `json: "-"` // 제거
	Tags     []string `json: "tags, omitempty"`
}

func main() {
	fmt.Println("JSON입니다.")
	// EncodeJson()
	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"Name": "ReactJS BootCamp",
		"Price": 299,
		"Platform": "LearnCodeOnline.in",
		"Tags": [
						"web-dev",
						"js"
		]
}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON은 검증되었습니다.")
		// 두번째 인자에 참조할 주소를 전달합니다.
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON은 유효하지 않습니다.")
	}
}
```

하지만 자주 사용하는 유스케이스들이 있습니다. 키와 값으로 자료를 추가하고 싶을 때가 있습니다.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json: "coursename"`
	Price    int
	Platform string   `json: "website"`
	password string   `json: "-"` // 제거
	Tags     []string `json: "tags, omitempty"`
}

func main() {
	fmt.Println("JSON입니다.")
	// EncodeJson()
	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"Name": "ReactJS BootCamp",
		"Price": 299,
		"Platform": "LearnCodeOnline.in",
		"Tags": [
						"web-dev",
						"js"
		]
}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON은 검증되었습니다.")
		// 두번째 인자에 참조할 주소를 전달합니다.
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON은 유효하지 않습니다.")
	}

	// 키와 값같은 해쉬테이블 자료형 응용
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
```

> Key is Name and value is ReactJS BootCamp and type is: string
> Key is Price and value is 299 and type is: float64
> Key is Platform and value is LearnCodeOnline.in and type is: string
> Key is Tags and value is [web-dev js] and type is: []interface {}

이렇게 터미널에 출력됩니다.

바이트, struct, 해쉬테이블 등으로 응용하는 법을 배웠습니다.

# A long video on MOD in golang

https://www.youtube.com/watch?v=O8uUGEobo-Q

이번에 다룰 주제는 조금 깁니다. go mod입니다. `go mod init 이름` 이 명령으로 생성되는 mod 파일을 이해해봅니다.

이것은 툴링입니다. go build 프로덕션 준비된 수준으로 만들어주는 것처럼 go mod도 golang을 다양하게 활용할 수 있게 해줍니다.

모든 것을 다룰 수 없어도 개괄적으로 다룰 수 있습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello mod in golang")
}
```

mod 선언 없이도 동작은 가능하지만 권장하지 않습니다.

go mod 는 github 호스팅 이름을 권장합니다.

```sh
go mod init github.com/arch-spatula/리포이름
```

이런 방식으로 선언하는 것을 권장합니다.

> go: creating new go.mod: module github.com/arch-spatula/golang-test
> go: to add module requirements and sums: go mod tidy

그러면 이렇게 피드백을 줍니다.

`go mod tidy`은 나중에 다룹니다.

```go
module github.com/arch-spatula/golang-test

go 1.20
```

파일을 열면 이렇게 작성되어 있습니다. 버전과 이런저런 정보가 보입니다. 버전관리는 시멘틱 버저닝을 이해해야 합니다.

메이저업데이트.마이너업데이트.패치

이렇게 되어있습니다. golang은 다양한 의존성이 있지만 다 가려져있습니다. 이렇게 된 것은 비교적 최근입니다.

go modules를 활용하면서 바뀐 부분들입니다. 2019년부터 golang를 사용하면서 시작된 것입니다.

https://go.dev/ref/mod

여기서 자세한 사항이 있습니다.

go tool chan으로 활용하는 것도 가능합니다.

```sh
go get -u github.com/gorilla/mux
```

이 명령으로 원격리포에서 관련 파일을 다운 받을 수 있습니다.

```txt
module github.com/arch-spatula/golang-test

go 1.20

require github.com/gorilla/mux v1.8.0 // indirect
```

go.sum 파일도 생성됩니다. github 리포를 확인하고 리포의 변화를 확인합니다.

GOPATH="/본인디렉토리/go"에서 의존성을 관리하고 제어합니다리

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
}

func greeter() {
	fmt.Println("mod 사용자")
}

// 백엔드의 흔한 문법입니다.
func serveHome(w http.ResponseWriter, r *http.Request) {
	// r은 요청이고 w는 응답입니다.
	w.Write([]byte("<h1>hello golang</h1>"))
}
```

이렇게 됩니다. serveHome은 받아온 라이브러에서 활용한 패턴입니다.

https://github.com/gorilla/mux에서 받아온 것을 선언하는 것으로 자동적으로 가져옵니다.

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
}

func greeter() {
	fmt.Println("mod 사용자")
}

// 백엔드의 흔한 문법입니다.
func serveHome(w http.ResponseWriter, r *http.Request) {
	// r은 요청이고 w는 응답입니다.
	w.Write([]byte("<h1>hello golang</h1>"))
}
```

이렇게 작성할 수 있습니다.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// log.Fatal로 에러처리
	log.Fatal(http.ListenAndServe(":3000", r))
}

func greeter() {
	fmt.Println("mod 사용자")
}

// 백엔드의 흔한 문법입니다.
func serveHome(w http.ResponseWriter, r *http.Request) {
	// r은 요청이고 w는 응답입니다.
	w.Write([]byte("<h1>hello golang</h1>"))
}
```

```sh
go build .
```

```sh
go run main.go
```

http://localhost:3000/

이렇게 되면 서버가 동작하는 것을 확인할 수 있습니다.

go mod는 계산 비용이 비쌉니다. 그래서 한번 사용할 때마다 주의합니다.

```sh
go mod tidy
```

```sh
go mod verify
```

이 명령을 실행하면 모든 모듈을 검증합니다. 이렇게 되면 사용가능 여부를 확인합니다.

go list는 의존 중인 모듈을 표시하지만 부정확합니다.

```sh
go list -m all
```

이 프로젝트 중에 의존 중인 모듈을 확인하기 가장 효율적입니다.

```sh
go list -m -versions github.com/gorilla/mux
```

이 명령으로 이용가능한 버전들을 확인할 수 있습니다.

go mod tidy는 사용안 하는 패키지를 삭제하고 삭제했던 패키지를 다시 설치합니다.

```sh
go mod why github.com/gorilla/mux
```

이 명령은 의존 중인 모듈을 알려줍니다.

```sh
go mod graph
```

go.mod파일을 편집하는 법입니다.

```sh
go mod edit -go 1.16 # golang 버전 변경
go mod edit -module 1.16 # module 버전 변경
```

버전 변경

```sh
go mod vendor
```

node_module처럼 의존하는 파일을 다운 받습니다. 하지만 vender 폴더를 활용해야 합니다.

```sh
go run -mod=vender main.go
```

그래서 `go mod vender` 명령을 하고난 뒤부터는 이렇게 명령해야 합니다.

구형은 업데이트하고 모듈을 활용하기를 권장합니다.

# Building API in golang - Models

https://www.youtube.com/watch?v=TrouRtr5xWE

이제는 완전한 API를 만들 수 있습니다. 이번에는 완전한 API를 만들지만 DB 통신은 없습니다.

이번에는 하나의 폴더로 이해하기 위해 사용할 것이고 다음에 관리하게 좋게 나누는 법을 다룹니다.

```sh
go mod init github.com/arch-spatula/golang-test
go get -u github.com/gorilla/mux
touch main.go
```

이렇게 환경을 생성합니다.

강의 판매 플랫폼을 만든다고 가정합니다.

```go
package main

// Model for course - file

type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrise int    `json:"price"`
	// 포인터로 타입을 지정하기
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {

}
```

이렇게 모델을 작성합니다.

```go
package main

// Model for course - file

type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrise int    `json:"price"`
	// 포인터로 타입을 지정하기
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// 모의 DB
var course []Course

// 미들웨어 헬퍼 -file
func (c *Course) IsEmpty() bool {
	// id와 강의명이 비어있는지 검증합니다.
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}
```

# Sending a API json response for all courses in golang

https://www.youtube.com/watch?v=V3S3hKrIdQo

이번에는 컨트롤러를 다룹니다. `r.HandleFunc("/", serveHome).Methods("GET")`이런 코드를 다루었습니다.

```go
// 백엔드의 흔한 문법입니다.
func serveHome(w http.ResponseWriter, r *http.Request) {
	// r은 요청이고 w는 응답입니다.
	w.Write([]byte("<h1>hello golang</h1>"))
}
```

이렇게 라우트가 정의되고 이 라우트에 따라 핸들링하는 로직을 작성합니다. 상황을 핸들링하는 컨트롤러를 먼저 정의하고 다음에 다음에 라우트를 정의합니다. 컨트롤러도 컨트롤러마다 다른 폴더로 정리합니다.

seeding은 임시 데이터를 넣고 테스트할 때 자주 하는 행위입니다. 지금은 seeding하기 전입니다.

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for course - file

type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrise int    `json:"price"`
	// 포인터로 타입을 지정하기
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// 모의 DB
var course []Course

// 미들웨어 헬퍼 -file
func (c *Course) IsEmpty() bool {
	// id와 강의명이 비어있는지 검증합니다.
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// controllers

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	// 순서와 두번째 인자는 포인터라는 것을 잘 외우도록 합니다.
	w.Write([]byte("<h1>API 생성</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("모든 강좌를 가져옵니다.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}
```

# Get one course based on request id in golang

https://www.youtube.com/watch?v=1J1sgKriM-o

이번에는 하나의 요소만 가져오는 법을 배웁니다.

```go

// Course라고만 명명하는 것이 좋은 컨벤션입니다. 요청 유형을 굳이 작성하지 않습니다.
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	// 선형탐색으로 존재하는 강좌를 찾아냅니다.
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	// 요청으로
	params := mux.Vars(r)
}

```

강의 진행을 위해서 getOneCourse라고 명명할 뿐입니다. 현업에서는 더 단순하게 명명합니다.

```go
// Course라고만 명명하는 것이 좋은 컨벤션입니다. 요청 유형을 굳이 작성하지 않습니다.
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	// 선형탐색으로 존재하는 강좌를 찾아냅니다.
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	// 요청으로
	params := mux.Vars(r)

	// 응답을 강의로 돌려줍니다.
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("강의가 없습니다.")
}
```

# Add a course controller in golang

https://www.youtube.com/watch?v=vMpBUkIMLzY

DB에 추가될 때 체커 같은 다양한 로직들을 추가해줘야 하지만 단순하게만 다루겠습니다.

```go
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// 유효성 검증을 진행합니다. 비어있을 때를 대응합니다.
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please and some data")
	}

	// {}이런 형식의 요청을 대응합니다.
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("JSON에 데이터가 없습니다.")
		return
	}
	// uid와 문자열을 생성합니다. 데이터를 뒤에 붙입니다.

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) // 정수를 문자열로 변환합니다.
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}
```

# Update a course controller in golang

https://www.youtube.com/watch?v=TQpB_QSgmXU

이번에는 update 연산을 다룹니다.

```go

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

```

이 로직을 재사용합니다.

```go
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one Course")
	w.Header().Set("Content-Type", "application/json")

	// first - garb id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		// 업데이트할 row를 찾은 경우입니다.
		if course.CourseId == params["id"] {
			// 삭제 연산
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			// id를 선택하고 덮어씁니다.
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	// TODO: 발견하지 못한 경우를 처리합니다.
}
```

업데이트 로직을 이렇게 처리합니다.

# Delete a course controller in golang

https://www.youtube.com/watch?v=b1nJBaCzkHI

업데이트에서 삭제를 다루었기 때문에 간단하게 만들어볼 것입니다.

```go
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one Course")
	w.Header().Set("Content-Type", "application/json")

	// first - garb id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		// 업데이트할 row를 찾은 경우입니다.
		if course.CourseId == params["id"] {
			// 삭제 연산
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			break
		}
	}
}
```

# Handling routes and testing routes in golang

https://www.youtube.com/watch?v=NedcxRgo4DY

데이터 seeding과 testing입니다.

편의를 위해 이렇게 복사합니다.

```go
courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file

type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"price"`
	// 포인터로 타입을 지정하기
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// 모의 DB
var courses []Course

// 미들웨어 헬퍼 -file
func (c *Course) IsEmpty() bool {
	// id와 강의명이 비어있는지 검증합니다.
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":3000", r))
}

// controllers

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	// 순서와 두번째 인자는 포인터라는 것을 잘 외우도록 합니다.
	w.Write([]byte("<h1>API 생성</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("모든 강좌를 가져옵니다.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Course라고만 명명하는 것이 좋은 컨벤션입니다. 요청 유형을 굳이 작성하지 않습니다.
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	// 선형탐색으로 존재하는 강좌를 찾아냅니다.
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	// 요청으로
	params := mux.Vars(r)

	// 응답을 강의로 돌려줍니다.
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("강의가 없습니다.")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// 유효성 검증을 진행합니다. 비어있을 때를 대응합니다.
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please and some data")
	}

	// {}이런 형식의 요청을 대응합니다.
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("JSON에 데이터가 없습니다.")
		return
	}
	// uid와 문자열을 생성합니다. 데이터를 뒤에 붙입니다.

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) // 정수를 문자열로 변환합니다.
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one Course")
	w.Header().Set("Content-Type", "application/json")

	// first - garb id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		// 업데이트할 row를 찾은 경우입니다.
		if course.CourseId == params["id"] {
			// 삭제 연산
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			// id를 선택하고 덮어씁니다.
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	// TODO: 발견하지 못한 경우를 처리합니다.
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one Course")
	w.Header().Set("Content-Type", "application/json")

	// first - garb id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		// 업데이트할 row를 찾은 경우입니다.
		if course.CourseId == params["id"] {
			// 삭제 연산
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			break
		}
	}
}
```

```url
http://localhost:3000/course/2
```

이렇게 요청의 응답이 동작합니다.

```json
{
  "coursename": "Java Spring",
  "price": 199,
  "Author": {
    "fullname": "Jake the Backend",
    "website": "jake.com"
  }
}
```

이렇게 post 요청을 보내고

```json
[
  {
    "courseid": "2",
    "coursename": "ReactJS",
    "author": {
      "fullname": "Hitesh Choudhary",
      "website": "lco.dev"
    }
  },
  {
    "courseid": "4",
    "coursename": "MERN Stack",
    "author": {
      "fullname": "Hitesh Choudhary",
      "website": "go.dev"
    }
  },
  {
    "courseid": "88",
    "coursename": "Java Spring",
    "author": {
      "fullname": "Jake the Backend",
      "website": "jake.com"
    }
  }
]
```

전체 응답을 받은 것을 확인할 수 있습니다.

이정도도 꽤 쉽지만 ORM을 활용해서 더 쉽게 다룰 수 있게 해줍니다.

# MongoDB setup for API in golang

https://www.youtube.com/watch?v=EaZo7pcueBE

이번장부터는 데이터베이스를 다루기 시작합니다. CRUD API를 만들 수 있고 이제는 Model을 제어하는 법을 활용합니다.

데이터 베이스랑 상호작용하는 법을 배우고 ORM을 다루는 법도 다룹니다. ORM을 활용하면 물론 어플리케이션이 느려지기는 합니다.

`gorilla/mux`는 계속 활용합니다.

```sh
go get -u github.com/gorilla/mux
```

MongoDB는 아틀라스를 사용합니다.

각각 관리하는 방법이 다릅니다.

https://github.com/mongodb/mongo-go-driver

```sh
go get go.mongodb.org/mongo-driver/mongo
```

이제는 인스턴스

root에는 항상 go확장자 파일 1개만 있어야 합니다.

# Defining models for netflix in golang

https://www.youtube.com/watch?v=laeZI6UdDNg

```sh
mkdir model
mkdir controller
mkdir router
```

이렇게 폴더를 만들어줍니다.

# Making a connection to database in golang

https://www.youtube.com/watch?v=GszGvj6eBZY

```go
package controller

import "go.mongodb.org/mongo-driver/mongo"

// 이부분은 커밋에 올라가면 안 됩니다. 환경변수 설정해야 합니다.
const connectionString = "mongodb+srv://사용자명:비번@cluster0.crqb9un.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const colName = "watchList"

// 이부분이 제일 중요합니다.
var collection *mongo.Collection
```

https://pkg.go.dev/context

```go
package controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 이부분은 커밋에 올라가면 안 됩니다.
const connectionString = "mongodb+srv://사용자명:비번@cluster0.crqb9un.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const colName = "watchList"

// 이부분이 제일 중요합니다.
var collection *mongo.Collection

// mongoDB랑 연결하기
// 함수명을 init이라고 하면 첫 실행에만 동작합니다.
func init() {
	// 클라이언트 옵션입니다.
	clientOptions := options.Client().ApplyURI(connectionString)

	mongo.Connect(context.TODO(), clientOptions)
}
```

context.TODO는 특이합니다. 스스로 공부하기 권장합니다.

```go
package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 이부분은 커밋에 올라가면 안 됩니다.
const connectionString = "mongodb+srv://사용자명:비번@cluster0.crqb9un.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const colName = "watchList"

// 이부분이 제일 중요합니다.
var collection *mongo.Collection

// mongoDB랑 연결하기
// 함수명을 init이라고 하면 첫 실행에만 동작합니다.
func init() {
	// 클라이언트 옵션입니다.
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongoDB 연결성공")

	collection = client.Database(dbName).Collection(colName)

	// collection 래퍼런스이지만 인스턴스가 준비되었으면
	fmt.Println("컬랙션이 준비되었습니다.")
}
```

# Insert data in mongodb in golang

https://www.youtube.com/watch?v=K2nC_WpPVQE

이제는 VScode와 다른 파일을 다루는 법을 배웁니다.

혹시 계속 경고가 나오면 `go get go.mongodb.org/mongo-driver/mongo`을 다시 터미널에 명령합니다.

go tools를 VScode에 설정하도록 합니다.

```go
// mongodb helper file

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	// 무엇이 삽입되었는지 확인합니다.
	fmt.Println("DB에 자료 1개 삽입", inserted.InsertedID)

}
```

이렇게 DB에 추가하는 법입니다.

# Update a record in mongodb in golang

https://www.youtube.com/watch?v=00YTOQTxUEc

업데이트는 상당히 단순합니다. 첫째는 필터입니다. 다음은 플레그를 설정합니다. 무엇을 set할지 정하는 것입니다.

```go
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}
```

golang에서는 주로 이렇게 작성합니다.

# Delete one and delete many in mongodb in golang

https://www.youtube.com/watch?v=qeZueI6d1Wg

mongoDB는 언어마다 모두 비슷합니다. 필터를 설정하고 삭제 요청하면 됩니다.

하나 삭제와 전체 삭제 모두 다룹니다.

```go
func deleteOenMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted count: ", deleteCount)
}
```

여러개 row를 삭제하는 법입니다.

```go
func deleteAllMovie() int64 {
	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("삭제된 영화개수: ", deleteCount.DeletedCount)
	return deleteCount.DeletedCount
}
```

# Get all collection in mongodb in golang

https://www.youtube.com/watch?v=1P5b4McGJQQ

```go
func getAllMovie() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var movies []primitive.M
  // mongoDB가 제공하는 객체를 순회합니다.
	for cursor.Next(context.Background()) {
		var movie bson.M

		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

```

# Get all movies from DB in golang

https://www.youtube.com/watch?v=Ycc9sXgDCNQ

```go
// 실제 controller
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	allMovies := getAllMovie()
	json.NewEncoder(w).Encode(allMovies)
}
```

이렇게 getAll에 대한 controller를 작성합니다.

# Mark movie as watched in golang

https://www.youtube.com/watch?v=XdYpKfIgIQg

```go
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Method", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
```

# Delete 1 and all movie in golang

https://www.youtube.com/watch?v=Qh1H-6Ebz7E

메서드 2개만 더 작성하면 라우팅을 작성해도 됩니다.

```go
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Method", "Delete")

	params := mux.Vars(r)
	deleteOenMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-control-Allow-Method", "Delete")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
```

다음에는 라우팅을 작성하고 테스트를 준비합니다.

# Creating routes and testing API in golang

https://www.youtube.com/watch?v=uG2fpX2ktNE

이번에는 라우팅을 시험합니다.

# Concurrency and goroutines in golang

https://www.youtube.com/watch?v=V-0ifUKCkBI

여기서부터는 golang을 다루는 방식이 다릅니다. golang 언어의 특수성을 다룹니다. 이전까지는 주 활용 분야인 백엔드 적용이었습니다.

사고방식이 조금 다릅니다. 동시성과 평행성입니다.

예를 들어 식사 중에 알림이 옵니다. 하지만 또 덮습니다. 동시성은 하나의 한 작업만 하는 경우와 잠시 작업을 보류하고 다시 착수하는 경우가 있습니다. 이것은 동시성입니다. 여러개의 과제를 수행하지만 하나의 과제만 들고 있는 것입니다.

아마 밥먹다가 알림을 확인하고 다시 밥먹다가 에어컨을 킬 수 있습니다.

평행성은 여러명의 사람이 동시에 있으면 누구는 에어컨을 키고 알림을 볼 사람은 알림만 봅니다.

쓰레드는 운영체제가 담당하지만 goroutines는 go runtime이 담당합니다. go runtime은 운영체제와 무관하게 더 많은 처리할 수 있습니다. 클라우드에서는 쓰레드 제한이 없어서 많이 끌어오는 것이 가능합니다. 이런 이유로 클라우드 엔지니어들이 많이 활용합니다.

운영체제의 쓰레드는 1MB입니다. go runtime의 쓰레드(?)는 2KB 정도 확보할 수 있습니다.

golang에서 자주 인용되는 문구입니다.

> Do not communicate by sharing memory; instead, share memory by commuicating

이 문구는 정말 자주 보게 될 것인 원칙입니다.

```go
package main

import "fmt"

func main() {
	greeter("Hello")
	greeter("world")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}
```

Hello
Hello
Hello
Hello
Hello
Hello
world
world
world
world
world
world

순서대로 실행됩니다.

goroutien을 실행하려면 go 키워드를 앞에 붙이면 됩니다.

```go
package main

import "fmt"

func main() {
	go greeter("Hello")
	go greeter("world")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}
```

하지만 이렇게 실행하면 출력이 없습니다.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	greeter("world")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(5 * time.Millisecond)
		fmt.Println(s)
	}
}
```

world
Hello
Hello
world
world
Hello
Hello
world
world
Hello
Hello
world

이렇게 출력됩니다. 하지만 권장하지 않습니다.

https://pkg.go.dev/sync

이 패키지를 활용할 것을 권장합니다.

# Wait groups in golang

https://www.youtube.com/watch?v=FiTbXvc08wE

복수의 서버에서 자주 사용하게 될 패턴입니다. 여러 마이크로 서비스에 활용하게 될 것입니다.

3가지 각각 엔드포인트에 응용하는 것도 가능합니다.

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
	}
	for _, web := range websitelist {
		getStatusCode(web)
	}
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Println("%d 200 status code for %s ", res.StatusCode, endpoint)
}
```

> %d 200 status code for %s 200 https://lco.dev
> %d 200 status code for %s 200 https://go.dev
> %d 200 status code for %s 200 https://google.com

이런 응답을 받습니다.

하지만 순서대로 처리합니다. 비동기적으로 동시에 처리하면 훨씬더 효율적이라는 것을 알 수 있습니다.

https://pkg.go.dev/sync#WaitGroup.Add

이 부분을 참고하도록 합니다. 참고로 지금 코드 예시는 적절하지 않습니다. 나중에 포인터로 올바르게 다루는 법을 배웁니다.

```go
package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		// 비동기 작업을 시작합니다.
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(endpoint string) {
	// 처리가 완료된 것을 마지막에 처리합니다.
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Println("%d 200 status code for %s ", res.StatusCode, endpoint)
}
```

동시에 처리가 되었습니다.

# Mutex in golang

https://www.youtube.com/watch?v=v7DKtAluvYA

이제는 가끔은 특별한 요구사항을 대응합니다.

```go
package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		// 비동기 작업을 시작합니다.
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	// 처리가 완료된 것을 마지막에 처리합니다.
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		signals = append(signals, endpoint)
		fmt.Println("%d 200 status code for %s ", res.StatusCode, endpoint)
	}
}
```

지금은 다르게 동작합니다. goroutines에서는 go runtime에서 제어합니다. 하지만 메모리 lock을 제어하는 것은 운영체제입니다. 하나의 메모리에 동시에 작성하려고 하면 지금 발생하는 문제입니다.

이럴 때 사용하는 것이 Mutex입니다. 메모리를 잠궈둡니다. 실행이 끝나기 전까지 다른 덮어쓰기를 막는 기능입니다.

```go
package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

// 모두 각각 포인터로 제어해야 합니다.
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		// 비동기 작업을 시작합니다.
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	// 처리가 완료된 것을 마지막에 처리합니다.
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Println("%d 200 status code for %s ", res.StatusCode, endpoint)
	}
}
```

지금같은 케이스에서는 엄청난 성능개선은 없습니다. 하지만 규모가 커지면 바로 볼 수 있습니다.

# Race Condition in golang

https://www.youtube.com/watch?v=4TCNUa-5wmI

이번에는 레이즈 컨디션을 다룹니다. golang의 tooling으로 제어합니다.

현실 세계에서 레이즈 컨디션을 처리하는 법을 배웁니다. 공식 문서를 보면 오히려 혼란 스럽습니다. 복수의 쓰레드가 존재합니다. 하나의 쓰레드로 같은 메모리를 작성하려고 하면 생기는 문제입니다.

언젠가 무조건은 필요한 개념입니다.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("레이즈 컨디션")

	wg := &sync.WaitGroup{}

	var score = []int{0}
	// 즉시 실행함수입니다.
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		fmt.Println("One")
		score = append(score, 1)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Two")
		score = append(score, 2)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Three")
		score = append(score, 3)
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println(score)
}
```

> 레이즈 컨디션
> Three
> One
> Two
> [0 3 1 2]

상호 독립적인 루틴이 동작하는 순서는 다릅니다. 실행할 때마다 순서가 바뀔 수 있습니다.

정상동작하는 것처럼 보이기는 합니다. 하지만 여기서 실수할 가능성이 높습니다.

```sh
go run --race .
```

실행하면 발생하는 문제를 확인해볼 수 있습니다.

그래서 뮤텍스를 활용합니다.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("레이즈 컨디션")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}
	// 즉시 실행함수입니다.
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
```

다시 명령해보면 문제가 없다고 피드백을 줍니다. 하지만 더 개선할 수 있습니다.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("레이즈 컨디션")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	mut.RLock()
	var score = []int{0}
	mut.RUnlock()

	// 즉시 실행함수입니다.
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
```

읽기 쓰기 모두 뮤텍스를 적용하는 코드입니다. 하지만 잘못된 코드입니다. 읽을 때 lock을 해야 합니다.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - LearnCodeonline.in")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	//wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Four")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
```

이렇게 작성해야 올바릅니다.

# Channels and Deadlock in golang

https://www.youtube.com/watch?v=i5HEE5Ikv3w

이번에는 golang의 채널을 다룹니다. golang의 복잡한 부분을 다룹니다. chanel이 특수한 점은 이해하기 어렵습니다. 초보자들에게 어려운 부분입니다.

go 루틴이 각각 메모리를 접근한다고 했습니다. go 루틴은 각각 독립적으로 통신과 전연변수로 처리하고 싶은데 못 할 수 있습니다.

채널은 각각의 go 루틴이 상호작용할 수 있게 해줍니다. 쓰레드 실행 종료랑 무관하게 실행할 수 있습니다. 채널은 초보자들에게 상당히 어려운 개념입니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("채널")

	myCh := make(chan int)

	// golang의 화살표연산자는 할당합니다. 반대방향은 없습니다.
	myCh <- 5
	fmt.Println(<-myCh)

}
```

이렇게 실행하면 에러를 돌려줍니다.

> 채널
> fatal error: all goroutines are asleep - deadlock!

> goroutine 1 [chan send]:
> main.main()

deadlock 에러입니다. golang의 고전적인 에러입니다. 채널은 고루틴을 배우고 메모리 접근을 이해하고서 활용할 수 있습니다.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("채널")

	// 마지막 인자는 버퍼입니다.
	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}
	// golang의 화살표연산자는 할당합니다. 반대방향은 없습니다.
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		// 메모리상 할당했던 5를 가져옵니다.
		fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		// myCh에 5를 할당합니다.
		myCh <- 5
		myCh <- 6
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
```

다른 고루틴에서 할당된 값을 참조할 수 있습니다.

하지만 문제가 있습니다.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("채널")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
```

> 채널
> 0

이렇게 실행하면 0입니다.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("채널")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 0 // 0을 메모리에 할당합니다.
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
```

이렇게 할당해도 문제가 됩니다. 0이 어떻게 참조된 것인지 모른다는 문제입니다.

하지만 언어의 설계자들이 이미지 다 고안했습니다.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("채널")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen) // 채널의 열고 닫힘 상태를 구분할 수 있습니다.
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
```

채널이 열고 닫혀있는지에 따라 0을 구분할 수 있습니다.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("채널")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// 받기 전용입니다.
		close(myCh) // 이렇게 되면 에러를 돌려줄 것입니다.
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// 보내기 전용입니다.
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
```

현실에서는 버퍼도 자주 활용할 일이 없을 것입니다.

# Math, crypto and random number in golang

https://www.youtube.com/watch?v=yPeYZ0QuspU

이번에는 특수한 부분만 다룹니다.

무작위 숫자를 만드는 법입니다.

초보자에게 가장 혼란스러운 것은 randomnumber입니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Math")
	var myNumberOne int = 2
	var myNumberTwo int64 = 4

	fmt.Println("합: ", myNumberOne+int(myNumberTwo)) // 자동완성으로 형변환이 발생합니다.
}
```

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Math")
	fmt.Println(rand.Intn(5))
}
```

무작위적으로 생성되지 않습니다. 시드를 제공해야 합니다.

컴퓨터에서 무작위적인 것은 없습니다. 그래서 무작위적이게 시간을 활용합니다.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Math")
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5)+1)
}
```

이렇게 무작위적인 숫자처럼 보이게 구현했습니다.

이제는 암호학적 정확성이 필요해지는 상황이 발생합니다.

```go
package main

import (
	"fmt"
	"math/big"
	"crypto/rand"
)

func main() {
	fmt.Println("Math")
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
```

더욱더 확실하게 난수를 생성하는 방법입니다.
