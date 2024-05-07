all: target/x86_64-apple-darwin/debug/librust_greet.a
	go run src/greet.go

target/x86_64-apple-darwin/debug/librust_greet.a: src/lib.rs Cargo.toml
	cargo build --target x86_64-apple-darwin

clean:
	rm -rf target
