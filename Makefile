CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := build/jsons.a build/jsons.so

.PHONY: all
all: $(LIBS)

build/jsons.a: jsons.go
	$(CGO) $(STATIC) -o build/jsons.a $<

build/jsons.so: jsons.go
	$(CGO) $(SHARED) -o build/jsons.so $<

.PHONY: clean
clean:
	find build -type f \( -name '*.h' -o -name '*.so' -o -name '*.a' \) -delete
