all: sapin sapin.js


sapin: ./cmd/sapin/main.go ./sapin.go
	go build -o $@ ./cmd/sapin


sapin.js: ./cmd/sapinjs/main.go ./sapin.go
	gopherjs build -o $@ ./cmd/sapinjs/main.go


clean:
	rm -rf sapin sapin.js sapin.js.map
