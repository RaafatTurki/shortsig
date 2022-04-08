DOCS_DIR = ./docs

doc_gen:
	pandoc --standalone --to man ${DOCS_DIR}/shortsig.1.md -o ${DOCS_DIR}/shortsig.1

.PHONY: doc_gen
