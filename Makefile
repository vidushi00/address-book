generate_address_book_grpc:
	protoc \
	--go_out=address-book \
	--go_opt=paths=source_relative \
	--go-grpc_out=address-book \
	--go-grpc_opt=paths=source_relative \
	address-book.proto