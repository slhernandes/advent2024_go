MOD_DIRS=$(shell find . -type d | awk -F\"/\" '{print $2}' | xargs)

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

.PHONY: $(MOD_DIRS) init_% deinit_%
.SILENT:
