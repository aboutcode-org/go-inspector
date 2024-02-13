# -*- coding: utf-8 -*-
#
# Copyright (c) nexB Inc. and others. All rights reserved.
# ScanCode is a trademark of nexB Inc.
# SPDX-License-Identifier: Apache-2.0
# See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
# See https://github.com/nexB/scancode-plugins for support or download.
# See https://aboutcode.org for more information about nexB OSS projects.
#

import json
import logging
import os

from commoncode import command
from commoncode import fileutils
from commoncode.functional import flatten

from src.go_inspector.plugin import collect_and_parse_symbols

TEST_DATA_DIR = os.path.join(os.path.dirname(__file__), "data")


def test_goresym_with_windows_exe():
    go_binary = os.path.join(TEST_DATA_DIR, "windows.exe")
    goresym_output = collect_and_parse_symbols(go_binary)
    assert goresym_output is None


def test_goresym_with_elf():
    go_binary = os.path.join(TEST_DATA_DIR, "arm_gentoo_elf")
    goresym_output = collect_and_parse_symbols(go_binary)
    assert goresym_output is None


def test_goresym_with_goresym_lin():
    go_binary = os.path.join(TEST_DATA_DIR, "GoReSym_lin")
    goresym_output = collect_and_parse_symbols(go_binary)
    # write data to a file
    with open(os.path.join(TEST_DATA_DIR, "gore_sym_output.json")) as f:
        expected_output = json.load(f)
        assert expected_output == goresym_output
