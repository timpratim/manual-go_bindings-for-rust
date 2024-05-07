all: target/aarch64-apple-darwin/debug/librust_greet.a
	go run src/greet.go

target/aarch64-apple-darwin/debug/librust_greet.a: src/lib.rs Cargo.toml
	cargo build --target aarch64-apple-darwin

clean:
	rm -rf target
