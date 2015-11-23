SOURCES :=	$(shell find . -name "*.go")
VERSION :=	$(shell cat .goxc.json | jq -c .PackageVersion | sed 's/"//g')
GOENV ?=	GO15VENDOREXPERIMENT=1
GO ?=		$(GOENV) go
GODEP ?=	$(GOENV) godep


all: sapin sapin.js


.PHONY: build
build: sapin


.PHONY: install
install:
	$(GO) install ./cmd/sapin


.PHONY: test
test:
	$(GODEP) restore
	$(GO) get -t .
	$(GO) test -v .


.PHONY: godep-save
godep-save:
	$(GODEP) save $(shell go list ./... | grep -v /vendor/)


.PHONY: cover
cover:
	rm -rf profile.out
	$(GO) test -covermode=count -coverpkg=. -coverprofile=profile.out .


.PHONY: convey
convey:
	$(GO) get github.com/smartystreets/goconvey
	goconvey -cover -port=9031 -workDir="$(shell realpath .)" -depth=-1


sapin: ./cmd/sapin/main.go $(SOURCES)
	$(GO) build -o $@ ./cmd/sapin


sapin.js: ./cmd/sapinjs/main.go $(SOURCES)
	$(GO) get github.com/gopherjs/gopherjs
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


.PHONY: release
release:
	goxc


.PHONY: build-docker
build-docker: contrib/docker/.docker-container-built
	@echo "now you can 'docker push moul/sapin'"


dist/latest/sapin_latest_linux_386: $(SOURCES)
	mkdir -p dist
	rm -f dist/latest
	(cd dist; ln -s $(VERSION) latest)
	goxc -bc="linux,386" xc
	cp dist/latest/sapin_$(VERSION)_linux_386 $@


contrib/docker/.docker-container-built: dist/latest/sapin_latest_linux_386
	cp $< contrib/docker/sapin
	docker build -t moul/sapin:latest contrib/docker
	docker tag -f moul/sapin:latest moul/sapin:$(shell echo $(VERSION) | sed 's/\+/plus/g')
	docker run -it --rm moul/sapin --size=1
	docker inspect --type=image --format="{{ .Id }}" moul/sapin > $@.tmp
	mv $@.tmp $@
