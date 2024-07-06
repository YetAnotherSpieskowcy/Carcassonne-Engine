import os
import subprocess

import setuptools
from setuptools.command.build_ext import build_ext as _build_ext

PY_PACKAGE_DIR = "python_bindings/carcassonne_engine"
GO_NAMESPACE = "github.com/YetAnotherSpieskowcy/Carcassonne-Engine"
GO_BASE_PACKAGE = f"{GO_NAMESPACE}/pkg"
GO_MAIN_PACKAGE_NAME = "engine"
GO_EXCLUDED_PACKAGES = (
    # avoid listing main package twice
    GO_MAIN_PACKAGE_NAME,
    # fortunately we shouldn't need it but this package is problematic
    # due to use of generics: https://github.com/go-python/gopy/issues/283
    "stack",
)
GO_MAIN_PACKAGE = f"{GO_BASE_PACKAGE}/{GO_MAIN_PACKAGE_NAME}"
GO_PACKAGES = tuple(
    f"{GO_NAMESPACE}/{pkg}"
    for pkg, _, _ in os.walk("pkg")
    if pkg != "pkg" and pkg[4:] not in GO_EXCLUDED_PACKAGES
)


class BinaryDistribution(setuptools.Distribution):
    def has_ext_modules(_):
        return True


def generate_and_build_extension_files():
    subprocess.check_call(
        (
            "gopy",
            "build",
            "-output",
            PY_PACKAGE_DIR,
            GO_MAIN_PACKAGE,
            *GO_PACKAGES,
        )
    )
    with open(f"{PY_PACKAGE_DIR}/__init__.py", "w", encoding="utf-8") as fp:
        fp.write("from .engine import *\n")


# TODO(Jackenmen): figure out correct place to generate / build extension files
class build_ext(_build_ext):
    def build_extensions(self) -> None:
        generate_and_build_extension_files()
        super().build_extensions()


# this is just a (slightly slower) workaround
generate_and_build_extension_files()


setuptools.setup(
    include_package_data=True,
    distclass=BinaryDistribution,
)
