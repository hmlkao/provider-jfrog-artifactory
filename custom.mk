
# 1: registry/org 2: repo
define xpkg.set-extensions
xpkg.set-extensions.$(1).$(2): $(UP)
	@$(INFO) Setting extensions for package $(1)/$(2):$(VERSION)
	@$(UP) alpha xpkg append \
		--extensions-root=./extensions \
		$(1)/$(2):$(VERSION)
	@$(OK) Extensions set for package $(1)/$(2):$(VERSION)
xpkg.set-extensions: xpkg.set-extensions.$(1).$(2)
endef
$(foreach r,$(XPKG_REG_ORGS), $(foreach x,$(XPKGS), $(eval $(call xpkg.set-extensions,$(r),$(x)))))

# Only set extensions for specific branches
set-extensions: ; @:
ifneq ($(filter main master release-%,$(BRANCH_NAME)),)
set-extensions: $(foreach r,$(XPKG_REG_ORGS), $(foreach x,$(XPKGS),xpkg.set-extensions.$(r).$(x)))
endif
