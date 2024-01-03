# This file uses the following variables or tasks from
#
# Depends on go.mk
#   - GO_ROOT_PACKAGE
GOA_ROOT_PACKAGE=$(GO_ROOT_PACKAGE)
GOA_SERVICES_DIR=./services

GOA_GEN_DIR=./gen
$(GOA_GEN_DIR): goa_gen

.PHONY: goa_gen
goa_gen:
	goa gen $(GOA_ROOT_PACKAGE)/design -o $(GOA_SERVICES_DIR) && \
	$(MAKE) -C ../modifiers apisvr_services_gen

.PHONY: goa_example
goa_example:
	goa example $(GOA_ROOT_PACKAGE)/design -o $(GOA_SERVICES_DIR) && \
	$(MAKE) -C ../modifiers apisvr_services_cmd
