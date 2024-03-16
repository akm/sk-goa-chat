.PHONY: git_check
git_check: git_check_uncommited_changes git_check_untracked_files

.PHONY: git_check_uncommited_changes
git_check_uncommited_changes:
	@git diff --exit-code > /dev/null && \
		echo "no uncommited changes" || \
		( echo "uncommited changes exists" && git diff && exit 1 )

GIT_UNTRACKED_FILES := $(shell git ls-files . --exclude-standard --others)

.PHONY: git_check_untracked_files
git_check_untracked_files:
	@if [ "$(GIT_UNTRACKED_FILES)" = "" ]; then \
	  echo "No untracked file" ; \
	else \
	  echo "There is untracked file(s): $(GIT_UNTRACKED_FILES)" && git status && exit 1 ; \
	fi
