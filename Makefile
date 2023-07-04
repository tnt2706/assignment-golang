.PHONY:	start gen-cal

start:
	air

gen-cal:
	./pkg/calculator/proto/*.proto \
    --go_out=. \
    --go_opt=paths=source_relative \
    --proto_path=.