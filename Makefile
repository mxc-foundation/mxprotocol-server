.PHONY: subdirs all clean
SUBDIRS=m2m-wallet

subdirs:
	@for subdir in $(SUBDIRS); \
	do \
	echo "build in $$subdir"; \
	( cd $$subdir && make build) || exit 1; \
	done

all: subdirs

clean:
	@for subdir in $(SUBDIRS); \
	do \
	echo "cleaning in $$subdir"; \
	( cd $$subdir && make clean) || exit 1; \
	done

