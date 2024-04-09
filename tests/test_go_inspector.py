# -*- coding: utf-8 -*-
#
# Copyright (c) nexB Inc. and others. All rights reserved.
# ScanCode is a trademark of nexB Inc.
# SPDX-License-Identifier: Apache-2.0
# See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
# See https://github.com/nexB/g-inspector for support or download.
# See https://aboutcode.org for more information about nexB OSS projects.
#

import json
import os

import pytest

from go_inspector.plugin import collect_and_parse_symbols

TEST_DATA_DIR = os.path.join(os.path.dirname(__file__), "data")


def test_collect_and_parse_symbols_with_plain_windows_exe():
    go_binary = os.path.join(TEST_DATA_DIR, "plain_windows.exe")
    with pytest.raises(Exception) as e:
        collect_and_parse_symbols(go_binary)


def test_collect_and_parse_symbols_with_plain_elf():
    go_binary = os.path.join(TEST_DATA_DIR, "plain_arm_gentoo_elf")
    with pytest.raises(Exception) as e:
        collect_and_parse_symbols(go_binary)


@pytest.mark.parametrize(
    "exe_path",
    [
        "basic/app_lin_exe",
        "basic/app_mac_exe",
        "basic/app_win_exe",
    ],
)
def test_collect_and_parse_symbols_with_mini_go_app_linux(exe_path):
    go_binary = os.path.join(TEST_DATA_DIR, exe_path)
    goresym_output = collect_and_parse_symbols(go_binary)
    with open(os.path.join(TEST_DATA_DIR, f"{exe_path}-goresym.json")) as f:
        expected_output = json.load(f)
        assert expected_output == goresym_output
