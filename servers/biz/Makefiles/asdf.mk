.PHONY: asdf_reshim
asdf_reshim:
ifdef GITHUB_ACTION
	@echo "SKIP asdf reshim"
else
	asdf reshim
endif
