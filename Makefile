MOD_DIRS=$(shell find . -type d | grep day | xargs)
MOD_DIRS_DEBUG=$(shell find . -type d | grep day | sed "s/$$/_debug/g" | xargs)

deinit_%:
	$(eval DIRNAME=$(subst deinit_,,$(@)))
	./deinit_module.sh $(DIRNAME)
	go work edit -dropuse=./$(DIRNAME)

init_%:
	$(eval DIRNAME=$(subst init_,,$(@)))
	./init_module.sh $(DIRNAME)
	cd $(DIRNAME) && mmv 'template*' '$(DIRNAME)#1'
	cd $(DIRNAME) && go mod init aoc/$(DIRNAME)
	go work use ./$(DIRNAME)

$(MOD_DIRS):
	cd $@ && $(MAKE) run

$(MOD_DIRS_DEBUG):
	$(eval DIRNAME=$(subst _debug,,$(@)))
	cd $(DIRNAME) && $(MAKE) debug


.PHONY: $(MOD_DIRS) init_% deinit_% $(MOD_DIRS_DEBUG)
.SILENT:
