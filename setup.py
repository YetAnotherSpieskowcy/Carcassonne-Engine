import os
import subprocess

import setuptools
from setuptools.command.build import build

PY_PACKAGE_DIR = "python_bindings/carcassonne_engine/_bindings"
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
    # nothing depends on performance tests
    f"game{os.sep}performancetests",
    f"engine{os.sep}request_performance_tests",
    "end_tests",
    f"end_tests{os.sep}four_player_game_test",
    f"end_tests{os.sep}two_player_game_test",
)
GO_MAIN_PACKAGE = f"{GO_BASE_PACKAGE}{os.sep}{GO_MAIN_PACKAGE_NAME}"


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
            f"{GO_NAMESPACE}{os.sep}{pkg}"
            for pkg, _, _ in os.walk(GO_PKG_DIR)
            if pkg[4:] not in GO_EXCLUDED_PACKAGES
        ][1:]

        try:
            subprocess.check_call(
                (
                    "gopy",
                    "build",
                    "-output",
                    PY_PACKAGE_DIR,
                    GO_MAIN_PACKAGE,
                    *go_packages,
                )
            )
        finally:
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
