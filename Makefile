refact:
	go mod edit -module github.com/noolingo/proto
	-- rename all imported module
	find . -type f -name '*.go' \
  	-exec sed -i -e 's,github.com/MelnikovNA/sha256password,github.com/noolingo/sha256password,g' {} \;