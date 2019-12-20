BAZEL ?= bazel
PWD := $(shell echo ${PWD})

CUSTOM_DAYS := day4 day18
DAYS := $(filter-out $(CUSTOM_DAYS),$(wildcard day*))

.PHONY: $(DAYS)
$(DAYS):
	$(BAZEL) run //cmd/$@:$@ -- $(PWD)/$@/input.txt

.PHONY: day4
day4:
	$(BAZEL) run //cmd/day4:day4 -- 246540 787419

.PHONY: day18
day18:
	$(BAZEL) run //cmd/day18:day18 -- $(PWD)/day18/input.txt $(PWD)/day18/input2.txt

.PHONY: test
test:
	$(BAZEL) test //... --test_output=all

.PHONY: gen
gen:
	$(BAZEL) run //:gazelle

.PHONY: deps
deps:
	$(BAZEL) run //:gazelle -- update-repos --from_file=go.mod