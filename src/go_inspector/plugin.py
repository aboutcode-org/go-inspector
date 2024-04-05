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

import attr
from commoncode import command
from commoncode import fileutils
from commoncode.cliutils import SCAN_GROUP
from commoncode.cliutils import PluggableCommandLineOption
from plugincode.scan import ScanPlugin
from plugincode.scan import scan_impl
from typecode import contenttype
from typecode.contenttype import get_type

"""
Extract symbols information from source code files with ctags.
"""
LOG = logging.getLogger(__name__)

from os.path import abspath
from os.path import dirname
from os.path import join

"""
https://github.com/mandiant/GoReSym/archive/refs/tags/v2.6.4.tar.gz
Download this zip at:
https://github.com/mandiant/GoReSym/releases/download/v2.6.4/GoReSym.zip
Extract goresym for linux
And then change it to executable 
chmod u+x src/go_inspector/bin/GoReSym_lin
"""


def get_goresym_location():
    curr_dir = dirname(abspath(__file__))
    return join(curr_dir, "bin", "GoReSym_lin")


@scan_impl
class GoSymbolScannerPlugin(ScanPlugin):
    """
    Scan a Go binary for symbols using GoReSym.
    """

    resource_attributes = dict(
        go_symbols=attr.ib(default=attr.Factory(dict), repr=False),
    )

    options = [
        PluggableCommandLineOption(
            ("--go-symbol",),
            is_flag=True,
            default=False,
            help="Collect Go symbols.",
            help_group=SCAN_GROUP,
            sort_order=100,
        ),
    ]

    def is_enabled(self, go_symbol, **kwargs):
        return go_symbol

    def get_scanner(self, **kwargs):
        return collect_and_parse_symbols


def is_macho(location):
    """
    Return True if the file at ``location`` is macho,
    otherwise False
    """
    t = get_type(location)
    return t.filetype_file.lower().startswith("mach-o") or t.mimetype_file.lower().startswith(
        "application/x-mach-binary"
    )


def is_executable_binary(location):
    """
    Return True if the file at ``location`` is an executable binary.
    """

    if not os.path.exists(location):
        return False

    if not os.path.isfile(location):
        return False

    type = contenttype.Type(location)

    if not (type.is_elf or type.is_winexe or is_macho(location=location)):
        return False

    return True



def collect_and_parse_symbols(location, check_type=True, **kwargs):
    """
    Run GoReSym and return a mapping of symbols of interest for the Go binary file
    at ``location``.
    If ``check_type`` is True, the file is checked.
    """
    if check_type and not is_executable_binary(location):
        # print("Not an executable binary")
        return

    goresym_args = ["-p", location]
    goresym_temp_dir = fileutils.get_temp_dir()
    envt = {"TMPDIR": goresym_temp_dir}

    try:
        rc, stdo, err = command.execute(
            cmd_loc=get_goresym_location(),
            args=goresym_args,
            env=envt,
            to_files=True,
        )

        if rc != 0:
            raise Exception(open(err).read())

        with open(stdo) as syms:
            symbols = json.load(syms)
            files = symbols.get("Files") or []
            files.sort()
            return dict(
                go_symbols=dict(build_info=symbols.get("BuildInfo") or {}, file_paths=files or [])
            )

    finally:
        fileutils.delete(goresym_temp_dir)
