# SPDX-License-Identifier: Apache-2.0
#
# Copyright (c) nexB Inc. and others. All rights reserved.
# ScanCode is a trademark of nexB Inc.
# SPDX-License-Identifier: Apache-2.0
# See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
# See https://github.com/aboutcode-org/skeleton for support or download.
# See https://aboutcode.org for more information about nexB OSS projects.
#

# Python version can be specified with `$ PYTHON_EXE=python3.x make conf`
PYTHON_EXE?=python3
VENV=venv
ACTIVATE?=. ${VENV}/bin/activate;

ARCH := $(shell uname -m)
ifeq ($(ARCH),aarch64)
	PLAT_NAME := "manylinux2014_aarch64"
	GOARCH := "arm64"
else
	PLAT_NAME := "manylinux1_x86_64"
	GOARCH := "amd64"
endif

conf:
	@echo "-> Install dependencies"
	./configure

dev:
	@echo "-> Configure and install development dependencies"
	./configure --dev

doc8:
	@echo "-> Run doc8 validation"
	@${ACTIVATE} doc8 --quiet docs/ *.rst

valid:
	@echo "-> Run Ruff format"
	@${ACTIVATE} ruff format
	@echo "-> Run Ruff linter"
	@${ACTIVATE} ruff check --fix

check:
	@echo "-> Run Ruff linter validation (pycodestyle, bandit, isort, and more)"
	@${ACTIVATE} ruff check
	@echo "-> Run Ruff format validation"
	@${ACTIVATE} ruff format --check
	@$(MAKE) doc8
	@echo "-> Run ABOUT files validation"
	@${ACTIVATE} about check etc/

clean:
	@echo "-> Clean the Python env"
	./configure --clean

test:
	@echo "-> Run the test suite"
	${VENV}/bin/pytest -vvs

docs:
	rm -rf docs/_build/
	@${ACTIVATE} sphinx-build docs/source docs/_build/

docs-check:
	@${ACTIVATE} sphinx-build -E -W -b html docs/source docs/_build/
	@${ACTIVATE} sphinx-build -E -W -b linkcheck docs/source docs/_build/

build:
	rm -f src/go_inspector/bin/GoReSym_lin
	python setup.py clean --all sdist
	GOOS=linux GOARCH=$(GOARCH) python setup.py clean --all bdist_wheel --python-tag py3 --plat-name $(PLAT_NAME)

.PHONY: conf dev check valid clean test docs docs-check build
