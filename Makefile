all: rust_build
	go run src/greet.go

rust_build: src/lib.rs Cargo.toml
	cargo build --target aarch64-apple-darwin

clean:
	rm -rf target
