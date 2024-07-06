import os
import subprocess

import setuptools
from setuptools.command.build import build

PY_PACKAGE_DIR = "python_bindings/carcassonne_engine"
GO_MOD = "go.mod"
GO_PKG_DIR = "pkg"
GO_NAMESPACE = "github.com/YetAnotherSpieskowcy/Carcassonne-Engine"
GO_BASE_PACKAGE = f"{GO_NAMESPACE}/{GO_PKG_DIR}"
GO_MAIN_PACKAGE_NAME = "engine"
GO_EXCLUDED_PACKAGES = (
    # avoid listing main package twice
    GO_MAIN_PACKAGE_NAME,
    # fortunately we shouldn't need it but this package is problematic
    # due to use of generics: https://github.com/go-python/gopy/issues/283
    "stack",
)
GO_MAIN_PACKAGE = f"{GO_BASE_PACKAGE}/{GO_MAIN_PACKAGE_NAME}"


class BinaryDistribution(setuptools.Distribution):
    def has_ext_modules(self):
        return True


class BuildGoCommand(setuptools.Command):
    command_name = "build_go"

    def initialize_options(self) -> None:
        pass

    def finalize_options(self) -> None:
        pass

    def get_source_files(self) -> list[str]:
        source_files = [GO_MOD]
        for dirname, _, files in os.walk(GO_PKG_DIR):
            for filename in files:
                source_files.append(os.path.join(dirname, filename))
        return source_files

    def run(self) -> None:
        go_packages = [
            f"{GO_NAMESPACE}/{pkg}"
            for pkg, _, _ in os.walk(GO_PKG_DIR)
            if pkg[4:] not in GO_EXCLUDED_PACKAGES
        ][1:]

        init_filepath = os.path.join(PY_PACKAGE_DIR, "__init__.py")
        with open(init_filepath, "rb") as fp:
            init_contents = fp.read()
        try:
            subprocess.check_call(
                (
                    "gopy",
                    "build",
                    # though this would be nicer to work with on Python side,
                    # it is buggy and converts only some occurrences to camel_case
                    # breaking the whole library in the process
                    # "-rename=true",
                    "-output",
                    PY_PACKAGE_DIR,
                    GO_MAIN_PACKAGE,
                    *go_packages,
                )
            )
        finally:
            with open(init_filepath, "wb") as fp:
                fp.write(init_contents)
            try:
                os.remove(os.path.join(PY_PACKAGE_DIR, "build.py"))
            except FileNotFoundError:
                pass


class CustomBuild(build):
    sub_commands = [("build_go", None)] + build.sub_commands


setuptools.setup(
    distclass=BinaryDistribution,
    cmdclass={"build": CustomBuild, "build_go": BuildGoCommand},
)
