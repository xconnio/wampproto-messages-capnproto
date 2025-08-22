import importlib.util
import os

from setuptools import setup

hook_path = os.path.join(os.path.dirname(__file__), "generate_schemas.py")
spec = importlib.util.spec_from_file_location("setup_hook", hook_path)
setup_hook = importlib.util.module_from_spec(spec)
spec.loader.exec_module(setup_hook)

setup(cmdclass={"build_py": setup_hook.CustomBuildPy})
