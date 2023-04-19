MAKEFLAGS       += --warn-undefined-variables
SHELL           := /bin/bash -o pipefail

# PUBLIC is set to "--public" when making the publicly exported repo. Used by our CI.
PUBLIC ?= ""

export GO111MODULE = on

ci: setup generate diff-check

default: generate spelling
generate: codegen format

.PHONY: codegen
codegen:
codegen:
	rm -rf jsonschemas
	rego jsonschema -d specs -o jsonschemas $(PUBLIC)
	elegen folder -d specs -o codegen --split-output -g openapi3 $(PUBLIC) || exit 1
	rm -rf openapi3
	elegen folder -d specs -o codegen $(PUBLIC) || exit 1
	mv custom_validations.go custom_validations.go.keep
	mv custom_validations_test.go custom_validations_test.go.keep
	rm -rf ./*.go
	mv custom_validations.go.keep custom_validations.go
	mv custom_validations_test.go.keep custom_validations_test.go
	mv codegen/elemental/*.go ./
	mv codegen/openapi3 openapi3
	rm -rf codegen
	data=$$(rego doc -d specs || exit 1) && echo -e "$${data}" > doc/documentation.md

.PHONY: diff-check
diff-check:
	git update-index -q --really-refresh
	git diff-index --quiet HEAD -- || (git diff && false)

format: format-specs format-type format-validation format-parameter
format-specs:
	for f in specs/*.abs; do \
		rego format < $$f > $$f.formatted && \
		mv $$f.formatted $$f; \
	done
	for f in specs/*.spec; do \
		rego format < $$f > $$f.formatted && \
		mv $$f.formatted $$f; \
	done

format-type: target = "specs/_type.mapping"
format-type:
	rego format -m typemapping < $(target) > $(target).formatted
	mv $(target).formatted $(target)

format-validation: target = "specs/_validation.mapping"
format-validation:
	rego format -m validationmapping < $(target) > $(target).formatted
	mv $(target).formatted $(target)

format-parameter: target = "specs/_parameter.mapping"
format-parameter:
	rego format -m parametermapping < $(target) > $(target).formatted
	mv $(target).formatted $(target)

spelling:
	docker run --rm -v $$PWD:/workdir gcr.io/prismacloud-cns/markdown-spellcheck:latest "doc/*.md" -r -a -n --en-us

codecgen:
	rm -f values_codecgen.go ; codecgen -o values_codecgen.go *.go;
	cd types && rm -f values_codecgen.go ; codecgen -o values_codecgen.go *.go;

update-spectools: setup
setup:
	@ if [ -x $$APO_ROOT/bin/update-spectools ]; then \
			$$APO_ROOT/bin/update-spectools ; \
	else \
		go install go.aporeto.io/regolithe/cmd/rego@master ; \
		go install go.aporeto.io/elemental/cmd/elegen@master ; \
	fi
