refact:
	go mod edit -module github.com/noolingo/sha256password
	-- rename all imported module
	find . -type f -name '*.go' \
  	-exec sed -i -e 's,github.com/noolingo/proto,github.com/noolingo/sha256password,g' {} \;