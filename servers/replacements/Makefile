.PHONY: setup
setup:
	gem install fukowl -v 0.1.1

.PHONY: run_ruby_in_dir
run_ruby_in_dir: setup
	find ${TARGET_DIR} -name '*.rb' | xargs -I{} ruby {}

.PHONY: apisvr_services_cmd
apisvr_services_cmd:
	TARGET_DIR=../replacements/apisvr/services/cmd $(MAKE) run_ruby_in_dir

.PHONY: apisvr_services_gen
apisvr_services_gen:
	TARGET_DIR=../replacements/apisvr/services/gen $(MAKE) run_ruby_in_dir

.PHONY: biz_models
biz_models: 
	TARGET_DIR=../replacements/biz/models $(MAKE) run_ruby_in_dir
