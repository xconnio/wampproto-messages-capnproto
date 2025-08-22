import os
import shutil
import glob

from setuptools.command.build_py import build_py as _build_py


class CustomBuildPy(_build_py):
    def run(self):
        schema_dst = os.path.join("wamp_msgs_capnp", "schemas")
        schema_src = os.path.abspath(os.path.join("..", "capnproto"))

        os.makedirs(schema_dst, exist_ok=True)

        for file in glob.glob(os.path.join(schema_src, "*.capnp")):
            shutil.copy(file, schema_dst)

        for capnp_file in glob.glob(os.path.join(schema_dst, "*.capnp")):
            with open(capnp_file, "r", encoding="utf-8") as f:
                lines = f.readlines()
            with open(capnp_file, "w", encoding="utf-8") as f:
                for line in lines:
                    if line.startswith("using Go") or line.lstrip().startswith("$Go."):
                        f.write("#" + line)
                    else:
                        f.write(line)

        super().run()
