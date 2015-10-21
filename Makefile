SOURCES = sapin.go bonus.go

all: sapin sapin.js


.PHONY: build
build: sapin


.PHONY: test
test:
	go test -v .


.PHONY: cover
cover:
	rm -rf profile.out
	go test -covermode=count -coverpkg=. -coverprofile=profile.out .



.PHONY: convey
convey:
	go get github.com/smartystreets/goconvey
	goconvey -cover -port=9031 -workDir="$(pwd)" -depth=1


sapin: ./cmd/sapin/main.go $(SOURCES)
	go build -o $@ ./cmd/sapin


sapin.js: ./cmd/sapinjs/main.go $(SOURCES)
	go get github.com/gopherjs/gopherjs
	gopherjs build -o $@ ./cmd/sapinjs/main.go


.PHONY: clean
clean:
	rm -rf sapin sapin.js sapin.js.map


.PHONY: goapp_serve
goapp_serve: cmd/appspot/static/terminal.css
	goapp serve ./cmd/appspot/app.yaml


.PHONY: goapp_deploy
goapp_deploy: cmd/appspot/static/terminal.css
	goapp deploy -application sapin-as-a-service ./cmd/appspot/app.yaml


cmd/appspot/static/terminal.css:
	wget https://raw.githubusercontent.com/buildkite/terminal/master/assets/terminal.css -O $@
