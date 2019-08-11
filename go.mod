module code.qlteam.com

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.3.0
	github.com/willerhe/webbase v0.0.0-20190625231052-22029a293ec2
)

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.40.0
	github.com/willerhe/webbase => /home/home/workspace/my/webbase
	go.uber.org/zap => github.com/uber-go/zap v1.10.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/image => github.com/golang/image v0.0.0-20190622003408-7e034cad6442
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190607214518-6fa95d984e88
	golang.org/x/net => github.com/golang/net v0.0.0-20190514140710-3ec191127204
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190516110030-61b9204099cb
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190517183331-d88f79806bbd
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.7.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.1
	google.golang.org/genproto => github.com/googleapis/go-genproto v0.0.0-20190620144150-6af8c5fc6601
	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.1
	rsc.io/binaryregexp => github.com/rsc/binaryregexp v0.2.0
)
