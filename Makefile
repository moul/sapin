all: sapin sapin.js


sapin: ./cmd/sapin/main.go ./sapin.go
	go build -o $@ ./cmd/sapin


sapin.js: ./cmd/sapinjs/main.go ./sapin.go
	go get github.com/gopherjs/gopherjs
	gopherjs build -o $@ ./cmd/sapinjs/main.go


clean:
	rm -rf sapin sapin.js sapin.js.map


goapp_serve:
	goapp serve ./cmd/appspot/app.yaml


goapp_deploy:
	goapp deploy -application sapin-as-a-service ./cmd/appspot/app.yaml
