.PHONY: copyright.verify
copyright.verify: tools.verify.addlicense
	@echo "===================================> Verifying all files"
	@addlicense --check -f $(ROOT_DIR)/scripts/copyright.txt $(ROOT_DIR) --skip-dirs=third_party,vendor,_output

.PHONY: copyright.add
copyright.add: tools.verify.addlicense
	@addlicense -v -f $(ROOT_DIR)/scripts/copyright.txt $(ROOT_DIR) --skip-dirs=third_party,vendor,_output