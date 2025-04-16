# ====================================================================================
# Setup Project

include ../../../build/makelib/common.mk

# ====================================================================================
#  Options

include ../../../build/makelib/imagelight.mk

# ====================================================================================
# Targets

# Dynamically set RUNTIME variable
RUNTIME := $(shell if which docker >/dev/null 2>&1; then echo docker; elif which podman >/dev/null 2>&1; then echo podman; else echo none; fi)

# Ensure RUNTIME is valid
ifeq ($(RUNTIME),none)
$(ERROR) Neither docker nor podman is available in PATH
endif

img.build:
	@$(INFO) $(RUNTIME) build $(IMAGE)
	@$(MAKE) BUILD_ARGS="--load" img.build.shared
	@$(OK) $(RUNTIME) build $(IMAGE)

img.publish:
	@$(INFO) Skipping image publish for $(IMAGE)
	@echo Publish is deferred to xpkg machinery
	@$(OK) Image publish skipped for $(IMAGE)

img.build.shared:
	@cp Dockerfile $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp terraformrc.hcl $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp -r $(OUTPUT_DIR)/bin/ $(IMAGE_TEMP_DIR)/bin || $(FAIL)
	@$(RUNTIME) buildx build $(BUILD_ARGS) \
		--platform $(IMAGE_PLATFORMS) \
		--build-arg TERRAFORM_VERSION=$(TERRAFORM_VERSION) \
		--build-arg TERRAFORM_PROVIDER_SOURCE=$(TERRAFORM_PROVIDER_SOURCE) \
		--build-arg TERRAFORM_PROVIDER_VERSION=$(TERRAFORM_PROVIDER_VERSION) \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_NAME=$(TERRAFORM_PROVIDER_DOWNLOAD_NAME) \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX=$(TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX) \
		--build-arg TERRAFORM_NATIVE_PROVIDER_BINARY=$(TERRAFORM_NATIVE_PROVIDER_BINARY) \
		-t $(IMAGE) \
		$(IMAGE_TEMP_DIR) || $(FAIL)

img.promote:
	@$(INFO) Skipping image promotion from $(FROM_IMAGE) to $(TO_IMAGE)
	@echo Promote is deferred to xpkg machinery
	@$(OK) Image promotion skipped for $(FROM_IMAGE) to $(TO_IMAGE)
